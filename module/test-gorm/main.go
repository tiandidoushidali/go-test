package main

import (
	"fmt"
	sqql "github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"go-test/library/database"
	"net"
	"sync"
	"time"
)
var i = 10
func init() {
	i = 11
}

type User struct {
	ID int `json:"id" gorm:"column:id;default"`
	Phone string `json:"phone" gorm:"column:phone;default"`
	Platform string `json:"platform" gorm:"column:platform;default"`
	UserName string `json:"username" gorm:"column:username"`
}

func (t User) TableName() string {
	return "user"
}

var (
	// rfc1918
	rfc1918 = []*net.IPNet{
		{IP: net.IPv4(0x0A, 0x00, 0x00, 0x00), Mask: net.CIDRMask(0x08, 32)}, // 10.0.0.0/8
		{IP: net.IPv4(0xAC, 0x10, 0x00, 0x00), Mask: net.CIDRMask(0x0C, 32)}, // 172.16.0.0/12
		{IP: net.IPv4(0xC0, 0xA8, 0x00, 0x00), Mask: net.CIDRMask(0x10, 32)}, // 192.168.0.0/16
	}

	PATH = "/grape/comet"
)

func main() {
	//ip := net.ParseIP("114.114.114.114, 10.0.0.1")
	//if ip.IsLoopback() {
	//	fmt.Println("---LOOPBK---")
	//}
	//for _, cidr := range rfc1918 {
	//	if cidr.Contains(ip) {
	//		fmt.Println("----PRIVATE----")
	//	}
	//}
	//
	//fmt.Println("----PUBLIC----")
	//fmt.Println("-----ip------", ip)
	//return
	ormDb := database.NewMySQL(&sqql.Config{
		//DSN:          "root:root123456@tcp(127.0.0.1:3306)/test?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		//DSN:          "test_write:u2RuMevJ5kGSmMLs@tcp(118.31.236.23:3306)/med_calendar?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		DSN: "daipei:Medlinker_246@tcp(118.31.236.23:3306)/med_primary?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
	})
	fmt.Println("---", decimal.NewFromInt(1234567).Mul(decimal.NewFromFloat(0.01)).String())
	var u *User
	fmt.Println("--uuuuu1--", &u)
	u = new(User)
	fmt.Println("--uuuuu2--", &u)
	uu1 := &u
	uu2 := u
	fmt.Println("--uu1--", &uu1, &uu2, &u, uu1)
	var u1 *User
	fmt.Println("-------uuuuuu1----", &u1)
	u1 = get2()
	fmt.Println("-------uuuuuu2----", &u1)
	var u2 *User = get2()
	fmt.Println("-------uuuuuu----", &u2)

	uu := User{}
	ormDb.Model(User{}).Order("pk desc").Limit(10).Find(&uu)
	fmt.Println("------uu-----", uu)

	var uus []*User
	ormDb.Model(User{}).Order("pk desc").Limit(10).Find(&uus)
	for _, u := range uus {
		fmt.Println("----uuuuu----", *u)
	}
	fmt.Println("----len----", len(uus))
	return

	rows, err := ormDb.Model(User{}).Select("id, cellphone").Limit(0).Rows()

	rows.Next()
	for rows.Next() {
		user := new(User)
		rows.Scan(&user.ID, &user.Phone)

		fmt.Println("======", user.ID, user.Phone)
	}
	return
	get(ormDb)
	get(ormDb)
	list, err := get(ormDb)
	fmt.Println("--", list, len(list))
	for i := 0; i < len(list); i++ {
		//fmt.Println("--", list[i].ID)
	}
	time.Sleep(2 * time.Second)
	return

	//us := make([]*User, 0)
	var us []*User

	err = ormDb.Model(&User{}).Where("id = ?" , -1).Offset(0).Limit(20).Find(&us).Error
	fmt.Println(err, us)
	for i := 0; i < len(us); i++ {
		fmt.Println("--", us[i].ID)
	}
}

var once sync.Once
func get(ormDb *gorm.DB) (us []*User, err error) {
	err = ormDb.Model(&User{}).Offset(0).Limit(20).Find(&us).Error

	once.Do(func() {
		fmt.Println("=====111====", time.Now().UnixNano())
	})
	return
}

var u *User

func get2() *User {
	once.Do(func() {
		u = new(User)
		return
	})
	return u
}
