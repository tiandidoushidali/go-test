package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

func main() {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	syncProducer, err := sarama.NewSyncProducer([]string{"0.0.0.0:9092", "0.0.0.0:9093"}, conf)
	if err != nil {
		panic(err)
	}

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

	var wg sync.WaitGroup
	for i := 0;i < 10; i ++ {
		wg.Add(1)
		go func(i int, msg sarama.ProducerMessage) {
			msg.Timestamp = time.Now()
			part, offset, err := syncProducer.SendMessage(&msg)
			if err != nil {
				fmt.Println("---err----", err)
				panic(err)
			}
			fmt.Println("part:", part, offset)
			wg.Done()
			fmt.Println("---j----", i)
		}(i, msg)
	}
	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("after wait")
}
