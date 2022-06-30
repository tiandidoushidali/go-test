package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/serialx/hashring"
	"github.com/speps/go-hashids"
	"go-test/grape/protocol/message"
	"path"
)

func main() {
	fmt.Println("path base :", path.Base("/aa/bb/cc/"))
	_ = `
CICAgJi07s
OpFRJVCNHB
pSUQACJMaH
R0cDovL2F2
YXRhci1maW
xlLnFhLm1l
ZGxpbmtlci
5jb20vOTdm
YTk2ZjMtZT
k0MC0xMWVi
LWFkN2UtYW
VlYWUxOWRm
ZDMxLmpwZy
IFCNHGwwEq
BQoD5pG4og
ETMTUzNjU4
ODg4NTU4Mz
IwMDI1Ng`
data := "CICAgJi07sOpFRJVCNHBpSUQACJMaHR0cDovL2F2YXRhci1maWxlLnFhLm1lZGxpbmtlci5jb20vOTdmYTk2ZjMtZTk0MC0xMWViLWFkN2UtYWVlYWUxOWRmZDMxLmpwZyIFCNHGwwEqBQoD5pG4ogETMTUzNjU4ODg4NTU4MzIwMDI1Ng"
data = "CICAgOqd2tGyFRJZCK7oriUaDeWMu+WKqS3popzojaMiQ2h0dHBzOi8vcHViLW1lZC1hdmF0YXIubWVkbGlua2VyLmNvbS9hc3Npc3RhbnQvZG9jdG9yLWFzc2lzdGFudC5wbmcaBAgBEAl6LwgAEAEaJDM3ZGJlOGMzLTZiNGMtNDQ0YS1iOWFhLTBhODk4NDc0ZjBmOSoDNC4zogETMTU0MTcxNjMxNDYxNDEzNjgzMg=="
//"CICAgJi07sOpFRJVCNHBpSUQACJMaHR0cDovL2F2YXRhci1maWxlLnFhLm1lZGxpbmtlci5jb20vOTdmYTk2ZjMtZTk0MC0xMWViLWFkN2UtYWVlYWUxOWRmZDMxLmpwZyIFCNHGwwEqBQoD5pG4ogETMTUzNjU4ODg4NTU4MzIwMDI1"
//"CICAgJi07sOpFRJVCNHBpSUQACJMaHR0cDovL2F2YXRhci1maWxlLnFhLm1lZGxpbmtlci5jb20vOTdmYTk2ZjMtZTk0MC0xMWViLWFkN2UtYWVlYWUxOWRmZDMxLmpwZyIFCNHGwwEqBQoD5pG4ogETMTUzNjU4ODg4NTU4MzIwMDI1Ng"
fmt.Println("--data--", len(data))

	decoded, err := base64.StdEncoding.DecodeString(data)
 	if err != nil {
 		fmt.Println("---err-----", err, len(data))
	}
	encoded := base64.StdEncoding.EncodeToString(decoded)
	fmt.Println("---src---", data)
	fmt.Println("---tar---", encoded)
	msg := &message.Message{}
	proto.Unmarshal(decoded, msg)

	m, _ := json.Marshal(msg)
	fmt.Println("msg:", string(m))

	return
	hd := hashids.NewData()
	hd.Alphabet = hashids.DefaultAlphabet
	hd.Salt = "<)Toyota~86(>"
	tk, _ := hashids.NewWithData(hd).EncodeInt64([]int64{11111111, 2222222, 3333333, 4444444})
	ids, _ := hashids.NewWithData(hd).DecodeInt64WithError(tk)
	fmt.Println("tk:", tk, ids)

	hr := hashring.New([]string{"comet-0-1", "comet-0-2", "comet-1-1", "comet-1-2"})
	hr.AddNode("comet-1-1")
	hr.AddNode("comet-1-2")

	for i := 0; i < 100; i ++ {

		node, _ := hr.GetNode(fmt.Sprintf("%d", i))
		node2, _ := hr.GetNode(fmt.Sprintf("%d", 100 - i))
		fmt.Println(fmt.Sprintf("%d:%s:%d:%s", i, node, 100 - i, node2))
	}

}
