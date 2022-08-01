package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/lvdbing/books/blog-service/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	// Id
	Id uint32 `json:"id"`
	// CreatedOn 创建时间
	CreatedOn int32 `json:"created_on"`
	// CreatedBy 创建人
	CreatedBy string `json:"created_by"`
	// ModifiedOn 修改时间
	ModifiedOn int32 `json:"modified_on"`
	// ModifiedBy 修改人
	ModifiedBy string `json:"modified_by"`
	// DeletedOn 删除时间
	DeletedOn int32 `json:"deleted_on"`
	// IsDel 是否删除 0为未删除 1为已删除
	IsDel int8 `json:"is_del"`
}

func NewDBEngine(dbSetting *setting.SettingDatabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Callback().Create().Before("gorm:create").Register("update_create_time", updateTimeStampWhileCreate)
	db.Callback().Update().Before("gorm:update").Register("update_modify_time", updateTimeStampWhileModify)

	return db, nil
}

func updateTimeStampWhileCreate(db *gorm.DB) {
	now := time.Now().Unix()
	timeFields := []string{"CreatedOn", "ModifiedOn"}
	for _, field := range timeFields {
		timeField := db.Statement.Schema.LookUpField(field)
		if timeField == nil {
			continue
		}

		switch db.Statement.ReflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				if _, isZero := timeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue.Index(i)); isZero {
					timeField.Set(db.Statement.Context, db.Statement.ReflectValue.Index(i), now)
				}
			}
		case reflect.Struct:
			if _, isZero := timeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue); isZero {
				timeField.Set(db.Statement.Context, db.Statement.ReflectValue, now)
			}
		}

	}
}

func updateTimeStampWhileModify(db *gorm.DB) {
	if _, ok := db.Statement.Get("gorm:update_column"); !ok {
		db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
