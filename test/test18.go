package main

import (
	"fmt"
	"go-test/library/utils"
	"net/http"
	"sync"
	"time"
)

func ttt(c chan<- *string) {
	a := "1"
	c <- &a
	b := "2"
	c <- &b
}
func main() {

	url := "https://api-qa.medlinker.com/order-interface/v1/jindong/jd_api_callback/order_detail?app_key=5F2ED914B148CD7FD3E3983DAC157BD1&method=com.jd.health.ares.open.platform.export.service.MedicineOrderExportService.queryOrderDetailsFromThirdPart&v=2.0&sign=BAB46760F24F9111FEDAAA72D0153179&timestamp=2022-07-14+16%3A38%3A58"
	url = "https://api-qa.medlinker.com/order-interface/v1/jindong/jd_api_callback/sync_sku_stock?360buy_param_json=%7B%22skuId%22%3A%223187739%22%2C%22sourceTimestamp%22%3A1658290923862%2C%22status%22%3A1%2C%22stockSize%22%3A1000%7D"
	//var b bytes.Reader
	var wg sync.WaitGroup
	start := time.Now().Unix()
	fmt.Println("---start---", time.Now().Format(utils.FORMAT_DAY_HIS))
	for i := 0; i < 3000; i ++ {
		wg.Add(1)
		go func(i int) {
			tr := &http.Transport{
				MaxIdleConns: 10,
				IdleConnTimeout: 1 * time.Second,
				DisableCompression: true,
			}
			c := http.Client{Transport: tr}
			req, _ := http.NewRequest("GET", url, nil)
			res, err := c.Do(req)
			//res, err := http.Post(url, "json", &b)
			fmt.Println("------", err, res)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("---end---", time.Now().Format(utils.FORMAT_DAY_HIS))
	fmt.Println("---spend---", time.Now().Unix()-start)
}
