package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func main() {
	fmt.Println("----", 1<<0x0F)
	// 11111111 11111111
	// 00000000 11111111
	// 00000001 00000000
	by := bytes.NewBuffer([]byte{byte(255), byte(255)}).Bytes()
	fmt.Println("-----src---", by)
	binary.BigEndian.PutUint16(by, 514)
	fmt.Println("===", by)


	var typ byte = byte(1)
	var content []byte = []byte("张三你好a马\n128")
	var length = bytes.NewBuffer([]byte{})

	fmt.Println("-----c---", content)
	binary.Write(length, binary.BigEndian, uint16(len(content)))


	bys := bytes.NewBuffer([]byte{}).Bytes()
	bys = append(bys, typ)
	bys = append(bys, length.Bytes()...)
	bys = append(bys, content...)

	r := make([]byte, 10)
	n, _ := io.ReadFull(bytes.NewBuffer(bys), r)
	fmt.Println("======", n, r)


	fmt.Println("----bys----", bys)
	sbys := bytes.NewBuffer([]byte{})
	err := binary.Write(sbys, binary.BigEndian, bys)
	if err != nil {
		fmt.Println("----err----", err)
	}

	fmt.Println("----sbys----", sbys.Bytes())

	var typr uint8
	var lengthr uint16
	binary.Read(sbys, binary.BigEndian, &typr)
	binary.Read(sbys, binary.BigEndian, &lengthr)
	var contentr []byte = make([]byte, lengthr)
	binary.Read(sbys, binary.BigEndian, &contentr)

	buf := bytes.NewBuffer(contentr)
	fmt.Println("--", typr, "--", lengthr, "--", buf.String())

}
