package main

import (
	"fmt"
	sqql "github.com/go-kratos/kratos/pkg/database/sql"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/jinzhu/gorm"
	"go-test/library/database"
	"sync"
	"time"
)

type CanalTestModel struct {
	ID int `json:"id" gorm:"column:id;default"`
	Name string `json:"name" gorm:"column:name;default"`
	Description string `json:"description" gorm:"column:description;default"`
	Status int `json:"status" gorm:"column:status;default"`
	CreatedAt xtime.Time `json:"created_at" gorm:"column:created_at;default"`
	UpdatedAt xtime.Time `json:"updated_at" gorm:"column:updated_at;default"`
}

func (t *CanalTestModel) TableName() string {
	return "canal_test"
}

type Project struct {
	ID int `json:"id" gorm:"column:id;default"`
	ProjectCode string `json:"project_code" gorm:"column:project_code;default"`
	ProjectRankType int64 `json:"project_rank_type" gorm:"column:project_rank_type;default"`
	ProjectType int64 `json:"project_type" gorm:"column:project_type;default"`
}

func (p Project) TableName() string {
	return "project"
}

type ProjectTask struct {
	ID int `json:"id" gorm:"column:id;default"`
	ProjectId int64 `json:"project_id" gorm:"column:project_id;default"`
	ProjectCode string `json:"project_code" gorm:"column:project_code;default"`
	ParentProjectCode string `json:"parent_project_code" gorm:"column:parent_project_code;default"`
	ItemCode string `json:"item_code" gorm:"column:item_code;default"`
}

func (t *ProjectTask) TableName() string {
	return "project_task"
}

type ProjectNode struct {
	ID int `json:"id" gorm:"column:id;default"`
	ProjectId int64 `json:"project_id" gorm:"column:project_id;default"`
	ProjectCode string `json:"project_code" gorm:"column:project_code;default"`
}

func (t *ProjectNode) TableName() string {
	return "project_node"
}

type ProjectRisk struct {
	ID int `json:"id" gorm:"column:id;default"`
	ProjectId int64 `json:"project_id" gorm:"column:project_id;default"`
	ProjectCode string `json:"project_code" gorm:"column:project_code;default"`
}

func (t *ProjectRisk) TableName() string {
	return "project_risk"
}

type User struct {
	ID int `json:"id" gorm:"column:id;default"`
	Phone string `json:"phone" gorm:"column:phone;default"`
	Platform string `json:"platform" gorm:"column:platform;default"`
}

func (t *User) TableName() string {
	return "user"
}

type UserFocusProject struct {
	ID int `json:"id" gorm:"column:id;default"`
	ProjectId int64 `json:"project_id" gorm:"column:project_id;default"`
	ProjectCode string `json:"project_code" gorm:"column:project_code;default"`
}

func (t *UserFocusProject) TableName() string {
	return "user_focus_project"
}


