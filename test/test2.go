package main

import (
	"bytes"
	"fmt"
	"time"
)

type Event struct {
	PuserId int
	TodoId int
	OpType int
	Reason string
}

func main() {
	sql := "insert into `todo_op_log` (`puser_id`,`todo_id`, `op_type`, `reason`, `created_at`) values"
	var buffer bytes.Buffer

	var inserts = make([]*Event, 0)
	inserts = append(inserts, &Event{
		PuserId: 1,
		TodoId:  1,
		OpType:  1,
		Reason:  "水电费水电费都是1",
	})
	inserts = append(inserts, &Event{
		PuserId: 2,
		TodoId:  2,
		OpType:  2,
		Reason:  "水电费水电费都是2",
	})
	var err error
	if _, err = buffer.WriteString(sql); err != nil {
		return
	}
	for i, e := range inserts {
		if i == len(inserts)-1 {
			buffer.WriteString(fmt.Sprintf("(%d,%d,%d, '%s', '%s');", e.PuserId, e.TodoId, e.OpType, e.Reason, time.Now().Format("2006-01-02 15:04:05")))
		} else {
			buffer.WriteString(fmt.Sprintf("(%d,%d,%d, '%s', '%s'),", e.PuserId, e.TodoId, e.OpType, e.Reason, time.Now().Format("2006-01-02 15:04:05")))
		}
	}

	fmt.Println(buffer.String())
}
