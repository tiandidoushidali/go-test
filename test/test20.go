package main

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/stat/sys/cpu"
	"github.com/jinzhu/copier"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Student20 struct {
	id int
	name string
	class *Class20
	classCh chan int
}

type Class20 struct {
	id int
	name string
	count int

	rw sync.RWMutex
}

func (c *Class20) add() {

	c.rw.Lock()
	defer c.rw.Unlock()
	c.count ++
}

func CalFeeUint32(src uint32, feeRate float32) uint32 {
	old := fmt.Sprintf("%0.0f", float32(src)*feeRate)
	dest, _ := strconv.ParseUint(old, 10, 32)
	return uint32(dest)
}

type Student11 struct {
	Id int64	`json:"id1"`
	Name string
}

type Student2 struct {
	Id2 int64	`json:"id1"`
	Name string
}



func main() {
	var ids []uint32 = make([]uint32, 0)
	for i := 0; i < 200; i ++ {
		ids = append(ids, uint32(i))
	}

	iids := ids

	var ids2 []uint

	copier.Copy(&ids2, &iids)
	fmt.Println("-----", ids2)

	return
	var ids3 = make([]uint, len(ids))
	t1 := time.Now().UnixNano()
	copier.Copy(&ids2, &ids)
	t2 := time.Now().UnixNano()
	fmt.Println("--------", t1, t2, t2-t1, len(ids2), time.Now().Unix())

	t1 = time.Now().UnixNano()
	var j int = 0
	for  ; j < len(ids); j ++ {
		ids3[j] = uint(ids[j])
	}
	t2 = time.Now().UnixNano()
	fmt.Println("--------", t1, t2, t2-t1, len(ids3))
	return
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("---er----", err)
		}

		panic("222")
	}()

	return
	stu1 := Student11{
		Id:   1,
		Name: "a",
	}
	stu2 := Student2{}

	copier.Copy(&stu2, &stu1)

	fmt.Println("-----", stu2, stu1)

	return

	var avaAmount int32 = 3000000
	var aaa int32 = 9991564
	fmt.Println(avaAmount - aaa)

	return
	err := fmt.Errorf("sku fetcher fetch query panic: %d -- %+v", 0)
	if err != nil {
		fmt.Println("---", err)
	}
	fmt.Println(err)

	return
	var aa string = "abc"
	bbbb := (*string)(unsafe.Pointer(&aa))

	cla201 := &Class20{
	}

	fmt.Println("----ccc0----", unsafe.Pointer(cla201))
	ccc := uintptr(unsafe.Pointer(cla201))
	runtime.GC()
	time.Sleep(1 * time.Second)
	fmt.Println("----ccc----", ccc)
	ccc = ccc + unsafe.Offsetof(cla201.id)
	fmt.Println("----ccc2----", ccc)
	dd := (*int)(unsafe.Pointer(ccc))
	*dd = 1111
	ccc = ccc + unsafe.Offsetof(cla201.name)
	fmt.Println("----ccc2----", ccc)
	ee := (*string)(unsafe.Pointer(ccc))
	*ee = "张三"
	fmt.Println(cla201)
	return

	fmt.Println("---", bbbb)
	fmt.Println(time.Second << 3)
	
	return
	var m int32 = 0
	bb := atomic.CompareAndSwapInt32(&m, 0, 1)

	fmt.Println("----m----", m, bb)


	return

	cla20 := &Class20{}
	for i := 0; i < 1000; i ++ {
		go func() {
			cla20.add()
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("total:", cla20.count)
	return
	a, b := 1, 2
	switch a {
	case 1:
		switch b {
		case 1:
			fmt.Println("b1")
			return
		default:
			fmt.Println("b2")
			return
		}
	}
	return
	stu := new(Student20)
	ch1 := make(chan int, 1)
	stu.classCh = ch1
	stu.classCh <- 10
	fmt.Println("----", <-ch1, <-stu.classCh)
	return
	fmt.Println(cpu.GetInfo())
	return
	num, totalCount := 0, 20
	ch := make(chan int, 10)

	ch<-5
	ch<-6
	close(ch)
	fmt.Println(<-ch, <-ch, <-ch)

	return

	go func(num int) {
		for {
			select {
			case v, ok := <-ch:
				fmt.Println("---v----", v, ok, num)
				time.Sleep(100 * time.Millisecond)
				//time.Sleep(100*time.Millisecond)
				if !ok {
					fmt.Println("ch close")
					return
				}
				num ++
				if num >= totalCount {
					fmt.Println("--close ch--")

				}
			}
		}
	}(num)

	for i := 0; i < totalCount; i ++ {
		fmt.Println("----i----", i)
		ch <- i
	}
	close(ch)
	fmt.Println("ch has close")

	time.Sleep(3 * time.Second)

}
