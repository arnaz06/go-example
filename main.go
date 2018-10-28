package main

import(
	"./Config"
	"./Models"
	"./Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main(){
	Config.DB, err = gorm.Open("mysql", "root:laskar45@tcp(127.0.0.1:3306)/articles?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("status: ",err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Article{})
	r:= Routers.SetupRouter()
	r.Run(":8081")
}

