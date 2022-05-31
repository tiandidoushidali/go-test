package utils

import (
	"bytes"
	"fmt"
	"math"
	"time"
)

// 转换患者年龄
// @param birthday 出生日期
// @param reference 计算的参照时间
// 规则：
// 3个月以下显示天数，比如56天。
// 3个月（含）到3岁显示几岁几个月，如果整数岁不显示月份。
// 3岁以上（含）直接显示岁数。
func BirthdayToPatientAgeWithReference(birthday time.Time, reference time.Time) string {
	// 相差的月份
	diffMonth := ((reference.Year() - birthday.Year()) * 12) + int(reference.Month()-birthday.Month())
	if diffMonth < 3 {
		// 3个月内
		return fmt.Sprintf("%d天", reference.YearDay()-birthday.YearDay())
	} else if diffMonth < 3*12 {
		dy := int(math.Floor(float64(diffMonth / 12)))
		dm := diffMonth - (dy * 12)

		// 日期还没有到，则月份要少一个月
		if reference.Day() < birthday.Day() {
			dm = dm - 1
		}

		var rs bytes.Buffer
		if dy > 0 {
			rs.WriteString(fmt.Sprintf("%d岁", dy))
		}

		if dm > 0 {
			rs.WriteString(fmt.Sprintf("%d个月", dm))
		}

		return rs.String()
	} else {
		dy := int(math.Floor(float64(diffMonth / 12)))
		return fmt.Sprintf("%d岁", dy)
	}
}

// 转换患者年龄，以当前时间为参照点
func BirthdayToPatientAge(birthday time.Time) string {
	return BirthdayToPatientAgeWithReference(birthday, time.Now())
}
