package models

import "fmt"

type Recipe struct {
	RecipeID         int    `db:"recipeID"`
	Title            string `db:"title"`
	Steps            string `db:"steps"`
	ParentRecipeID   int    `db:"parentRecipeID"`
	IsRescaled       bool   `db:"isRescaled"`
	CreatedDate      string `db:"createdDate"`
	LastModifiedDate string `db:"lastModifiedDate"`
}

type OutputRecipe struct {
	RecipeID         int    `db:"recipeID"`
	Title            string `db:"title"`
	Steps            string `db:"steps"`
	ParentRecipeID   int    `db:"parentRecipeID"`
	IsRescaled       bool   `db:"isRescaled"`
	CreatedDate      string `db:"createdDate"`
	LastModifiedDate string `db:"lastModifiedDate"`
	SubrecipeList    []Recipe
}

type IngredientCategory struct {
	CategoryID   int    `db:"categoryID"`
	RecipeID     int    `db:"recipeID"`
	CategoryName string `db:"categoryName"`
}

type RecipeIngredient struct {
	RecipeIngredientID int    `db:"recipeIngredientID"`
	CategoryID         int    `db:"categoryID"`
	RecipeID           int    `db:"recipeID"`
	IngredientID       int    `db:"ingredientID"`
	Amount             string `db:"amount"`
	AmountUnit         string `db:"amountUnit"`
}

type IngredientInfo struct {
	IngredientID   int    `db:"ingredientID"`
	IngredientName string `db:"ingredientName"`
	Cost           int    `db:"cost"`
	Amount         int    `db:"amount"`
	AmountUnit     string `db:"amountUnit"`
}

type RescaleInfo struct {
	RecipeID     int `db:"recipeID"`
	IngredientID int `db:"ingredientID"`
	NewAmount    int `db:"newAmount"`
}

type InputRecipeIngredient struct {
	IngredientName string `db:"ingredientName"`
	Amount         string `db:"amount"`
	AmountUnit     string `db:"amountUnit"`
}

type InputIngredient struct {
	CategoryName string
	Ingredients  []InputRecipeIngredient
}

type OutputRecipeIngredient struct {
	RecipeIngredientID int
	CategoryID         int
	IngredientID       int
	IngredientName     string
	Amount             string
	AmountUnit         string
	Cost               float64
}

type OutputIngredient struct {
	CategoryName string
	Ingredients  []OutputRecipeIngredient
	TotalCost    float64
}

type InputRecipe struct {
	Title            string
	Steps            string
	ParentRecipeID   int
	IsRescaled       bool
	InputIngredients []InputIngredient
	RescaleInfo      RescaleInfo
}

type InputRecipeUpdate struct {
	RecipeID         int
	Title            string
	Steps            string
	IsRescaled       bool
	InputIngredients []InputIngredient
	RescaleInfo      RescaleInfo
}

type InputIngredientInfo struct {
	IngredientName string `db:"ingredientName"`
	Cost           int    `db:"cost"`
	Amount         int    `db:"amount"`
	AmountUnit     string `db:"amountUnit"`
}

type RecipeDetail struct {
	RecipeID          int
	Title             string
	Steps             string
	ParentRecipeID    int
	IsRescaled        bool
	CreatedDate       string
	LastModifiedDate  string
	Ingredients       []OutputIngredient
	RecipeRescaleInfo RescaleInfo
	TotalCost         float64
}

type conversionInfo struct {
	SourceUnit string
	TargetUnit string
	Scale      int
}

var Conversion_info = []conversionInfo{
	{"kg", "g", 1000},
	{"tsp", "g", 5},
	{"tsp", "ml", 5},
	{"tbsp", "ml", 15},
	{"tbsp", "g", 15},
	{"g", "gr", 1},
	{"ml", "cc", 1},
	{"g", "cc", 1},
}

func PrintRecipeDetail(recipe_detail RecipeDetail) {
	fmt.Println("RecipeID:", recipe_detail.RecipeID)
	fmt.Println("Title: ", recipe_detail.Title)
	fmt.Printf("Steps: \n%s\n", recipe_detail.Steps)
	fmt.Println("ParentRecipeID: ", recipe_detail.ParentRecipeID)
	fmt.Println("IsRescaled: ", recipe_detail.IsRescaled)
	fmt.Println("CreatedDate: ", recipe_detail.CreatedDate)
	fmt.Println("LastModifiedDate: ", recipe_detail.LastModifiedDate)
	PrintIngredients(recipe_detail.Ingredients)
	if recipe_detail.IsRescaled {
		PrintRecipeRescaleInfo(recipe_detail.RecipeRescaleInfo)
	}
}

func PrintIngredients(recipe_ingredients []OutputIngredient) {
	fmt.Println("Ingredients:")
	fmt.Println("=========================")
	for _, ingredient_group := range recipe_ingredients {
		fmt.Println("Category Name:", ingredient_group.CategoryName)
		fmt.Println("=========================")
		for _, recipe_ingredient := range ingredient_group.Ingredients {
			fmt.Println("RecipeIngredientID:", recipe_ingredient.RecipeIngredientID)
			fmt.Println("IngredientID:", recipe_ingredient.IngredientID)
			fmt.Println("IngredientName: ", recipe_ingredient.IngredientName)
			fmt.Println("Amount: ", recipe_ingredient.Amount)
			fmt.Println("AmountUnit: ", recipe_ingredient.AmountUnit)
			fmt.Printf("Cost: %.2f\n", recipe_ingredient.Cost)
			fmt.Println("=========================")
		}
	}
}

func PrintRecipeRescaleInfo(rescale_info RescaleInfo) {
	fmt.Println("RescaleInfo:")
	fmt.Println("=========================")
	fmt.Println("RecipeID:", rescale_info.RecipeID)
	fmt.Println("IngredientID:", rescale_info.IngredientID)
	fmt.Println("NewAmount:", rescale_info.NewAmount)
	fmt.Println("=========================")
}

func PrintRecipeList(recipe_list []Recipe) {
	fmt.Println("RECIPES")
	fmt.Println("=========================")
	for _, recipe := range recipe_list {
		fmt.Println("RecipeID:", recipe.RecipeID)
		fmt.Println("Title:", recipe.Title)
		fmt.Println("Steps:", recipe.Steps)
		fmt.Println("ParentRecipeID: ", recipe.ParentRecipeID)
		fmt.Println("IsRescaled: ", recipe.IsRescaled)
		fmt.Println("CreatedDate: ", recipe.CreatedDate)
		fmt.Println("LastModifiedDate: ", recipe.LastModifiedDate)
		fmt.Println("=========================")
	}
}
