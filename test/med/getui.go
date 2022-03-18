package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 获取个推token

	var appId, appKey, masterSecret = "iav63gJVAy8InyoTgLKBD1", "pUos3ATg4D7N7pdeFWNbW", "xMXpSywNBUAj2y292QmHp4"
	url := fmt.Sprintf("https://restapi.getui.com/v2/%s/auth", appId)


	ts := time.Now().Unix() * 1000 + int64(time.Now().Nanosecond() / 1000000)
	hs := sha256.New()
	hs.Write([]byte(fmt.Sprintf("%s%d%s", appKey, ts, masterSecret)))
	params := map[string]interface{} {
		"sign":fmt.Sprintf("%x", hs.Sum(nil)),
		"timestamp": ts,
		"appkey": appKey,
	}
	js, _ := json.Marshal(params)
	fmt.Println("-----json-----", bytes.NewBuffer(js).String())
	res, err := http.Post(url, "application/json", bytes.NewBuffer(js))
	if err != nil {
		fmt.Println("----err-----", err)
		return
	}

	bf := bytes.NewBuffer(make([]byte, 0))
	bf.ReadFrom(res.Body)
	fmt.Println("----res-----", res.StatusCode, bf.String())

}
