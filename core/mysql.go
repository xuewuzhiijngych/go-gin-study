package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//数据库配置
const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "hr"
)

var Db *gorm.DB

// init 初始化mysql 连接
func init() {
	sec, errs := Cfg.GetSection("database")
	if errs != nil {
		log.Fatal(2, "Fail to get section 'database': %v", errs)
	}

	user := sec.Key("USER")
	password := sec.Key("PASSWORD")
	host := sec.Key("HOST")
	name := sec.Key("NAME")
	network := sec.Key("NETWORK")
	port := sec.Key("PORT")

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, password, network, host, port, name)
	//dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/hr")
	db, err := gorm.Open("mysql", dsn)
	db.SingularTable(true)
	defer db.Close()

	if err != nil {
		panic(err)
	}
}
