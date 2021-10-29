package main

import (
	"context"
	"fmt"
	_ "github.com/influxdata/influxdb1-client/v2"
	influxHttpClientV2 "github.com/influxdata/influxdb/client/v2"
	influxClientV2 "github.com/influxdata/influxdb-client-go/v2"
	test_influxdb "go-test/module/test-influxdb"
	"math/rand"
	"time"
)
// influxdb influxdb
const (
	MyDB          = "test"
	username      = "test"
	password      = "test"
	MyMeasurement = "cpu_usage"
)

const (
	YMD_FORMAT = "2006-01-02"
	YMDH_FORMAT = "2006-01-02 15"
	YMDHI_FORMAT = "2006-01-02 15:04"
	YMDHIS_FORMAT = "2006-01-02 15:04:05"
)

var client influxClientV2.Client
var httpClient influxHttpClientV2.Client

const token = "yeFRlcIrJrZxG4W_xspaecLWt7YMkwow-ZSJrXC4OfHjWLKuSm4iK93ZmkQ1uMAcAJpveL_oVNhCVGmhakJ8TA=="
const bucket = "test"
const org = "test"

//func init() {
//	rand.Seed(time.Now().UnixNano())
//	cli, err := influxHttpClientV2.NewHTTPClient(influxHttpClientV2.HTTPConfig{
//		Addr:               "http://localhost:8086",
//		Username:           username,
//		Password:           password,
//	})
//	if err != nil {
//		fmt.Println("influxHttpClientV2 NewHTTPClient fail:", err)
//	}
//	httpClient = cli
//	fmt.Println("influxHttpClientV2 cli:", cli)
//	bps, err := influxHttpClientV2.NewBatchPoints(influxHttpClientV2.BatchPointsConfig{
//		Database:         "test",
//		RetentionPolicy:  "1",
//	})
//	err = cli.WriteCtx(context.Background(), bps)
//	if err != nil {
//		fmt.Println("cli.WriteCtx fail reason:", err)
//	}
//
//
//	client = influxClientV2.NewClient("http://localhost:8086", token)
//	// always close client at the end
//	//defer client.Close()
//}
func main() {
	fmt.Println("client:", client)
	// insert

	//writePoints()
	//queryPoints()
	//writePoints2()
	//writePoints3()
	//queryPoints3()
	//httpQueryPoints()

	influxDB := test_influxdb.TestInfluxDb{}
	influxDB.Handle()
}

func httpQueryPoints() {
	cmd := fmt.Sprintf("select count(*) from %s limit %d", "test", 10)
	q := influxHttpClientV2.Query{
		Command:         cmd,
		Database:        "test",
		RetentionPolicy: "1",
		Precision:       "ns",
		Chunked:         false,
		ChunkSize:       0,
		Parameters:      nil,
	}
	res, err := httpClient.Query(q)
	if err != nil {
		fmt.Println("httpClient.Query error:", err)
		if res.Error() != nil {
			fmt.Println("httpClient.Query response error", res.Error())
		}
		fmt.Println("res results:", res.Results)
	}

}
func writePoints() {
	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)
	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 22.5, 45.0))
	// Flush writes
	writeAPI.Flush()
}

func writePoints2() {
	p := influxClientV2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point asynchronously
	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)
	writeAPI.WritePoint(p)
	writeAPI.Flush()

}

func writePoints3() {
	// create point using fluent style
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
	writeAPI := client.WriteAPI(org, bucket)
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
}

func queryPoints() {
	query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"stat\")", bucket)
	// Get query client
	queryAPI := client.QueryAPI(org)
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("key: %v, value: %v\n", result.Record().Field(), result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}

func queryPoints3() {
	query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h)|> limit(n:4) |> filter(fn: (r) => r._measurement == \"test\")", bucket)
	// Get query client
	queryAPI := client.QueryAPI(org)
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Iterate over query response
		values := make(map[string]interface{})
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				//fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("key: %v, value: %vn", result.Record().Field(), result.Record().Value())
			values[result.Record().Field()] = result.Record().Value()
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}