package controller

import (
	"fmt"
	"net/http"
	"../model"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error



func Ping(){
	db, err= gorm.Open("mysql", "root:laskar45@/articles")
	
		if err != nil{
			panic("Faild to Connect to DB")
		}
	defer db.Close()
}

func AllArticle(w http.ResponseWriter, r *http.Request){
	Ping()
	var articles []model.Article
	db.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

func NewArticle(w http.ResponseWriter, r *http.Request){
	Ping()
	// decoder := json.NewDecoder(r.Body)
	var article model.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil{
		panic(err)
	}
	if err:= db.Create(&article); err != nil{
		panic(err)
	}
	fmt.Println(Article)
}

func Article(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Article Endpoint Hit")
}

func UpdateArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update Article Endpoint Hit")
}

func DeleteArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete Article Endpoint Hit")
}
// func SearchArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Search Article Endpoint Hit")
// }
