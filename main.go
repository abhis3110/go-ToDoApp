package main

import (
	"log"
	"os"

	"github.com/abhis3110/go-todoApp/config"
	"github.com/abhis3110/go-todoApp/migration"
	//"github.com/abhis3110/go-todoApp/model"
	"github.com/abhis3110/go-todoApp/route"
	"github.com/gin-gonic/gin"
)

func init() {
	db := config.CreatePostgraceDB()
	migration.Migrate(db)
	//db.AutoMigrate(&model.Task{})
	//db.AutoMigrate(&model.User{})
}

func main(){
	gin.SetMode(gin.ReleaseMode)

	router:=route.SetupRoutes()

	port:=os.Getenv("PORT")
	if port =="" {
		port="8080"
	}

	if err := router.Run(":"+port);err!=nil{
		log.Panicf("error: %s", err)
	}
}