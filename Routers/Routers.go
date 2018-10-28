package Routers

import(
	"../Controllers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r:= gin.Default()
	r.Use(static.Serve("/image", static.LocalFile("Public", false)))
	v1 := r.Group("/v1")
	{
		v1.GET("articles", controller.AllArticle)
		v1.GET("article/:id", controller.Article)
		v1.DELETE("/article/:id", controller.DeleteArticle)
		v1.PUT("/article/:id", controller.UpdateArticle)
		v1.POST("article", controller.CreateArticle)
	}
	return r
}