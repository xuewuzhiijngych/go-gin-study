package controller

import (
	"github.com/jinzhu/gorm"
	"yyy/router"
)

var R = router.SetupRouter()
var Db *gorm.DB
