package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	conf := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"0.0.0.0:9093"}, conf)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	partConsumer, err := consumer.ConsumePartition("test2", 0, 0)
	if err != nil {
		panic(err)
	}
	defer partConsumer.Close()

	for v := range partConsumer.Messages() {
		fmt.Println(fmt.Sprintf("---msg---k:%s, v:%s t:%v", string(v.Key), string(v.Value), v.Timestamp))
	}
}
