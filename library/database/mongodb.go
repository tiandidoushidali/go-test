package database

import (
	"context"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Mongodb struct {
	Client *mongo.Client

	database string
	collection string
	options *options.FindOptions
	oneOption *options.FindOneOptions
}

func NewMongodbClient() (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	//client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	//client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://meddev:1UF0tcZ2V9eEzxCb@47.98.48.43:27017/admin"))
	//client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://medv3:hahhaloiwewe11@10.80.179.145:47027,10.80.179.145:47027"))
	//client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://imuser:dP9sYjtgmnbvhdp6k@dds-bp17d2991f9f02a41703.mongodb.rds.aliyuncs.com:3717"))
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://imuser:dP9sYjtgmnbvhdp6k@10.10.106.68:27017"))
	if err != nil {
		log.Errorf("[library|database|mongodb] NewMongodbClient mongo.Connect error:%+v", err)

	}

	return
}

func (m *Mongodb) DbList(ctx context.Context) {
	rets, err := m.Client.ListDatabases(ctx, bson.D{})
	if err != nil {
		log.Errorf("[library|database|mongodb] DbList m.Client.ListDatabases error:%+v", err)
	}
	log.Infof("[library|database|mongodb] DbList m.Client.ListDatabases return:%d, %+v", rets.TotalSize, rets.Databases)
}

func (m *Mongodb) DbListNames(ctx context.Context) {
	rets, err := m.Client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		log.Errorf("[library|database|mongodb] DbList m.Client.ListDatabaseNames error:%+v", err)
	}
	log.Infof("[library|database|mongodb] DbList m.Client.ListDatabaseNames return:%+v", rets)
}

func (m *Mongodb) Database(database string) *Mongodb{
	m.database = database

	return m
}

func (m *Mongodb) Collection(collection string) *Mongodb {
	if m.options == nil {
		m.initOptions()
	}
	if m.oneOption == nil {
		m.initOneOptions()
	}
	m.collection = collection

	return m
}

// 初始化 options

func (m *Mongodb) initOptions() {
	m.options = &options.FindOptions{}
}

func (m *Mongodb) initOneOptions() {
	m.oneOption = &options.FindOneOptions{}
}

func (m *Mongodb) SkipOption(val int64) *Mongodb {
	m.options.Skip = &val
	return m
}

func (m *Mongodb) LimitOption(val int64) *Mongodb {
	m.options.Limit = &val
	return m
}

func (m *Mongodb) SortOption(val interface{}) *Mongodb {
	m.options.Sort = val	// 1：正序 -1：倒序
	return m
}

func (m *Mongodb) HitOption(hit interface{}) *Mongodb {
	m.options.Hint = hit
	return m
}

func (m *Mongodb) InsertOne(ctx context.Context, document interface{}) (resp *mongo.InsertOneResult, err error){
	resp, err = m.Client.Database(m.database).Collection(m.collection).InsertOne(ctx, document)
	if err != nil {
		log.Infof("InsertOne InsertOne fail", err)
	}
	return
}

func (m *Mongodb) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (ret *mongo.UpdateResult, err error) {
	opts := options.Update().SetUpsert(true)
	ret, err = m.Client.Database(m.database).Collection(m.collection).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Infof("UpdateOne UpdateOne fail", err)
	}
	return
}

func (m *Mongodb) UpdateMany(ctx context.Context, filter interface{}, update interface{}) (ret *mongo.UpdateResult, err error) {
	ret, err = m.Client.Database(m.database).Collection(m.collection).UpdateMany(ctx, filter, update)
	if err != nil {
		log.Errorf("[library|database|mongodb] UpdateMany m.Client.UpdateMany error:%+v", err)
	}

	return
}

func (m *Mongodb) Find(ctx context.Context, filter interface{}) (cursor *mongo.Cursor, err error) {
	cursor, err = m.Client.Database(m.database).Collection(m.collection).Find(ctx, filter, m.options)
	if err != nil {
		log.Errorf("[library|database|mongodb] FindOne m.Client.Find error:%+v", err)
		return
	}
	return
}

func (m *Mongodb) FindOne(ctx context.Context, filter interface{}) (ret bson.Raw, err error){
	singleResult := m.Client.Database(m.database).Collection(m.collection).FindOne(ctx, filter, m.oneOption)
	err = singleResult.Err()
	if err == mongo.ErrNoDocuments {
		log.Errorf("[library|database|mongodb] FindOne m.Client.FindOne 数据不存在 error:%+v", err)
		return
	}
	if err != nil {
		log.Errorf("[library|database|mongodb] FindOne m.Client.FindOne error:%+v", err)
		return
	}
	ret, err = singleResult.DecodeBytes()
	if err != nil {
		log.Errorf("[library|database|mongodb] FindOne singleResult.DecodeBytes() error:%+v", err)
		return
	}
	log.Infof("[library|database|mongodb] FindOne params(%+v) return(%+v)", filter, ret)
	return
}

func (m *Mongodb) Count(ctx context.Context, filter interface{}) int64 {
	count, err := m.Client.Database(m.database).Collection(m.collection).CountDocuments(ctx, filter, options.Count().SetMaxTime(5 * time.Second))
	if err != nil {
		log.Errorf("[library|database|mongodb] Count error:%+v", err)
		return 0
	}
	log.Infof("[library|database|mongodb] Count params(%+v) count:%d", filter, count)
	return count
}

func (m *Mongodb) CreateIndex(ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"foo",1}},
		Options: options.Index().SetName("index_foo"),
	}
	m.Client.Database(m.database).Collection(m.collection).Indexes().CreateOne(ctx, indexModel)
}

func (m *Mongodb) CreateIndexes(ctx context.Context, ) {
	indexModels := make([]mongo.IndexModel, 0)
	m.Client.Database(m.database).Collection(m.collection).Indexes().CreateMany(ctx, indexModels)
}

func (m *Mongodb) GetIndexes(ctx context.Context) {
	view := m.Client.Database(m.database).Collection(m.collection).Indexes()
	var batchSize int32 = 2
	cursor, err := view.List(ctx, &options.ListIndexesOptions{
		BatchSize: &batchSize,
		MaxTime:   nil,
	})
	if err != nil {
		log.Errorf("[library|database|mongodb] GetIndexes view.List error:%+v", err)
		return
	}
	var results []bson.D
	if err := cursor.All(ctx, &results); err != nil {
		log.Errorf("[library|database|mongodb] GetIndexes cursor.All error:%+v", err)
		return
	}

	for _, v := range results {
		log.Infof("[library|database|mongodb] GetIndexes results range:k->v", v)
	}
}

type conn struct {
	addr string
}

type Collection struct {
	db   *conn
	orgC *mongo.Collection
}


// 获取im msg_new connection
func (m *Mongodb) GetImMsgNewConnection() *Collection {
	//orgC := m.Database("im").Collection("msg_new")
	//return &Collection{
	//	orgC: orgC,
	//	db:   c.conn,
	//}
	return  nil
}