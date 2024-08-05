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
	Logger   *log.Logger
}

func (e *TEngine) Connect() error {
	e.Logger.Printf("engine connect %s-%d\n", e.Host, e.Port)
	// log.Printf("engine connect %s-%d\n", e.Host, e.Port)
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
	e.Logger.Printf("Send %s\n", data)
	_, err := e.conn.Write(data)
	return err
}

func (e *TEngine) Receive() ([]byte, error) {
	buffer := make([]byte, 1024)
	n, err := e.conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	e.Logger.Printf("Recv %s\n", buffer)
	return buffer[:n], nil
}

func (e *TEngine) SendAndRecv(bytes []byte) ([]byte, error) {

	err := e.Send(bytes)
	if err != nil {
		return nil, err
	}

	// 接收响应
	buffer, err := e.Receive()
	if err != nil {
		return nil, err
	}

	// 返回接收到的数据
	return buffer, nil
}
