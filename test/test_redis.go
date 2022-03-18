package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"time"
)

func main() {
	flag.Parse()
	paladin.Init()
	var (
		chatRc struct {
			RedisChatConfig *redis.Config
		}
	)

	fmt.Println("----")
	var e error

	if e = paladin.Get("redis.txt").UnmarshalTOML(&chatRc); e != nil {
		panic(e)
	}

	ctx := context.Background()
	pl := redis.NewPool(chatRc.RedisChatConfig, redis.DialDatabase(9))
	conn := pl.Get(ctx)
	var scanIndex int64 = 0

	var results = make([]string, 0)
	for {
		rel, e := redis.Values(conn.Do("SCAN", scanIndex, "match", "global:user:notalk:set:user:*", "count", 10000))
		//rel, e := conn.Do("GET", "global:user:notalk:set:user:93114183")
		if e != nil {
			fmt.Println("-----e-----", e)
			return
		}
		count, e := redis.Int64(rel[0], e)
		values, e := redis.Strings(rel[1], e)
		scanIndex = count
		fmt.Println("------len-----", scanIndex, values)
		if len(values) > 0 {
			results = append(results, values...)
		}
		if scanIndex == 0 {
			break
		}
		time.Sleep(10 * time.Microsecond)
	}

	fmt.Println("-----results------", len(results), results)

}
