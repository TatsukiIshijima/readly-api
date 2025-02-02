package router

import (
	"github.com/gin-gonic/gin"
	"readly/controller"
)

func Setup(
	authMiddleware gin.HandlerFunc,
	bc controller.BookController,
	uc controller.UserController,
) *gin.Engine {
	router := gin.Default()

	root := router.Group("/")
	{
		v1 := root.Group("v1")
		{
			v1.POST("/signup", uc.SignUp)
			v1.POST("/signin", uc.SignIn)

			books := v1.Group("/books").Use(authMiddleware)
			{
				books.POST("", bc.Register)
				books.DELETE("", bc.Delete)
			}
		}
	}
	return router
}
