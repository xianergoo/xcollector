package receiver

import "bytes"

type TSxPackage struct {
	DstAddr byte
	SrcAddr byte
	Length  uint
	Command byte
	Data    []byte
}

func SxExamBuffer(buf []byte) int {
	return 6
}

// SxPackageToBytes 将 TSxPackage 转换为字节切片
func SxPackageToBytes(pkg TSxPackage) []byte {
	var buf bytes.Buffer
	buf.WriteByte(pkg.DstAddr)
	buf.WriteByte(pkg.SrcAddr)
	buf.WriteInt32(int32(pkg.Command))
	return buf.Bytes()
}

// BytesToSxPackage 从字节切片中解析出 TSxPackage
func BytesToSxPackage(bytes []byte) TSxPackage {
	return TSxPackage{
		DstAddr: bytes[0],
		SrcAddr: bytes[1],
		Command: int(bytes[2]) | int(bytes[3])<<8 | int(bytes[4])<<16 | int(bytes[5])<<24,
	}
}
