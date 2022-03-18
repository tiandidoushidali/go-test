package utils

import (
	"fmt"
	"time"
)

const (
	FORMAT_DAY = "2006-01-02"
	FORMAT_YMDHIS = "2006-01-02 15:04:05"
)
func GetCurrentMondayDate(tim  int64) time.Time {
	var t time.Time
	if tim > 0 {
		t = time.Unix(tim, 0)
	} else {
		t = time.Now()
	}
	d := int(t.Weekday())
	if t.Weekday() == time.Sunday {
		d = 7
	}
	return t.AddDate(0, 0, -d + 1)
}
// 星期六 星期日 要上班的特殊日期
var supplementClassDayMap = map[int][]string{
	2021:[]string{},
	2022:[]string{
		"2022-01-29", "2022-01-30",
		"2022-04-02", "2022-04-24",
		"2022-05-07",
		"2022-10-08", "2022-10-09",
	},
}
// 周一到周五 是节假日的日期
var classDayMap = map[int][]string{
	2021:[]string{},
	2022:[]string{
		"2022-01-03", "2022-01-31",
		"2022-02-01", "2022-02-02", "2022-02-03", "2022-02-04",
		"2022-04-04", "2022-04-05",
		"2022-05-02", "2022-05-03", "2022-05-04",
		"2022-06-03",
		"2022-09-12",
		"2022-10-03", "2022-10-04", "2022-10-05", "2022-10-06", "2022-10-07",
	},
}


// 是否是节假日
func IsHolidayYear(year int, tim interface{}) bool {
	if _, ok := supplementClassDayMap[year]; !ok {
		return false
	}
	if _, ok := classDayMap[year]; !ok {
		return false
	}
	// 星期六 星期日 要上班的特殊日期
	supplementClassDay := supplementClassDayMap[year]
	// 周一到周五 是节假日的日期
	classDay := classDayMap[year]
	var ttim time.Time
	switch tim.(type) {
	case string:
		var err error
		ttim, err = time.ParseInLocation(FORMAT_DAY, tim.(string), time.Local)
		if err != nil {
			return false
		}
	case int64:
		ttim = time.Unix(tim.(int64), 0)
	case time.Time:
		ttim = tim.(time.Time)
	default:
		return false
	}
	wkDay := ttim.Weekday()
	fmt.Println("--------", wkDay)
	if wkDay == time.Saturday || wkDay == time.Sunday { // 星期六 星期日
		// 是否会补班
		for _, v := range supplementClassDay {
			if ttim.Format(FORMAT_DAY) == v {
				return false
			}
		}
		return true
	} else {
		// 是否是假期
		for _, v := range classDay {
			if ttim.Format(FORMAT_DAY) == v {
				return true
			}
		}
		return false
	}
}

// 获取月份第一天和最后一天
func GetMonthStartEndDay(tim interface{}) (start time.Time, end time.Time) {
	if tim == nil {
		return
	}
	var ttime time.Time
	switch tim.(type) {
	case time.Time:
		ttime = tim.(time.Time)
	case string:
		times := tim.(string)
		ttime, _ = time.ParseInLocation(FORMAT_DAY, times, time.Local)
	default :
		return
	}
	start = time.Date(ttime.Year(), ttime.Month(), 1, 0, 0, 0, 0, time.Local)
	end = time.Date(ttime.Year(), ttime.Month()+1, 0, 0, 0, 0, 0, time.Local)
	return
}
