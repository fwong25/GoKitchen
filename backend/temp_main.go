package main

import (
	"fmt"
	"strings"
)

type ingredient_t struct {
	IngredientName string
	Amount         float32
	Unit           string
}

type ingredient_cost_t struct {
	IngredientName string
	Amount         float32
	Cost           float32
}

var costs = []ingredient_cost_t{
	ingredient_cost_t{"Ripe bananas", 150, 10000},
	ingredient_cost_t{"Cake flour", 1000, 13000},
	ingredient_cost_t{"Butter", 200, 14000},
	ingredient_cost_t{"Sugar", 1000, 17500},
	ingredient_cost_t{"Eggs", 1, 2000},
	ingredient_cost_t{"Baking powder", 45, 8500},
	ingredient_cost_t{"Cocoa powder", 90, 39000},
	ingredient_cost_t{"Almond", 100, 14200},
	ingredient_cost_t{"Sugar powder", 500, 11000},
}

func scaleIngredients(ingredients []ingredient_t, scaledIngredient string, newAmount float32) (scaledIngredients []ingredient_t) {
	var scale float32
	for _, ingredient := range ingredients {
		if ingredient.IngredientName == scaledIngredient {
			scale = newAmount / ingredient.Amount
		}
	}

	for id, _ := range ingredients {
		ingredients[id].Amount = ingredients[id].Amount * scale
	}
	return ingredients
}

func printIngredients(title string, ingredients []ingredient_t) {
	var total_cost float32
	fmt.Println(strings.ToUpper(title))
	fmt.Println("======================================================")
	fmt.Printf("Ingredient           Weight          Cost\n")
	fmt.Println("======================================================")
	for _, ingredient := range ingredients {
		cost := getIngredientCost(ingredient)
		total_cost += cost
		amount_str := fmt.Sprintf("%.2f %s", ingredient.Amount, ingredient.Unit)
		fmt.Printf("%-20s%15s%10.2f\n", ingredient.IngredientName, amount_str, cost)
		// fmt.Println(ingredient.IngredientName, "\t", ingredient.Amount, " ", ingredient.Unit, "\t", cost)
	}
	fmt.Println("======================================================")
	fmt.Println("Total cost: ", total_cost)
}

func getIngredientCost(ingredient ingredient_t) (ingredient_cost float32) {
	for _, entry := range costs {
		if entry.IngredientName == ingredient.IngredientName {
			scale := ingredient.Amount / entry.Amount
			return entry.Cost * scale
		}
	}
	return 0
}

func main() {
	// ingredients := []ingredient_t{
	// 	{"Ripe bananas", 70, "g"},
	// 	{"Cake flour", 125, "g"},
	// 	{"Butter", 100, "g"},
	// 	{"Sugar", 80, "g"},
	// 	{"Eggs", 2, "pc"},
	// 	{"Baking powder", 4, "g"},
	// 	{"Banana slice", 10, "slices"},
	// }

	// soesIngredients := []ingredient_t{
	// 	{"Cake flour", 120, "g"},
	// 	{"Butter", 100, "g"},
	// 	{"Butter", 100, "g"},
	// 	{"Eggs", 3, "pc"},
	// 	{"Salt", 2, "g"},
	// 	{"Water", 240, "cc"},
	// 	{"Baking powder", 2, "g"},
	// }

	// flaIngredients := []ingredient_t{
	// 	{"Cake flour", 100, "g"},
	// 	{"Butter", 50, "g"},
	// 	{"Maizena", 40, "g"},
	// 	{"Eggs", 2, "pc"},
	// 	{"Sugar", 100, "g"},
	// 	{"Susu kental manis", 3, "sdm"},
	// 	{"Water", 500, "cc"},
	// }

	cookiesIngredients := []ingredient_t{
		{"Butter", 230, "g"},
		{"Sugar powder", 190, "g"},
		{"Eggs", 1, "pc"},
		{"Cocoa powder", 46, "g"},
		{"Cake flour", 385, "g"},
		{"Almond", 76, "g"},
	}

	// scaledIngredients := scaleIngredients(ingredients, "Ripe bananas", 150)
	// printIngredients("Banana cake", scaledIngredients)

	// printIngredients("Kue sus", soesIngredients)
	// printIngredients("Fla", flaIngredients)

	printIngredients("Choco almond cookies", cookiesIngredients)
}
