package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-test/grape/protocol/message"
	"go-test/library/database"
	"go-test/library/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/golang/protobuf/proto"
	"strings"
	"time"
)

// 获取历史消息
func main() {
	getMessageDate()
	return
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

	getHistoryMessage(ctx, mongodb, 8260672, 93940229, 1, 0, 100)
}

// 获取群历史消息
/**
sort    int          0:升序 1:降序
 */
func getHistoryMessage(ctx context.Context, mongodb *database.Mongodb, groupId int64, userId int64, sort int64, start, limit int64) {

	opt := new(options.FindOptions)
	sortt := bson.D{{"_id", -1}}
	if sort == 0 {
		sortt = bson.D{{"_id", 1}}
	}

	opt.SetSort(sortt).SetSkip(int64(start)).SetLimit(int64(limit + 1))
	opt.SetHint(bson.M{"mapId": 1})

	//filter := bson.D{{"mapId", groupId}}
	//objId, err := primitive.ObjectIDFromHex("5adf374a56d85a0049475944")
	//fmt.Println("----objId----", objId)
	//filter := bson.D{{"_id", objId}}
	//filter := bson.D{{"_id", objId}}
	filter := bson.M{"mapId": fmt.Sprintf("h:%d",groupId)}
	cur, err := mongodb.Database("im").Collection("msg_new").
		//SortOption(bson.D{{"_id", -1}}).SkipOption(start).
		LimitOption(limit + 1).
		Find(ctx, filter)
	if err != nil {
		fmt.Println("getHistoryMessage err :", err)
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		r := bson.M{}
		err := cur.Decode(&r)
		if err != nil {
			fmt.Println("----cur.next errr----", err)
			continue
		}

		parseMessage(r["value"].(string))
		//fmt.Println("----r-----", r)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("-errr--", err)
	}
}

func parseMessage(data string) {
	if strings.Contains(data, "COOegLbNud6hFRKGAQixgJwfGgbnjovkuL0iRWh0dHBzOi8vcHViLW1lZC1hdmF0YXIubWVkbGlua2VyLmNvbS9jZWIzMTEwY") {
		fmt.Println("----match-----", data)
	}
	decoded, _ := base64.StdEncoding.DecodeString(data)
	msg := &message.Message{}
	proto.Unmarshal(decoded, msg)

	fmt.Println("--msg---", msg.Id)
	m, _ := json.Marshal(msg)
	if msg.Id == 1532202218919300963 {
		fmt.Println("----msg time----", utils.Parse(1532202218919300963), time.Unix(utils.Parse(1532202218919300963)/1000, 0).Format(utils.FORMAT_DAY_HIS))
		fmt.Println("msg:", string(m))
	}
}

func getMessageDate() string {
	unix := utils.Parse(1351065226280816640)
	fmt.Println("---unix----", (unix - unix % 1000) / 1000)

	ret := time.Unix((unix - unix % 1000) / 1000, unix - unix % 1000).Format(utils.FORMAT_DAY_HIS)
	fmt.Println("---ret---", ret)
	return ret
}