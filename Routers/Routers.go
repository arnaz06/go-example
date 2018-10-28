package Routers

import(
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r:= gin.Default()
	r.Static("/image","../Public")
	v1 := r.Group("/v1")
	{
		v1.GET("articles", controller.AllArticle)
		v1.GET("article/:id", controller.Article)
		// r.DELETE("/article/:title", controller.DeleteArticle)
		// r.PUT("/article/:title", controller.UpdateArticle)
		v1.POST("article", controller.CreateArticle)
	}
	return r
}