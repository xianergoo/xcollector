package receiver

import (
	"fmt"
	"sync"
	"time"
)

type TTransceiver struct {
	TEngine
	TPackageQueue
	Enabled bool
}

type TPackageQueue struct {
	queue []TSxPackage
	mu    sync.Mutex
}

func (q *TPackageQueue) Enqueue(pkg TSxPackage) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, pkg)
}

// Count 返回队列中元素的数量
func (q *TPackageQueue) Count() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.queue)
}

// Clear 清空队列
func (q *TPackageQueue) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = q.queue[:0]
}

func (t *TTransceiver) Execute() {
	t.queue = make([]TSxPackage, 10)
	for {
		err := t.collectData()
		if err != nil {
			//
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (t *TTransceiver) collectData() (err error) {
	hdrMaskBits := byte(1) // starting mask bits is 01
	hdrMaskSize := byte(1)
	lastAddress := byte(0)

	for hdrMaskSize < 8 {
		sndPackage := TSxPackage{
			DstAddr: (hdrMaskBits << (8 - hdrMaskSize)) | (^hdrMaskBits & 0xFF),
			SrcAddr: 1,
			Command: 4,
		}

		rcBytes := []byte{}
		if sndPackage.DstAddr >= 2 {
			rcBytes, err = t.SendAndRecv(sndPackage.Data)
			if err != nil {
				return err
			}
		}

		if len(rcBytes) == 0 {
			if sndPackage.DstAddr == 0xFF {
				return nil
			}
			hdrMaskSize++
			hdrMaskBits = 2 * (hdrMaskBits + 1)
			continue
		}

		for len(rcBytes) > 0 {
			pkgSize := SxExamBuffer(rcBytes)
			if pkgSize <= 0 {
				break
			}

			if pkgSize > 0 {
				t.Enqueue(BytesToSxPackage(rcBytes[:pkgSize]))
				if len(rcBytes) >= pkgSize {
					rcBytes = append(rcBytes[pkgSize:], nil...)
					rcBytes = rcBytes[:len(rcBytes)-pkgSize]
				}
			}
		}

		if Count() > 0 {
			fmt.Printf("Result is %d\n", Count())
			break
		}

		lastAddress = sndPackage.DstAddr
		hdrMaskSize++
		hdrMaskBits = 2 * hdrMaskBits

	}
	if hdrMaskSize == 8 {
		fmt.Printf("Result is %d\n", -(lastAddress & 0x7F))
	}

}

func (t *TTransceiver) handleCollision(device_id int) {
	// 	with FEngine.Readers[deviceId] do begin
	//     TextOut(3, 1, 'Duplicated ID' + IntToStr(deviceId));
	//     Beep(3, 1);
	//     LogOff;
	//     Serial := '';
	//     Worker := '';
	//     WorkerCard := '';
	//   end;
}

func (t *TTransceiver) processPackage() {
	for packages.Count() > 0 {
		pkg := packages.Dequeue()

		reader, exists := e.Readers[pkg.SrcAddr]
		if !exists {
			reader = &TSxReader{
				ID:         pkg.SrcAddr,
				Serial:     "",
				Worker:     "",
				WorkerCard: "",
			}
			e.Readers[pkg.SrcAddr] = reader
		}

		newId := byte(0)

		if pkg.SerialNum != "" {
			if reader.Serial == "" {
				reader.Serial = pkg.SerialNum
			} else if (reader.ID == e.FirstReaderId) || (reader.Serial != pkg.SerialNum) {
				newId = e.GetFreeReaderId()
			}
		} else {
			if reader.Serial == "" {
				reader.Serial = reader.GetSerial()
			}
		}

		if newId > 0 {
			reader.SetDeviceId(pkg.SerialNum, newId)
			reader = e.Readers[newId]
			reader.Serial = pkg.SerialNum
			reader.Worker = ""
			reader.WorkerCard = ""
		}

		reader.DeleteData(pkg.DataStr[:10])

		_workerCard := "2E0055300A00"
		_workerNo := "123"

		if pkg.DataId == "" {
			if pkg.DataStr != _workerCard {
				reader.Beep(3, 0)
				reader.TextOut(0, 1, "Invalid User Card")
				continue
			}

			if reader.Worker == "" {
				reader.Worker = _workerNo
				reader.WorkerCard = pkg.DataStr
			} else if reader.Worker != _workerNo {
				errorMsg := fmt.Sprintf("Duplicate reader[%d] error %s", reader.ID, "Unknow Reason")
				e.Logger.Warning(errorMsg)
				reader.TextOut(0, 1, errorMsg)
				reader.Beep(1, 0)
				reader.LogOff()
				continue
			}

			e.Logger.Debug("Reader[%d] Logon:[%s-%s]", reader.ID, reader.Worker, reader.WorkerCard)
			reader.LogOn(reader.Worker)
			continue
		} else {
			if reader.Worker == "" {
				reader.Worker = pkg.DataId
				reader.WorkerCard = _workerCard
			} else if reader.Worker != pkg.DataId {
				errorMsg := fmt.Sprintf("Duplicate reader[%d] error %s", reader.ID, "F1 Login")
				e.Logger.Warning(errorMsg)
				reader.TextOut(0, 1, errorMsg)
				reader.Beep(3, 0)
				reader.LogOff()
				continue
			}
		}

		if pkg.DataStr == reader.WorkerCard {
			e.Logger.Debug("Reader[%d] LogOff:[%s-%s]", reader.ID, reader.Worker, reader.WorkerCard)
			reader.Worker = ""
			reader.WorkerCard = ""
			reader.Beep(1, 0)
			reader.LogOff()
		} else {
			reader.Beep(3, 0)
			reader.TextOut(0, 1, "Invalid Card")
			e.Logger.Debug("Reader[%d] SCANCARD:[%s]", reader.ID, pkg.DataStr)
		}
	}
}
}
