package receiver

import (
	"fmt"
	"log"
	"net"
)

type TEngine struct {
	conn     net.TCPConn
	Host     string
	Port     uint
	Location string
}

func (e *TEngine) Connect() error {
	log.Printf("engine connect %s-%d\n", e.Host, e.Port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", e.Host, e.Port))
	if err != nil {
		return err
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		// 如果断言失败，则关闭连接并返回错误
		conn.Close()
		return fmt.Errorf("connection is not a TCP connection")
	}

	// 将 tcpConn 赋值给 e.Conn
	e.conn = *tcpConn
	return nil

}

func (e *TEngine) Send(data []byte) error {
	_, err := e.conn.Write(data)
	return err
}

func (e *TEngine) Receive() ([]byte, error) {
	buffer := make([]byte, 1024)
	n, err := e.conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}

func (e *TEngine) SendAndRecv(bytes []byte) ([]byte, error) {
	// 这里只是一个示例，您可以替换为实际的发送和接收逻辑
	return bytes, nil
}
