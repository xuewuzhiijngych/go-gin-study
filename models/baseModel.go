package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//数据库配置
const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "golang"
)

var Db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	var err error
	Db, err = gorm.Open("mysql", dsn)

	//Db.SingularTable(true)
	//Db.LogMode(true)
	//Db.DB().SetMaxIdleConns(10)
	//Db.DB().SetMaxOpenConns(100)

	if err != nil {
		panic(err)
	}

	//defer Db.Close()
}
