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
data := "CICAgJid07sOpFRJVCNHBpSUQACJMaHR0cDovL2F2YXRhci1maWxlLnFhLm1lZGxpbmtlci5jb20vOTdmYTk2ZjMtZTk0MC0xMWViLWFkN2UtYWVlYWUxOWRmZDMxLmpwZyIFCNHGwwEqBQoD5pG4ogETMTUzNjU4ODg4NTU4MzIwMDI1Ng"
data = "CL+ZgMaYz5DUFRJJCOmr2kYaCeaigeaZqOaZqCI1aHR0cHM6Ly9wdWItbWVkLWF2YXRhci5pbWdzLm1lZGxpbmtlci5uZXQvbWFsZW5ldy5wbmcoISKbAgjy4tMGGgnmooHmmajmmagiNWh0dHBzOi8vcHViLW1lZC1hdmF0YXIuaW1ncy5tZWRsaW5rZXIubmV0L21hbGVuZXcucG5nIklodHRwczovL3B1Yi1tZWQtYXZhdGFyLm1lZGxpbmtlci5jb20vNDY5OTdmYzMtNGIzMi00MzliLWIyOTAtYjhlMGE3MmYwNDljIkNodHRwczovL3B1Yi1tZWQtYXZhdGFyLm1lZGxpbmtlci5jb20vYXNzaXN0YW50L2RvY3Rvci1hc3Npc3RhbnQucG5nKAMwxaj7FzgEQKGNBko1aHR0cHM6Ly9wdWItbWVkLWF2YXRhci5pbWdzLm1lZGxpbmtlci5uZXQvbWFsZW5ldy5wbmeiARMxNTYwNTcwMzYwNjM2NjQwNDQ3Kh0KG+mjjuWNjue7neS7o+S6n+W+heino+WGs+WQgw=="
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

	hd := hashids.NewData()
	hd.Alphabet = hashids.DefaultAlphabet
	hd.Salt = "<)Toyota~86(>"
	tk, _ := hashids.NewWithData(hd).EncodeInt64([]int64{11111111, 2222222, 3333333, 4444444})
	ids, _ := hashids.NewWithData(hd).DecodeInt64WithError(tk)
	fmt.Println("tk:", tk, ids)

	idds, _ := hashids.NewWithData(hd).DecodeInt64WithError("JoJaoxf9")
	fmt.Println("----idds----", idds)
	return
	hr := hashring.New([]string{"comet-0-1", "comet-0-2", "comet-1-1", "comet-1-2"})
	hr.AddNode("comet-1-1")
	hr.AddNode("comet-1-2")

	for i := 0; i < 100; i ++ {

		node, _ := hr.GetNode(fmt.Sprintf("%d", i))
		node2, _ := hr.GetNode(fmt.Sprintf("%d", 100 - i))
		fmt.Println(fmt.Sprintf("%d:%s:%d:%s", i, node, 100 - i, node2))
	}

}
