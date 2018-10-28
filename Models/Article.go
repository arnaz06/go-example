package Models

import (
	"../Config"
	_ "github.com/go-sql-driver/mysql"
)




func GetAllArticles(a *[]Article) (err error){
	if err = Config.DB.Find(a).Error; err != nil{
		return err
	}
	return nil
}

func CreateArticle(a *Article) (err error) {
	if err = Config.DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}



