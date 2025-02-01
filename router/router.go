package router

import (
	"github.com/gin-gonic/gin"
	"readly/controller"
)

func Setup(bc controller.BookController, uc controller.UserController) *gin.Engine {
	router := gin.Default()

	readly := router.Group("/readly")
	{
		readly.POST("signup", uc.SignUp)
		readly.POST("signing", uc.SignIn)
		books := readly.Group("/books")
		{
			books.POST("", bc.Register)
			books.DELETE("", bc.Delete)
		}
	}

	return router
}
