package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"yyy/controller"
	_ "yyy/controller"
	model "yyy/models"
)

func init() {
	//core.StsrtServer()
}

func main() {
	//sec, err := config.Cfg.GetSection("database")
	//if err != nil {
	//	log.Fatal(2, "Fail to get section 'database': %v", err)
	//}
	//name := sec.Key("NAME")
	//log.Println(name)

	//db.Link()

	model.Insert()

	r := controller.R
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
