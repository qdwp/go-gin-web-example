package route

import (
	"example.com/api"
	v1 "example.com/api/v1"
	"example.com/interceptor"

	"github.com/gin-gonic/gin"
)

var RouterEngine *gin.Engine

func GetRouterEngine() *gin.Engine {
	if RouterEngine == nil {
		RouterEngine = InitRouterEngine()
	}
	return RouterEngine
}

func InitRouterEngine() *gin.Engine {

	router := gin.New()
	// router.Use(gin.Logger())
	router.Use(gin.Recovery())

	root := router.Group("/")
	root.Use(interceptor.CorsInterceptor())
	root.Use(interceptor.ResultInterceptor())

	root.GET("/", api.Welcome)
	root.GET("/ping", api.Ping)

	v1group := root.Group("/v1")
	v1group.POST("/post", v1.PostCall)

	return router
}
