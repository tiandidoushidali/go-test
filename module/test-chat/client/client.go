package main

import (
	"bufio"
	"fmt"
	"go-test/module/test-chat/proto"
	"io"
	"net"
	"os"
)

//func main() {
//	conn, err := net.Dial("tcp", ":9999")
//	if err != nil {
//		panic("dail tcp fail")
//	}
//	// 从命令行读入
//	bufReader := bufio.NewReader(os.Stdin)
//
//	go func() {
//		for {
//			line, isPrefix, err := bufReader.ReadLine()
//			if err != nil {
//				fmt.Println("read line err", err)
//				continue
//			}
//			fmt.Println(fmt.Sprintf("read data isPrefix %t is %s", isPrefix, string(line)))
//			//buf := bytes.NewBuffer(line)
//			//buf.WriteTo(conn)
//
//			connect := proto.NewConnection(conn)
//			connect.Send(proto.MESSAGE, line)
//		}
//	}()
//
//	for {
//		bytes := make([]byte, 10)
//		if _, err := io.ReadFull(conn, bytes); err != nil {
//			fmt.Println("read server msg fail ", err)
//			continue
//		}
//		//if bytes, err = ioutil.ReadAll(conn); err != nil {
//		//	fmt.Println("read server msg fail ", err)
//		//	continue
//		//}
//		time.Sleep(1 * time.Second)
//		fmt.Println("receive server msg ", string(bytes), len(bytes), time.Now())
//	}
//}

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		panic("dail tcp fail")
	}
	// 从命令行读入
	bufReader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			line, _, err := bufReader.ReadLine()
			if err != nil {
				fmt.Println("read line err", err)
				continue
			}
			fmt.Println(fmt.Sprintf("send to server %s", string(line)))

			connect := proto.NewConnection(conn)
			connect.Send(proto.MESSAGE, line)
		}
	}()

	for {
		pkg := new(proto.Packet)
		err := pkg.Unpack(conn)
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Println("-------client pkg error disconnect ------")
				return
			}
			fmt.Println("receive server pkg err ", err)
			continue
		}

		fmt.Println(fmt.Sprintf("receive server pkg : %d, %d, %s", pkg.Type, pkg.GetLength(), string(pkg.Payload)))
	}
}