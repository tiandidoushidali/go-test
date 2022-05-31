package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	nethttp "net/http"
	"os"
	"strings"
	"time"
)

var deptSql []string
var userSql []string
var deptMap map[int64]Dept
var deptBaseMap map[int64]string // <deptId, deptCode>
var dbDeptMap map[int64]*Dept
var oneDeptMap map[int64]string

var accessToken string = "c01797aaa8ba398e8159090a9080549b"

/**
::::::::::warn:::::::::
注意校验多次拉取数据的钉钉返回的部门顺序，这会影响部门的编号生成
*/
/**
获取钉钉token
curl https://oapi.dingtalk.com/gettoken\?appkey\=dingfdjkfkda9lp4p4ws\&appsecret\=3iVrXwYJyAgPT4KVo9jmb67-KZe4dv9uVwIEq-X2TQE44QS2dbNBPK3DvK4O3uWB
*/
// 初始化数据库中的部门
var initDeptId int64 = 0
var initDeptCode string = ""
var initLevel int64 = 1

type Dept struct {
	DeptId     int64  `json:"dept_id" gorm:"column:dept_id"`
	Name       string `json:"name" gorm:"column:name"`
	ParentId   int64  `json:"parent_id" gorm:"column:parent_id"`
	Level      int64  `json:"level" gorm:"column:level"`
	DeptCode   string `json:"dept_code" gorm:"column:dept_code"`
	ParentCode string `json:"parent_dept_code" gorm:"column:parent_dept_code"`
}

type CUser struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	DeptId   int64  `json:"dept_id"`
	DeptCode string `json:"dept_code"`
	IsLeader int64  `json:"is_leader"`
	Email    string `json:"email"`
}

var dingdingDeptList []*Dept
var dingdingUserList []*CUser

func init() {
	deptMap = make(map[int64]Dept)
	dingdingDeptList = make([]*Dept, 0)
	dingdingUserList = make([]*CUser, 0)

	deptBaseMap = make(map[int64]string)

	oneDeptMap = map[int64]string{
		484112605: "001-001",
		541166612: "001-006",
		500378762: "001-007",
		501897712: "001-008",
		483871885: "001-00B",
		484337051: "001-00C",
		483543962: "001-00D",
		483435892: "001-00E",
		483585923: "001-00F",
		483512885: "001-010",
		484163568: "001-014",
		583214290: "001-016",
		583354137: "001-017",
		583398159: "001-018",
		583645108: "001-019",
		583608165: "001-01A",
		583010694: "001-01B",
		583773101: "001-01C",
		583433149: "001-01D",
		583216492: "001-01E",
	}
	//initDbDept()
}

func initDbDept() {
	dsn := "test_write:u2RuMevJ5kGSmMLs@tcp(118.31.236.23:3306)/med_calendar?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	var depts []*Dept
	err = db.Table("dept").Find(&depts).Error
	if err != nil && !db.RecordNotFound() {
		panic(err)
	}
	for _, v := range depts {
		dbDeptMap[v.DeptId] = v
	}
}
func main() {
	start := time.Now().Unix()
	// 获取钉钉的部门数据
	getDingDingDept(1, 0, "001")
	// 生成部门sql
	genDeptSql()
	// 获取钉钉用户数据
	getDingDingUser()
	// 生成用户sql
	genUserSql()

	fmt.Println("------spend time------", time.Now().Unix()-start)
}

