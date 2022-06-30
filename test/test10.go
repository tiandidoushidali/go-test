package main

import (
	"fmt"
	"sync"
	"time"
)

type TTTT struct {
	Id int64
	sync.Mutex
}


func (t *TTTT) setId(id int64) {
	fmt.Println("------in-----", time.Now())
	t.Lock()
	defer t.Unlock()
	t.Id = id

	time.Sleep(1 * time.Second)

	fmt.Println("----time----", id, time.Now())
}
func main() {
	m := sync.Mutex{}
	rwm := sync.RWMutex{}
	sum := 0
	ch := make(chan struct{}, 4)
	for i := 0 ; i < 100; i ++ {
		go func(i int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("-----err-----", err)
				}
			}()
			ch <- struct{}{}
			m.Lock()
			defer m.Unlock()
			//rwm.Lock()
			//defer rwm.RUnlock()
			//fmt.Println("-----i-----", i)
			sum += i
		}(i)
	}
	time.Sleep(1 * time.Second)
	go func() {
		rwm.RLock()
		defer rwm.RUnlock()
		fmt.Println("==========")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("----sum-----", sum)
	return
	for {
		fmt.Println("------")
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("---time now ---", time.Now())
		}
	}
	return
	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("----delay 10 second----")
	})

	time.Sleep(5 * time.Second)

	var val = 10

	fmt.Printf("--change to %b %o %d %X", val, val, val, val)

	var val2 = 0x00A

	fmt.Printf("--change to %b %o %d %X", val2, val2, val2, val2)

	fmt.Println("----", fmt.Sprintf("%04X", val2 + 16))
	return
	t := new(TTTT)

	var i int64 = 0
	eg := &sync.WaitGroup{}
	for i < 10 {
		eg.Add(1)
		go func(i int64) {
			t.setId(i)
			eg.Done()
		}(i)

		i ++
	}

	eg.Wait()

	fmt.Println("-----", t.Id)
}
