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

	r.POST("/recipes", controllers.RecipeCreate)
	r.GET("/recipes", controllers.RecipeIndex)
	r.GET("/recipes/:id", controllers.RecipeShow)
	
	r.Run()

}
