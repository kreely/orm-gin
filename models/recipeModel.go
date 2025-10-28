package models

import "gorm.io/gorm"

type RecipeIngredient struct {
	gorm.Model
	RecipeID int
	Amount float32
	Unit string // cup, bag, teaspoon, bag, etc...
	Description string
}

type Recipe struct {
	gorm.Model
	Name string
	Servings uint
	Steps []RecipeStep `gorm:"foreignKey:RecipeID"`
	Ingredients []RecipeIngredient `gorm:"foreignKey:RecipeID"`
}

type RecipeStep struct {
	gorm.Model
	RecipeID int
	StepNumber uint
	DirectionText string
}


