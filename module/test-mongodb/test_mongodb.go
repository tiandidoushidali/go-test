package test_mongodb

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"go-test/library/database"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)
type TestMongodb struct {

}

func (t *TestMongodb)Handle() (err error) {

	client, err := database.NewMongodbClient()
	if err != nil {
		log.Errorf("[module|test-mongodb|test_mongodb] Handle database.NewMongodbClient():%+v", err)
		return
	}
	mongodb := &database.Mongodb{
		Client: client,
	}
	ctx := context.Background()
	nowms := time.Now().UnixNano() / 1e6
	mongodb.DbList(ctx)
	//mongodb.DbListNames(ctx)
	//mongodb.Database("test").Collection("body_info").InsertOne(ctx, bson.D{{"user_id", 111}, {"bmi", bson.A{bson.D{{"date", 111}, {"value", 197.04}}, bson.D{{"date", 112}, {"value", 198.04}}}}})
	//objId, err := primitive.ObjectIDFromHex("6152d075210828624c95295e")
	//raw, _ := mongodb.Database("test").Collection("body_info").FindOne(ctx, bson.D{{"_id", objId}})
	//var obj struct{
	//	Id string `json:"_id" bson:"_id"`
	//	UserId string `json:"user_id" bson:"user_id"`
	//	Bmi []struct{
	//		Date string `json:"date" bson:"date"`
	//		Value string `json:"value" bson:"value"`
	//	}	`json:"bmi"`
	//}
	//str := `{"_id":  {"$oid":"6152d075210828624c95295e"},"user_id": "111","bmi": [{"date": "111","value": "197.04"},{"date": {"$numberInt":"112"},"value": "198.04"}]}`
	//err = json.Unmarshal([]byte(str), &obj)
	//fmt.Println("errrrrror:", err)
	//fmt.Println("rawrawrawrawrawrawraw:", obj)
	//if els, err := raw.Elements(); err == nil {
	//	for _, v := range els {
	//		fmt.Println("elselselselselselsels:", v)
	//	}
	//}
	//mongodb.Database("test").Collection("body_info").UpdateOne(ctx, bson.D{{"_id", objId}}, bson.D{{"$set", bson.D{{"user_id", 111}}}})
	// doc：https://docs.mongodb.com/v3.0/core/index-unique/#unique-constraint-across-separate-documents
	// 多个元素或者索引数组中的嵌入式文档 可以拥有同样的值
	mongodb.Database("patient_archives").Collection("patient_archives_assay").InsertOne(ctx, bson.D{
		{"illId", 27537},
		{"illName", "慢喉病"},
		{"whetherRelapse", 1},
		{"relapseTime", 1622476800},
		{"treatmentProcess", bson.A{
			bson.D{
				{"id", 1532},
				{"treatmentEffect", "可以"},
				{"treatmentEffectTime", 1632672000},
				{"treatmentEnd", nil},
				{"treatmentStart", 1632067200},
				{"treatmentPlan", []string{"默认方案", "默认方案2"}},
			},
			bson.D{
				{"id", 1533},
				{"treatmentEffect", "可以"},
				{"treatmentEffectTime", 1632758400},
				{"treatmentEnd", 1632672000},
				{"treatmentStart", 1631462400},
				{"treatmentPlan", []string{"默认方案", "默认方案2"}},
			},
		}},
	})
	cursor, _ := mongodb.Database("patient_archives").Collection("patient_archives_assay").Find(ctx, bson.D{{"illId", 27537}})
	var objs  []struct{
		Id string `json:"id" bson:"_id"`
		IllId int64 `json:"illId" bson:"illId"`
		TreatmentProcess []struct{
			Id int64 `json:"id" bson:"id"`
			TreatmentPlan []string `json:"treatmentPlan" bson:"treatmentPlan"`
		}
	}
	if err := cursor.All(ctx, &objs); err != nil {
		// 出异常
		log.Errorf("[module|test-mongodb|test_mongodb] Handle error:%+v", err)
	}
	for _, v := range objs {
		fmt.Println("------v", v.Id, v.IllId, v.TreatmentProcess[0].Id, v.TreatmentProcess[0].TreatmentPlan)
		for _, v2 := range v.TreatmentProcess {
			fmt.Println("======v", v.IllId, v2.Id)
		}
	}

	//objId, _ := primitive.ObjectIDFromHex("616405a7bd3f824dd5ed3a76")
	//mongodb.Database("test").Collection("patient_archives_assay").UpdateOne(ctx, bson.D{{"illId", 27537}}, bson.D{{"$pull", bson.D{{"treatmentProcess", bson.D{{"id", 1529}}}}}})
	//mongodb.Database("test").Collection("patient_archives_assay").UpdateMany(ctx, bson.D{{"illId", 27537}}, bson.D{{"$pull", bson.D{{"treatmentProcess", bson.D{{"id", 1528}}}}}})
	//mongodb.Database("test").Collection("patient_archives_assay").UpdateMany(ctx, bson.D{{"illId", 27537}}, bson.D{{"$push", bson.D{{"treatmentProcess", bson.D{{"id", 1528}, {"treatmentEffect", "可以"}, {"treatmentEffectTime", 1632758400}, {"treatmentEnd", 1632672000}, {"treatmentStart", 1631462400}, {"treatmentPlan", []string{"默认方案", "默认方案2"}}}}}}})


	//mongodb.Database("test").Collection("body_info").Find(ctx, bson.D{{"user_id", 111}})
	//mongodb.Database("test").Collection("test").InsertOne(ctx, bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}})
	//mongodb.Database("test").Collection("test").FindOne(ctx, bson.D{{"foo", "bar"}})
	//mongodb.Database("test").Collection("test").SkipOption(0).LimitOption(2).SortOption(bson.D{{"_id", -1}}).Find(ctx, bson.D{{"foo", "bar"}})
	//mongodb.Database("test").Collection("test").CreateIndex(ctx)
	mongodb.Database("test").Collection("patient_archives_assay").GetIndexes(ctx)
	//mongodb.Database("med_api_gateway_admin").Collection("gateway_network_traffic").GetIndexes(ctx)
	//var start, end int64 = 1632625453147, 1632711853147
	//fmt.Println("---start--", start % 1000, time.Unix((start - start % 1000) / 1000, start % 1000))
	//fmt.Println("---end--", end % 1000, time.Unix((end - end % 1000) / 1000, end % 1000))

	//mongodb.Database("med_api_gateway_admin").Collection("gateway_network_traffic").SkipOption(3).LimitOption(10).Find(ctx, bson.D{{"gt", time.Unix(start % 1000, start - start % 1000)}, {"lt", time.Unix(end % 1000, end - end % 1000)}})
	//mongodb.Database("med_api_gateway_admin").Collection("gateway_network_traffic").SkipOption(0).LimitOption(1).Find(ctx, bson.D{{"network_traffic_id", "d0d8988a-1f6a-11ec-b2bf-0aff024567b3"}})
	//mongodb.Database("med_api_gateway_admin").Collection("gateway_network_traffic").Count(ctx, bson.D{{"created_at", bson.D{{"$gt", time.Unix((start - start % 1000) / 1000, start % 1000)}, {"$lt", time.Unix((end - end % 1000) / 1000, end % 1000)}}}})
	fmt.Printf("花费时间:%d ms", time.Now().UnixNano() / 1e6 - nowms)
	return
}
