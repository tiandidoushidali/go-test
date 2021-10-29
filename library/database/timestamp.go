package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"reflect"
)

// updateTimeStampForCreateCallback will set `ctime`, `mtime` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := gorm.NowFunc()

		if createdAtField, ok := scope.FieldByName("ctime"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}
		if createdAtField, ok := scope.FieldByName("mtime"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}
		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(now)
			}
		}
		if createdAtField, ok := scope.FieldByName("CreateTime"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdateTime"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(now)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("mtime", gorm.NowFunc())
	}

	if _, ok := scope.Get("gorm:update_column"); !ok {
		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			val := updatedAtField.Field.Type().Kind()
			if val >= reflect.Int && val <= reflect.Uint64 {
				scope.SetColumn("UpdatedAt", gorm.NowFunc().Unix())
			}else{
				scope.SetColumn("UpdatedAt", gorm.NowFunc())
			}
		}
	}
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdateTime", time.Now())
	}
}
