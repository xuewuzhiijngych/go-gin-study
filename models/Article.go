package model

import "github.com/jinzhu/gorm"

// Article 文章
type Article struct {
	gorm.Model
	ID      int    `gorm:"type:uint(10);not null;AUTO_INCREMENT"`
	Title   string `gorm:"type:varchar(100);not null"`
	CateId  int    `gorm:"not null"`
	Content string `gorm:"type:text(0);not null"`
}

func ArticleCreate() {
	Db.AutoMigrate(&Article{})
}
