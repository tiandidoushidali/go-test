package main

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/serialx/hashring"
	"go-test/grape/protocol/message"
	"path"
	"github.com/speps/go-hashids"
)

func main() {
	fmt.Println("path base :", path.Base("/aa/bb/cc/"))
	data := "CICg7ZzC86G0FKIBEzE0NzA1NzQzODMwMTg0OTYwMDASYgiNqow+GhblgaXlurfnrqHnkIbluIgt5bCP6I6OIkFodHRwczovL3B1Yi1tZWQtYXZhdGFyLm1lZGxpbmtlci5jb20vYXNzaXN0YW50L2hlYWx0aC1tYW5hZ2VyLnBuZygDItYBCLCv9wMaCeeOi+agkeWIqSI7aHR0cHM6Ly9wdWItbWVkLWF2YXRhci5tZWRsaW5rZXIuY29tLzE1OTY3MDIxNTg5OHY0cGMxMC5qcGciQWh0dHBzOi8vcHViLW1lZC1hdmF0YXIubWVkbGlua2VyLmNvbS9hc3Npc3RhbnQvaGVhbHRoLW1hbmFnZXIucG5nKAIw35P4NTgHSjtodHRwczovL3B1Yi1tZWQtYXZhdGFyLm1lZGxpbmtlci5jb20vMTU5NjcwMjE1ODk4djRwYzEwLmpwZyq4Agq1Aua4qemmqOaPkOekuu+8jOeWq+aDhei/mOS4jeeos+Wumu+8jOimgeWBmuWlvemYsuaKpO+8jOWLpOa0l+aJi++8jOWLpOmAmumjju+8jOWkluWHuuaXtuW4puWlveWPo+e9qe+8jOWBmuWlveiHqui6q+mYsuaKpOOAggoK55ar5oOF5Yqg5bm05bC+5b+r6YCS5rS+6YCB5Y+v6IO95Lya5pyJ5bu26L+f77yM6K6w5b6X5o+Q5YmN5aSH6I2v6YG/5YWN6I2v5ZOB5LiN6Laz5a+86Ie05pat6I2v5ZOmfiAKCuWmgumcgOW8gOWkhOaWue+8jOeCueWHu+acrOadoea2iOaBr+iBlOezu+aCqOeahOWBpeW6t+euoeeQhuW4iOWNj+WKqeaCqOi0reiNr+OAgg=="
	decoded, _ := base64.StdEncoding.DecodeString(data)
	msg := &message.Message{}
	proto.Unmarshal(decoded, msg)
	fmt.Println("msg:", msg, msg.From.Hospital)

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
