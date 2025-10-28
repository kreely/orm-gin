package main

import (
	"github.com/kreely/orm-gin/initialisers"
	"github.com/kreely/orm-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/kreely/orm-gin/utilities"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB("test.db")
	utilities.printFiles()
}

func main() {

	r := gin.Default()

	r.POST("/recipes", controllers.RecipeCreate)
	r.PUT("/recipes/:id", controllers.RecipeUpdate)
	r.GET("/recipes", controllers.RecipeIndex)
	r.GET("/recipes/:id", controllers.RecipeShow)
	
	r.Run()

}
