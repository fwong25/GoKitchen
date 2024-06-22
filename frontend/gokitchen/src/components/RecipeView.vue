<template>
    <subpage-template>
        <small class="text-muted">Created on {{ selected_recipe.CreatedDate }}, </small>
        <small class="text-muted">last modified on {{ selected_recipe.LastModifiedDate}}</small>
        <br>
        <h3 class="mb-1 d-inline"  id="title_p"></h3>
        <form class="float-right d-inline" style="float: right;">
            <button type="submit" class="btn" title="Delete recipe" @click.stop.prevent="deleteRecipe()">
                <i class="text-danger float-right">
                    <font-awesome-icon icon="trash-alt" size="lg" class="text-danger float-right"/>
                </i>
            </button>
        </form>
        <form class="float-right d-inline" style="float: right;">
            <button type="submit" class="btn"title="Edit recipe" @click.stop.prevent="editRecipe()">
                <i class="text-danger float-right">
                    <font-awesome-icon icon="pencil-alt" size="lg" class="text-danger float-right"/>
                </i>
            </button>
        </form>
        <form v-if="selected_recipe.ParentRecipeID == 0" class="float-right d-inline" style="float: right;">
            <button type="submit" class="btn"  title="Add new subrecipe" @click.stop.prevent="addNewSubRecipe()">
                <i class="text-danger float-right">
                    <font-awesome-icon icon="plus" size="lg" />
                </i>
            </button>
        </form>

        <br>
        
        <div v-for="ingredient_category in selected_recipe.Ingredients">
            <br>
            <h5>{{ ingredient_category.CategoryName }}</h5>
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <!-- <th scope="col">Ingredient</th> -->
                        <th scope="col">Ingredients</th>
                        <th scope="col">Amount</th>
                        <th scope="col">Cost</th>
                    </tr>
                </thead>
                <tbody>
                    <!-- <tr>
                        <th scope="row">1</th>
                        <td>Flour</td>
                        <td>100g</td>
                        <td>1000</td>
                    </tr> -->
                    <tr v-for="ingredient in ingredient_category.Ingredients">
                        <th scope="row">{{ ingredient.id }}</th>
                        <td>{{ ingredient.IngredientName }}</td>
                        <td>{{ ingredient.Amount }} {{ ingredient.AmountUnit }}</td>
                        <td v-if="ingredient.Cost == 0">-</td>
                        <td v-else>Rp. {{ ingredient.Cost }}</td>
                    </tr>
                    <tr>
                        <th scope="row"></th>
                        <th>Total Cost</th>
                        <td></td>
                        <td>Rp. {{ ingredient_category.TotalCost }}</td>
                        <!-- <td>Rp. {{ selected_recipe.TotalCost }}</td> -->
                    </tr>
                </tbody>
            </table>
        </div>

        <h5>Steps</h5>
        <p class="mb-1" style="white-space: pre-line" id="steps_p"></p>
    </subpage-template>
</template>

<script>
import { host_info } from './../main.js'
export default {
    data() {
        return {
            ingredient_order_id: 0,
            selected_recipe: {
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
            }
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
                    this.selected_recipe = data.recipe

                    const title_p = document.getElementById("title_p");
                    const steps_p = document.getElementById("steps_p");
                    title_p.innerHTML = this.selected_recipe.Title;
                    steps_p.innerHTML = this.selected_recipe.Steps;

                    if (this.selected_recipe.Ingredients != null) {
                        for (let i = 0; i < this.selected_recipe.Ingredients.length; ++i) {
                            if (this.selected_recipe.Ingredients[i].Ingredients != null)
                                for (let j = 0; j < this.selected_recipe.Ingredients[i].Ingredients.length; ++j) {
                                    this.selected_recipe.Ingredients[i].Ingredients[j]["id"] = j+1;
                                }
                        }
                    }     
                })
            })
        },
        deleteRecipe() {
            fetch('http://' + host_info.backend_host + '/delete_recipe?' + 
                    new URLSearchParams({
                        recipe_id: this.$route.query.recipe_id
                    }), 
                    {
                    method: 'DELETE',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    }
                })
                .then((response) => {
                    console.log(response)

                    if (this.selected_recipe.ParentRecipeID === 0) {
                        this.$router.push({ path: '/'});
                    } else {
                        this.$router.push({ path: 'recipe_view',
                            query: { recipe_id: this.selected_recipe.ParentRecipeID }
                        });
                    }
                })
        },
        editRecipe() {
            this.$router.push({ path: 'recipe_edit',
                query: { recipe_id: this.$route.query.recipe_id }
            });
        },
        addNewSubRecipe() {
            this.$router.push({ path: 'recipe_add_new', query: { parent_recipe_id: this.$route.query.recipe_id } });
        }
    },
    beforeMount() {
        this.fetchSelectedRecipe(this.$route.query.recipe_id)
    }
}
</script>

<style>
</style>