// package main

package db_mgmt

import (
	"database/sql"
	"fmt"
	"models"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// func main() {
// 	ConnectDB()
// 	// insertIngredientInfo("Egg", 2000, 1, "pcs")

// 	// parent_recipe_id := 0
// 	// title := "Banana pound cake"
// 	// steps := "1. Mix butter and sugar \n" +
// 	// 	"2. Add in egg mixture little by little \n" +
// 	// 	"3. Add in flour and baking soda \n" +
// 	// 	"4. Add in mashed banana \n" +
// 	// 	"5. After all ingredients are mixed well, pour into baking mold \n"
// 	// is_rescaled := false
// 	// input_recipe_ingredients := []models.InputRecipeIngredient{
// 	// 	{IngredientName: "Banana", Amount: "70", AmountUnit: "g"},
// 	// 	{IngredientName: "Cake Flour", Amount: "125", AmountUnit: "g"},
// 	// 	{IngredientName: "Unsalted butter", Amount: "100", AmountUnit: "g"},
// 	// 	{IngredientName: "Sugar", Amount: "80", AmountUnit: "g"},
// 	// 	{IngredientName: "Egg", Amount: "2", AmountUnit: "pcs"},
// 	// 	{IngredientName: "Baking powder", Amount: "1", AmountUnit: "tsp"},
// 	// }
// 	// rescale_info := models.RescaleInfo{
// 	// 	RecipeID:     0,
// 	// 	IngredientID: 0,
// 	// 	NewAmount:    0,
// 	// }

// 	// recipe_id := InsertRecipe(parent_recipe_id, title, steps, is_rescaled, input_recipe_ingredients, rescale_info)

// 	// ========================================================================

// 	// parent_recipe_id := 7
// 	// title := "Banana pound cake"
// 	// steps := "1. Mix butter and sugar \n" +
// 	// 	"2. Add in egg mixture little by little \n" +
// 	// 	"3. Add in flour and baking soda \n" +
// 	// 	"4. Add in mashed banana \n" +
// 	// 	"5. After all ingredients are mixed well, pour into baking mold \n"
// 	// is_rescaled := true
// 	// input_recipe_ingredients := []models.InputRecipeIngredient{
// 	// 	{IngredientName: "Banana", Amount: "70", AmountUnit: "g"},
// 	// 	{IngredientName: "Cake Flour", Amount: "125", AmountUnit: "g"},
// 	// 	{IngredientName: "Unsalted butter", Amount: "100", AmountUnit: "g"},
// 	// 	{IngredientName: "Sugar", Amount: "80", AmountUnit: "g"},
// 	// 	{IngredientName: "Egg", Amount: "2", AmountUnit: "pcs"},
// 	// 	{IngredientName: "Baking powder", Amount: "1", AmountUnit: "tsp"},
// 	// }
// 	// rescale_info := models.RescaleInfo{
// 	// 	RecipeID:     2,
// 	// 	IngredientID: 17,
// 	// 	NewAmount:    1,
// 	// }

// 	// recipe_id := InsertRecipe(parent_recipe_id, title, steps, is_rescaled, input_recipe_ingredients, rescale_info)

// 	// recipe_detail := GetRecipeDetail(recipe_id)
// 	// models.PrintRecipeDetail(recipe_detail)

// 	// recipe_list := GetRecipeList()
// 	// models.PrintRecipeList(recipe_list)

// 	// DeleteRecipe(11)

// 	CloseDB()
// }

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "12345678"
	dbname   = "GoKitchenDB"
)

var db *sql.DB
var err error

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error when connecting to database", err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error when pinging database", err)
	}

	fmt.Println("Successfully connected to database!")
}

func CloseDB() {
	db.Close()
}

