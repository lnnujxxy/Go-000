package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gogeekbang/internal/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(databaseSetting config.Database) *gorm.DB {
	db, err := gorm.Open(databaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Database,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		panic(fmt.Sprintf( "mysql:NewDB databaseSetting = %+v", databaseSetting))
	}

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db
}
