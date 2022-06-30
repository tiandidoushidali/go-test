package mongodb

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/go-kratos/kratos/pkg/net/trace"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	_family          = "mongo_client"
	_slowLogDuration = time.Millisecond * 250
)

type conn struct {
	addr string
}

type Collection struct {
	db   *conn
	orgC *mongo.Collection
}

func (coll *Collection) Name(c context.Context) string {
	return coll.orgC.Name()
}
func (coll *Collection) Database(c context.Context) *mongo.Database {
	return coll.orgC.Database()
}

func (coll *Collection) BulkWrite(c context.Context, models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions) (ret *mongo.BulkWriteResult, err error) {
	cb := TraceHook(c, coll.db, "BulkWrite")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("BulkWrite"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:BulkWrite")
	}()
	ret, err = coll.orgC.BulkWrite(c, models, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection:BulkWrite", "BulkWrite")
	}
	return
}

func (coll *Collection) InsertOne(c context.Context, document interface{},
	opts ...*options.InsertOneOptions) (ret *mongo.InsertOneResult, err error) {
	cb := TraceHook(c, coll.db, "InsertOne")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("InsertOne"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:InsertOne")
	}()
	ret, err = coll.orgC.InsertOne(c, document, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.InsertMany", "InsertOne")
	}
	return
}

func (coll *Collection) InsertMany(c context.Context, documents []interface{},
	opts ...*options.InsertManyOptions) (ret *mongo.InsertManyResult, err error) {
	cb := TraceHook(c, coll.db, "InsertMany")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("InsertMany"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:InsertMany")
	}()
	ret, err = coll.orgC.InsertMany(c, documents, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "ping", "ping")
	}
	return
}

func (coll *Collection) DeleteOne(c context.Context, filter interface{},
	opts ...*options.DeleteOptions) (ret *mongo.DeleteResult, err error) {
	cb := TraceHook(c, coll.db, "DeleteOne")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("DeleteOne"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:DeleteOne")
	}()
	ret, err = coll.orgC.DeleteOne(c, filter, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.DeleteOne", "DeleteOne")
	}
	return
}

func (coll *Collection) DeleteMany(c context.Context, filter interface{},
	opts ...*options.DeleteOptions) (ret *mongo.DeleteResult, err error) {
	cb := TraceHook(c, coll.db, "DeleteMany")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("DeleteMany"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:DeleteMany")
	}()
	ret, err = coll.orgC.DeleteMany(c, filter, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.DeleteMany", "DeleteMany")
	}
	return
}

