package interfaces

import (
	"db_mgmt"
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case ("/list_recipe"):
		listRecipe(w, r)
	case ("/get_recipe"):
		getRecipe(w, r)
	case ("/list_ingredient_info"):
		listIngredientInfo(w, r)
	case ("/insert_recipe"):
		insertRecipe(w, r)
	case ("/insert_ingredient_info"):
		insertIngredientInfo(w, r)
	case ("/update_recipe"):
		updateRecipe(w, r)
	case ("/update_ingredient_info"):
		updateIngredientInfo(w, r)
	case ("/delete_recipe"):
		deleteRecipe(w, r)
	case ("/delete_ingredient_info"):
		deleteIngredientInfo(w, r)
	default:
		fmt.Fprintf(w, "bad address 404")
	}
}

func listIngredientInfo(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "GET")

	if r.Method == "GET" {
		ingredient_list := db_mgmt.GetIngredientInfoList()
		var result, err = json.Marshal(map[string]interface{}{"ingredient_list": ingredient_list})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func listRecipe(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "GET")

	if r.Method == "GET" {
		recipe_list := db_mgmt.GetRecipeList()
		var result, err = json.Marshal(map[string]interface{}{"recipe_list": recipe_list})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func getRecipe(w http.ResponseWriter, r *http.Request) {
	recipe_id := r.FormValue("recipe_id")

	fmt.Println("Get recipe with recipeID: ", recipe_id)
	recipe_id_int, _ := strconv.Atoi(recipe_id)

	setHeader(w, "GET")

	if r.Method == "GET" {
		recipe_detail := db_mgmt.GetRecipeDetail(recipe_id_int)
		var result, err = json.Marshal(map[string]interface{}{"recipe": recipe_detail})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
func insertRecipe2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Insert recipe called")
	setHeader(w, "POST")

	contentType := r.Header.Get("Content-type")

	if contentType == "" {
		return
	}

	fmt.Println(contentType)

	var result, _ = json.Marshal(map[string]interface{}{"recipe_id": 0})
	w.Write(result)
}
func insertRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Insert recipe called")
	setHeader(w, "POST")

	contentType := r.Header.Get("Content-type")

	if contentType == "" {
		return
	}

	fmt.Println("Content type:")
	fmt.Println(contentType)

	var t models.InputRecipe
	var ingredients []models.InputRecipeIngredient

	if contentType == "application/json" {
		fmt.Println("content type is application/json")
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		t.Title = r.FormValue("title")
		t.Steps = r.FormValue("steps")
		t.ParentRecipeID, _ = strconv.Atoi(r.FormValue("parentRecipeID"))
		t.IsRescaled, _ = strconv.ParseBool(r.FormValue("isRescaled"))

		fmt.Println("recipeIngredients in bytes")
		fmt.Println([]byte(r.FormValue("recipeIngredients")))
		// json.Unmarshal([]byte(r.FormValue("recipeIngredients")), &t.InputRecipeIngredients)
		if err := json.Unmarshal([]byte(r.FormValue("recipeIngredients")), &ingredients); err != nil {
			panic(err)
		}

		json.Unmarshal([]byte(r.FormValue("rescaleInfo")), &t.RescaleInfo)
	}

	printInputRecipeIngredients(ingredients)
	PrintRescaleInfo(t.RescaleInfo)

	fmt.Println("Insert recipe with the following parameters:")
	fmt.Println("parentRecipeID: ", t.ParentRecipeID)
	fmt.Println("Title: ", t.Title)
	fmt.Println("Steps: ", t.Steps)
	fmt.Println("is_rescaled: ", t.IsRescaled)

	if r.Method == "POST" {
		recipe_id := db_mgmt.InsertRecipe(t)
		fmt.Print("New record note ID is: ", recipe_id)

		var result, err = json.Marshal(map[string]interface{}{"recipe_id": recipe_id})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func printInputRecipeIngredients(input_ingredient []models.InputRecipeIngredient) {
	fmt.Println("===============================")
	for id, ingredient := range input_ingredient {
		fmt.Println("Ingredient ID: ", id)
		fmt.Println("Ingredient name: ", ingredient.IngredientName)
		fmt.Println("Amount: ", ingredient.Amount)
		fmt.Println("AmountUnit: ", ingredient.AmountUnit)
		fmt.Println("===============================")
	}
}

func PrintRescaleInfo(rescale_info models.RescaleInfo) {
	fmt.Println("RescaleInfo:")
	fmt.Println("=========================")
	fmt.Println("RecipeID:", rescale_info.RecipeID)
	fmt.Println("IngredientID:", rescale_info.IngredientID)
	fmt.Println("NewAmount:", rescale_info.NewAmount)
	fmt.Println("=========================")
}

func insertIngredientInfo(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "POST")

	contentType := r.Header.Get("Content-type")

	var t models.InputIngredientInfo

	fmt.Println(contentType)

	if contentType == "application/json" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		t.IngredientName = r.FormValue("ingredientName")
		t.Cost, _ = strconv.Atoi(r.FormValue("cost"))
		t.Amount, _ = strconv.Atoi(r.FormValue("amount"))
		t.AmountUnit = r.FormValue("amountUnit")
	}

	fmt.Println("Insert recipe ingredient with the following parameters:")
	fmt.Println("ingredientName: ", t.IngredientName)
	fmt.Println("cost: ", t.Cost)
	fmt.Println("amount: ", t.Amount)
	fmt.Println("amountUnit: ", t.AmountUnit)

	if r.Method == "POST" {
		ingredient_id := db_mgmt.InsertIngredientInfo(t)
		fmt.Print("New record ingredient ID is: ", ingredient_id)

		var result, err = json.Marshal(map[string]interface{}{"ingredient_id": ingredient_id})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}

func updateRecipe(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "PUT")
	contentType := r.Header.Get("Content-type")

	if contentType == "" {
		return
	}
	fmt.Println("Update Recipe")
	fmt.Println(contentType)

	var t models.InputRecipeUpdate

	if contentType == "application/json" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		t.RecipeID, _ = strconv.Atoi(r.FormValue("recipeID"))
		t.Title = r.FormValue("title")
		t.Steps = r.FormValue("steps")
		t.IsRescaled, _ = strconv.ParseBool(r.FormValue("isRescaled"))
		json.Unmarshal([]byte(r.FormValue("recipeIngredients")), &t.InputRecipeIngredients)
		json.Unmarshal([]byte(r.FormValue("rescaleInfo")), &t.RescaleInfo)
	}

	fmt.Println("Update recipe with the following parameters:")
	fmt.Println("RecipeID: ", t.RecipeID)
	fmt.Println("Title: ", t.Title)
	fmt.Println("Steps: ", t.Steps)
	fmt.Println("is_rescaled: ", t.IsRescaled)

	if r.Method == "PUT" {
		db_mgmt.UpdateRecipe(t)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func updateIngredientInfo(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "PUT")
	contentType := r.Header.Get("Content-type")

	var t models.IngredientInfo

	if contentType == "application/json" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		t.IngredientID, _ = strconv.Atoi(r.FormValue("ingredientID"))
		t.IngredientName = r.FormValue("ingredientName")
		t.Cost, _ = strconv.Atoi(r.FormValue("cost"))
		t.Amount, _ = strconv.Atoi(r.FormValue("amount"))
		t.AmountUnit = r.FormValue("amountUnit")
	}

	fmt.Println("Update recipe ingredient with the following parameters:")
	fmt.Println("ingredientID: ", t.IngredientID)
	fmt.Println("ingredientName: ", t.IngredientName)
	fmt.Println("cost: ", t.Cost)
	fmt.Println("amount: ", t.Amount)
	fmt.Println("amountUnit: ", t.AmountUnit)

	if r.Method == "PUT" {
		db_mgmt.UpdateIngredientInfo(t)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "DELETE")

	recipe_id := r.FormValue("recipe_id")
	fmt.Println("Delete note with note id:", recipe_id)

	if r.Method == "DELETE" {
		note_id_int, _ := strconv.Atoi(recipe_id)
		db_mgmt.DeleteRecipe(note_id_int)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func deleteIngredientInfo(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "DELETE")

	ingredient_id := r.FormValue("ingredient_id")
	fmt.Println("Delete note with recipe id:", ingredient_id)

	if r.Method == "DELETE" {
		ingredient_id_int, _ := strconv.Atoi(ingredient_id)
		db_mgmt.DeleteIngredientInfo(ingredient_id_int)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func setHeader(w http.ResponseWriter, allow_methods string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", allow_methods)
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	w.WriteHeader(http.StatusOK)
}
