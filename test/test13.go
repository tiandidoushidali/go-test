package main

import (
	"fmt"
	sqql "github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/shopspring/decimal"
	"go-test/library/database"
	"go-test/library/utils"
	"strconv"
	"time"
)

type Test struct{
	Amount int64 `gorm:"column:amount"`
	Fee int64 `gorm:"column:fee"`
}

type Test2 struct{
	Amount int64 `gorm:"column:amount"`
	Fee int64 `gorm:"column:fee"`
}

func (t Test) TableName() string {
	return "test"
}

func (t *Test2) TableName() string {
	return "test2"
}

func CalFee(src int64, feeRate float64) int64 {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", feeRate), 64)

	decimalValue := decimal.NewFromFloat(num)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(src))

	return decimalValue.BigInt().Int64()
}

//func CalFee2(src int64, feeRate float64) int64 {
//	old := fmt.Sprintf("%0.0f", float64(src)*feeRate)
//	dest, _ := strconv.ParseUint(old, 10, 64)
//	return int64(dest)
//}

//func CalFee(src uint32, feeRate float64) uint32 {
//	old := fmt.Sprintf("%0.0f", float64(src)*feeRate)
//	dest, _ := strconv.ParseUint(old, 10, 64)
//	return uint32(dest)
//}

// 487500 30712 != 30713 30712.5 30712 518212 518212.47
func main() {
	var totalCount int64 = 100
	taskChan := make(chan int, totalCount)
	goChanNum := make(chan struct{}, 15)


	var num int64
	for {
		select {
			case val, ok := <-taskChan:
				if !ok { // 通道已经关闭
					break
				}
				num ++
				if num >= totalCount { // 最后一个执行完，则关闭通道
					close(taskChan)
				}
				goChanNum <- struct{}{}
				go func(val int) error {
					defer func() {
						<-goChanNum
						if err := recover(); err != nil {
							fmt.Println("err:", err)
						}
					}()

					fmt.Println("val:", val)
					return nil
				}(val)
		}
	}

	var a int64 = 1527462172165146427 >> 22
	var u int64 = a+1288834974657
	fmt.Println( 3 >> 1, time.Unix(u/ 1000, 0).Format(utils.FORMAT_YMDHIS))
	return
	//num, e := strconv.ParseFloat("4.2", 64)
	//fmt.Println(e)
	//fmt.Println("num:", CalFee(100, num))
	//
	//return
	//
	//v := 67.6
	//fmt.Println(int(v * 100), int64(6760))
	//fmt.Println(CalFee(487500, 0.063))
	//return
	ormDb := database.NewMySQL(&sqql.Config{
		DSN:          "root:root@tcp(172.16.68.79:3306)/default?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		//DSN:          "test_write:u2RuMevJ5kGSmMLs@tcp(118.31.236.23:3306)/med_calendar?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		//DSN: "daipei:Medlinker_246@tcp(118.31.236.23:3306)/med_primary?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
	})



	var results1 []*Test
	var results2 []*Test2

	ormDb.Model(Test{}).Find(&results1)
	ormDb.Table("test2").Find(&results2)

	re1map := make(map[int64]int64)
	re2map := make(map[int64]int64)
	for _, v := range results1 {
		fmt.Println("vvv", v.Amount, v.Fee)
		re1map[v.Amount] = v.Fee
	}
	for _, v := range results2 {
		re2map[v.Amount] = v.Fee
	}

	for _, v := range results1 {
		if re1map[v.Amount] != re2map[v.Amount] {
			fmt.Println(v.Amount, re1map[v.Amount], "!=", re2map[v.Amount], float32(v.Amount) *0.063, CalFee(v.Amount, 0.063), CalFee(v.Amount, 1+0.063), float32(v.Amount) * 1.063)
		}
	}

}
