package tcpclient

import (
	"fmt"
	"net"
)

type ClientManager struct{}

func NewClientManager() *ClientManager {
	return &ClientManager{}
}

func (cm *ClientManager) Connect(host string, port int) (*net.TCPConn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	return conn.(*net.TCPConn), nil
}

func (cm *ClientManager) Send(conn *net.TCPConn, data []byte) error {
	_, err := conn.Write(data)
	return err
}

func (cm *ClientManager) Receive(conn *net.TCPConn) ([]byte, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}
