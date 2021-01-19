package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"yyy/controller"
	_ "yyy/controller"
	"yyy/tool"
)

func init() {

}

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	app := controller.R
	if err := app.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
