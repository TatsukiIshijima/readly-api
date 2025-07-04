package router

import (
	"github.com/gin-gonic/gin"
	"readly/image/server"
)

func Setup(
	authMiddleware gin.HandlerFunc,
	imageValidationMiddleware gin.HandlerFunc,
	imageServer server.ImageServer,
) *gin.Engine {
	router := gin.Default()

	root := router.Group("/")
	{
		v1 := root.Group("v1")
		{
			imgGroup := v1.Group("")
			imgGroup.Use(authMiddleware, imageValidationMiddleware)
			imgGroup.POST("/image/upload", imageServer.Upload)
		}
	}
	return router
}
