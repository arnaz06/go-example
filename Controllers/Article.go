package Controllers

import(
	"../ApiHelpers"
	"../Models"
	"net/http"
	"fmt"
	"../Config"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)




func AllArticle(c *gin.Context){
	var article []Models.Article
	err := Config.DB.Find(&article).Error
	if err != nil{
		ApiHelpers.RespondJSON(c, 404, article)
	}
	ApiHelpers.RespondJSON(c, 200, article)
}

func CreateArticle(c *gin.Context) {
	file, err := c.FormFile("thumbnail")
	if err !=nil{
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	}
	if err:= c.SaveUploadedFile(file, "Public/"+file.Filename); err != nil{
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	}
	var article Models.Article
	article.Title= c.PostForm("title")
	article.Content= c.PostForm("content")
	baseUrlImage := os.Getenv("BASE_IMAGE_URL")
	article.Thumbnail= baseUrlImage+file.Filename
	err = Config.DB.Create(&article).Error
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

func UpdateArticle(c *gin.Context){
	file, err := c.FormFile("thumbnail")
	if err != nil{
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	}
	if err:= c.SaveUploadedFile(file, "Public/"+file.Filename); err != nil{
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	}
	id := c.Params.ByName("id")
	var article Models.Article
	err = Config.DB.Where("id = ?", id).First(&article).Error
	if err != nil{
		ApiHelpers.RespondJSON(c, 404, article)
	}
	article.Title= c.PostForm("title")
	article.Content= c.PostForm("content")
	baseUrlImage := os.Getenv("BASE_IMAGE_URL")
	article.Thumbnail= baseUrlImage+file.Filename
	err = Config.DB.Save(&article).Error
	if err != nil{
		ApiHelpers.RespondJSON(c,404, article)
	}else{
		ApiHelpers.RespondJSON(c,200, article)
	}
}

func DeleteArticle(c *gin.Context){
	id := c.Params.ByName("id")
	var article Models.Article
	err := Config.DB.Where("id = ?", id).Delete(&article).Error
	if err != nil{
		ApiHelpers.RespondJSON(c,404, article)
	}else{
		ApiHelpers.RespondJSON(c,200, gin.H{"id #"+id:"deleted"})
	}
}
