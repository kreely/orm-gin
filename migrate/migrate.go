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
	// Creating this one constraint seems to initialise both the steps and ingredients
	// constraint.
	initialisers.DB.Migrator().CreateConstraint(&models.Recipe{}, "fk_recipe_recipe_steps")
	
	initialisers.DB.AutoMigrate(&models.RecipeStep{})	
	initialisers.DB.AutoMigrate(&models.RecipeIngredient{})
	initialisers.DB.AutoMigrate(&models.Recipe{})	
}
