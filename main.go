package main

import (
	"github.com/kreely/orm-gin/initialisers"
	"github.com/kreely/orm-gin/controllers"	
	"github.com/gin-gonic/gin"

)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB("test.db")
}

func main() {

	r := gin.Default()

	r.GET("/", controllers.PostsCreate)

	r.Run()

}
