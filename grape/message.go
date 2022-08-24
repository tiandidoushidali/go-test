package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-test/grape/protocol/message"
	"go-test/library/database"
	"go-test/library/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"regexp"
	"strconv"
	"time"
)

// 获取历史消息
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

	//getHistoryMessage(ctx, mongodb, 8260672, 93940229, 1, 0, 100)
	getHistoryMessage(ctx, mongodb, 13917434, 0, 1, 0, 100)
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

//func getUserChatList(ctx context.Context, mongodb *database.Mongodb, userId int64, sort int64, start, limit int64) {
//	chatRelKey := fmt.Sprintf("c:%d", userId)
//	opt := new(options.FindOptions)
//	//->orderBy('msgId', 'desc')
//	sort := bson.D{{"msgId", -1}}
//	// ->get(['msgId', 'mapId']);
//	projection := bson.D{
//		{"_id", 0},
//		{"msgId", 1},
//		{"mapId", 1},
//	}
//	// ->where('key', $chatRelKey)  ->limit(200)
//	opt.SetProjection(projection).SetSort(sort).SetLimit(200)
//	cur, err := mongodb.Database("im").Collection("msg_new").
//		LimitOption(10).Find(ctx, )
//	cursor, err := dao.GetChatListCollection(ctx).Find(ctx, bson.M{"key": chatRelKey}, opt)
//	filter := bson.M{"mapId": fmt.Sprintf("h:%d",groupId)}
//	cur, err := mongodb.Database("im").Collection("msg_new").
//		//SortOption(bson.D{{"_id", -1}}).SkipOption(start).
//		LimitOption(limit + 1).
//		Find(ctx, filter)
//	if err != nil {
//		fmt.Println("getHistoryMessage err :", err)
//		return
//	}
//}
//
func parseMessage(data string) {
	decoded, _ := base64.StdEncoding.DecodeString(data)
	msg := &message.Message{}
	proto.Unmarshal(decoded, msg)

	typeMap := map[int32]string{
		1: "医生用户",
		2: "机构用户",
		3: "手机注册用户",
		4: "游客(仅redis)",
		5: "微信",
		6: "QQ",
		11: "营销平台患者",
		33: "第三方账号",
	}
	typeName := typeMap[msg.GetFrom().GetType()]

	fmt.Printf("------msg:%s-------- \n", msg.GetIdStr())
	fmt.Printf("userId:%d, name:%s, type:(%d, %s), reference:%d, id:%d group:%d date:%s content:%s data:%+v \n",
		msg.From.Id, msg.GetFrom().GetName(), msg.GetFrom().GetType(), typeName,
		msg.GetFrom().GetReference(), msg.GetId(), msg.GetGroup().Id, getMessageDate(int64(msg.GetId())), msg.GetText().GetContent(),
		msg.GetData())

	m, _ := json.Marshal(msg)
	if msg.Id == 1532202218919300963 {
		fmt.Println("----msg time----", utils.Parse(1532202218919300963), time.Unix(utils.Parse(1532202218919300963)/1000, 0).Format(utils.FORMAT_DAY_HIS))
		fmt.Println("msg:", string(m))
	}
}

func getMessageDate(id int64) string {
	unix := utils.Parse(id)
	//unixS := (unix - unix % 1000) / 1000
	//fmt.Println("---unix----", unixS, time.Unix(unixS, 0).Format(utils.FORMAT_DAY_HIS))

	ret := time.Unix(unix / 1000, unix % 1000).Format(utils.FORMAT_DAY_HIS)
	//fmt.Println("---ret---", ret)
	return ret
}

//func getMessageDate(id int64) string {
//	unix := utils.Parse(id)
//	//unixS := (unix - unix % 1000) / 1000
//	//fmt.Println("---unix----", unixS, time.Unix(unixS, 0).Format(utils.FORMAT_DAY_HIS))
//
//	ret := time.Unix((unix - unix % 1000) / 1000, unix - unix % 1000).Format(utils.FORMAT_DAY_HIS)
//	//fmt.Println("---ret---", ret)
//	return ret
//}

// 转换8进制utf-8字符串到中文
// eg: `\346\200\241` -> 怡
func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	fmt.Println("====", in, s)
	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {

			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			fmt.Println("--==--", s, b, i, string(b[1:]))
			return []byte{byte(i)}
		})
	return string(out)
}