package receiver

import (
	"encoding/binary"
	"time"
)

type TSxPackage struct {
	DstAddr byte
	SrcAddr byte
	Length  uint
	Command byte
	Data    []byte
}

type TSxPackageData struct {
	Time     time.Time
	Serial   string
	WorkerNo string
	Data     string
}

const (
	commmand_length = 5
	SYNC_BYTE       = 0x54
)

func SxExamBuffer(buf []byte) int {
	return 6
}

// SxPackageToBytes 将 TSxPackage 转换为字节切片
func SxPackageToBytes(pkg TSxPackage) ([]byte, error) {
	var buf []byte
	data_length := len(pkg.Data)
	comlete_length := data_length + commmand_length
	buf = make([]byte, comlete_length)
	buf[0] = SYNC_BYTE
	buf[1] = pkg.DstAddr
	buf[2] = pkg.SrcAddr
	buf[3] = byte(data_length)
	buf[4] = pkg.Command
	copy(buf[5:], pkg.Data)

	// crc := crc16.Calculate(buf[:comlete_length], crc16.X25)
	crc := uint16(12244)
	crcBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(crcBytes, crc)
	copy(buf[comlete_length:comlete_length+2], crcBytes)

	return buf, nil
}

// BytesToSxPackage 从字节切片中解析出 TSxPackage
func BytesToSxPackage(bytes []byte) TSxPackage {
	return TSxPackage{
		DstAddr: bytes[0],
		SrcAddr: bytes[1],
		// Command: int(bytes[2]) | int(bytes[3])<<8 | int(bytes[4])<<16 | int(bytes[5])<<24,
	}
}

func BytesToSxPackageData(bytes []byte) (TSxPackageData, error) {
	return TSxPackageData{
		Time:     time.Now(),
		Serial:   "",
		WorkerNo: "",
		// Command: int(bytes[2]) | int(bytes[3])<<8 | int(bytes[4])<<16 | int(bytes[5])<<24,
	}, nil
}
