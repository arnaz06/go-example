package main

import(
	"./Config"
	"./Models"
	"./Routers"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var err error

func main(){
	err = godotenv.Load()
	if err != nil{
		fmt.Println("error load .env file")
	}
	urlDB := os.Getenv("DB_URL")
	portApp := os.Getenv("PORT_APP")
	Config.DB, err = gorm.Open("mysql", urlDB)
	if err != nil{
		fmt.Println("status: ",err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Article{})
	r:= Routers.SetupRouter()
	r.Run(portApp)
}

