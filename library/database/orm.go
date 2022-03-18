package database

import (
	"fmt"
	mmysql "github.com/go-sql-driver/mysql"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/go-kratos/kratos/pkg/log"

	// database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Pager struct {
	Limit  uint32
	Offset uint32
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	if len(v) > 2 {
		if _, ok := v[2].(*mmysql.MySQLError); ok {
			mErr := v[2].(*mmysql.MySQLError)
			if mErr.Number >= 1000 && mErr.Number < 2000 {
				msg := fmt.Sprintf("%v", gorm.LogFormatter(v...))
				msg = strings.Replace(msg, fmt.Sprintf("Error %d:", mErr.Number), fmt.Sprintf("SqlError %d:", mErr.Number), 1)
				log.Error(msg)
				return
			}
		}
	}
	fmt.Println("=========", v[0], v[1], v[2])
	log.Info("%d, %v", len(v), gorm.LogFormatter(v...))
}

func init() {
	gorm.ErrRecordNotFound = ecode.NothingFound
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *sql.Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		log.Error("orm: open error(%v)", err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	db.SetLogger(ormLog{})
	db.LogMode(true)

	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	return
}
