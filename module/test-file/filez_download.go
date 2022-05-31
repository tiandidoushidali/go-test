package main

import (
	"github.com/go-kratos/kratos/pkg/log"
	"io"
	"net/http"
	"os"
)

func main() {
	url, dist := "https://pub-temp-file-qa.medlinker.com/yr/du-5000.xlsx", "/tmp/2022/04/12/17c0d9b0-ba36-11ec-a910-acde48001122.xlsx"
	resp, err := http.Get(url)
	if err != nil {
		log.Error("download fail error params(url:%s, dist:%s) err(%+v) ", url, dist, err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(dist)
	if err != nil {
		log.Error("create file error params(url:%s, dist:%s) err(%+v) ", url, dist, err)
		return
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Error("write resp body to fail error params(url:%s, dist:%s) err(%+v) ", url, dist, err)
		return
	}
	return
}
