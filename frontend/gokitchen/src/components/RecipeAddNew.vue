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

            <br>
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
                            <input type="text" class="form-control" name="ingredient_name" placeholder="Ingredient name" v-model="new_ingredient_name">
                        </td>
                        <td>
                            <input type="text" class="form-control" name="amount" placeholder="Amount" v-model="new_amount">
                            <input type="text" class="form-control" name="amountUnit" placeholder="Unit" v-model="new_amount_unit">
                        </td>
                        <td>
                            <!-- <form class="float-right d-inline" style="float: right;"> -->
                            <form class="float-right d-inline">
                                <button type="submit" class="btn btn-sm" @click.stop.prevent="addIngredient()">
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

                    <tr v-for="ingredient in recipe.RecipeIngredients">
                        <th scope="row">{{ ingredient.id }}</th>

                        <td v-if="edit_ingredient_order_id == ingredient.id">
                            <input type="text" class="form-control" name="ingredient_name" placeholder="Ingredient name" v-model="edit_ingredient_name">
                        </td>
                        <td v-else>{{ ingredient.IngredientName }}</td>

                        <td v-if="edit_ingredient_order_id == ingredient.id">
                            <input type="text" class="form-control" name="amount" placeholder="Amount" v-model="edit_amount">
                            <input type="text" class="form-control" name="amountUnit" placeholder="Unit" v-model="edit_amount_unit">
                        </td>
                        <td v-else>{{ ingredient.Amount }} {{ ingredient.AmountUnit }}</td>
                        
                        <td v-if="subrecipe_mode == 'rescale' && rescale_ing_id != ingredient.id - 1">
                            {{ ingredient.NewAmount }}
                        </td>
                        <td v-else-if="subrecipe_mode == 'rescale'">
                            <input type="number" class="form-control" size="4" name="amount" placeholder="Amount" v-model="rescale_amount">
                        </td>

                        <td v-if="subrecipe_mode == 'rescale'">
                            {{ ingredient.AmountUnit }}
                        </td>

                        <td v-if="edit_ingredient_order_id == ingredient.id">
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
                                <button type="submit" class="btn" title="Delete ingredient" @click.stop.prevent="deleteIngredient(ingredient.id)">
                                    <i class="text-danger float-right">
                                        <font-awesome-icon icon="trash-alt" size="lg" class="text-danger float-right"/>
                                    </i>
                                </button>
                            </form>
                            <form class="float-right d-inline" style="float: right;">
                                <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="editIngredient(ingredient.id)">
                                    <i class="text-danger float-right">
                                        <font-awesome-icon icon="pencil-alt" size="s" class="text-info float-right"/>
                                    </i>
                                </button>
                            </form>
                        </td>
                        <td v-else-if="subrecipe_mode == 'rescale' && rescale_ing_id == ingredient.id - 1">
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
                                <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="selectRescaleIngredient(ingredient.id)">
                                    <i class="text-danger float-right">
                                        <font-awesome-icon icon="pencil-alt" size="s" class="text-info float-right"/>
                                    </i>
                                </button>
                            </form>
                        </td>

                    </tr>
                </tbody>
            </table>

            <h5>Steps</h5>
            <p data-placeholder="Insert recipe steps here..." class="mb-1" id="steps_p" style="white-space: pre-line" contenteditable="true"></p>
            
        </form>
    </subpage-template>
</template>
    
