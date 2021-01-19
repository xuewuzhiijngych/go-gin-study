package model

type Admin struct {
	username string `gorm:"not null;unique"`
	password string `gorm:"not null;unique"`
	status   uint   `gorm:"default:1"`
}

func Insert() {
	admin := Admin{username: "Jinzhu", password: "ycycycy"}
	DB.Create(&admin)
}
