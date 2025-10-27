package controllers

import (
	"log"
	"context"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/kreely/orm-gin/initialisers"
	"github.com/kreely/orm-gin/models"
	//	"encoding/json"
)

func RecipeIndex(c *gin.Context) {
	// Get back all of the recipes.
	var recipes []models.Recipe
	initialisers.DB.Find(&recipes)

	c.JSON(200, gin.H{"recipes": recipes })
}

func RecipeShow(c *gin.Context) {
	// Get back the data for one recipe.
	
	// Get ID of recipe.
	recipe_id := c.Param("id")
	var recipe models.Recipe
	c.Bind(&recipe)
	initialisers.DB.First(&recipe, recipe_id)
	initialisers.DB.Where("recipe_steps.recipe_id = ?", recipe_id).Find(&recipe.Steps)
	initialisers.DB.Where("recipe_ingredients.recipe_id = ?", recipe_id).Find(&recipe.Ingredients)
	log.Println("Recipe_id: ", recipe_id)
	//	b, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	log.Println("Recipe ", recipe)
	log.Println("This is recipe.Steps... ", recipe.Steps)


	c.JSON(200, gin.H{"recipe": recipe })
}

func RecipeUpdate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Get ID of recipe.
	recipe_id := c.Param("id")
	
	var recipe models.Recipe
		
	// Find these so that they can be "deleted" when the recipe is updated.
	// GORM doesn't actually delete them.. they get marked as deleted.
	//	var recipeSteps []models.RecipeStep
	//	var recipeIngredients []models.RecipeIngredient

	c.Bind(&recipe)
	initialisers.DB.First(&recipe, recipe_id)
	// Delete by ID
	res, err := gorm.G[models.RecipeStep](initialisers.DB).
		Where("recipe_id = ?", recipe_id).Delete(ctx)
	log.Println("res: ", res, "err: ", err)
	
	res2, err2 := gorm.G[models.RecipeIngredient](initialisers.DB).
		Where("recipe_id = ?", recipe_id).Delete(ctx)
	log.Println("res2: ", res2, "err2: ", err2)
	
	initialisers.DB.Model(&recipe).Updates(recipe)
	initialisers.DB.Model(&recipe.Steps).Create(recipe.Steps)
	initialisers.DB.Model(&recipe.Ingredients).Create(recipe.Ingredients)

	c.JSON(200, gin.H{
		"recipe": recipe,
	})
	
}

func RecipeCreate(c *gin.Context) {
	var recipe models.Recipe
	c.Bind(&recipe)
	//	log.Println(recipe)
	
	// Post the recipe.
	result := initialisers.DB.Create(&recipe)

	if result.Error != nil {
		log.Println(result.Error)
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"recipe": recipe })
}
