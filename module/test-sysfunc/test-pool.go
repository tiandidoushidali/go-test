package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"sync"
	"time"
	"unicode/utf8"
	"unsafe"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024*1024)
		return &b
	},
}


var commonPool = sync.Pool{New: func() interface{} {
	return "1"
}}
func main() {
	fmt.Println("----", commonPool.Get())
	t := fmt.Sprintf("%s 23:59:59", time.Now().Format("2006-01-02"))
	ts, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	fmt.Println("---", ts, time.Now())
	bytes := make([]byte, 1024)
	s := []byte("中国人123")
	str := "中国人123"
	fmt.Println("-----", unsafe.Sizeof(bytes), unsafe.Sizeof(s), len([]rune(str)), utf8.RuneCountInString(str))
	return
	//type poolSt struct {
	//	Id int64
	//	Name string
	//}
	//pool := sync.Pool{New: func() interface{} {
	//	return new(poolSt)
	//}}
	//
	//for i := 0; i < 10; i ++ {
	//	st := pool.Get().(*poolSt)
	//	st.Id = time.Now().UnixNano()
	//	st.Name = fmt.Sprintf("name-%d-name", st.Id)
	//	pool.Put(st)
	//
	//	runtime.Gosched()
	//	fmt.Println("---", st.Id, st.Name, &st)
	//}

	//var pool = sync.Pool{
	//	New: func() interface{} {
	//		return "123"
	//	},
	//}
	//
	//t := pool.Get().(string)
	//fmt.Println(t)
	//
	//pool.Put("3211")
	//pool.Put("3212")
	//pool.Put("3213")
	//pool.Put("3214")
	//
	//runtime.GC()
	//time.Sleep(1 * time.Second)
	//
	//t2 := pool.Get().(string)
	//fmt.Println(t2)
	//
	//runtime.GC()
	//time.Sleep(1 * time.Second)
	//
	//t2 = pool.Get().(string)
	//fmt.Println(t2)

	runtime.GOMAXPROCS(4)
	// 不使用对象池
	eg, _ :=  errgroup.WithContext(context.Background())
	a := time.Now().UnixNano()
	count := 1000
	for i := 0; i < 1000; i++ {
		for j := 0; j < count; j++ {
			func(i, j int) {
				eg.Go(func() error {
					obj := make([]byte, 1024*1024)
					_ = obj
					return nil
				})
			}(i, j)
		}
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("----", err)
	}
	b := time.Now().UnixNano()

	// 使用对象池
	//eg, _ =  errgroup.WithContext(context.Background())
	c := time.Now().UnixNano()
	//for i := 0; i < 1000; i++ {
	//	for j := 0; j < count; j++ {
	//		func(i, j int) {
	//			eg.Go(func() error {
	//				obj := bytePool.Get().(*[]byte)
	//				_ = obj
	//				bytePool.Put(obj)
	//				return nil
	//			})
	//		}(i, j)
	//	}
	//}
	//if err := eg.Wait(); err != nil {
	//	fmt.Println("----", err)
	//}
	d := time.Now().UnixNano()
	fmt.Println("without pool ", b-a, "ns")
	fmt.Println("with    pool ", d-c, "ns")
}