func genDeptSql() {
	deptSql = make([]string, 0)
	for _, item := range dingdingDeptList {
		isql := fmt.Sprintf("( '%d', '%s', '%s', %d, '%d', '%s', 1),", item.DeptId, item.Name, item.DeptCode, item.Level, item.ParentId, item.ParentCode)
		deptSql = append(deptSql, isql)

		deptBaseMap[item.DeptId] = item.DeptCode
	}
	deptFileName := fmt.Sprintf("./dept_sql%d.sql", time.Now().Unix())
	deptFile, _ := os.Create(deptFileName)
	deptWriter := bufio.NewWriter(deptFile)
	deptInsert := "INSERT INTO `med_calendar`.`dept`(`dept_id`, `name`, `dept_code`, `level`, `parent_id`, `parent_dept_code`, `state`) VALUES "
	deptInsert = fmt.Sprintf("%s%s", deptInsert, strings.Join(deptSql, "\n"))
	deptWriter.WriteString(fmt.Sprintf("%s%s", "delete from dept where id > 1 and not (dept_id >= 200000001 and dept_id < 300000000);", "\n")) // 删除除默认数据外的所有部门
	deptWriter.WriteString(fmt.Sprintf("%s%s", deptInsert[:len(deptInsert)-1], ";"))
	deptWriter.Flush()
	deptFile.Close()
}

func genUserSql() {
	userSql = make([]string, 0)
	userIds := make([]string, 0)
	for _, item := range dingdingUserList {
		isql := fmt.Sprintf("('%s', '%s', '%d', 1, '%s', %d, '%s'),", item.UserId, item.Name, item.DeptId, item.DeptCode, item.IsLeader, item.Email)
		userSql = append(userSql, isql)

		userIds = append(userIds, fmt.Sprintf("'%s'", item.UserId))
	}

	userFileName := fmt.Sprintf("./user_sql%d.sql", time.Now().Unix())
	userFile, _ := os.Create(userFileName)
	userWriter := bufio.NewWriter(userFile)
	userInsert := "INSERT INTO `med_calendar`.`user`(`user_id`, `name`, `dept_id`, `state`, `dept_code`, `is_leader`, `email`) VALUES "
	userInsert = fmt.Sprintf("%s%s", userInsert, strings.Join(userSql, "\n"))
	userIdsStr := strings.Join(userIds, ",")
	userWriter.WriteString(fmt.Sprintf("delete from `med_calendar`.`user` where user_id in (%s);%s", userIdsStr, "\n"))
	userWriter.WriteString(fmt.Sprintf("%s%s", userInsert[:len(userInsert)-1], ";"))
	userWriter.Flush()
	userFile.Close()
}

func getDingDingDept(deptId int64, level int64, deptCode string) {
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
	req, err := nethttp.NewRequest("POST", fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=%s", accessToken), bytes.NewBuffer(bt))
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
	for idx, item := range res.Result {
		deptItem := &Dept{
			DeptId:     item.DeptId,
			Name:       item.Name,
			ParentId:   item.ParentId,
			Level:      level + 1,
			ParentCode: deptCode,
			DeptCode:   fmt.Sprintf("%s-%s", deptCode, fmt.Sprintf("%03X", idx+1)),
		}
		if oneDeptCode, ok := oneDeptMap[item.DeptId]; ok {
			deptItem.DeptCode = oneDeptCode
		}
		// 如果部门id
		dingdingDeptList = append(dingdingDeptList, deptItem)
		//getUser(item.DeptId, pcode)
		fmt.Println("----------进行中-------", deptItem.DeptId, deptItem.Level, deptItem.DeptCode)
		getDingDingDept(deptItem.DeptId, deptItem.Level, deptItem.DeptCode)
	}
}

func getDingDingUser() {
	for _, item := range dingdingDeptList {
		getUserFromDept(item)
	}
}

func getUserFromDept(dept *Dept) {
	type List struct {
		Leader bool   `json:"leader"`
		UserId string `json:"userid"`
		Name   string `json:"name"`
		Email  string `json:"email"`
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
		"dept_id": dept.DeptId,
		"cursor":  0,
		"size":    100,
	}
	bt, _ := json.Marshal(params)
	req, err := nethttp.NewRequest("POST", fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/user/list?access_token=%s", accessToken), bytes.NewBuffer(bt))
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
		user := &CUser{
			UserId:   item.UserId,
			Name:     item.Name,
			DeptId:   dept.DeptId,
			DeptCode: dept.DeptCode,
			IsLeader: isLeader,
			Email:    item.Email,
		}
		dingdingUserList = append(dingdingUserList, user)
	}
}
