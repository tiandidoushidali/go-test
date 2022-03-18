package proto

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"
)

//  0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 ...
// +---------------+-------------------------------+---------------+...
// |    type(8)    |          length(16)           |     payload    ...

type Code uint8

const (
	CONNECT Code = 0x00 + iota
	CONNECTED
	RECONNECT
	PING
	PONG
	MESSAGE
	COMMAND
	CLOSE
)

type Packet struct {
	Type Code
	Payload []byte
	length uint16
}

var pool *sync.Pool = &sync.Pool{New: func() interface{} {
	return new(Packet)
}}

// 数据打包
func (p *Packet) Pack() ([]byte, error) {
	switch p.Type {
	case CONNECT, CONNECTED, RECONNECT, PING, PONG, MESSAGE, COMMAND, CLOSE:
		length := len(p.Payload)

		b := bytes.NewBuffer(p.Payload)
		// 消息限制为64k
		if length > 1 << 16 {
			str := fmt.Sprintf("message payload len(%d) overflow 64k", len(b.Bytes()))
			return nil, errors.New(str)
		}
		p.length = uint16(length)
		return p.ToBytes(), nil

	default:
		return nil, errors.New(fmt.Sprintf("unsupport protocol %d ", p.Type))
	}
	return nil, nil
}

// 拆解数据包
func (p *Packet) Unpack(r io.Reader) error {
	switch p.Type {
	case CONNECT, CONNECTED, RECONNECT, PING, PONG, MESSAGE, COMMAND:
		if err := binary.Read(r, binary.BigEndian, &p.Type); err != nil {
			return err
		}
		var length uint16
		if err := binary.Read(r, binary.BigEndian, &length); err != nil {
			return err
		}
		p.length = length
		// 非常重要，读取指定字节到满
		p.Payload = make([]byte, p.length)
		if err := binary.Read(r, binary.BigEndian, &p.Payload); err != nil {
			return err
		}

	default:
		return errors.New(fmt.Sprintf("unsupport protocol"))
	}
	return nil
}

func (p *Packet) GetLength() uint16 {
	return p.length
}

func (p *Packet) ToBytes() []byte {
	var typeBuf, lenBuf, payloadBuf bytes.Buffer
	binary.Write(&typeBuf, binary.BigEndian, p.Type)
	binary.Write(&lenBuf, binary.BigEndian, p.length)
	binary.Write(&payloadBuf, binary.BigEndian, p.Payload)

	msgBytes := append(typeBuf.Bytes(), append(lenBuf.Bytes(), payloadBuf.Bytes()...)...)
	return msgBytes
}