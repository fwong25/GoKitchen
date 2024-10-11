<template>
<div class="col-auto col-md-3 col-lg-3 col-xl-3 px-sm-2 px-0 bg-dark">
    <div class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
        <button style="font-size:16.0pt" type="submit" class="text-white mt-2 mb-2 fs-5 d-none d-sm-inline astext" @click.stop.prevent="backToHome()">
            <span style="font-size:16.0pt" class="mt-2 mb-2 fs-5 d-none d-sm-inline">GoKitchen</span>
        </button>
        <div class="input-group mb-2">
            <form>
                <div class="input-group-append text-info">
                    <span class="input-group-text bg-dark py-0">
                        <button type="submit" class="btn btn-sm text-info" @click.stop.prevent="addNewRecipe()">
                            <font-awesome-icon icon="plus-circle" size="lg"/>
                        </button>
                    </span>
                </div>
            </form>
        </div>
        <button type="submit" class="astext" @click.stop.prevent="viewIngredientList()">
            <i class="fs-4 bi-house"></i> <span class="ms-1 d-none d-sm-inline text-info">Ingredients Info</span>
        </button>
        <ul class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start" id="menu">
            <li class="nav-item" v-for="recipe in recipe_list">
                <!-- <form  class="d-inline"> -->
                    <!-- <nobr> -->
                        <!-- <i v-if="recipe.ParentRecipeID > 0">-</i> -->
                        <button type="submit" class="astext" @click.stop.prevent="viewRecipe(recipe.RecipeID, recipe.ParentRecipeID)">
                            <i class="fs-4 bi-house"></i>
                            <span v-if="$store.getters.viewRecipeId == recipe.RecipeID" class="ms-1 d-none d-sm-inline text-warning">{{ recipe.Title }}</span>
                            <span v-else class="ms-1 d-none d-sm-inline text-white">{{ recipe.Title }}</span>
                        </button>
                        <button  v-if="recipe.SubrecipeList.length != 0" type="submit" class="astext" @click.stop.prevent="toggleSubrecipeList(recipe.RecipeID)">
                            &nbsp; &nbsp;<font-awesome-icon icon="caret-down" size="s" class="text-white"/>
                        </button>

                        <div v-if="$store.getters.subrecipeListId == recipe.RecipeID || $store.getters.viewRecipeId == recipe.RecipeID" class="dropdown-container" v-for="subrecipe in recipe.SubrecipeList">
                        <li class="dropdown-item nav-item">
                            <!-- <form  class="d-inline"> -->
                                <nobr>
                                    <!-- <i>-</i> -->
                                    <button type="submit" class="astext text-start" @click.stop.prevent="viewRecipe(subrecipe.RecipeID, subrecipe.ParentRecipeID)">
                                        <i class="fs-4 bi-house"></i><span class="text-wrap ms-1 d-none d-sm-inline text-white">{{ subrecipe.Title }}</span>
                                    </button>
                                </nobr>
                            <!-- </form> -->
                        </li>
                        </div>
                    <!-- </nobr> -->
                <!-- </form> -->
            </li>
        </ul>
    </div>
</div>
</template>

<script>
import { host_info } from './../main.js'
// import { subrecipe_list_id } from './../main.js'

export default {
    data() {
        return {
            recipe_list: [],
            subrecipe_list_id: 0
        }
    },
    methods: {
        fetchRecipeList() {
            fetch('http://' + host_info.backend_host + '/list_recipe', {
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
                    this.recipe_list = data.recipe_list
                })
            })
        },
        viewIngredientList() {
            this.$router.push({ path: 'ingredient_info_view' });
        },
        addNewRecipe() {
            this.$router.push({ path: 'recipe_add_new', query: { parent_recipe_id: 0 } });
        },
        viewRecipe(recipe_id, parent_recipe_id) {
            if (parent_recipe_id != 0) {
                this.$store.dispatch('setSubrecipeListId', parent_recipe_id)
            }
            else {
                this.$store.dispatch('setSubrecipeListId', recipe_id)
            }
            this.$store.dispatch('setViewRecipeId', recipe_id)
            this.$router.push({ path: 'recipe_view', query: { recipe_id: recipe_id } });

        },
        backToHome() {
            this.$router.push({ path: '/' });
        },
        toggleSubrecipeList(recipe_id) {
            // if(this.subrecipe_list_id != recipe_id)
            //     this.subrecipe_list_id = recipe_id
            // else
            //     this.subrecipe_list_id = 0

            if(this.$store.getters.subrecipeListId != recipe_id)
                this.$store.dispatch('setSubrecipeListId', recipe_id)
            else
                this.$store.dispatch('setSubrecipeListId', 0)

            console.log(this.$store.getters.subrecipeListId)
        }
    },
    beforeMount() {
        this.fetchRecipeList()
    }
}
</script>

<style>
.dropdown-container {
  background-color: #262626;
}
.dropdown-item {
  padding-left: 15px;
}
</style>