<script>
    export default {
        data() {
        return {
            new_ingredient_name: '',
            new_amount: '',
            new_amount_unit: '',

            edit_ingredient_order_id: 0,
            edit_ingredient_name: '',
            edit_amount: '',
            edit_amount_unit: '',

            ingredient_order_id: 0,
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
            rescale_amount: 0,
            
        }  
    },
    methods: {
        fetchSelectedRecipe(recipe_id) {
            fetch('http://localhost:8000/get_recipe?' + 
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
                            this.recipe_detail.Ingredients[i]["id"] = i+1;
                        }
                        this.ingredient_order_id = this.recipe_detail.Ingredients.length
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
                        ingredientName: this.recipe.RecipeIngredients[i].IngredientName,
                        amount: this.recipe.RecipeIngredients[i].Amount,
                        amountUnit: this.recipe.RecipeIngredients[i].AmountUnit
                    }
                    ingredients.push(ingredient)
                }

            console.log(ingredients)

            fetch('http://localhost:8000/insert_recipe', 
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
                    recipeIngredients: ingredients,
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
        addIngredient() {
            this.ingredient_order_id += 1
            var ingredient = {
                id: this.ingredient_order_id,
                IngredientName: this.new_ingredient_name,
                Amount: this.new_amount,
                AmountUnit: this.new_amount_unit
            }
            this.recipe.RecipeIngredients.push(ingredient)

            this.new_ingredient_name = ''
            this.new_amount = ''
            this.new_amount_unit = ''
        },
        editIngredient(edit_ingredient_order_id) {
            this.edit_ingredient_order_id = edit_ingredient_order_id

            this.edit_ingredient_name = this.recipe.RecipeIngredients[edit_ingredient_order_id-1].IngredientName
            this.edit_amount = this.recipe.RecipeIngredients[edit_ingredient_order_id-1].Amount
            this.edit_amount_unit = this.recipe.RecipeIngredients[edit_ingredient_order_id-1].AmountUnit
        },
        deleteIngredient(ingredient_id) {
            ingredient_id -= 1
            this.recipe.RecipeIngredients.splice(ingredient_id, 1)
            for(let i = ingredient_id; i < this.recipe.RecipeIngredients.length; ++i) {
                this.recipe.RecipeIngredients[i].id -= 1
            }
        },
        confirmEditIngredient() {
            this.recipe.RecipeIngredients[this.edit_ingredient_order_id-1].IngredientName = this.edit_ingredient_name
            this.recipe.RecipeIngredients[this.edit_ingredient_order_id-1].Amount = this.edit_amount
            this.recipe.RecipeIngredients[this.edit_ingredient_order_id-1].AmountUnit = this.edit_amount_unit

            this.edit_ingredient_order_id = 0
        },
        cancelEditIngredient() {
            this.edit_ingredient_order_id = 0
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
                    this.recipe.RecipeIngredients[i].NewAmount = this.recipe.RecipeIngredients[i].Amount
                }

                this.recipe.RecipeRescaleInfo.RecipeID = this.recipe_detail.RecipeID
                this.recipe.RecipeRescaleInfo.IngredientID = this.recipe.RecipeIngredients[0].RecipeIngredientID
                if(isNaN(this.recipe.RecipeIngredients[0].Amount))
                    this.recipe.RecipeRescaleInfo.NewAmount = -1
                else
                    this.recipe.RecipeRescaleInfo.NewAmount = parseInt(this.recipe.RecipeIngredients[0].Amount)
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

                this.new_ingredient_name = ''
                this.new_amount = '',
                this.new_amount_unit = ''

                this.edit_ingredient_order_id = 0
                this.edit_ingredient_name = ''
                this.edit_amount = ''
                this.edit_amount_unit = ''

                this.ingredient_order_id = 0
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
        selectRescaleIngredient(ingredient_id) {
            if(!isNaN(this.recipe.RecipeIngredients[ingredient_id-1].NewAmount)) {
                this.rescale_ing_id = ingredient_id-1
                this.rescale_amount = this.recipe.RecipeIngredients[ingredient_id-1].NewAmount
            }
        },
        confirmRescaleIngredient() {
            // this.recipe.RecipeIngredients[this.rescale_ing_id].NewAmount = this.rescale_amount

            var scale = this.rescale_amount / this.recipe.RecipeIngredients[this.rescale_ing_id].Amount
            for(let i = 0; i < this.recipe.RecipeIngredients.length; ++i) {
                this.recipe.RecipeIngredients[i].NewAmount *= scale
            }

            this.recipe.RecipeRescaleInfo.RecipeID = this.recipe_detail.RecipeID
            this.recipe.RecipeRescaleInfo.IngredientID = this.recipe.RecipeIngredients[this.rescale_ing_id].RecipeIngredientID
            this.recipe.RecipeRescaleInfo.NewAmount = this.rescale_amount

            this.rescale_ing_id = -1
        },
        cancelRescaleIngredient() {
            this.rescale_ing_id = -1
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

            this.new_ingredient_name = ''
            this.new_amount = ''
            this.new_amount_unit = ''

            this.edit_ingredient_order_id = 0
            this.edit_ingredient_name = ''
            this.edit_amount = ''
            this.edit_amount_unit = ''
        }
    },
    beforeMount() {
        console.log(this.$route.query.parent_recipe_id)
        if (this.$route.query.parent_recipe_id != 0)
            this.fetchSelectedRecipe(this.$route.query.parent_recipe_id)
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