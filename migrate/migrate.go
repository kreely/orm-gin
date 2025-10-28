package main

import (
	"log"
	"github.com/kreely/orm-gin/initialisers"
	"github.com/kreely/orm-gin/models"
)

func init() {
	initialisers.LoadEnvVariables() 
	initialisers.ConnectToDB("/data/test.db")
}

func main() {
	// Creating this one constraint seems to initialise both the steps and ingredients
	// constraint.
	log.Println("About to migrate...")

	initialisers.DB.AutoMigrate(&models.Recipe{})	
	initialisers.DB.AutoMigrate(&models.RecipeStep{})	
	initialisers.DB.AutoMigrate(&models.RecipeIngredient{})

	err := initialisers.DB.Migrator().CreateConstraint(&models.Recipe{}, "fk_recipe_recipe_steps")
	if err != nil {
		log.Fatal("Unable to setup constraints!!!", err)
	}
	
}
