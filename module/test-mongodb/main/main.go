package main

import (
	"context"
	"fmt"
	"go-test/library/database"
	"time"
)

func main() {
	client, err := database.NewMongodbClient()
	if err != nil {
		fmt.Println("-----err-----", err)
		return
	}
	mongodb := &database.Mongodb{
		Client: client,
	}
	ctx := context.Background()
	nowms := time.Now().UnixNano() / 1e6
	mongodb.DbList(ctx)
	fmt.Println("----nowms----", nowms)
}
