package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"sync"
	"sync/atomic"
)

//func CalFee(src int64, feeRate float64) int64 {
//	old := fmt.Sprintf("%0.0f", float64(src)*feeRate)
//	dest, _ := strconv.ParseUint(old, 10, 64)
//	return int64(dest)
//}

func CalFee(src int64, feeRate float64) int64 {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", feeRate), 64)

	decimalValue := decimal.NewFromFloat(num)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(src))

	return decimalValue.BigInt().Int64()
}

func main() {


	var sum int64 = 0
	var sumT int64 = 0
	var sumF int64 = 0
	var i int64 = 1
	sql := "insert into test values"
	for i = 1; i <= 5000; i ++ {
		fee := CalFee(i * 100, 0.063)
		sum += i*100 + fee
		fmt.Println("i:", i, " t:", i * 100, " fee:", fee, " f:",  float32(i * 100) * 0.063)
		sumT += i * 100
		sumF += fee
		//fmt.Println(fmt.Sprintf("%d,%d, %.2f", i*100, fee, float32(i * 100) * 0.063))
		sql = sql + fmt.Sprintf("(%d,%d),", i*100, fee)
	}
	sql = sql[:len(sql)-1]
	//fmt.Println("sql:", sql)
	fmt.Println("sumT:", sumT, " sumF:", sumF)
	total := sum + 2636603517
	dif := total - CalFee(29100, 1.063)
	fmt.Println("sum:", sum, " total:", total, " dif:", dif)
	// 13289865.55
	// 13289863.05
	return
	var wg sync.WaitGroup
	var n int64

	for i := 0; i < 100000; i ++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			atomic.AddInt64(&n, int64(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("--n:",n)
}
