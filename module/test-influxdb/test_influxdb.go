package test_influxdb

import (
	"context"
	"fmt"
	influxClientV2 "github.com/influxdata/influxdb-client-go/v2"

	//"github.com/influxdata/influxdb-client-go/v2"
	//influxClientV2 "github.com/influxdata/influxdb-client-go/v2"
	influxHttpClientV2 "github.com/influxdata/influxdb/client/v2"
	"go-test/library/database"
	"math/rand"
	"time"
)

const (
	YMD_FORMAT = "2006-01-02"
	YMDH_FORMAT = "2006-01-02 15"
	YMDHI_FORMAT = "2006-01-02 15:04"
	YMDHIS_FORMAT = "2006-01-02 15:04:05"
)

const bucket = "test"
const org = "test"

type TestInfluxDb struct {

}

func (t *TestInfluxDb) Handle() {
	t.queryAPI()
}

//const FluxLangQueryDomainQPS = `
//from(bucket: "%s")
//|> range(start: time(v: %d), stop: time(v: %d))
//|> filter(fn: (r) => (r._field == "response_status" or r._field == "request_time_milli_second") and (r.domain_name == "%s"))
//|> pivot(columnKey: ["_field"], rowKey: ["_time"], valueColumn: "_value")
//|> map(fn: (r) => ({
//       r with
//           request_time: time(v: r["request_time_milli_second"] * 1000000),
//   }))
//|> window(every: 1s, timeColumn: "request_time", startColumn: "request_time", stopColumn: "request_time")
//|> group(columns: ["request_time"])
//|> count(column: "response_status")
//|> map(fn: (r) => ({
//   request_time: r["request_time"],
//   qps: r["response_status"]
// }))
//|> yield()
//`

const FluxLangQueryDomainQPS = `
from(bucket:"%s")
|> range(start:-7d)
|> filter(fn: (r) => (r._field == "response_status" or r._field == "request_time_milli_second") and (r.domain_name == "%s"))
|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
|> window(every: 4h)
|> group(columns: ["_stop"])
|> count(column: "response_status")
|> map(fn: (r) => ({
   request_time: r["_stop"],
   qps: r["response_status"]
 }))
|> yield()
`

const ContinueQueryDomainQps = `
create CONTINUOUS QUERY cq_domain_qps_1s ON "api_gateway" begin select count(response_status) from "autogen" end
`

// 创建新记录


const INFLUXDB_USERNAME = "gateway_wd"
const INFLUXDB_PASSWORD = "K1QB37DDy07BSP1o"

var client database.InfluxdbClient
func init() {
	rand.Seed(time.Now().UnixNano())
	cli := database.NewInfluxdbClient("http://localhost:8086", INFLUXDB_USERNAME, INFLUXDB_PASSWORD)

	fmt.Println("influxHttpClientV2 cli:", cli)
	bps, err := influxHttpClientV2.NewBatchPoints(influxHttpClientV2.BatchPointsConfig{
		Database:         "test",
		RetentionPolicy:  "1h",
	})
	err = cli.HttpClient.WriteCtx(context.Background(), bps)
	if err != nil {
		fmt.Println("cli.WriteCtx fail reason:", err)
	}
}

func (t *TestInfluxDb) queryAPI() {

	client := influxClientV2.NewClient("https://ts-bp1hd043zt801hg8x.influxdata.tsdb.aliyuncs.com:8086", fmt.Sprintf("%s:%s", INFLUXDB_USERNAME, INFLUXDB_PASSWORD))
	//fluxQuery := fmt.Sprintf(FluxLangQueryDomainQPS, "api_gateway/autogen", 1632191493051, 1632277893051, "api-qa.medlinker.com")
	fluxQuery := fmt.Sprintf(FluxLangQueryDomainQPS, "api_gateway/autogen", "api-qa.medlinker.com")

	start := time.Now().Unix()
	ret, err := client.QueryAPI("").Query(context.Background(), fluxQuery)
	fmt.Println("----------fluxquery--expend---", time.Now().Unix()-start)
	if err != nil {
		panic(err)
	}
	for ret.Next() {
		if ret.TableChanged() {
			//fmt.Printf("table: %s\n", ret.TableMetadata().String())
		}
		// read result
		values := ret.Record().Values()
		fmt.Printf("row: %s\n", values)
	}
	if ret.Err() != nil {
		//fmt.Printf("Query error: %s\n", ret.Err().Error())
	}
	fmt.Println("end")
	client.Close()
}


func (t *TestInfluxDb) continueQueryApi() {
	//client := influxdb2.NewClient("https://ts-bp1hd043zt801hg8x.influxdata.tsdb.aliyuncs.com:8086", fmt.Sprintf("%s:%s", INFLUXDB_USERNAME, INFLUXDB_PASSWORD))
	////fluxQuery := fmt.Sprintf(FluxLangQueryDomainQPS, "api_gateway/autogen", 1632191493051, 1632277893051, "api-qa.medlinker.com")
	//fluxQuery := fmt.Sprintf(FluxLangQueryDomainQPS, "api_gateway/autogen", "api-qa.medlinker.com")

}

func (t *TestInfluxDb) writePoints() {
	apps := []string{"user-interface", "panel-interface", "order-interface"}
	status := []int{200, 401, 429, 499, 500, 502, 503, 504}
	curApp := apps[rand.Intn(2)]
	httpStatus := status[rand.Intn(7)]
	bps, _ := influxHttpClientV2.NewBatchPoints(influxHttpClientV2.BatchPointsConfig{
		Precision:        "ns",
		Database:         "test",
		RetentionPolicy:  "1h",
	})
	tags := map[string]string{
		"reqTimeS": time.Now().Format(YMDHIS_FORMAT),
		"app": "curApp",
	}
	fields := map[string]interface{}{
		"avg":"23.2",
		"max": "45",
		"method": "GET",
		"service": curApp,
		"path": fmt.Sprintf("/%s/tag/list", curApp),
		"reqTime": time.Now().Unix(),
		"resTime": rand.Intn(10),
		"http_status": fmt.Sprintf("%d", httpStatus),
	}
	pt, _ := influxHttpClientV2.NewPoint("test", tags, fields, time.Now())
	bps.AddPoint(pt)
	client.HttpClient.Write(bps)
}

func (t *TestInfluxDb) writePoints2() {
	apps := []string{"user-interface", "panel-interface", "order-interface"}
	status := []int{200, 401, 429, 499, 500, 502, 503, 504}
	curApp := apps[rand.Intn(2)]
	httpStatus := status[rand.Intn(7)]
	p := influxClientV2	.NewPointWithMeasurement("test").
		AddTag("reqTimeS", time.Now().Format(YMDHIS_FORMAT)).
		AddTag("app", curApp).
		AddField("avg", 23.2).
		AddField("max", 45).
		AddField("method", "GET").
		AddField("service", curApp).
		AddField("path", fmt.Sprintf("/%s/tag/list", curApp)).
		AddField("reqTime", time.Now().Unix()).
		AddField("resTime", rand.Intn(10)).
		AddField("http_status", httpStatus).
		SetTime(time.Now())
	var url, token, org, bucket string = "", "", "", ""
	client := influxClientV2.NewClient(url, token)
	writeAPI := client.WriteAPI(org, bucket)

	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()

}