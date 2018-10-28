package controller

import(
	"../ApiHelpers"
	"../Models"
	"fmt"
	"github.com/gin-gonic/gin"
)




func AllArticle(c *gin.Context){
	var article []Models.Article
	err := Models.GetAllArticles(&article)
	if err != nil{
		ApiHelpers.RespondJSON(c, 404, article)
	}
	ApiHelpers.RespondJSON(c, 200, article)
}

func CreateArticle(c *gin.Context) {
	var article Models.Article
	fmt.Println(c.PostForm("title"))
	article.Title= c.PostForm("title")
	article.Content= c.PostForm("content")
	article.Thumbnail= c.PostForm("thumbnail")
	err := Models.CreateArticle(&article)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, article)
	} else {
		ApiHelpers.RespondJSON(c, 201, article)
	}
}



// func Article(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Article Endpoint Hit")
// }

// func UpdateArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Update Article Endpoint Hit")
// }

// func DeleteArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Delete Article Endpoint Hit")
// }


// func SearchArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Search Article Endpoint Hit")
// }
