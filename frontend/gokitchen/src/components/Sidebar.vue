<template>
<div class="col-auto col-md-3 col-lg-3 col-xl-3 px-sm-2 px-0 bg-dark">
    <div class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
        <span style="font-size:16.0pt" class="mt-2 mb-2 fs-5 d-none d-sm-inline">GoKitchen</span>
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
                        <i v-if="recipe.ParentRecipeID > 0">-</i>
                        <button type="submit" class="astext" @click.stop.prevent="viewRecipe(recipe.RecipeID)">
                            <i class="fs-4 bi-house"></i> <span class="ms-1 d-none d-sm-inline text-white">{{ recipe.Title }}</span>
                            <!-- <font-awesome-icon icon="plus-circle" size="s" class="fa fa-caret-down"/> -->
                        </button>

                        <li class="dropdown-container nav-item" v-for="subrecipe in recipe.SubrecipeList">
                            <!-- <form  class="d-inline"> -->
                                <nobr>
                                    <!-- <i>-</i> -->
                                    <button type="submit" class="astext" @click.stop.prevent="viewRecipe(subrecipe.RecipeID)">
                                        <i class="fs-4 bi-house"></i> <span class="ms-1 d-none d-sm-inline text-white">{{ subrecipe.Title }}</span>
                                    </button>
                                </nobr>
                            <!-- </form> -->
                        </li>
                    <!-- </nobr> -->
                <!-- </form> -->
            </li>
        </ul>
    </div>
</div>
</template>

<script>
export default {
    data() {
        return {
            recipe_list: []
        }
    },
    methods: {
        fetchRecipeList() {
            fetch('http://localhost:8000/list_recipe', {
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
                    console.log(data)
                    console.log(this.recipe_list[1].Title)
                    console.log(this.recipe_list[1].SubrecipeList)
                })
            })
        },
        viewIngredientList() {
            this.$router.push({ path: 'ingredient_info_view' });
        },
        addNewRecipe() {
            this.$router.push({ path: 'recipe_add_new', query: { parent_recipe_id: 0 } });
        },
        viewRecipe(recipe_id) {
            this.$router.push({ path: 'recipe_view', query: { recipe_id: recipe_id } });
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
  padding-left: 15px;
}
</style>