package controllers

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/kreely/orm-gin/initialisers"
	"github.com/kreely/orm-gin/models"


)

func RecipeIndex(c *gin.Context) {
	var recipes []models.Recipe
	initialisers.DB.Find(&recipes)

	c.JSON(200, gin.H{"recipes": recipes })
}

func RecipeShow(c *gin.Context) {
	// Get ID of recipe.
	id := c.Param("id")
	var recipe []models.Recipe
	initialisers.DB.First(&recipe, id)

	c.JSON(200, gin.H{"recipe": recipe })
}

func RecipeUpdate(c *gin.Context) {

}

func RecipeCreate(c *gin.Context) {
	// Get data off req body
	
	type Ingredient struct {
		Amount uint `json:"Amount" binding:"required"`
		Description string  `json:"Description" binding:"required"`
	}

	type Recipe struct {
		Name string `json:"Name"`
		Servings uint `json:"Servings"`
		Ingredients []models.RecipeIngredient `form:"Ingredients" binding:"dive"`
		Steps []models.RecipeStep `form:"Steps" binding:"dive"`
	}

	var recipe Recipe

	c.Bind(&recipe)
	//	log.Println(recipe)

	//log.Println(recipe.Ingredients)
	//log.Println("Name ", recipe.Name, "Servings ", recipe.Servings)

	// Populate the post variable with the values bound from JSON...
	post := models.Recipe{Name: recipe.Name,
		Servings: recipe.Servings,
		Ingredients: recipe.Ingredients,
		Steps: recipe.Steps,
	}
	// Post the record.
	result := initialisers.DB.Create(&post)

	if result.Error != nil {
		log.Println(result.Error)
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"recipe": post })
	
}
