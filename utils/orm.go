package utils

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

var (
	Orm *gorm.DB
)

func SetupOrm() {
	var err error

	Orm, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetConfig("mysql::user"),
		GetConfig("mysql::password"),
		GetConfig("mysql::host"),
		GetConfig("mysql::db")))

	if err != nil {
		Log.Panicln(err)
	}
	// 自动创建/修改表结构
	//Orm.AutoMigrate(&models.Users{})
}