func GetRecipeDetail(recipe_id int) models.RecipeDetail {
	recipe := models.RecipeDetail{} // Initialize struct to hold retrieved data

	queryStat := "SELECT \"recipeID\", title, steps, \"parentRecipeID\", \"isRescaled\", \"createdDate\", \"lastModifiedDate\" FROM recipes WHERE \"recipeID\" = " + strconv.Itoa(recipe_id) + ";"
	// fmt.Println(queryStat)

	rows, _ := db.Query(queryStat)
	defer rows.Close()

	if rows == nil {
		fmt.Println("Can't retrieve note with ID: ", recipe_id)
		return recipe
	}

	rows.Next()
	err := rows.Scan(&recipe.RecipeID, &recipe.Title, &recipe.Steps, &recipe.ParentRecipeID, &recipe.IsRescaled, &recipe.CreatedDate, &recipe.LastModifiedDate) // Scan the current row into the "place" variable
	if err != nil {
		fmt.Println("Error when scanning rows", err)
	}

	if recipe.IsRescaled {
		fmt.Println("recipe is rescaled")
		recipe.RecipeRescaleInfo = getRescaleInfo(recipe.RecipeID)
		fmt.Println("getRescaleInfo done")
		recipe.Ingredients, recipe.TotalCost = getRescaledIngredients(recipe.ParentRecipeID, recipe.RecipeRescaleInfo)
		fmt.Println("getRescaledIngredients done")
	} else {
		recipe.Ingredients, recipe.TotalCost = getRecipeIngredients(recipe.RecipeID)
	}

	return recipe
}

func GetRecipeList() []models.OutputRecipe {
	place := models.OutputRecipe{} // Initialize a User struct to hold retrieved data
	queryStat := "SELECT \"recipeID\", title, steps, \"parentRecipeID\", \"isRescaled\", \"createdDate\", \"lastModifiedDate\" FROM recipes WHERE \"parentRecipeID\" = 0 ORDER BY \"recipeID\""

	rows, _ := db.Query(queryStat)
	defer rows.Close()

	recipe_list := []models.OutputRecipe{}

	for rows.Next() {
		err := rows.Scan(&place.RecipeID, &place.Title, &place.Steps, &place.ParentRecipeID, &place.IsRescaled, &place.CreatedDate, &place.LastModifiedDate) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}

		// get ingredients
		if isSubrecipeExists(place.RecipeID) {
			fmt.Println("In getRecipeList(): Get subnotes for note", place.RecipeID)
			subrecipe_list := getSubrecipes(place.RecipeID)
			// recipe_list = append(subrecipe_list, recipe_list...)
			place.SubrecipeList = subrecipe_list
		} else {
			place.SubrecipeList = []models.Recipe{}
		}

		recipe_list = append([]models.OutputRecipe{place}, recipe_list...) // ... unpack into note_list[0], note_list[1], ...
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return recipe_list
}

func isSubrecipeExists(parent_recipe_id int) (is_subnote_exists bool) {
	queryStat := "SELECT EXISTS (SELECT 1 FROM recipes WHERE \"parentRecipeID\" = " + strconv.Itoa(parent_recipe_id) + ")"

	rows, _ := db.Query(queryStat)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&is_subnote_exists) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
		fmt.Println("subnote of note ", parent_recipe_id, ":", is_subnote_exists)
	}
	return is_subnote_exists
}

func getSubrecipes(parent_recipe_id int) []models.Recipe {
	place := models.Recipe{}
	queryStat := "SELECT \"recipeID\", title, steps, \"parentRecipeID\", \"isRescaled\", \"createdDate\", \"lastModifiedDate\" FROM recipes WHERE \"parentRecipeID\" = " + strconv.Itoa(parent_recipe_id) + " ORDER BY \"recipeID\""

	rows, _ := db.Query(queryStat)
	defer rows.Close()

	recipe_list := []models.Recipe{}

	for rows.Next() {
		err := rows.Scan(&place.RecipeID, &place.Title, &place.Steps, &place.ParentRecipeID, &place.IsRescaled, &place.CreatedDate, &place.LastModifiedDate) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
		// todo: get ingredients, if isRescaled then get rescaled ingredients
		recipe_list = append([]models.Recipe{place}, recipe_list...)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return recipe_list
}

func getRescaleInfo(recipe_id int) (rescale_info models.RescaleInfo) {
	queryStat := "SELECT \"recipeID\", \"ingredientID\", \"newAmount\" FROM \"rescaleInfo\" WHERE \"recipeID\" = " + strconv.Itoa(recipe_id)
	rows, _ := db.Query(queryStat)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&rescale_info.RecipeID, &rescale_info.IngredientID, &rescale_info.NewAmount) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return rescale_info
}

