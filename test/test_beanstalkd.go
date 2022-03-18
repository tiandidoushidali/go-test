package main

import (
	"fmt"
	"github.com/beanstalkd/go-beanstalk"
	"reflect"
	"strings"
	"time"
)

func main() {
	str := "abcdefg"
	fmt.Println(str[0:2])
	fmt.Println(str[0:2])
	fmt.Println(reflect.TypeOf(str))
	return
	t := 1
	switch t {
	case 1:fmt.Println("11")
	case 2:fmt.Println("22")
	}
	return
	tubeName := "med-im"
	c, err := beanstalk.Dial("tcp", "beanstalkd-1.beanstalkd.app.svc.cluster.local:11300")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true

	substr := "timeout"
	for {
		//从队列中取出
		id, body, err := c.Reserve(1 * time.Second)
		if err != nil {
			if !strings.Contains(err.Error(), substr) {
				fmt.Println(" [Consumer] [", c.Tube.Name, "] err:", err, " id:", id)
			}
			continue
		}
		fmt.Println(" [Consumer] [", c.Tube.Name, "] job:", id, " body:", string(body))

		//从队列中清掉
		err = c.Delete(id)
	}
}

