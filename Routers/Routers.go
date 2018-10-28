package Routers

import(
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r:= gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("articles", controller.AllArticle)
		// r.GET("/article/:title", controller.Article)
		// r.DELETE("/article/:title", controller.DeleteArticle)
		// r.PUT("/article/:title", controller.UpdateArticle)
		v1.POST("article", controller.CreateArticle)
	}
	return r
}