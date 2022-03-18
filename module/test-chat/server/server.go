package main

import (
	"bufio"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go-test/module/test-chat/proto"
	"io"
	"net"
	"os"
)

var queue chan *proto.Packet

var hub map[interface{}]interface{}

func init() {
	queue = make(chan *proto.Packet, 128) // 限制channel 容量为128
}

func main() {
	// 监听端口
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9999")
	//if err != nil {
	//	panic(fmt.Sprintf("resolve tcp addr err %v", err))
	//}
	tcpLis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(fmt.Sprintf("tcp listen err %v", err))
	}

	for {
		conn, err := tcpLis.Accept()
		if err != nil {
			fmt.Println(fmt.Sprintf("tcp listen accept err %v", err))
			continue
		}

		fmt.Println("server receive connection")

		go handleConn(conn)

		// 监听输入
		go writeLoop(conn)

	}

	fmt.Println("----over------")

}


//func handleConn(conn net.Conn) {
//	for {
//		var buf []byte = make([]byte, 128)
//		n, err := io.ReadFull(conn, buf)
//		if err != nil {
//			fmt.Println("io readFull conn to buf err", err)
//			return
//		}
//
//		fmt.Println("receive msg from client ", string(buf))
//
//		buff := bytes.NewBuffer(append([]byte("hello world "), buf...))
//
//		fmt.Println("send msg to client", n, time.Now(), buff)
//
//		buff.WriteTo(conn)
//
//		//if n > 0 {
//		//	// 不同消息进行不同处理
//		//	switch string(buf) {
//		//	case " V1": // 建立链接
//		//		userWrapper(conn, buf)
//		//	default:
//		//
//		//	}
//		//}
//	}
//}
//func handleConn(conn net.Conn) {
//	for {
//		pkg := new(proto.Packet)
//		err := pkg.Unpack(conn)
//		if err != nil {
//			fmt.Println("server unpack packet err ", err)
//			continue
//		}
//
//		fmt.Println(fmt.Sprintf("receive client pkg : %d, %d, %s", pkg.Type, pkg.GetLength(), string(pkg.Payload)))
//
//		pkg.Payload = append(pkg.Payload, []byte("hello world ")...)
//
//		fmt.Println(fmt.Sprintf("send to client pkg : %d, %d, %s", pkg.Type, pkg.GetLength(), string(pkg.Payload)))
//
//		connect := proto.NewConnection(conn)
//		connect.Send(proto.MESSAGE, pkg.Payload)
//	}
//}

func handleConn(conn net.Conn) {
	var err error
	defer func() {
		if er := recover(); er != nil {
			fmt.Println("---关闭连接1----", er)
			conn.Close()
		}
		if err != nil {
			conn.Close()
		}
	}()
	for {
		pkg := new(proto.Packet)
		err = pkg.Unpack(conn)
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Println("-------server pkg error disconnect------")
				return
			}
			fmt.Println("server unpack packet err ", err)
			continue
		}

		fmt.Println(fmt.Sprintf("receive client pkg : %d, %d, %s", pkg.Type, pkg.GetLength(), string(pkg.Payload)))

		//pkg.Payload = append(pkg.Payload, []byte("hello world ")...)
		//
		//fmt.Println(fmt.Sprintf("send to client pkg : %d, %d, %s", pkg.Type, pkg.GetLength(), string(pkg.Payload)))
		//
		//connect := proto.NewConnection(conn)
		//connect.Send(proto.MESSAGE, pkg.Payload)
	}
}

func writeLoop(conn net.Conn) {

	connect := proto.NewConnection(conn)
	connect.OnConnected = func(connection *proto.Connection) {
		uuid := uuid.NewV1().String()
		hub[uuid] = connection
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("bufio readline error", err)
			continue
		}

		fmt.Println(fmt.Sprintf("send to client %s", string(line)))
		connect.Send(proto.MESSAGE, line)
	}
}

func userWrapper(conn net.Conn, msg []byte) *proto.Connection {
	return proto.NewConnection(conn)
}