package main

import (
	"fmt"
	errgroup2 "golang.org/x/sync/errgroup"
	"runtime"
	"time"
)


var name string
type Student struct {
	Name string
}

func main() {
	s := &Student{}
	test3(s)
	time.Sleep(5 * time.Second)
	mem := runtime.MemProfileRecord{
	}
	fmt.Println("mem allocate bytes", mem.AllocBytes, mem.AllocObjects)
	fmt.Println("--=======--", runtime.NumGoroutine())
}

func test() {
	//g := sync.WaitGroup{}
	ch := make(chan struct{})
	defer close(ch)
	n := 100
	for i := 0; i < n; i ++ {
		//g.Add(1)

		test2(i)

		//go func(i int) {
		//	defer func() {
		//		g.Done()
		//		<-ch
		//	}()
		//	fmt.Println("----:", i, "---", runtime.NumGoroutine())
		//	return
		//}(i)
	}
	//g.Wait()
}

func test2(i int) {
	for i := 0; i < 10; i ++ {
		fmt.Println("i value ", i)
		if i == 8 {
			return
		}
	}
	fmt.Println("1111")
	errgroup := errgroup2.Group{}
	errgroup.Go(func() error {
		defer func() {
			fmt.Println("test2 goroutine 11 value ", i)
		}()
		fmt.Println("test2 goroutine value ", i)
		return nil
	})

	errgroup.Go(func() error {
		defer func() {
			fmt.Println("panic test2 goroutine tine2 value go ", i)
			if err := recover(); err != nil {

				fmt.Println("panic err test2 goroutine tine2 value go ", i, err)
			}
		}()
		panic("panic test2 goroutine tine2 value " + fmt.Sprintf("%d", i))
		fmt.Println("test2 goroutine tine2 value ", i)
		return nil
	})

	errgroup.Wait()
}
func test3(s *Student) {
	for n := 1000; n > 0; n -- {
		go func() {
			name = fmt.Sprintf("name:%d", n )
			s.Name = name
		}()
	}
}