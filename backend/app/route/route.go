package routes

import (
	"test-sharing-vision/app/controllers"
	cache "test-sharing-vision/platform/cache"
	"time"

	cachex "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	rdb, _ := cache.RedisConnection()
	redisStore := persist.NewRedisStore(rdb)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Group("/api/v1")
	r.GET("/articles", controllers.GetArticles)
	r.GET("/articles/:id", cachex.CacheByRequestURI(redisStore, 5*time.Minute), controllers.GetArticleById)
	r.DELETE("/articles/:id", controllers.DeleteArticleById)
	r.PUT("/articles/:id", controllers.UpdateArticleById)
	r.POST("/articles", controllers.CreateArticle)
	return r
}