func getIngredientCategories(recipe_id int) (ingredient_categories []models.IngredientCategory) {
	place := models.IngredientCategory{}
	queryStat := "SELECT \"categoryID\", \"recipeID\", \"categoryName\"  FROM \"ingredientCategory\" WHERE \"recipeID\" = " + strconv.Itoa(recipe_id) + " ORDER BY \"categoryID\""

	rows, _ := db.Query(queryStat)
	defer rows.Close()

	// ingredient_list := []models.OutputRecipeIngredient{}

	for rows.Next() {
		err := rows.Scan(&place.CategoryID, &place.RecipeID, &place.CategoryName) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
		ingredient_categories = append([]models.IngredientCategory{place}, ingredient_categories...)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return ingredient_categories
}

func getRecipeIngredients(recipe_id int) (ingredient_output_list []models.OutputIngredient, total_cost float64) {
	ingredient_categories := getIngredientCategories(recipe_id)
	recipe_cost := 0.0
	for _, ingredient_category := range ingredient_categories {
		total_cost = 0
		// get ingredients with category_id
		ingredient_list := []models.OutputRecipeIngredient{}
		place := models.OutputRecipeIngredient{}

		ingredient_output_place := models.OutputIngredient{}

		queryStat := "SELECT \"recipeIngredientID\", \"categoryID\", \"ingredientID\", amount, \"amountUnit\"  FROM \"recipeIngredients\" WHERE \"categoryID\" = " + strconv.Itoa(ingredient_category.CategoryID) + " ORDER BY \"recipeIngredientID\""

		rows, _ := db.Query(queryStat)
		defer rows.Close()

		// ingredient_list := []models.OutputRecipeIngredient{}

		for rows.Next() {
			err := rows.Scan(&place.RecipeIngredientID, &place.CategoryID, &place.IngredientID, &place.Amount, &place.AmountUnit) // Scan the current row into the "place" variable
			if err != nil {
				fmt.Println("Error when scanning rows", err)
			}

			ingredient_info := getIngredientInfo(place.IngredientID)
			place.IngredientName = ingredient_info.IngredientName
			place.Cost = calculateIngredientCost(place, ingredient_info)
			total_cost += place.Cost

			ingredient_list = append([]models.OutputRecipeIngredient{place}, ingredient_list...)
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			fmt.Println("Error during note iteration", err)
		}

		ingredient_output_place.Ingredients = ingredient_list
		ingredient_output_place.CategoryName = ingredient_category.CategoryName
		ingredient_output_place.TotalCost = total_cost
		ingredient_output_list = append([]models.OutputIngredient{ingredient_output_place}, ingredient_output_list...)

		recipe_cost += total_cost
	}

	return ingredient_output_list, total_cost

}

func getRecipeIngredientFromId(recipe_ingredient_id int) (recipe_ingredient models.RecipeIngredient) {
	// get ingredients with recipe_id
	queryStat := "SELECT \"recipeIngredientID\", \"recipeID\", \"ingredientID\", amount, \"amountUnit\"  FROM \"recipeIngredients\" WHERE \"recipeIngredientID\" = " + strconv.Itoa(recipe_ingredient_id) + " ORDER BY \"recipeIngredientID\""

	rows, _ := db.Query(queryStat)
	defer rows.Close()

	rows.Next()
	err := rows.Scan(&recipe_ingredient.RecipeIngredientID, &recipe_ingredient.RecipeID, &recipe_ingredient.IngredientID, &recipe_ingredient.Amount, &recipe_ingredient.AmountUnit) // Scan the current row into the "place" variable
	if err != nil {
		fmt.Println("Error when scanning rows", err)
	}
	return recipe_ingredient
}

func getRescaledIngredients(parent_recipe_id int, rescale_info models.RescaleInfo) (ingredients []models.OutputIngredient, recipe_cost float64) {
	recipe_cost = 0
	total_cost := 0.0
	rescaled_ingredient := getRecipeIngredientFromId(rescale_info.IngredientID)
	fmt.Println("getRecipeIngredientFromId done")
	// get parent recipe
	// get recipeIngredients with recipeID = parentRecipe.ID or the recipe ID itself
	parent_ingredients, _ := getRecipeIngredients(rescaled_ingredient.RecipeID)
	fmt.Println(rescaled_ingredient.RecipeIngredientID)
	fmt.Println(rescaled_ingredient.RecipeID)
	fmt.Println(rescaled_ingredient.IngredientID)
	fmt.Println("getRecipeIngredients done")
	// calculate ratio
	original_amount, err := strconv.ParseFloat(strings.TrimSpace(rescaled_ingredient.Amount), 64)
	if err != nil {
		fmt.Println("getRescaledIngredient(): Failed to parse original amount")
	}
	scale := float64(rescale_info.NewAmount) / original_amount
	// multiply with all ingredients (amount and cost)
	for id, parent_ingredient := range parent_ingredients {
		total_cost = 0
		for ingredient_id, ingredient := range parent_ingredient.Ingredients {
			amount_float, _ := strconv.ParseFloat(ingredient.Amount, 64)
			if err == nil {
				amount_float = amount_float * scale
				parent_ingredients[id].Ingredients[ingredient_id].Amount = fmt.Sprintf("%.2f", amount_float)

				parent_ingredients[id].Ingredients[ingredient_id].Cost = ingredient.Cost * scale
				total_cost += parent_ingredients[id].Ingredients[ingredient_id].Cost
			}
		}
		parent_ingredients[id].TotalCost = total_cost
	}

	return parent_ingredients, recipe_cost
}

func getIngredientInfo(ingredient_id int) (ingredient_info models.IngredientInfo) {
	queryStat := "SELECT \"ingredientID\", \"ingredientName\", cost, amount, \"amountUnit\" FROM \"ingredientInfo\" WHERE \"ingredientID\" = " + strconv.Itoa(ingredient_id)
	// fmt.Println(queryStat)
	rows, _ := db.Query(queryStat)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ingredient_info.IngredientID, &ingredient_info.IngredientName, &ingredient_info.Cost, &ingredient_info.Amount, &ingredient_info.AmountUnit) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return ingredient_info
}

