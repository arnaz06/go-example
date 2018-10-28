package controller

import(
	"../ApiHelpers"
	"../Models"
	// "net/http"
	"fmt"
	"../Config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)




func AllArticle(c *gin.Context){
	var article []Models.Article
	err := Config.DB.Find(article).Error
	if err != nil{
		ApiHelpers.RespondJSON(c, 404, article)
	}
	ApiHelpers.RespondJSON(c, 200, article)
}

func CreateArticle(c *gin.Context) {
	var article Models.Article
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("get form err : 	%s",err.Error()))
	// 	return 
	// }
	// Files := form.File["files"]
	// fmt.Println(Files)
	
	// file := c.FormFile("thumbnail")
	// fmt.Println(file)
	article.Title= c.PostForm("title")
	article.Content= c.PostForm("content")
	article.Thumbnail= c.PostForm("thumbnail")
	err := Config.DB.Create(&article).Error
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, article)
	} else {
		ApiHelpers.RespondJSON(c, 201, article)
	}
}

func Article(c *gin.Context){
	id := c.Params.ByName("id")
	var article Models.Article
	err := Config.DB.Where("id = ?",id).First(&article).Error
	fmt.Println(article)
	if err != nil{
		ApiHelpers.RespondJSON(c,404, article)
	}else{
		ApiHelpers.RespondJSON(c,200, article)
	}
}



// func UpdateArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Update Article Endpoint Hit")
// }

// func DeleteArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Delete Article Endpoint Hit")
// }


// func SearchArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Search Article Endpoint Hit")
// }
