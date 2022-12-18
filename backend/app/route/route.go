package routes

import (
	"test-sharing-vision/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/article/:id", controllers.GetArticleById)
	r.GET("/articles/:limit/:offset", controllers.GetArticles)
	r.DELETE("/article/:id", controllers.DeleteArticleById)
	r.PUT("/article/", controllers.UpdateArticleById)
	r.POST("/article/", controllers.CreateArticle)
	return r
}
