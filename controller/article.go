package controller

import (
	"fmt"
	"net/http"
	"github.com/arnaz06/go-example/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func AllArticle(w http.ResponseWriter, r *http.Request){
	config.ping()


}

func NewArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "New Article Endpoint Hit")
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
