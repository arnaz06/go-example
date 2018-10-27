package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Article struct{
	ID int64	`gorm:"not null" json:"id"`
	Title string	`gorm:"not null" json:"title"`
	Content string	`gorm:"not null" json:"content"`
	Thumbnail string	`gorm:"not null" json:"thumbnail"`

}

func InitialArticleMigration(){
		db, err= gorm.Open("mysql", "root:laskar45@/articles")
	
		if err != nil{
			fmt.Println(err.Error())
			panic("Faild to Connect to DB")
		}

	defer db.Close()
	db.AutoMigrate(&Article{})
}
