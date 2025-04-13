package router

import (
	"github.com/gin-gonic/gin"
	"readly/server"
)

func Setup(
	authMiddleware gin.HandlerFunc,
	imageServer server.ImageServer,
) *gin.Engine {
	router := gin.Default()

	root := router.Group("/")
	{
		v1 := root.Group("v1")
		{
			//v1.POST("/image/upload", imageServer.Upload).Use(authMiddleware)
			v1.POST("/image/upload", imageServer.Upload)
		}
	}
	return router
}
