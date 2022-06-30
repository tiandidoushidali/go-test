package main

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/speps/go-hashids"
	"math"
	"strconv"
	"strings"
	"time"
)

func walk(d []int) ([]int, []int) {
	// 先只处理奇数
	if len(d) == 0 {
		return []int{}, []int{}
	}

	is := make([]int, 0) // 奇数位置
	js := make([]int, 0) // 偶数位置
	for idx, val := range d {
		if (idx+1) % 2 == 1 {
			is = append(is, val)
		} else {
			js = append(js, val)
		}
	}

	fmt.Println(is)

	return walk(js)


}

func testt( f func(i int64) int64 ) int64 {
	return f(1)
}

func testt2() int {
	i := 0
	defer func() {
		i ++
		fmt.Println("---i----", i)
		return
	}()
	return func() int {
		fmt.Println("---re----", i)
		i ++
		return i
	}()
}
//go:generate echo hello
func main() {
	channel := make(chan string)
	channel <- "a"
	fmt.Println(<-channel)
	return
	s := int64(math.Pow(2, 63))
	fmt.Println("s", s)
	n := 0
	st := time.Now().UnixNano()
	for s > 0 {
		if s & 1 ==  1 {
			n ++
		}
		s >>= 1
	}
	et := time.Now().UnixNano()

	s = 999999999999999
	m := 0
	for s > 0 {
		m ++
		s &= s - 1
	}
	fmt.Println("--n--", n, m, et - st)
	return
	//cs := make(chan int, 10)
	//go func() {
	//	for i := 0; i < 50; i ++ {
	//		cs <- i
	//		time.Sleep(10 * time.Millisecond)
	//	}
	//	close(cs)
	//}()
	//
	//for {
	//	select {
	//	case v, ok := <-cs:
	//		if !ok {
	//			fmt.Println("---ok----", ok, v)
	//			time.Sleep(20 * time.Millisecond)
	//			return
	//		}
	//		fmt.Println("---ok----", ok, v)
	//		time.Sleep(20 * time.Millisecond)
	//	}
	//}
	//
	//return
	//m := sync.Mutex{}
	//var wg sync.WaitGroup
	//var n int64 = 0
	//for i := 0; i < 100000; i ++ {
	//	wg.Add(1)
	//	go func(i int ) {
	//		m.Lock()
	//		defer m.Unlock()
	//
	//		n += int64(i)
	//
	//		wg.Done()
	//	}(i)
	//}
	//
	//wg.Wait()
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//defer cancel()
	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("时间到了")
	//		goto END
	//	default:
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("睡眠了1天了")
	//	}
	//}
	//
	//END:
	//
	//fmt.Println("----n----", n )
	//
	//fmt.Println("--testt2()--", testt2())
	//
	//
	//return
	//a := []int{1, 2, 3, 4, 5}
	//bc := []int{0, 0, 0, 0, 0}
	//for i, v := range a {
	//	n := 1
	//	for _, vv := range a {
	//		if vv == v {
	//			continue
	//		}
	//		n *= vv
	//	}
	//	bc[i] = n
	//}
	//fmt.Println("--n--", bc)
	//return
	//arr := []int{1,8,3,4,1,7,2,2,1,6,1,3,1,5,0,1,9}
	//walk(arr)
	//
	//return
	//var t int64
	//var i int64
	//for i = 0; i < 10000; i ++ {
	//	go func(i int64) {
	//		atomic.AddInt64(&t, i)
	//	}(i)
	//}
	//time.Sleep(5 * time.Second)
	//fmt.Println("----t----", t)
	HashDecoder := hashids.NewWithData(&hashids.HashIDData{Alphabet: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",Salt: "<)Toyota~86(>"})
	dd, err1 := HashDecoder.DecodeInt64WithError("gjWxyQTXu2")
	if err1 != nil {
		fmt.Println("---err----", err1)
	}
	fmt.Println("-----dd-----", dd)

	str1 := "0 0 36 57 51 48 55 49 98 98 102 45 100 99 52 55 45 52 57 50 57 45 98 48 98 52 45 55 102 99 57 55 55 97 55 102 54 50 49"
	ss := strings.Split(str1, " ")

	b := make([]byte, 0)
	b = append(b, 1)
	for _, v := range ss {
		i, _ := strconv.ParseInt(v, 10, 64)
		b = append(b, uint8(i))
	}
	fmt.Println("---b----", string(b))
	str := "abcd"
	fmt.Println("--str---", []byte(str))
	return
	bytes := []byte{4, 0, 0}
	fmt.Println("b:", string(bytes))
	var err error
	err = errors.New("调用返回错误编号[1700] -- 描述 [调用中心端查询国家局人员信息发生错误，调用WebApi路由Remoting服务时出错！出错信息：【国家电子就医凭证-解析二维码】出错，出错编码为500049,出错信息电子凭证二维1码无效]")
	if strings.Contains(err.Error(), "出错信息电子凭证二维码无效") {
		fmt.Println("----info----", fmt.Sprintf("%s", "医保二维码失效，请重新截图上传"))
	} else {
		fmt.Println("---er----", err.Error())
	}
	return
	d, _ := decimal.NewFromString("10")
	fmt.Println(d.Mul(decimal.NewFromInt(10)).String())
	return
	fmt.Println(time.Now().Add(-3 * 24 * time.Hour).Format("2006-01-02 15:04:05"))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("----r----", r)
			return
		}
		if err != nil {
			fmt.Println("----r2----", err)
			return
		}
	}()
	//panic(1111111)
	err = errors.New("1111")
	return
}
