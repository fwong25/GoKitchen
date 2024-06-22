<template>
    <subpage-template>
        <form>
            <div class="input-group mb-2">
                <input v-if="this.$route.query.parent_recipe_id == 0" type="text" class="form-control" placeholder="Edit mode: Add new recipe" readonly>
                <input v-else type="text" class="form-control" placeholder="Edit mode: Add new subrecipe" readonly>
                <div class="input-group-append text-info">
                    <span class="input-group-text bg-white py-0">
                        <button type="submit" class="btn btn-sm text-info" @click.stop.prevent="insertRecipe()">
                            <h6>
                                <font-awesome-icon icon="plus-circle" size="lg"/>
                            </h6>
                        </button>
                    </span>
                </div>
                <div class="input-group-append text-info">
                    <span class="input-group-text bg-white py-0">
                        <button type="submit" class="btn btn-sm text-info"  @click.stop.prevent="cancelAddRecipe">
                            <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                        </button>
                    </span>
                </div>
            </div>

            <h3 data-placeholder="Insert recipe title here..."class="mb-1" id="title_p" contenteditable="true" ></h3>
            <div  v-if="this.$route.query.parent_recipe_id != 0">
                <br>
                <ul v-if="this.$route.query.parent_recipe_id != 0" class="nav nav-pills">
                    <li class="nav-item">
                        <a v-bind:class="modify_recipe_class" href="#" @click.stop.prevent="modifyRecipe()">
                                Modify original recipe
                        </a>
                    </li>
                    <li class="nav-item">
                        <a v-bind:class="rescale_recipe_class" href="#"  @click.stop.prevent="rescaleRecipe()">
                            Rescale original recipe
                        </a>
                    </li>
                    <li class="nav-item">
                        <a v-bind:class="blank_recipe_class" href="#" @click.stop.prevent="blankRecipe()">
                            Create from blank recipe
                        </a>
                    </li>
                </ul>
            </div>

            <button  v-if="subrecipe_mode != 'rescale'" type="submit" class="astext" @click.stop.prevent="addIngredientGroup()">
                <i class="fs-4 bi-house"></i> <span class="ms-1 d-none d-sm-inline text-info">+ Add Ingredient Group</span>
            </button>
            <div  v-if="this.input_ingredient_group_mode" class="input-group mb-2">
                <input type="text" class="form-control" placeholder="Insert Ingredient Group Title..." v-model="new_ingredient_group_name">
                <div class="input-group-append text-info">
                    <span class="input-group-text bg-white py-0">
                        <button type="submit" class="btn btn-sm text-info" @click.stop.prevent="insertIngredientCategory(new_ingredient_group_name)">
                            <h6>
                                <font-awesome-icon icon="plus-circle" size="lg"/>
                            </h6>
                        </button>
                    </span>
                </div>
                <div class="input-group-append text-info">
                    <span class="input-group-text bg-white py-0">
                        <button type="submit" class="btn btn-sm text-info"  @click.stop.prevent="cancelAddIngredientGroup()">
                            <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                        </button>
                    </span>
                </div>
            </div>
            <br>

            <br>
            <div v-for="ingredient_group in recipe.RecipeIngredients">
                <div v-if="subrecipe_mode == 'rescale'">
                    <h5> {{ ingredient_group.CategoryName }}</h5>
                </div>
                <div v-else-if="edit_ingredient_group_id != ingredient_group.id">
                    <h5 class="d-inline"> {{ ingredient_group.CategoryName }}</h5>
                    <form class="float-right d-inline" style="float: right;">
                        <button type="submit" class="btn" title="Delete ingredient group" @click.stop.prevent="deleteIngredientGroup(ingredient_group.id)">
                            <i class="float-right">
                                <font-awesome-icon icon="trash-alt" size="lg" class="float-right"/>
                            </i>
                        </button>
                    </form>
                    <form class="float-right d-inline" style="float: right;">
                        <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="editIngredientGroup(ingredient_group.id)">
                            <i class="float-right">
                                <font-awesome-icon icon="pencil-alt" size="s" class="float-right"/>
                            </i>
                        </button>
                    </form>
                </div>
                <div v-else>
                    <div class="input-group mb-2">
                        <input v-model="ingredient_group.edit_group_name" type="text" class="form-control">
                        <div class="input-group-append text-info">
                            <span class="input-group-text bg-white py-0">
                                <button type="submit" class="btn btn-sm text-info" @click.stop.prevent="confirmEditIngredientGroup(ingredient_group.id)">
                                    <h6>
                                        <font-awesome-icon icon="plus-circle" size="lg"/>
                                    </h6>
                                </button>
                            </span>
                        </div>
                        <div class="input-group-append text-info">
                            <span class="input-group-text bg-white py-0">
                                <button type="submit" class="btn btn-sm text-info"  @click.stop.prevent="cancelEditIngredientGroup()">
                                    <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                                </button>
                            </span>
                        </div>
                    </div>
                </div>
                

                <table class="table">
                    <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Ingredient</th>

                            <th v-if="subrecipe_mode == 'rescale'" scope="col">Original Amount</th>
                            <th v-else scope="col">Amount</th>

                            <th v-if="subrecipe_mode == 'rescale'">New Amount</th>
                            <th v-if="subrecipe_mode == 'rescale'">Unit</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-if="subrecipe_mode != 'rescale'">
                            <th scope="row"></th>
                            <td>
                                <input type="text" class="form-control" name="ingredient_name" placeholder="Ingredient name" v-model="ingredient_group.new_ingredient_name">
                            </td>
                            <td>
                                <input type="text" class="form-control" name="amount" placeholder="Amount" v-model="ingredient_group.new_amount">
                                <input type="text" class="form-control" name="amountUnit" placeholder="Unit" v-model="ingredient_group.new_amount_unit">
                            </td>
                            <td>
                                <!-- <form class="float-right d-inline" style="float: right;"> -->
                                <form class="float-right d-inline">
                                    <button type="submit" class="btn btn-sm" @click.stop.prevent="addIngredient(ingredient_group.id)">
                                        <font-awesome-icon icon="plus-circle" size="xl" class="text-info"/>
                                    </button>
                                </form>
                            </td>
                        </tr>
                        <!-- <tr>
                            <th scope="row">1</th>
                            <td>Flour</td>
                            <td>100 g</td>
                            <td></td>
                        </tr> -->

                        <tr v-for="ingredient in ingredient_group.Ingredients">
                            <th scope="row">{{ ingredient.id }}</th>

                            <td v-if="edit_ingredient_order_id == ingredient.id && edit_ingredient_category_id == ingredient_group.id">
                                <input type="text" class="form-control" name="ingredient_name" placeholder="Ingredient name" v-model="edit_ingredient_name">
                            </td>
                            <td v-else>{{ ingredient.IngredientName }}</td>

                            <td v-if="edit_ingredient_order_id == ingredient.id && edit_ingredient_category_id == ingredient_group.id">
                                <input type="text" class="form-control" name="amount" placeholder="Amount" v-model="edit_amount">
                                <input type="text" class="form-control" name="amountUnit" placeholder="Unit" v-model="edit_amount_unit">
                            </td>
                            <td v-else>{{ ingredient.Amount }} {{ ingredient.AmountUnit }}</td>
                            
                            <td v-if="subrecipe_mode == 'rescale' && (rescale_ing_id != ingredient.id - 1 || rescale_cat_id != ingredient_group.id)">
                                {{ ingredient.NewAmount }}
                            </td>
                            <td v-else-if="subrecipe_mode == 'rescale'">
                                <input type="number" class="form-control" size="4" name="amount" placeholder="Amount" v-model="rescale_amount">
                            </td>

                            <td v-if="subrecipe_mode == 'rescale'">
                                {{ ingredient.AmountUnit }}
                            </td>

                            <td v-if="edit_ingredient_order_id == ingredient.id && edit_ingredient_category_id == ingredient_group.id">
                                <form class="float-right d-inline" style="float: right;">
                                    <button type="submit" class="btn" title="Cancel edit ingredient" @click.stop.prevent="cancelEditIngredient()">
                                        <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                                    </button>
                                </form>
                                <form class="float-right d-inline" style="float: right;">
                                    <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="confirmEditIngredient()">
                                        <i class="text-danger float-right">
                                            <font-awesome-icon icon="check" size="xl" class="text-info float-right"/>
                                        </i>
                                    </button>
                                </form>
                            </td>
                            <td v-else-if="subrecipe_mode != 'rescale'">
                                <form class="float-right d-inline" style="float: right;">
                                    <button type="submit" class="btn" title="Delete ingredient" @click.stop.prevent="deleteIngredient(ingredient.id, ingredient_group.id)">
                                        <i class="text-danger float-right">
                                            <font-awesome-icon icon="trash-alt" size="lg" class="text-danger float-right"/>
                                        </i>
                                    </button>
                                </form>
                                <form class="float-right d-inline" style="float: right;">
                                    <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="editIngredient(ingredient.id, ingredient_group.id)">
                                        <i class="text-danger float-right">
                                            <font-awesome-icon icon="pencil-alt" size="s" class="text-info float-right"/>
                                        </i>
                                    </button>
                                </form>
                            </td>
                            <td v-else-if="subrecipe_mode == 'rescale' && rescale_ing_id == ingredient.id - 1 && rescale_cat_id == ingredient_group.id">
                                <form class="float-right d-inline" style="float: right;">
                                    <button type="submit" class="btn" title="Cancel edit ingredient" @click.stop.prevent="cancelRescaleIngredient()">
                                        <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                                    </button>
                                </form>
                                <form class="float-right d-inline" style="float: right;">
                                    <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="confirmRescaleIngredient()">
                                        <i class="text-danger float-right">
                                            <font-awesome-icon icon="check" size="xl" class="text-info float-right"/>
                                        </i>
                                    </button>
                                </form>
                            </td>
                            <td v-else>
                                <form class="float-right d-inline" style="float: right;">
                                    <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="selectRescaleIngredient(ingredient.id, ingredient_group.id)">
                                        <i class="text-danger float-right">
                                            <font-awesome-icon icon="pencil-alt" size="s" class="text-info float-right"/>
                                        </i>
                                    </button>
                                </form>
                            </td>

                        </tr>
                    </tbody>
                </table>
            </div>

            <h5>Steps</h5>
            <p data-placeholder="Insert recipe steps here..." class="mb-1" id="steps_p" style="white-space: pre-line" contenteditable="true"></p>
            
        </form>
    </subpage-template>
