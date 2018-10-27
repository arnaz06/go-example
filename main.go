package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"./controller"
	"./model"
)


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage endpoint")
}



func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", controller.AllArticle).Methods("GET")
	myRouter.HandleFunc("/article/{title}", controller.Article).Methods("GET")
	myRouter.HandleFunc("/article/{title}", controller.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{title}", controller.UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/article", controller.NewArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8881", myRouter))	
}


func main(){
	fmt.Println("Starting APP in port 8881.......!")
	model.InitialArticleMigration()
	handleRequests()
}