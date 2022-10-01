package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	var err error
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("创建失败%v", err)
	}
	return db
}
