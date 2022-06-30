package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Cursor struct {
	orgR    *mongo.Cursor
	db      *conn
	Current bson.Raw
}

func (cursor *Cursor) Close(c context.Context) (err error) {
	cb := TraceHook(c, cursor.db, "Close")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog("Close", now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), cursor.db.addr, cursor.db.addr, "Close")
	}()
	err = cursor.orgR.Close(c)
	if err != nil {
		_metricReqErr.Inc(cursor.db.addr, cursor.db.addr, "Close", "Close")
	}
	return
}

func (cursor *Cursor) Next(c context.Context) bool {
	return cursor.orgR.Next(c)
}

func (cursor *Cursor) TryNext(c context.Context) bool {
	return cursor.orgR.TryNext(c)
}

func (cursor *Cursor) Decode(c context.Context, val interface{}) error {
	return cursor.orgR.Decode(val)
}

func (cursor *Cursor) Err() error {
	return cursor.orgR.Err()
}

func (cursor *Cursor) All(c context.Context, results interface{}) error {
	return cursor.orgR.All(c, results)
}

func (cursor *Cursor) RemainingBatchLength(c context.Context) int {
	return cursor.orgR.RemainingBatchLength()
}