</template>
    
<script>
import { host_info } from './../main.js'
    export default {
        data() {
        return {
            edit_ingredient_category_id: 0,
            edit_ingredient_order_id: 0,
            edit_ingredient_name: '',
            edit_amount: '',
            edit_amount_unit: '',

            ingredient_order_id: 0,
            ingredient_category_id: 0,
            recipe: {
                Title: '',
                Steps: '',
                ParentRecipeID: 0,
                IsRescaled: false,
                RecipeIngredients: [],
                RecipeRescaleInfo: {
                    recipeID: 0,
                    ingredientID: 0,
                    newAmount: 0
                }
            },

            recipe_detail: {
                RecipeID: 0,
                Title: '',
                Steps: '',
                ParentRecipeID: 0,
                IsRescaled: false,
                CreatedDate: '',
                LastModifiedDate: '',
                Ingredients: [],
                RecipeRescaleInfo: {
                    RecipeID: 0,
                    IngredientID: 0,
                    NewAmount: 0
                },
                TotalCost: 0
            },

            subrecipe_mode: 'modify',
            modify_recipe_class : "nav-link active",
            rescale_recipe_class: "nav-link",
            blank_recipe_class: "nav-link",

            rescale_ing_id: -1,
            rescale_cat_id: -1,
            rescale_amount: 0,
            
            input_ingredient_group_mode: false,
            new_ingredient_group_name: '',

            edit_ingredient_group_id: -1,
        }  
    },
    methods: {
        fetchSelectedRecipe(recipe_id) {
            fetch('http://' + host_info.backend_host + '/get_recipe?' + 
                new URLSearchParams({
                    recipe_id: recipe_id
                }), 
                {
                method: 'GET',
                mode: 'cors',
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json",
                    "Origin": "http://localhost:5173"
                }
            })
            .then((response) => {
                console.log(response)
                response.json().then((data) => {
                    console.log(data.recipe)
                    console.log(data.recipe.Title)
                    console.log(data.recipe.Steps)
                    this.recipe_detail = data.recipe

                    const title_p = document.getElementById("title_p");
                    const steps_p = document.getElementById("steps_p");
                    title_p.innerHTML = this.recipe_detail.Title;
                    steps_p.innerHTML = this.recipe_detail.Steps;

                    if (this.recipe_detail.Ingredients != null) {
                        for (let i = 0; i < this.recipe_detail.Ingredients.length; ++i) {
                            this.recipe_detail.Ingredients[i]["id"] = i;
                            this.recipe_detail.Ingredients[i].new_ingredient_name = '',
                            this.recipe_detail.Ingredients[i].new_amount = '',
                            this.recipe_detail.Ingredients[i].new_amount_unit = ''

                            this.recipe_detail.Ingredients[i].edit_group_name = ''

                            if (this.recipe_detail.Ingredients[i].Ingredients != null) {
                                for (let j = 0; j < this.recipe_detail.Ingredients[i].Ingredients.length; ++j) {
                                    this.recipe_detail.Ingredients[i].Ingredients[j]["id"] = j+1;
                                }
                            }
                        }
                        
                        this.ingredient_category_id = this.recipe_detail.Ingredients.length
                    }  
                    else {
                        this.recipe_detail.Ingredients = []
                        this.rescale_recipe_class = "nav-link disabled"
                    }

                    this.recipe.Title = this.recipe_detail.Title
                    this.recipe.Steps = this.recipe_detail.Steps
                    this.recipe.ParentRecipeID = this.recipe_detail.ParentRecipeID
                    this.recipe.IsRescaled = this.recipe_detail.IsRescaled
                    this.recipe.RecipeIngredients = this.recipe_detail.Ingredients
                    this.recipe.RecipeRescaleInfo = this.recipe_detail.RecipeRescaleInfo


                })
            })
        },
        insertRecipe() {
            const title_p = document.getElementById("title_p");
            const steps_p = document.getElementById("steps_p");
            this.recipe.Title = title_p.innerHTML;
            this.recipe.Steps = steps_p.innerHTML;
            this.recipe.ParentRecipeID = parseInt(this.$route.query.parent_recipe_id);

            var ingredients = []
            if (this.recipe.RecipeIngredients != null)
                for(let i = 0; i < this.recipe.RecipeIngredients.length; ++i) {
                    var ingredient = {
                        categoryName: this.recipe.RecipeIngredients[i].CategoryName,
                        ingredients: []
                    }
                    if (this.recipe.RecipeIngredients[i].Ingredients != null) {
                        for(let j = 0; j < this.recipe.RecipeIngredients[i].Ingredients.length; ++j) {
                            var ingredient_entry = {
                                ingredientName: this.recipe.RecipeIngredients[i].Ingredients[j].IngredientName,
                                amount: this.recipe.RecipeIngredients[i].Ingredients[j].Amount,
                                amountUnit: this.recipe.RecipeIngredients[i].Ingredients[j].AmountUnit
                            }
                            ingredient.ingredients.push(ingredient_entry)
                        }
                    }
                    ingredients.push(ingredient)
                }

            console.log(ingredients)

            fetch('http://' + host_info.backend_host + '/insert_recipe', 
            {
                method: 'POST',
                mode: 'cors',
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json",
                    "Origin": "http://localhost:5173"
                },
                body: JSON.stringify({
                    title: this.recipe.Title,
                    steps: this.recipe.Steps,
                    parentRecipeID: this.recipe.ParentRecipeID,
                    isRescaled: this.recipe.IsRescaled,
                    inputIngredients: ingredients,
                    rescaleInfo: this.recipe.RecipeRescaleInfo
                })
            })
            .then((response) => {
                console.log(response)
                response.json().then((data) => {
                    console.log(data)
                    console.log(data.recipe_id)
                    this.$router.push({ path: 'recipe_view',
                        query: { recipe_id: data.recipe_id }
                    });
                })
            })
        },
        cancelAddRecipe() {
            console.log(this.$route.query.parent_recipe_id)
            if (this.$route.query.parent_recipe_id == 0) {
                this.$router.push({ path: '/'});
            } else {
                this.$router.push({ path: 'recipe_view',
                    query: { recipe_id: this.$route.query.parent_recipe_id }
                });
            }
        },
        addIngredient(ingredient_category_id) {
            // this.ingredient_order_id += 1
            this.ingredient_order_id = this.recipe.RecipeIngredients[ingredient_category_id].Ingredients.length + 1
            var ingredient = {
                id: this.ingredient_order_id,
                IngredientName: this.recipe.RecipeIngredients[ingredient_category_id].new_ingredient_name,
                Amount: this.recipe.RecipeIngredients[ingredient_category_id].new_amount,
                AmountUnit: this.recipe.RecipeIngredients[ingredient_category_id].new_amount_unit
            }
            this.recipe.RecipeIngredients[ingredient_category_id].Ingredients.push(ingredient)

            this.recipe.RecipeIngredients[ingredient_category_id].new_ingredient_name = ''
            this.recipe.RecipeIngredients[ingredient_category_id].new_amount = ''
            this.recipe.RecipeIngredients[ingredient_category_id].new_amount_unit = ''
        },
        editIngredient(edit_ingredient_order_id, edit_ingredient_category_id) {
            this.edit_ingredient_group_id = -1

            this.edit_ingredient_order_id = edit_ingredient_order_id
            this.edit_ingredient_category_id = edit_ingredient_category_id

            this.edit_ingredient_name = this.recipe.RecipeIngredients[edit_ingredient_category_id].Ingredients[edit_ingredient_order_id-1].IngredientName
            this.edit_amount = this.recipe.RecipeIngredients[edit_ingredient_category_id].Ingredients[edit_ingredient_order_id-1].Amount
            this.edit_amount_unit = this.recipe.RecipeIngredients[edit_ingredient_category_id].Ingredients[edit_ingredient_order_id-1].AmountUnit
        },
        deleteIngredient(ingredient_id, category_id) {
            ingredient_id -= 1
            this.recipe.RecipeIngredients[category_id].Ingredients.splice(ingredient_id, 1)
            for(let i = ingredient_id; i < this.recipe.RecipeIngredients[category_id].Ingredients.length; ++i) {
                this.recipe.RecipeIngredients[category_id].Ingredients[i].id -= 1
            }
        },
        confirmEditIngredient() {
            this.recipe.RecipeIngredients[this.edit_ingredient_category_id].Ingredients[this.edit_ingredient_order_id-1].IngredientName = this.edit_ingredient_name
            this.recipe.RecipeIngredients[this.edit_ingredient_category_id].Ingredients[this.edit_ingredient_order_id-1].Amount = this.edit_amount
            this.recipe.RecipeIngredients[this.edit_ingredient_category_id].Ingredients[this.edit_ingredient_order_id-1].AmountUnit = this.edit_amount_unit

            this.edit_ingredient_order_id = 0
            this.edit_ingredient_category_id = 0
        },
        cancelEditIngredient() {
            this.edit_ingredient_order_id = 0
            this.edit_ingredient_category_id = 0
        },
        modifyRecipe() {
            console.log("modify recipe tab is selected")
            if(this.subrecipe_mode != "modify") {
                this.subrecipe_mode = "modify"
                this.recipe.IsRescaled = false
                
                this.modify_recipe_class = "nav-link active"
                this.rescale_recipe_class = "nav-link"
                this.blank_recipe_class = "nav-link"

                // this.fetchSelectedRecipe(this.$route.query.parent_recipe_id)
                this.resetRecipe()
            }
        },
        rescaleRecipe() {
            console.log("rescale recipe tab is selected")
            if(this.subrecipe_mode != "rescale") {
                this.resetRecipe()

                this.subrecipe_mode = "rescale"
                this.recipe.IsRescaled = true

                this.modify_recipe_class = "nav-link"
                this.rescale_recipe_class = "nav-link active"
                this.blank_recipe_class = "nav-link"

                this.rescale_ing_id = -1
                // this.rescale_amount = this.recipe.RecipeIngredients[0].Amount

                for(let i = 0; i < this.recipe.RecipeIngredients.length; ++i) {
                    for(let j = 0; j < this.recipe.RecipeIngredients[i].Ingredients.length; ++j) {
                        this.recipe.RecipeIngredients[i].Ingredients[j].NewAmount = this.recipe.RecipeIngredients[i].Ingredients[j].Amount
                    }
                }

                this.recipe.RecipeRescaleInfo.RecipeID = this.recipe_detail.RecipeID
                this.recipe.RecipeRescaleInfo.IngredientID = this.recipe.RecipeIngredients[0].RecipeIngredientID
                if(isNaN(this.recipe.RecipeIngredients[0].Ingredients[0].Amount))
                    this.recipe.RecipeRescaleInfo.NewAmount = -1
                else
                    this.recipe.RecipeRescaleInfo.NewAmount = parseInt(this.recipe.RecipeIngredients[0].Ingredients[0].Amount)
            }
        },
        blankRecipe() {
            console.log("blank recipe tab is selected")
            if(this.subrecipe_mode != "blank") {
                this.subrecipe_mode = "blank"
                this.recipe.IsRescaled = false

                this.modify_recipe_class = "nav-link"
                this.rescale_recipe_class = "nav-link"
                this.blank_recipe_class = "nav-link active"

                this.edit_ingredient_order_id = 0
                this.edit_ingredient_category_id = 0
                this.edit_ingredient_name = ''
                this.edit_amount = ''
                this.edit_amount_unit = ''

                this.ingredient_order_id = 0
                this.ingredient_category_id = 0
                this.recipe.Title = ''
                this.recipe.Steps = ''
                this.recipe.IsRescaled = false
                this.recipe.RecipeIngredients = []
                this.recipe.RecipeRescaleInfo = {
                    recipeID: 0,
                    ingredientID: 0,
                    newAmount: 0
                }
                const title_p = document.getElementById("title_p");
                const steps_p = document.getElementById("steps_p");
                title_p.innerHTML = this.recipe.Title;
                steps_p.innerHTML = this.recipe.Steps;
            }
        }, 
        selectRescaleIngredient(ingredient_id, category_id) {
            if(!isNaN(this.recipe.RecipeIngredients[category_id].Ingredients[ingredient_id-1].NewAmount)) {
                this.rescale_ing_id = ingredient_id-1
                this.rescale_cat_id = category_id
                this.rescale_amount = this.recipe.RecipeIngredients[category_id].Ingredients[ingredient_id-1].NewAmount
            }
        },
        confirmRescaleIngredient() {
            // this.recipe.RecipeIngredients[this.rescale_ing_id].NewAmount = this.rescale_amount

            var scale = this.rescale_amount / this.recipe.RecipeIngredients[this.rescale_cat_id].Ingredients[this.rescale_ing_id].Amount
            for(let i = 0; i < this.recipe.RecipeIngredients.length; ++i) {
                for(let j = 0; j < this.recipe.RecipeIngredients[i].Ingredients.length; ++j) {
                    this.recipe.RecipeIngredients[i].Ingredients[j].NewAmount *= scale
                }
            }

            this.recipe.RecipeRescaleInfo.RecipeID = this.recipe_detail.RecipeID
            this.recipe.RecipeRescaleInfo.IngredientID = this.recipe.RecipeIngredients[this.rescale_cat_id].Ingredients[this.rescale_ing_id].RecipeIngredientID
            this.recipe.RecipeRescaleInfo.NewAmount = this.rescale_amount

            this.rescale_ing_id = -1
            this.rescale_cat_id = -1
        },
        cancelRescaleIngredient() {
            this.rescale_ing_id = -1
            this.rescale_cat_id = -1
        },
        resetRecipe() {
            this.recipe.Title = this.recipe_detail.Title
            this.recipe.Steps = this.recipe_detail.Steps
            this.recipe.ParentRecipeID = this.recipe_detail.ParentRecipeID
            this.recipe.IsRescaled = this.recipe_detail.IsRescaled
            this.recipe.RecipeIngredients = this.recipe_detail.Ingredients
            this.recipe.RecipeRescaleInfo = this.recipe_detail.RecipeRescaleInfo

            const title_p = document.getElementById("title_p");
            const steps_p = document.getElementById("steps_p");
            title_p.innerHTML = this.recipe.Title;
            steps_p.innerHTML = this.recipe.Steps;

            this.edit_ingredient_order_id = 0
            this.edit_ingredient_category_id = 0
            this.edit_ingredient_name = ''
            this.edit_amount = ''
            this.edit_amount_unit = ''

            this.edit_ingredient_group_id = -1
        },
        insertIngredientCategory(category_name) {
            var ingredient_category = {
                CategoryName: category_name,
                Ingredients: [],
                id: this.ingredient_category_id,
                edit_group_name: ''
            }
            this.recipe.RecipeIngredients.push(ingredient_category)
            this.ingredient_category_id += 1

            this.cancelAddIngredientGroup();
        },
        addIngredientGroup() {
            this.input_ingredient_group_mode = true
        },
        cancelAddIngredientGroup() {
            this.input_ingredient_group_mode = false
            this.new_ingredient_group_name = ''
        },
        deleteIngredientGroup(ingredient_category_id) {
            this.recipe.RecipeIngredients.splice(ingredient_category_id, 1)
            for(let i = ingredient_category_id; i < this.recipe.RecipeIngredients.length; ++i) {
                this.recipe.RecipeIngredients[i].id -= 1
            }
        },
        editIngredientGroup(ingredient_category_id) {
            this.edit_ingredient_group_id = ingredient_category_id
            
            this.edit_ingredient_order_id = 0
            this.edit_ingredient_category_id = -1

            this.recipe.RecipeIngredients[ingredient_category_id].edit_group_name = this.recipe.RecipeIngredients[ingredient_category_id].CategoryName
        },
        confirmEditIngredientGroup(ingredient_category_id) {
            this.edit_ingredient_group_id = -1
            this.recipe.RecipeIngredients[ingredient_category_id].CategoryName = this.recipe.RecipeIngredients[ingredient_category_id].edit_group_name
        },
        cancelEditIngredientGroup() {
            this.edit_ingredient_group_id = -1
        }
    },
    
    beforeMount() {
        console.log(this.$route.query.parent_recipe_id)
        if (this.$route.query.parent_recipe_id != 0)
            this.fetchSelectedRecipe(this.$route.query.parent_recipe_id)

        // if (this.recipe.Ingredients == null || this.recipe.Ingredients.length == 0 ) {
        //     this.insertIngredientCategory('Default Ingredients')
        // }
    }
}
</script>

<style>
    p:empty:not(:focus)::before {
        content: attr(data-placeholder);
        color: grey;
    }
    h3:empty:not(:focus)::before {
        content: attr(data-placeholder);
        color: grey;
    }
</style>