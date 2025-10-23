package main

import (
	"github.com/kreely/orm-gin/initialisers"
	"github.com/kreely/orm-gin/models"
)

func init() {
	initialisers.LoadEnvVariables() 
	initialisers.ConnectToDB("test.db")
}

func main() {
	initialisers.DB.Migrator().CreateConstraint(&models.Recipe{}, "fk_recipe_recipe_steps")
	
	initialisers.DB.AutoMigrate(&models.RecipeSteps{})	
	initialisers.DB.AutoMigrate(&models.RecipeIngredient{})
	initialisers.DB.AutoMigrate(&models.Recipe{})	
}
