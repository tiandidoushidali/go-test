package proto

import (
	"bytes"
	"fmt"
	"net"
)

type Connection struct {
	conn net.Conn

	OnConnected func(connection *Connection)
}

func NewConnection(conn net.Conn) *Connection {
	cnet := &Connection{
		conn: conn,
	}
	return cnet
}

func (conn *Connection) Send(code Code, data []byte) {
	pkg := pool.Get().(*Packet)
	pkg.Type = code
	pkg.Payload = data

	defer func() {
		pool.Put(pkg)
	}()

	msgBytes, err := pkg.Pack()
	if err != nil {
		fmt.Println("msg pack err ", err)
		return
	}
	fmt.Println("------pkg------", msgBytes)
	buf := bytes.NewBuffer(msgBytes)
	if _ , err := buf.WriteTo(conn.conn); err != nil {
		fmt.Println("buf writeto w err ", err)
		return
	}
}

// 链接后的处理
func onConnect(conn *Connection) {


}