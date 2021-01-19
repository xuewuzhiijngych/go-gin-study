package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"yyy/tool"
)

var Db *gorm.DB

func init() {
	var cfg *tool.Config
	var err error

	cfg, err = tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	//数据库配置
	var (
		username = cfg.DbUserName
		password = cfg.DbPassWord
		network  = cfg.NETWORK
		server   = cfg.SERVER
		port     = cfg.DbPORT
		database = cfg.DataBase
	)

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, network, server, port, database)
	Db, err = gorm.Open("mysql", dsn)
	Db.SingularTable(true)

	if err != nil {
		panic(err)
	}
	//defer Db.Close()
}
