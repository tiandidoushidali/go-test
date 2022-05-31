package mysocket

import (
	"fmt"
	"net"
	"strconv"
	"syscall"
)

/**
@https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651451353&idx=2&sn=6f167344bc84fa9d5772d62c02e356dd&chksm=80bb312bb7ccb83d9dc1cdbaabc80c2e1da3e9d05a62a4f246e554931d5beca4f30211f5443a&scene=21#wechat_redirect
*/
type MySocket struct {
	FileDescriptor int
}

// 从socket 中读取信息
func (socket MySocket) Read(bytes []byte) (int, error) {

	if len(bytes) == 0 {
		return 0, nil
	}
	numBytesRead, err := syscall.Read(socket.FileDescriptor, bytes)
	if err != nil {
		numBytesRead = 0
	}
	return numBytesRead, err
}

// 向socket 中写入信息
func (socket MySocket) Write(bytes []byte) (int, error) {
	numBytesWritten, err := syscall.Write(socket.FileDescriptor, bytes)
	if err != nil {
		numBytesWritten = 0
	}
	return numBytesWritten, err
}

// Close 关闭socket
func (socket *MySocket) Close() error {
	return syscall.Close(socket.FileDescriptor)
}

func (socket *MySocket) String() string {
	return strconv.Itoa(socket.FileDescriptor)
}

// Listen 监听socket
func Listen(ip string, port int) (*MySocket, error) {

	socket := &MySocket{}

	socketFileDescriptor, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, fmt.Errorf("faild to create socket (%+v)", err)
	}
	socket.FileDescriptor = socketFileDescriptor

	socketAddress := &syscall.SockaddrInet4{
		Port: port,
	}
	copy(socketAddress.Addr[:], net.ParseIP(ip))

	// 将socket 地址bind 文件描述符
	if err = syscall.Bind(socket.FileDescriptor, socketAddress); err != nil {

		return nil, fmt.Errorf("failed to bind socket (%+v)", err)
	}
	if err = syscall.Listen(socket.FileDescriptor, syscall.SOMAXCONN); err != nil {
		return nil, fmt.Errorf("failed to listen on socket (%+v)", err)
	}
	return socket, nil
}

// 定义事件循环
type EventLoop struct {
	KqueueFileDescriptor int
	SocketFileDescriptor int
}

func NewEventLoop(s *MySocket) (*EventLoop, error) {
	kqueue, err := syscall.Kqueue()
	if err != nil {
		return nil, fmt.Errorf("failed to create kqueue file descriptor (%+v)", err)
	}
	changeEvent := syscall.Kevent_t{
		Ident:  uint64(s.FileDescriptor),
		Filter: syscall.EVFILT_READ,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE,
		Fflags: 0,
		Data:   0,
		Udata:  nil,
	}

	changeEventRegistered, err := syscall.Kevent(kqueue,
		[]syscall.Kevent_t{changeEvent},
		nil, nil)

	if err != nil || changeEventRegistered == -1 {
		return nil, fmt.Errorf("failed to register change event (%v)", err)
	}

	return &EventLoop{
		KqueueFileDescriptor: kqueue,
		SocketFileDescriptor: s.FileDescriptor,
	}, nil
}

func (eventLoop *EventLoop) Handle(handler Handler) {

}