func calculateIngredientCost(recipe_ingredient models.OutputRecipeIngredient, ingredient_info models.IngredientInfo) (cost float64) {
	// check if amount is valid
	recipe_amount, err := strconv.ParseFloat(strings.TrimSpace(recipe_ingredient.Amount), 64)
	if err != nil {
		return 0
	}

	if ingredient_info.AmountUnit == "unknown" {
		return 0
	}

	// check if amountUnit is valid, convert if needed
	if recipe_ingredient.AmountUnit != ingredient_info.AmountUnit {
		conv_done := false
		for _, conv_info := range models.Conversion_info {
			if (conv_info.SourceUnit == recipe_ingredient.AmountUnit) && (conv_info.TargetUnit == ingredient_info.AmountUnit) {
				recipe_amount = recipe_amount * float64(conv_info.Scale)
				conv_done = true
				break
			} else if (conv_info.SourceUnit == ingredient_info.AmountUnit) && (conv_info.TargetUnit == recipe_ingredient.AmountUnit) {
				ingredient_info.Amount = ingredient_info.Amount * conv_info.Scale
				conv_done = true
				break
			}
		}
		if !conv_done {
			fmt.Println("Ingredient Name: " + recipe_ingredient.IngredientName + ": Can't convert " + recipe_ingredient.AmountUnit + " to " + ingredient_info.AmountUnit + " when calculating cost")
			return 0
		}
	}

	scale := recipe_amount / float64(ingredient_info.Amount)
	return float64(ingredient_info.Cost) * scale
}

