package database

import (
	"fmt"
	influxHttpClientV2 "github.com/influxdata/influxdb/client/v2"
)

type InfluxdbClient struct {
	HttpClient influxHttpClientV2.HTTPClient
}

func NewInfluxdbClient(addr, userName, password string) (client *InfluxdbClient){
	httpClient, err := influxHttpClientV2.NewHTTPClient(influxHttpClientV2.HTTPConfig{
		Addr:               addr,
		Username:           userName,
		Password:           password,
	})
	if err != nil {
		fmt.Println("influxHttpClientV2 NewHTTPClient fail:", err)
		panic("influxdb new http client err" )
	}

	client = &InfluxdbClient{HttpClient: httpClient}

	return client
}
