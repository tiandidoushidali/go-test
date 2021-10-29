package main

import (
	"bytes"
	"fmt"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/time"
	"go-test/library/database"
)

type CanalTestModel struct {
	ID int `json:"id" gorm:"column:id;default"`
	Name string `json:"name" gorm:"column:name;default"`
	Description string `json:"description" gorm:"column:description;default"`
	Status int `json:"status" gorm:"column:status;default"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;default"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;default"`
}

func (t CanalTestModel) TableName() string {
	return "canal_test"
}
func main() {
	ormDb := database.NewMySQL(&sql.Config{
		DSN:          "root:root123456@tcp(127.0.0.1:3306)/test?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
	})

	//ctm := CanalTestModel{
	//	Status: 0,
	//}

	//var ctm CanalTestModel
	//ormDb.Model(CanalTestModel{}).First(&ctm)
	//fmt.Println("cccccc---", ctm)
	//db1 := ormDb.Model(CanalTestModel{}).Create(CanalTestModel{
	//	ID:        11,
	//	Status:    1,
	//})
	//if err := db1.Error; err != nil {
	//	fmt.Println("errrrrrrr---", err)
	//	return
	//}
	//db := ormDb.Model(CanalTestModel{}).Select("Status").Where("id = ?", 2).Updates(CanalTestModel{
	//	Status: 2,
	//})
	//if err := db.Error; err != nil {
	//	fmt.Println("db err ", err)
	//}
	//db = ormDb.Model(CanalTestModel{}).Updates(map[string]interface{}{"status": 0})

	var inserts = make([]*CanalTestModel, 0)
	inserts = append(inserts, &CanalTestModel{
		Name:        "name1",
		Description: "description1",
		Status: 1,
	})
	inserts = append(inserts, &CanalTestModel{
		Name:        "name2",
		Description: "description2",
		Status: 1,
	})
	var err error
	var buffer bytes.Buffer
	sql := "insert into canal_test(name, description) values"
	if _, err = buffer.WriteString(sql); err != nil {
		return
	}
	for i, e := range inserts {
		if i == len(inserts)-1 {
			buffer.WriteString(fmt.Sprintf("('%s', '%s');", e.Name, e.Description))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s', '%s'),", e.Name, e.Description))
		}
	}
	err = ormDb.Model(CanalTestModel{}).Exec(buffer.String()).Error
	fmt.Println("-----", err)
	//fmt.Println("ctmctm ", ctm)
}
