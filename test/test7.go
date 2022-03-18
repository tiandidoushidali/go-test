package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	nethttp "net/http"
	"os"
	"strings"
	"time"
)

var sql []string
var d map[int64]string

func main() {
	get1(1, "001", 1)
	file := fmt.Sprintf("./sql%d.txt", time.Now().Unix())
	f1, _ := os.Create(file)
	w1 := bufio.NewWriter(f1)
	w1.WriteString(strings.Join(sql, "\n"))
	w1.Flush()
	f1.Close()

}

func get1(deptId int64, parentCode string, level int64) {
	type Result struct {
		DeptId   int64  `json:"dept_id"`
		Name     string `json:"name"`
		ParentId int64  `json:"parent_id"`
	}
	type Res struct {
		Errcode int64     `json:"errcode"`
		Errmsg  string    `json:"errmsg"`
		Result  []*Result `json:"result"`
	}
	params := map[string]interface{}{
		"dept_id": deptId,
	}
	bt, _ := json.Marshal(params)
	req, err := nethttp.NewRequest("POST", "https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=32ed4ebb962f32068f4fbf8bb2da98b3", bytes.NewBuffer(bt))
	if nil != err {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &nethttp.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	var res *Res
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range res.Result {
		var pcode string
		if _, ok := d[item.DeptId]; ok {
			isql := fmt.Sprintf("update `med_calendar`.`dept` set state = 1 where dept_id = %d;", item.DeptId)
			sql = append(sql, isql)
			pcode = d[item.DeptId]
		} else {
			var ix int64
			for _, v := range d {
				if len(v)-len(parentCode) != 4 {
					continue
				}
				if strings.HasPrefix(v, parentCode) {
					ix++
				}
			}
			icode := fmt.Sprintf("%X", ix+1)
			if len(icode) == 1 {
				icode = "00" + icode
			} else if len(icode) == 2 {
				icode = "0" + icode
			}
			pcode = fmt.Sprintf("%s-%s", parentCode, icode)
			println(item.DeptId, item.Name, pcode)
			d[item.DeptId] = pcode
			isql := fmt.Sprintf("INSERT INTO `med_calendar`.`dept`(`dept_id`, `name`, `dept_code`, `level`, `parent_id`, `parent_dept_code`, `state`) VALUES ( '%d', '%s', '%s', %d, '%d', '%s', 1);", item.DeptId, item.Name, pcode, level, item.ParentId, parentCode)
			sql = append(sql, isql)
		}
		//getUser(item.DeptId, pcode)
		get1(item.DeptId, pcode, level+1)
	}
}

func getUser(deptId int64, parentCode string) {
	type List struct {
		Leader bool   `json:"leader"`
		UserId string `json:"userid"`
		Name   string `json:"name"`
		Email   string `json:"email"`
	}
	type Result struct {
		HasMore    bool    `json:"has_more"`
		NextCursor int64   `json:"next_cursor"`
		List       []*List `json:"list"`
	}
	type Res struct {
		Errcode int64   `json:"errcode"`
		Errmsg  string  `json:"errmsg"`
		Result  *Result `json:"result"`
	}
	params := map[string]interface{}{
		"dept_id": deptId,
		"cursor":  0,
		"size":    100,
	}
	bt, _ := json.Marshal(params)
	req, err := nethttp.NewRequest("POST", "https://oapi.dingtalk.com/topapi/v2/user/list?access_token=32ed4ebb962f32068f4fbf8bb2da98b3", bytes.NewBuffer(bt))
	if nil != err {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &nethttp.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	//println(string(body))
	var res *Res
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.Result == nil {
		return
	}
	for _, item := range res.Result.List {
		//println(item.UserId, item.Name, pcode)
		var isLeader int64
		if item.Leader {
			isLeader = 1
		}
		isql := fmt.Sprintf("INSERT INTO `med_calendar`.`user`(`user_id`, `name`, `dept_id`, `state`, `dept_code`, `is_leader`, `email`) VALUES ('%s', '%s', '%d', 1, '%s', %d, '%s');", item.UserId, item.Name, deptId, parentCode, isLeader, item.Email)
		sql = append(sql, isql)
	}
}