package receiver

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type TTransceiver struct {
	TEngine
	Readers      [127]TReader
	PackageQueue TPackageQueue
	Enabled      bool
}

type TPackageQueue struct {
	queue []TSxPackage
	mu    sync.Mutex
}

func (q *TPackageQueue) Dequeue() TSxPackage {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue[len(q.queue)-1]
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
	for {
		err := t.collectData()
		if err != nil {
			//
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (t *TTransceiver) getFreeReaderId() (id int, err error) {
	for _, reader := range t.Readers {
		if reader.Serial != "" {
			return int(reader.ID), nil
		}
	}
	return -1, errors.New("can not find valid device id")
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
				t.PackageQueue.Enqueue(BytesToSxPackage(rcBytes[:pkgSize]))
				if len(rcBytes) >= pkgSize {
					rcBytes = append(rcBytes[pkgSize:], nil...)
					rcBytes = rcBytes[:len(rcBytes)-pkgSize]
				}
			}
		}

		// if Count() > 0 {
		// 	fmt.Printf("Result is %d\n", Count())
		// 	break
		// }

		lastAddress = sndPackage.DstAddr
		hdrMaskSize++
		hdrMaskBits = 2 * hdrMaskBits

	}
	if hdrMaskSize == 8 {
		fmt.Printf("Result is %d\n", -(lastAddress & 0x7F))
	}
	return nil

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
	for t.PackageQueue.Count() > 0 {
		pkg := t.PackageQueue.Dequeue()

		reader := t.Readers[pkg.SrcAddr]
		if reader.Serial != "" {
			reader := &TReader{
				// ID:         pkg.SrcAddr,
				// Serial:     "",
				// Worker:     "",
				// WorkerCard: "",
			}
			t.Readers[pkg.SrcAddr] = *reader
		}

		rcvPackage, err := BytesToSxPackageData(pkg.Data)
		if err != nil {
			return
		}
		newId := -1
		if rcvPackage.Serial != "" {
			if reader.Serial == "" {
				reader.Serial = rcvPackage.Serial
			} else if (reader.ID == 2) || (reader.Serial != rcvPackage.Serial) {
				newId, err = t.getFreeReaderId()
			}
		} else {
			if reader.Serial == "" {
				reader.Serial = reader.GetSerial()
			}
		}

		if newId > 0 {
			reader.SetDeviceId(rcvPackage.Serial, newId)
			reader = t.Readers[newId]
			reader.Serial = rcvPackage.Serial
			reader.Worker = ""
			reader.WorkerCard = ""
		}

		reader.DeleteData(pkg.Data[:10])

		_workerCard := "2E0055300A00"
		_workerNo := "123"

		if rcvPackage.WorkerNo == "" {
			if rcvPackage.Data != _workerCard {
				reader.Beep(3, 0)
				reader.TextOut(0, 1, "Invalid User Card")
				continue
			}

			if reader.Worker == "" {
				reader.Worker = _workerNo
				reader.WorkerCard = rcvPackage.Data
			} else if reader.Worker != _workerNo {
				errorMsg := fmt.Sprintf("Duplicate reader[%d] error %s", reader.ID, "Unknow Reason")
				// e.Logger.Warning(errorMsg)
				reader.TextOut(0, 1, errorMsg)
				reader.Beep(1, 0)
				reader.LogOff()
				continue
			}

			// t.Logger.Debug("Reader[%d] Logon:[%s-%s]", reader.ID, reader.Worker, reader.WorkerCard)
			reader.LogOn(reader.Worker)
			continue
		} else {
			if reader.Worker == "" {
				reader.Worker = rcvPackage.WorkerNo
				reader.WorkerCard = _workerCard
			} else if reader.Worker != rcvPackage.WorkerNo {
				errorMsg := fmt.Sprintf("Duplicate reader[%d] error %s", reader.ID, "F1 Login")
				// e.Logger.Warning(errorMsg)
				reader.TextOut(0, 1, errorMsg)
				reader.Beep(3, 0)
				reader.LogOff()
				continue
			}
		}

		if rcvPackage.Data == reader.WorkerCard {
			// e.Logger.Debug("Reader[%d] LogOff:[%s-%s]", reader.ID, reader.Worker, reader.WorkerCard)
			reader.Worker = ""
			reader.WorkerCard = ""
			reader.Beep(1, 0)
			reader.LogOff()
		} else {
			reader.Beep(3, 0)
			reader.TextOut(0, 1, "Invalid Card")
			// e.Logger.Debug("Reader[%d] SCANCARD:[%s]", reader.ID, rcvPackage.Data)
		}
	}
}
