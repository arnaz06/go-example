package controller

import (
	"fmt"
	"net/http"
)

func AllArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "All Article Endpoint Hit")
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
