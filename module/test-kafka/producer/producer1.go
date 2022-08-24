package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)
func main() {
	//ch := make(chan int64, 5)
	//go func() {
	//	for v := range ch {
	//		fmt.Println("----v-----", v)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		ch <- time.Now().UnixNano()
	//		time.Sleep(500 * time.Millisecond)
	//		ch <- time.Now().UnixNano()
	//		time.Sleep(3 * time.Second)
	//	}
	//}()
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	syncProducer, err := sarama.NewSyncProducer([]string{"172.16.68.136:9092", "172.16.68.136:9093"}, conf)
	if err != nil {
		panic(err)
	}

	//msg := sarama.ProducerMessage{
	//	Topic:     "test2",
	//	Key:       sarama.StringEncoder(time.Now().String()),
	//	Value:     sarama.StringEncoder("test message:" + time.Now().String()),
	//	Headers:   nil,
	//	Metadata:  nil,
	//	Offset:    0,
	//	Partition: 0,
	//	Timestamp: time.Now(),
	//}

	var wg sync.WaitGroup
	for i := 0;i < 10; i ++ {
		wg.Add(1)
		//time.Sleep(10 * time.Millisecond)
		go func(i int) {
			msg := sarama.ProducerMessage{
				Topic:     "test2",
				Key:       sarama.StringEncoder(time.Now().String()),
				Value:     sarama.StringEncoder("test message:" + time.Now().String()),
				Headers:   nil,
				Metadata:  nil,
				Offset:    0,
				Partition: 0,
				Timestamp: time.Now(),
			}
			msg.Timestamp = time.Now().In(time.Local)
			msg.Key = sarama.StringEncoder(fmt.Sprintf("key:%d", time.Now().UnixNano()))
			part, offset, err := syncProducer.SendMessage(&msg)
			if err != nil {
				panic(err)
			}
			fmt.Println(fmt.Sprintf("part:%d, %d, %+v", part, offset, msg))
			wg.Done()
		}(i)
	}
	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("after wait")
}
