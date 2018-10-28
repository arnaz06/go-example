package Models

import (
	"github.com/jinzhu/gorm"
)



type Article struct{
	gorm.Model
	Title 		string	`json:"title"`
	Content 	string	`json:"content"`
	Thumbnail string	`json:"thumbnail"`
}

func (a *Article) TableName() string{
	return "articles"
}