func (coll *Collection) UpdateOne(c context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (ret *mongo.UpdateResult, err error) {
	cb := TraceHook(c, coll.db, "UpdateOne")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("UpdateOne"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:UpdateOne")
	}()
	ret, err = coll.orgC.UpdateOne(c, filter, update, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.UpdateOne", "UpdateOne")
	}
	return
}

func (coll *Collection) UpdateMany(c context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (ret *mongo.UpdateResult, err error) {
	cb := TraceHook(c, coll.db, "UpdateMany")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("UpdateMany"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:UpdateMany")
	}()
	ret, err = coll.orgC.UpdateMany(c, filter, update, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.UpdateMany", "UpdateMany")
	}
	return
}

func (coll *Collection) ReplaceOne(c context.Context, filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (ret *mongo.UpdateResult, err error) {
	cb := TraceHook(c, coll.db, "ReplaceOne")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("ReplaceOne"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:ReplaceOne")
	}()
	ret, err = coll.orgC.ReplaceOne(c, filter, replacement, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.ReplaceOne", "ReplaceOne")
	}
	return
}

func (coll *Collection) Aggregate(c context.Context, pipeline interface{},
	opts ...*options.AggregateOptions) (ret *Cursor, err error) {
	cb := TraceHook(c, coll.db, "Aggregate")
	if cb != nil {
		defer func() {
			cb(err)
		}()
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("Aggregate"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:Aggregate")
	}()
	r, err := coll.orgC.Aggregate(c, pipeline, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.Aggregate", "Aggregate")
	}
	ret = &Cursor{orgR: r}
	return
}

func (coll *Collection) CountDocuments(c context.Context, filter interface{},
	opts ...*options.CountOptions) (ret int64, err error) {
	cb := TraceHook(c, coll.db, "CountDocuments")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("CountDocuments"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:CountDocuments")
	}()
	ret, err = coll.orgC.CountDocuments(c, filter, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.CountDocuments", "CountDocuments")
	}
	return
}

func (coll *Collection) EstimatedDocumentCount(c context.Context, opts ...*options.EstimatedDocumentCountOptions) (ret int64, err error) {
	cb := TraceHook(c, coll.db, "EstimatedDocumentCount")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("EstimatedDocumentCount"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:EstimatedDocumentCount")
	}()
	ret, err = coll.orgC.EstimatedDocumentCount(c, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection.EstimatedDocumentCount", "EstimatedDocumentCount")
	}
	return
}

func (coll *Collection) Distinct(c context.Context, fieldName string, filter interface{},
	opts ...*options.DistinctOptions) (ret []interface{}, err error) {
	cb := TraceHook(c, coll.db, "Distinct")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("Distinct"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:Distinct")
	}()
	ret, err = coll.orgC.Distinct(c, fieldName, filter, opts...)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection:Distinct", "Distinct")
	}
	return
}

func (coll *Collection) FindOne(c context.Context, filter interface{}, opts ...*options.FindOneOptions) (ret *mongo.SingleResult) {
	cb := TraceHook(c, coll.db, "FindOne")
	if cb != nil {
		defer func() {
			if ret != nil {
				cb(ret.Err())
			} else {
				cb(nil)
			}
		}()
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("FindOne"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:FindOne")
	}()
	ret = coll.orgC.FindOne(c, filter, opts...)
	if ret != nil && ret.Err() != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection:FindOne", "FindOne")
	}
	return
}

func (coll *Collection) FindOneAndDelete(c context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) (ret *mongo.SingleResult) {
	cb := TraceHook(c, coll.db, "FindOneAndDelete")
	if cb != nil {

	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("FindOne"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:FindOneAndDelete")
	}()
	ret = coll.orgC.FindOneAndDelete(c, filter, opts...)
	if ret != nil && ret.Err() != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection:FindOneAndDelete", "FindOneAndDelete")
	}
	return
}

func (coll *Collection) FIndOneAndReplace(c context.Context) {

}

func (coll *Collection) FindOneAndUpdate(c context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) (ret *mongo.SingleResult) {
	cb := TraceHook(c, coll.db, "FindOneAndUpdate")
	if cb != nil {
		defer func() {
			if ret != nil {
				cb(ret.Err())
			} else {
				cb(nil)
			}
		}()
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("FindOneAndUpdate"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:Indexes")
	}()
	ret = coll.orgC.FindOneAndUpdate(c, filter, update, opts...)
	if ret.Err() != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection:FindOneAndUpdate", "FindOneAndUpdate")
	}
	return
}

func (coll *Collection) Watch(c context.Context) {

}

func (coll *Collection) Indexes(c context.Context) (ret mongo.IndexView) {
	cb := TraceHook(c, coll.db, "Indexes")
	if cb != nil {
		defer cb(nil)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("Indexes"), now)
	//耗时
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:Indexes")
	}()
	ret = coll.orgC.Indexes()

	return
}

func (coll *Collection) Drop(c context.Context) (err error) {
	cb := TraceHook(c, coll.db, "Drop")
	if cb != nil {
		defer cb(err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("Drop"), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:Drop")
	}()
	err = coll.orgC.Drop(c)
	if err != nil {
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection:Drop", "Drop")
	}
	return
}

//Find
func (coll *Collection) Find(c context.Context, filter interface{}, options ...*options.FindOptions) (ret *Cursor, err error) {
	if t, ok := trace.FromContext(c); ok {
		t = forkTrace(t, coll.db, "Collection:Find")
		t.SetTag(trace.String(trace.TagDBStatement, "Find"))
		defer t.Finish(&err)
	}
	now := time.Now()
	defer slowLog(fmt.Sprintf("Find  "), now)
	defer func() {
		_metricReqDur.Observe(int64(time.Since(now)/time.Millisecond), coll.db.addr, coll.db.addr, "Collection:Find")
	}()
	r, err := coll.orgC.Find(c, filter, options...)
	if err != nil {
		log.Errorc(c, "err:%v", err)
		_metricReqErr.Inc(coll.db.addr, coll.db.addr, "Collection:Find", "Find")
	}
	return &Cursor{orgR: r, db: coll.db}, err
}

func slowLog(statement string, now time.Time) {
	du := time.Since(now)
	if du > _slowLogDuration {
		log.Warn("%s slow log command: %s time:%v", _family, statement, du)
	}
}

func forkTrace(t trace.Trace, conn *conn, operationName string) trace.Trace {
	newSpan := t.Fork("", operationName)
	newSpan.SetTag(trace.TagString(trace.TagSpanKind, "client"))
	newSpan.SetTag(trace.TagString(trace.TagPeerAddress, conn.addr))
	newSpan.SetTag(trace.TagString(trace.TagDBType, "mongodb"))
	newSpan.SetTag(trace.TagString(trace.TagComment, "library/database/mongodb"))
	return newSpan
}
