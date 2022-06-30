package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/snowflake"
	_ "github.com/bwmarrin/snowflake"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/jinzhu/copier"
	"go-test/library/utils"
	"math"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
)

////go:generate mockery -name=CC
type CC interface {
	Get(key string) (data interface{}, err error)
}

type Bu int

const (
	A Bu = iota - 1
	B
	C
)

type AAAA struct {
	Id int64 `copier:"id" json:"id"`
}

type BBBB struct {
	ID   int64 `copier:"id json:"id"`
	Name string
}

func testA(a *AAAA) {
	a.Id = 2
}

func convert(a interface{}) {

	fmt.Println("---", reflect.TypeOf(a).Kind() == reflect.Map)
}

func main() {

	strr := []string{"a", "b", "c"}
	fmt.Println("----", strings.Join(strr, ","))
	return
	lsof_path, _ := exec.LookPath("lsof")
	lsof_cmd := exec.Command(lsof_path, "-i:8033")
	var out bytes.Buffer
	lsof_cmd.Stdout = &out
	lsof_cmd.Stderr = os.Stderr
	lsof_cmd.Run()

	fmt.Println(out.String())

	tttt, er2r := time.ParseInLocation("2006-01-02T15:04:05", "2022-04-27T22:30:01", time.Local)

	//b := xtime.Time(0).Time().Format(utils.FORMAT_YMDHIS)
	fmt.Println("-----", tttt, er2r)
	return
	//l, err2 := time.LoadLocationFromTZData(time.Local, []byte("2001-01-01T00:00:00Z"))

	//tttt, err1 := time.Parse(utils.FORMAT_YMDHIS, "2001-01-01T00:00:00Z")
	tttt, err1 := time.ParseInLocation(utils.FORMAT_YMDHIS, strings.ReplaceAll(strings.ReplaceAll("0001-01-01T00:00:00Z", "T", " "), "Z", ""), time.Local)
	fmt.Println("---", tttt, err1)
	return
	for i := 0; i <= 1024; i ++ {
		fmt.Println("---i--", i % 0x400)
	}
	// 01111111111 1023
	// 10000000000 1024
	// 10000000001 -1
	// 10000100000 -1 << 5
	// 11111111111 -1
	// 11111100000 -1 << 5
	// 00000011111 -1 << 5 ^ -1
	// 00000000001 1
	// 00000000001 1
	// 00000100000 1 << 5
	// 00000100001 1 << 5 ^ 1
	// 11111111111 -1 补码
	// 10000001000 -8 源码
	// 11111110111 -8 反码
	// 11111111000 -8 补码 反码 + 1
	// 11111111001
	NodeBits := 5
	fmt.Println("-----", uint64(1) << NodeBits ^ uint64(1), -1 << NodeBits, -1 << NodeBits ^ int64(-1), -1 << NodeBits, -1 ^ (-1 << NodeBits))
	var num int64 = 10231111
	fmt.Println("----", num % 0x400)
	_, err := snowflake.NewNode(num % 0x400)
    if err != nil {
    	panic(err)
	}
	return
	data := []int{
		1,2,3,4,5,6,7,8,9,10,
		11,12,13,14,15,16,17,18,19,20,
	}
	perNum := 3
	fmt.Println("------", float64(len(data)) / float64(perNum))
	le := int(math.Ceil(float64(len(data)) / float64(perNum)))
	modNum := len(data) % perNum
	if le == 0 {
		fmt.Println("为空")
		return
	}
	res := make([][]int, le)
	fmt.Println("---le----", le)
	for i := 0; i < le; i ++ {
		start := i * perNum
		end := start + perNum
		fmt.Println("--start--", start, "---end---", end)

		if i == le - 1 && modNum > 0 {
			end = start + modNum
		}
		res[i] = data[start:end]
	}
	fmt.Println("------", res)
	fmt.Println("---", math.Floor(1.3505), int(math.Ceil(1.2)), math.Ceil(1.000000001), 9%3)
	return
	bb := new(BBBB)
	str := `{"id":1, "name":"zs"}`
	fmt.Println("--bb1--", bb, &bb)
	json.Unmarshal([]byte(str), &bb)
	fmt.Println("--bb2--", bb)

	bbb := reflect.New(reflect.TypeOf(bb).Elem()).Interface()
	fmt.Println("--bbb1--", bbb, &bbb)
	json.Unmarshal([]byte(str), &bbb)
	fmt.Println("--bbb2--", reflect.ValueOf(bbb))

	fmt.Println("----", bbb, reflect.TypeOf(bb), reflect.TypeOf(bbb))
	return
	fmt.Println(fmt.Sprintf("----%.2f--", float64(119191996095) * 0.001))
	return
	tt := new(xtime.Time)
	*tt = xtime.Time(time.Now().Unix())
	fmt.Println("---t-----", tt, *tt)
	return
	//url := "https://panel-file-qa.medlinker.com/Ft6ViiXwKFTjGQttD7TMu-zkrtuO"
	url := "https://panel-file-qa.medlinker.com/FmW5ouRtcCjp9zV0LUIv8BTiTxzg"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("--err--", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("----header----", resp.Header.Values("Content-Type"))
return
	bytee := make([]byte, 10)
	n, err := resp.Body.Read(bytee)
	if err != nil {
		fmt.Println("------err-----", err)
	}
	fmt.Println("------nnnn----", n)
	fmt.Println("-----", utils.GetFileType(bytee))


	return

	return
	var aaaaaa []string
	aaaaaa = append(aaaaaa, "a", "b")
	zipIncFiles := make(map[uint][]string, 0)
	fmt.Println("----aaaaaa---", aaaaaa)
	return
	zipIncFiles[1] = append(zipIncFiles[1], "a")

	fmt.Println("---", zipIncFiles)
	return
	fmt.Println("---", strconv.FormatInt(1504390061741838399, 10))

	return
	type Project struct {
		ProjectId int64  `protobuf:"varint,1,opt,name=project_id,proto3`
		Name      string `json:"name"`
	}
	type ProjectEntity struct {
		ProjectId2 int    `protobuf:"varint,1,opt,name=project_id,proto3`
		Name       string `json:"name"`
	}
	p := &Project{
		ProjectId: 11,
		Name:      "张三",
	}
	var pp = new(ProjectEntity)
	copier.Copy(&pp, p)
	fmt.Println(pp)
	return

	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 9.826), 64)
	fmt.Println("----", value)
	return
	flag.Set("test", "test")
	flag.Parse()
	var t string
	flag.StringVar(&t, "test", "", "u")
	fmt.Println(t)
	fmt.Println("-----", A, B, C)
	return
	var aa [][]int64
	aa = append(aa, []int64{1, 2, 3})
	aa = append(aa, []int64{1, 2})
	fmt.Println("---", aa)
	return
	fmt.Println("---", 10>>2, 10/3, 10%3, float64(math.Ceil(10/3)))
	sha1Buf := make([]byte, 0, 21)
	sha1Buf = append(sha1Buf, 0x16)
	fmt.Println("-----sha1Buf----", string(sha1Buf), len(sha1Buf), cap(sha1Buf), sha1Buf[0])
	// 10001
	fmt.Println("----", 80000000>>22, 0x16, 1<<1, 17>>(1<<1))
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("----", arr[4:])
}
