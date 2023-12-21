package route

import (
	"github.com/abhis3110/go-todoApp/auth"
	"github.com/abhis3110/go-todoApp/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func SetupRoutes() *gin.Engine{
	router:=gin.Default()
	authMiddleware, err:=auth.SetUpAuth()

	if err!=nil{
		log.Fatal("JWT error"+err.Error())
	}

	router.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"Welcome to my todo app")
	})

	v1:=router.Group("/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)

		v1.POST("/register", controller.RegisterEndPoint)

		todo:=v1.Group("todo")
		{
			todo.POST("/create", authMiddleware.MiddlewareFunc(),controller.CreatTask)
			todo.GET("/all", authMiddleware.MiddlewareFunc(),controller.FetchAllTask)
			todo.GET("/get/:id", authMiddleware.MiddlewareFunc(),controller.FetchSingleTask)
			todo.PUT("/update/:id", authMiddleware.MiddlewareFunc(),controller.UpdateTask)
			todo.DELETE("/delete/:id", authMiddleware.MiddlewareFunc(),controller.DeleteTask)
		}
	}
	authorization:=router.Group("/auth")
	authorization.GET("/refresh_token", authMiddleware.RefreshHandler)
	return router
}