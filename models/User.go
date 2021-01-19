package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday int
	Email    string `gorm:"type:varchar(100)"`
}

func UserCreate() {
	Db.AutoMigrate(&User{})
}

func GetUserByID(id uint) *User {
	var user = new(User)
	user.ID = id
	Db.First(user)
	return user
}

func UserAdd(user *User) {
	Db.Create(&user)
}

func main() {

}
