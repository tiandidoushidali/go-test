package main

import (
	"fmt"
	"go-test/library/utils"
	"strings"
	"time"
)

func main() {

	mic := time.Now().UnixNano() / 1000000
	id := utils.GenerateSnowflake(mic, 0)
	fmt.Println("----", id, mic)
	fmt.Println("----", utils.Parse(id))
	return
	fmt.Println(fmt.Sprintf("%03X", 1))
	return
	type Item struct {
		DeptCode string
	}
	item := new(Item)
	item.DeptCode = "001-007-009-006"
	isPush := strings.HasPrefix(item.DeptCode, "001-009") ||
		strings.HasPrefix(item.DeptCode, "001-00A") ||
		strings.HasPrefix(item.DeptCode, "001-006") ||
		strings.HasPrefix(item.DeptCode, "001-004") ||
		strings.HasPrefix(item.DeptCode, "001-008") ||
		strings.HasPrefix(item.DeptCode, "001-001") ||
		strings.HasPrefix(item.DeptCode, "001-005") ||
		strings.HasPrefix(item.DeptCode, "001-010") ||
		strings.HasPrefix(item.DeptCode, "001-002") ||
		strings.HasPrefix(item.DeptCode, "001-003") ||
		strings.HasPrefix(item.DeptCode, "001-00B") ||
		strings.HasPrefix(item.DeptCode, "001-00E") ||
		strings.HasPrefix(item.DeptCode, "001-00D") ||
		strings.HasPrefix(item.DeptCode, "001-00C") ||
		strings.HasPrefix(item.DeptCode, "001-014") ||
		strings.HasPrefix(item.DeptCode, "001-00F")

	fmt.Println("-------", isPush)
	return
	t, _ := time.ParseInLocation(utils.FORMAT_DAY, "2022-01-10", time.Local)
	t2 := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.UTC().Year(), t.UTC().Month(), t.UTC().Day(), t.UTC().Hour(), t.UTC().Minute(), t.UTC().Second())
	fmt.Println("====", "2022-1-10T00:00:00+08:00" >= t2, t2)
	t3, _ := time.ParseInLocation(utils.FORMAT_YMDHIS, "2022-01-10 01:01:10", time.Local)
	fmt.Println("------", t3)
}