type CanalTest struct {
	Id int64 `gorm:"column:id;not null;default"`
	Name string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

func (t *CanalTest) TableName() string {
	return "canal_test"
}

func main() {
	ormDb := database.NewMySQL(&sqql.Config{
		DSN:          "root:root123456@tcp(127.0.0.1:3306)/test?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		//DSN:          "test_write:u2RuMevJ5kGSmMLs@tcp(118.31.236.23:3306)/med_calendar?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		//DSN: "daipei:Medlinker_246@tcp(118.31.236.23:3306)/med_primary?timeout=4s&readTimeout=4s&writeTimeout=4s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
	})

	//cananTests := make([]*CanalTest, 0)
	//cananTests = append(cananTests, &CanalTest{
	//	Name:        "name1",
	//	Description: "desc1",
	//})
	//
	//var projectIds []int64
	//ormDb.Model(&CanalTest{}).Create(&cananTests)
	//
	//fmt.Println("---------", projectIds)
	pts := make([]*User, 0)

	ormDb.Model(&User{}).Offset(0).Limit(20).Find(&pts)

	count := 0

	lock := sync.Mutex{}
	phone, platform := "18100000001", "a"
	for i := 0; i < 10; i ++ {
		go func(phone, platform string) {
			lock.Lock()
			var err error
			db := ormDb.Begin()
			defer func() {
				lock.Unlock()
				if err != nil {
					db.Rollback()
				}
			}()
			db.Model(&User{}).Where("phone = ?", phone).Where("platform = ?", platform).Count(&count)
			if count == 0 {
				db.Create(&User{
					Phone:    phone,
					Platform: platform,
				})
			}
			if err = db.Commit().Error; err != nil {
				fmt.Println("------err------", err)
			}
		}(phone, platform)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("---------", len(pts))
}

//func main() {
//	ormDb := database.NewMySQL(&sql.Config{
//	})
//
//	p := make([]Project, 0)
//	ormDb.Model(&Project{}).Find(&p)
//
//	ps := make([]*ProjectTask, 0)
//	ormDb.Model(&ProjectTask{}).Find(&ps)
//
//	fmt.Println("ps len", len(ps), "p", len(p))
//
//	fmt.Println("------")
//	// 将项目类型归类
//	projectTypeMap := make(map[int][]int64)
//	for k, v := range p {
//		cate := 0
//		switch v.ProjectType {
//		case 0:
//			switch v.ProjectRankType {
//			case 1:
//				cate = 1
//				break
//			case 2:
//				cate = 2
//				break
//			}
//		case 1:
//			cate = 3
//			break
//		case 2:
//			cate = 4
//			break
//		}
//		if _, ok := projectTypeMap[cate]; !ok {
//			projectTypeMap[cate] = append(projectTypeMap[cate], int64(v.ID))
//		}
//		projectTypeMap[cate] = append(projectTypeMap[cate], int64(v.ID))
//
//		//ormDb.Model(&ProjectTask{}).Where("project_code = ?", v.ProjectCode).Updates(map[string]interface{}{"project_id": v.ID})
//		fmt.Println("k:", k, "v:", v.ID, "name:", v.ProjectCode)
//		//func(k int, v Project, ormDb *gorm.DB) {
//		//	eg.Go(func() error {
//		//		ormDb.Model(&ProjectTask{}).Where("project_code = ?", v.ProjectCode).Updates(map[string]interface{}{"project_id": v.ID})
//		//		return nil
//		//	})
//		//} (k, v, ormDb)
//		//ormDb.Model(&ProjectTask{}).Where("(project_code = ? or item_code = ?) and (project_code != ? and item_code != ?)", v.ProjectCode, v.ProjectCode, "", "").Updates(map[string]interface{}{"project_id": v.ID})
//		//ormDb.Model(&ProjectNode{}).Where("project_code = ?", v.ProjectCode).Updates(map[string]interface{}{"project_id": v.ID})
//		//ormDb.Model(&ProjectRisk{}).Where("project_code = ?", v.ProjectCode).Updates(map[string]interface{}{"project_id": v.ID})
//		//ormDb.Table("project").Where("parent_project_code = ?", v.ProjectCode).Updates(map[string]interface{}{"parent_id": v.ID})
//		//ormDb.Model(&UserFocusProject{}).Where("project_code = ?", v.ProjectCode).Updates(map[string]interface{}{"project_id": v.ID})
//	}
//	j, _ := json.Marshal(projectTypeMap)
//	fmt.Println("=======", string(j))
//	updateProjectTaskCate(ormDb, projectTypeMap)
//}

func updateProjectTaskCate(ormDb *gorm.DB, mp map[int][]int64) {
	for cateId, projectIds := range mp {
		ormDb.Model(&ProjectTask{}).Where("project_id in (?)", projectIds).Updates(map[string]interface{}{"cate_id": cateId})
	}
}

//func main() {
//	ormDb := database.NewMySQL(&sql.Config{
//		DSN:          "root:root123456@tcp(127.0.0.1:3306)/test?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
//	})
//
//	//ctm := CanalTestModel{
//	//	Status: 0,
//	//}
//
//	//var ctm CanalTestModel
//	//ormDb.Model(CanalTestModel{}).First(&ctm)
//	//fmt.Println("cccccc---", ctm)
//	//db1 := ormDb.Model(CanalTestModel{}).Create(CanalTestModel{
//	//	ID:        11,
//	//	Status:    1,
//	//})
//	//if err := db1.Error; err != nil {
//	//	fmt.Println("errrrrrrr---", err)
//	//	return
//	//}
//	//db := ormDb.Model(CanalTestModel{}).Select("Status").Where("id = ?", 2).Updates(CanalTestModel{
//	//	Status: 2,
//	//})
//	//if err := db.Error; err != nil {
//	//	fmt.Println("db err ", err)
//	//}
//	//db = ormDb.Model(CanalTestModel{}).Updates(map[string]interface{}{"status": 0})
//
//	//var inserts = make([]*CanalTestModel, 0)
//	//inserts = append(inserts, &CanalTestModel{
//	//	Name:        "name1",
//	//	Description: "description1",
//	//	Status: 1,
//	//})
//	//inserts = append(inserts, &CanalTestModel{
//	//	Name:        "name2",
//	//	Description: "description2",
//	//	Status: 1,
//	//})
//	//var err error
//	//var buffer bytes.Buffer
//	//sql := "insert into canal_test(name, description) values"
//	//if _, err = buffer.WriteString(sql); err != nil {
//	//	return
//	//}
//	//for i, e := range inserts {
//	//	if i == len(inserts)-1 {
//	//		buffer.WriteString(fmt.Sprintf("('%s', '%s');", e.Name, e.Description))
//	//	} else {
//	//		buffer.WriteString(fmt.Sprintf("('%s', '%s'),", e.Name, e.Description))
//	//	}
//	//}
//	//err = ormDb.Model(CanalTestModel{}).Exec(buffer.String()).Error
//	//fmt.Println("-----", err)
//	//fmt.Println("ctmctm ", ctm)
//
//	om := ormDb.Begin()
//	res := make([]*CanalTestModel, 0)
//	if err := om.Model(&CanalTestModel{}).Find(&res).Error; err != nil {
//		if !gorm.IsRecordNotFoundError(err) {
//			bt, _ := json.Marshal(res)
//			fmt.Println("bt:", string(bt))
//		}
//	}
//	json1, _ := json.Marshal(res)
//	fmt.Println("res before:", string(json1))
//	om.Model(&CanalTestModel{}).Updates(map[string]interface{} {
//		"name" : "test",
//	})
//
//
//	mmm := make([]*CanalTestModel, 0)
//	mmm = append(mmm, &CanalTestModel{
//		Name:        "name1",
//	})
//	mmm = append(mmm, &CanalTestModel{
//		Name:        "name2",
//	})
//	fmt.Println("mmm len:", len(mmm))
//	om.Model(&CanalTestModel{}).Create(&CanalTestModel{
//		Name:        "name2",
//	})
//
//	if err := om.Model(&CanalTestModel{}).Find(&res).Error; err != nil {
//		if !gorm.IsRecordNotFoundError(err) {
//			bt, _ := json.Marshal(res)
//			fmt.Println("bt:", string(bt))
//		}
//	}
//
//	json1, _ = json.Marshal(res)
//	fmt.Println("res:", string(json1))
//
//	om.Commit()
//
//	type ss struct {
//		Name string	`json:"name"`
//		Desc string	`json:"desc"`
//	}
//
//	var s []*ss
//	jstr := `[{"name":"zs", "desc":"descccc"}, {"name":"zs1", "desc":"descccc1"}]`
//	json.Unmarshal([]byte(jstr), &s)
//	for i := 0; i < len(s); i ++ {
//		fmt.Println("jjjjj11111:", s[i].Name)
//	}
//}