func InsertRecipe(input_recipe models.InputRecipe) (recipe_id int) {
	currentTime := time.Now()
	created_date := currentTime.Format("2006-01-02")
	last_modified_date := currentTime.Format("2006-01-02")
	fmt.Println("Insert note:")
	fmt.Println(input_recipe.Title, created_date, last_modified_date)
	// fmt.Println("Parent_recipe_id: " + strconv.Itoa(parent_recipe_id))

	sqlStatement := `
		INSERT INTO recipes (title, steps, "parentRecipeID", "isRescaled", "createdDate", "lastModifiedDate")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING "recipeID"`

	err = db.QueryRow(sqlStatement, input_recipe.Title, input_recipe.Steps, input_recipe.ParentRecipeID, input_recipe.IsRescaled, created_date, last_modified_date).Scan(&recipe_id)
	if err != nil {
		fmt.Println("Error when inserting recipe", err)
	}

	if input_recipe.IsRescaled {
		insertRescaleInfo(recipe_id, input_recipe.RescaleInfo)
	} else {
		insertRecipeIngredients(recipe_id, input_recipe.InputIngredients)
	}

	return
}

func insertRescaleInfo(recipe_id int, rescale_info models.RescaleInfo) {
	sqlStatement := `
		INSERT INTO "rescaleInfo" ("recipeID", "ingredientID", "newAmount")
		VALUES ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, recipe_id, rescale_info.IngredientID, rescale_info.NewAmount)

	if err != nil {
		fmt.Println("Error when inserting rescaleInfo", err)
	}
}

func insertIngredientCategory(recipe_id int, category_name string) (category_id int) {
	sqlStatement := `
		INSERT INTO "ingredientCategory" ("recipeID", "categoryName")
		VALUES ($1, $2)
		RETURNING "categoryID"`

	err := db.QueryRow(sqlStatement, recipe_id, category_name).Scan(&category_id)

	if err != nil {
		fmt.Println("Error when inserting ingredientCategory", err)
	}

	return category_id
}

func insertRecipeIngredients(recipe_id int, input_ingredients []models.InputIngredient) {
	category_id := 0
	for _, input_ingredient := range input_ingredients {
		category_id = insertIngredientCategory(recipe_id, input_ingredient.CategoryName)
		for _, input_recipe_ingredient := range input_ingredient.Ingredients {
			ingredient_id := getIngredientIDFromName(input_recipe_ingredient.IngredientName)
			if ingredient_id == -1 {
				// register ingredient to database
				var t models.InputIngredientInfo
				t.IngredientName = input_recipe_ingredient.IngredientName
				t.Cost = 0
				t.Amount = 0
				t.AmountUnit = "unknown"

				ingredient_id = InsertIngredientInfo(t)
			}

			sqlStatement := `
				INSERT INTO "recipeIngredients" ("recipeID", "categoryID", "ingredientID", amount, "amountUnit")
				VALUES ($1, $2, $3, $4, $5)`

			_, err := db.Exec(sqlStatement, recipe_id, category_id, ingredient_id, input_recipe_ingredient.Amount, input_recipe_ingredient.AmountUnit)

			if err != nil {
				fmt.Println("Error when inserting rescaleInfo", err)
			}
		}

	}
}

func InsertIngredientInfo(input_ingredient_info models.InputIngredientInfo) (ingredient_info_id int) {
	sqlStatement := `
		INSERT INTO "ingredientInfo" ("ingredientName", cost, amount, "amountUnit")
		VALUES ($1, $2, $3, $4)
		RETURNING "ingredientID"`

	// query_stat := "INSERT INTO \"ingredientInfo\" (\"ingredientName\", cost, amount, \"amountUnit\") " +
	// 	"VALUES ('" + ingredient_name + "', " + strconv.Itoa(cost) + ", " + strconv.Itoa(amount) + ", '" + amountUnit + "') " +
	// 	"RETURNING \"ingredientID\""

	// fmt.Println(query_stat)

	err = db.QueryRow(sqlStatement, input_ingredient_info.IngredientName, input_ingredient_info.Cost, input_ingredient_info.Amount, input_ingredient_info.AmountUnit).Scan(&ingredient_info_id)
	// err = db.QueryRow(query_stat).Scan(&ingredient_info_id)
	if err != nil {
		fmt.Println("Error when inserting recipe", err)
	}

	fmt.Println("inserted ingredient_info_id :", ingredient_info_id)
	return ingredient_info_id
}

func getIngredientIDFromName(ingredient_name string) (ingredient_id int) {
	ingredient_id = -1
	queryStat := "SELECT \"ingredientID\" from \"ingredientInfo\" WHERE \"ingredientName\" = '" + ingredient_name + "'"

	rows, _ := db.Query(queryStat)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ingredient_id) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
	}
	return ingredient_id
}

func UpdateRecipe(input_recipe models.InputRecipeUpdate) {
	fmt.Println("Update recipe with recipe ID:", input_recipe.RecipeID)
	// todo: getRecipe
	recipe := GetRecipeDetail(input_recipe.RecipeID)
	// todo: check if isRescaled on old recipe
	if recipe.IsRescaled {
		if input_recipe.IsRescaled {
			// update rescaleInfo
			UpdateRescaleInfo(input_recipe.RescaleInfo)
		} else {
			// delete existing rescaleInfo
			DeleteRescaleInfo(input_recipe.RecipeID)
			insertRecipeIngredients(input_recipe.RecipeID, input_recipe.InputIngredients)
		}
	} else {

		if input_recipe.IsRescaled {
			// insert rescaleInfo
			insertRescaleInfo(input_recipe.RecipeID, input_recipe.RescaleInfo)
		} else {
			// delete existing recipeIngredients
			DeleteRecipeIngredients(input_recipe.RecipeID)
			// insert recipeIngredients
			insertRecipeIngredients(input_recipe.RecipeID, input_recipe.InputIngredients)
		}
	}

	// todo: update recipes
	currentTime := time.Now()
	last_modified_date := currentTime.Format("2006-01-02")

	sqlStatement := `
		UPDATE recipes
		SET title = $2, steps = $3, "isRescaled" = $4, "lastModifiedDate" = $5
		WHERE "recipeID" = $1`
	_, err := db.Exec(sqlStatement, input_recipe.RecipeID, input_recipe.Title, input_recipe.Steps, input_recipe.IsRescaled, last_modified_date)

	if err != nil {
		fmt.Println("Error when updating note", err)
	}

}

func UpdateRescaleInfo(rescale_info models.RescaleInfo) {
	fmt.Println("Update rescale info for recipe ID:", rescale_info.RecipeID)

	sqlStatement := `
		UPDATE "rescaleInfo"
		SET "ingredientID" = $2, "newAmount" = $3
		WHERE "recipeID" = $1`
	_, err := db.Exec(sqlStatement, rescale_info.RecipeID, rescale_info.IngredientID, rescale_info.NewAmount)

	if err != nil {
		fmt.Println("Error when updating note", err)
	}
}

func UpdateIngredientInfo(ingredient_info models.IngredientInfo) {
	fmt.Println("Update ingredient info for ingredient ID:", ingredient_info.IngredientID)

	sqlStatement := `
		UPDATE "ingredientInfo"
		SET "ingredientName" = $2, cost = $3, amount = $4, "amountUnit" = $5
		WHERE "ingredientID" = $1`
	_, err := db.Exec(sqlStatement, ingredient_info.IngredientID, ingredient_info.IngredientName, ingredient_info.Cost, ingredient_info.Amount, ingredient_info.AmountUnit)

	if err != nil {
		fmt.Println("Error when updating ingredient", err)
	}
}

func DeleteRecipe(recipe_id int) {
	fmt.Println("Delete recipe with ID: ", recipe_id)

	if isSubrecipeExists(recipe_id) {
		DeleteSubrecipes(recipe_id)
	}

	DeleteRescaleInfo(recipe_id)

	DeleteRecipeIngredients(recipe_id)

	sqlStatement := `
		DELETE FROM recipes
		WHERE "recipeID" = $1;`
	_, err := db.Exec(sqlStatement, recipe_id)

	if err != nil {
		fmt.Println("Error when deleting recipe", err)
	}
}

func DeleteSubrecipes(recipe_id int) {
	fmt.Println("Delete subrecipes of recipe ID: ", recipe_id)

	subrecipe_list := getSubrecipes(recipe_id)

	for _, recipe := range subrecipe_list {
		if recipe.IsRescaled {
			DeleteRescaleInfo(recipe.RecipeID)
		} else {
			DeleteRecipeIngredients(recipe.RecipeID)
		}
	}

	sqlStatement := `
		DELETE FROM recipes
		WHERE "parentRecipeID" = $1;`
	_, err := db.Exec(sqlStatement, recipe_id)

	if err != nil {
		fmt.Println("Error when deleting subrecipes", err)
	}
}

func DeleteRecipeIngredients(recipe_id int) {
	sqlStatement := `
		DELETE FROM "ingredientCategory"
		WHERE "recipeID" = $1;`
	_, err := db.Exec(sqlStatement, recipe_id)

	if err != nil {
		fmt.Println("Error when deleting ingredientCategory", err)
	}

	sqlStatement = `
		DELETE FROM "recipeIngredients"
		WHERE "recipeID" = $1;`
	_, err = db.Exec(sqlStatement, recipe_id)

	if err != nil {
		fmt.Println("Error when deleting recipeIngredient", err)
	}
}

func DeleteRescaleInfo(recipe_id int) {
	sqlStatement := `
		DELETE FROM "rescaleInfo"
		WHERE "recipeID" = $1;`
	_, err := db.Exec(sqlStatement, recipe_id)

	if err != nil {
		fmt.Println("Error when deleting rescaleInfo", err)
	}
}

func DeleteIngredientInfo(ingredient_id int) {
	sqlStatement := `
		DELETE FROM "ingredientInfo"
		WHERE "ingredientID" = $1;`
	_, err := db.Exec(sqlStatement, ingredient_id)

	if err != nil {
		fmt.Println("Error when deleting ingredientInfo", err)
	}
}

func GetIngredientInfoList() []models.IngredientInfo {
	place := models.IngredientInfo{} // Initialize a User struct to hold retrieved data
	queryStat := "SELECT \"ingredientID\", \"ingredientName\", \"cost\", \"amount\", \"amountUnit\" FROM \"ingredientInfo\" ORDER BY \"ingredientID\""

	rows, _ := db.Query(queryStat)
	defer rows.Close()

	ingredient_list := []models.IngredientInfo{}

	for rows.Next() {
		err := rows.Scan(&place.IngredientID, &place.IngredientName, &place.Cost, &place.Amount, &place.AmountUnit) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
		// ingredient_list = append([]models.IngredientInfo{place}, ingredient_list...) // ... unpack into note_list[0], note_list[1], ...
		ingredient_list = append(ingredient_list, place)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return ingredient_list
}
