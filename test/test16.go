package main

import (
	"fmt"
	"sync"
	"time"
)

type Connection struct {
	queue chan int64
	stop chan bool
}

func (conn *Connection) Send(v int64) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("send send to queue fail ", err)
		}
	}()
	conn.queue <- v
}

func (conn *Connection) Close() {
	fmt.Println("conn close")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("close send to stop fail ", err)
		}
	}()
	//conn.stop <- true
	close(conn.stop)
}

func (conn *Connection) ReadLoop() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ReadLoop send to stop fail ", err)
		}
	}()
	defer conn.Close()

	time.Sleep(3 * time.Second)
}

func (conn *Connection) WriteLoop() {
	n := 0
	for {
		select {
		case <-conn.stop:
			fmt.Println("conn.stop prepare stop for")
			time.Sleep(1 * time.Second)
			return
		case v, ok := <-conn.queue:
			if !ok {
				fmt.Println("conn.queue has closed")
				return
			}
			n ++
			fmt.Println("conn.queue value ", v, ok)
		}
	}
}

func (conn *Connection) IoLoop() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("IoLoop panic ", err)
		}
		//close(conn.stop)
		//close(conn.queue)
	}()
	// 模拟出错
	go conn.ReadLoop()

	conn.WriteLoop()
}

type ty int64
const (
	DEVELOP ty = 1
	PRE ty = 2
	PRO ty = 3
)
func (ty ty) GetTy() string {
	switch ty {
	case 1: return "1"
	default: return "a"
	}
}

func (ty ty) GetId() int {
	switch ty {
	case 1: return 1
	default: return 2
	}
}

type Class1 struct {
	name string
	stu Student1
}

type Student1 struct {
	id int
}



func main() {
	cc := Class1{name:"1班", stu: Student1{id: 1}}
	cc.name = "2班"
	cc.stu.id = 2
	fmt.Println("-------", cc)

	return
	//time.Sleep(5 * time.Second)
	fmt.Println("----", cc)

	return
	var tt *Connection
	fmt.Println(tt.queue)
	return
	var t ty = 1

	var i int64 = 1
	if t != ty(i) {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(t.GetTy(), PRO.GetTy(), ty(4).GetId())
	return

	conn := &Connection{
		queue: make(chan int64, 10),
		stop:  make(chan bool, 1),
	}

	// 每500毫秒发送一条数据到queue
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			conn.Send(time.Now().Unix())
		}
	}()

	conn.IoLoop()

	time.Sleep(20 * time.Second)

	return
	stop := make(chan int)
	rec := make(chan int64)
	var once sync.Once
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("close stop fail ", err)
			}
		}()
		time.Sleep(2 * time.Second)
		stop <- 1
		close(stop)

	}()
	time.Sleep(1 * time.Second)
	go func() {
		fmt.Println("111111")
		for {
			select {
			case v, ok := <- stop:
				if ok && v == 1 {
					fmt.Println("---no close----", v, ok )
					return
				}
				fmt.Println("---close----", v, ok )
				time.Sleep(1 * time.Second)
			default:
				time.Sleep(1 * time.Second)
				rec <- time.Now().Unix()

			}
		}
		//close(rec)
	}()

	//fmt.Println("<- rec", <- rec)
	for {
		select {
		case v, ok := <-rec:
			fmt.Println("--rec---", v, ok)
			time.Sleep(1 * time.Second)
		case v, ok := <-stop:
			if ok && v == 1 {
				fmt.Println("---close for -----")
				return
			}
			if !ok {
				fmt.Println("has close goroutine")
				return
			}
			fmt.Println("--stop--", v, ok )
			time.Sleep(3 * time.Second)
			once.Do(func() {
				close(rec)
			})
			//close(rec)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now())
		}
	}

	return
	// channel 初始化
	c := make(chan int, 10)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}

	// sender（写端）
	go func() {
		// 入队
		c <- 1
		// ...
		// 满足某些情况，则 close channel
		close(c)
	}()

	// receivers （读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ... 处理 channel 里的数据
			for v := range c {
				_ = v
			}
		}()
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
}
