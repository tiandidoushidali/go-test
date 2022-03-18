package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/copier"
	"go-test/library/utils"
	test_student_v1 "go-test/module/test-grpc/proto"
	"sort"
	"strconv"
	"strings"
	"time"
)

type tt struct {
	Name string
}
func test2222(t *tt) {
	t.Name = "张三123"
}

func test2233() string {
	defer func() string {
		fmt.Println("--1--", 222)
		return "222"
	}()
	defer func() string {
		fmt.Println("--2--", 333)
		return "333"
	}()
	return "test"
}
func main() {
	fmt.Println("-----", time.Now().AddDate(0, 0, -6))
	return
	//fmt.Println("--3--", test2233())
	//return
	ttim, err := time.ParseInLocation("2006-1-2", "2021-1-2", time.Local)
	if err != nil {
		fmt.Println("-----", err)
	}
	fmt.Println(utils.IsHolidayYear(2022, "2022-10-03"), ttim)
	return
	fmt.Println("-----", strings.Join([]string{"a", "b", "c"}, "，"))
	return
	ss := make([]string, 0, 4)
	ss = append(ss, "aa")
	ss = append(ss, "aa")
	ss = append(ss, "aa")
	ss = append(ss, "aa")
	ss = append(ss, "aa")
	fmt.Println("----", ss)
	return
	type Project struct {
		ProjectId int64 `json:"project_id"`
		Name string `json:"name"`
	}
	type ProjectEntity struct {
		ProjectId int `json:"name"`
		Name string `json:"name"`
	}
	p := &Project{
		ProjectId: 11,
		Name:      "张三",
	}
	var pp = new(ProjectEntity)
	copier.Copy(&pp, p)
	fmt.Println(pp)
	return
	fmt.Println("----", strings.HasPrefix("abcde", "abcde"))
	return
	req := &test_student_v1.StudentInfoReq{
		Id:                   1111,
	}
	var dd []byte
	dd, _ = proto.Marshal(req)
	s := base64.StdEncoding.EncodeToString(dd)
	bb, _ := base64.StdEncoding.DecodeString(s)
	req1 := new(test_student_v1.StudentInfoReq)
	proto.Unmarshal(bb, req1)
	fmt.Println(dd, s, req1.Id)
	return
	for i := 0; i < 5; i ++ {
		if i == 3 {
			return
		}
		fmt.Println("iiii", i)
	}
	fmt.Println(fmt.Sprintf("%.4f", 9.82446))
	return
	aa1 := new(tt)
	arr2 := &tt{Name: "s1"}
	copier.Copy(&aa1, arr2)
	fmt.Println(*aa1)
	return
	var a, b = 2, 3
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(a)/float64(b)), 64)
	fmt.Println(num)
	return
	arr := []int64{1, 2, 3, 4, 5}
	arr = arr[:9]
	fmt.Println("---", arr)
	return
	t := fmt.Sprintf("%s 23:59:59", time.Now().Format("2006-01-02"))
	ts, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	fmt.Println("---", ts.Unix()-time.Now().Unix())
	return
	fmt.Println(utils.IsHolidayYear(time.Now().Year()+1, "2022-10-07"))
	fmt.Println(utils.IsHolidayYear(time.Now().Year()+1, "2022-04-03"))
	fmt.Println(utils.IsHolidayYear(time.Now().Year()+1, "2022-04-02"))
	fmt.Println(utils.IsHolidayYear(time.Now().Year()+1, "2022-04-04"))
	fmt.Println(utils.IsHolidayYear(time.Now().Year()+1, "2022-04-11"))
	fmt.Println(utils.IsHolidayYear(time.Now().Year()+1, time.Now().AddDate(0, 0, 3)))
	return
	mp := make(map[string]int64)
	mp["a"] = mp["a"] + 1
	fmt.Println(mp["a"])
	return
	weekDay := utils.GetCurrentMondayDate(0)
	fmt.Println(weekDay.Format("2006-01-02"))
	return
	fmt.Println("---year---", time.Now().Year())
	str := "[\"123\", \"1234\"]"
	userIds := make([]string, 0)
	json.Unmarshal([]byte(str), &userIds)
	fmt.Println(userIds)
	if "2012-12-12 12:10:12" < "2012-12-12 12:11:12" {
		fmt.Println("111")
	}
	type Test2 struct {
		Status int64
	}
	type Test struct {
		Id int64
		Name string
		Test2 Test2
	}
	type PTest struct {
		Tt []*Test
	}

	tt := make([]*Test, 0)
	tt = append(tt, &Test{
		Id:   1,
		Name: "张三",
		Test2: Test2{Status: 1},
	})
	tt = append(tt, &Test{
		Id:   3,
		Name: "王武",
		Test2: Test2{Status: 3},
	})
	tt = append(tt, &Test{
		Id:   2,
		Name: "李四",
		Test2: Test2{Status: 2},
	})
	pt := &PTest{}
	pt.Tt = tt
	sort.SliceStable(pt.Tt, func(i, j int) bool {
		return pt.Tt[i].Test2.Status < pt.Tt[j].Test2.Status
	})
	j, _ := json.Marshal(pt.Tt)
	fmt.Println(string(j))
}
