package config

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error


func Ping(){
	db, err= gorm.Open("mysql", "root:laskar45@/articles")
	
		if err != nil{
			fmt.Println(err.Error())
			panic("Faild to Connect to DB")
		}

		defer db.Close()

}