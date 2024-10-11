<template>
    <subpage-template>
        <h3 class="mb-1 d-inline"  id="title_p">Ingredients</h3>

        <br>
        <table class="table">
            <thead>
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">Ingredient</th>
                    <th scope="col">Amount</th>
                    <th scope="col">Cost</th>
                    <th></th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th scope="row"></th>
                    <td>
                        <input type="text" class="form-control" name="ingredient_name" placeholder="Ingredient name" v-model="new_ingredient_name">
                    </td>
                    <td>
                        <input type="number" class="form-control" name="amount" placeholder="Amount" v-model="new_amount">
                        <input type="text" class="form-control" name="amountUnit" placeholder="Unit" v-model="new_amount_unit">
                    </td>
                    <td>
                        <input type="number" class="form-control" name="Cost" placeholder="Cost" v-model="new_cost">
                    </td>
                    <td>
                        <form class="float-right d-inline" style="float: right;">
                            <button type="submit" class="btn btn-sm" @click.stop="addIngredient()">
                                <font-awesome-icon icon="plus-circle" size="xl" class="text-info"/>
                            </button>
                        </form>
                    </td>
                    <td>
                    </td>
                </tr>
                <tr v-for="ingredient in ingredient_list">
                    <th scope="row">{{ ingredient.id }}</th>
                    <td v-if="edit_ingredient_order_id == ingredient.id">
                        <input type="text" class="form-control" name="ingredient_name" placeholder="Ingredient name" v-model="edit_ingredient_name">
                    </td>
                    <td v-else>{{ ingredient.IngredientName }}</td>

                    <!-- <td v-if="ingredient.Amount == 0 || ingredient.AmountUnit == 'unknown'">-</td>
                    <td v-else>{{ ingredient.Amount }} {{ ingredient.AmountUnit }}</td> -->

                    <td v-if="edit_ingredient_order_id == ingredient.id">
                        <input type="number" class="form-control" name="amount" placeholder="Amount" v-model="edit_amount">
                        <input type="text" class="form-control" name="amountUnit" placeholder="Unit" v-model="edit_amount_unit">
                    </td>
                    <td v-else>
                        <div v-if="ingredient.Amount == 0 || ingredient.AmountUnit == 'unknown'">-</div>
                        <div v-else>{{ ingredient.Amount }} {{ ingredient.AmountUnit }}</div>
                    </td>

                    <!-- <td v-if="ingredient.Cost == 0">-</td>
                    <td v-else>Rp. {{ ingredient.Cost }}</td> -->
                    
                    <td v-if="edit_ingredient_order_id == ingredient.id">
                        <input type="number" class="form-control" name="Cost" placeholder="Cost" v-model="edit_cost">
                    </td>
                    <td v-else>
                        <div v-if="ingredient.Cost == 0">-</div>
                        <div v-else>Rp. {{ ingredient.Cost }}</div>
                    </td>

                    <td v-if="edit_ingredient_order_id == ingredient.id">
                        <form class="float-right d-inline" style="float: right;">
                            <button type="submit" class="btn"title="Edit ingredient" @click.stop="confirmEditIngredient()">
                                <i class="text-danger float-right">
                                    <font-awesome-icon icon="check" size="xl" class="text-info float-right"/>
                                </i>
                            </button>
                        </form>
                    </td>
                    <td v-else>
                        <form class="float-right d-inline" style="float: right;">
                            <button type="submit" class="btn"title="Edit ingredient" @click.stop.prevent="editIngredient(ingredient.id)">
                                <i class="text-danger float-right">
                                    <font-awesome-icon icon="pencil-alt" size="s" class="text-info float-right"/>
                                </i>
                            </button>
                        </form>
                    </td>

                    <td v-if="edit_ingredient_order_id == ingredient.id">
                        <form class="float-right d-inline" style="float: right;">
                            <button type="submit" class="btn" title="Cancel edit ingredient" @click.stop.prevent="cancelEditIngredient()">
                                <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                            </button>
                        </form>
                    </td>
                    <td v-else>
                        <form class="float-right d-inline" style="float: right;">
                            <button type="submit" class="btn" title="Delete ingredient" @click.stop="deleteIngredient(ingredient.IngredientID)">
                                <i class="text-danger float-right">
                                    <font-awesome-icon icon="trash-alt" size="lg" class="text-danger float-right"/>
                                </i>
                            </button>
                        </form>
                    </td>
                    
                </tr>
            </tbody>
        </table>
    </subpage-template>
</template>

<script>
import { host_info } from './../main.js'
export default {
    data() {
        return {
            ingredient_order_id: 0,
            ingredient_list: [],

            new_ingredient_name: '',
            new_amount: null,
            new_amount_unit: '',
            new_cost: null,

            edit_ingredient_order_id: 0,

            edit_ingredient_name: '',
            edit_amount: null,
            edit_amount_unit: '',
            edit_cost: null,
        }  
    },
    methods: {
        fetchIngredientList() {
            fetch('http://' + host_info.backend_host + '/list_ingredient_info', 
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
                    console.log(data.ingredient_list)
                    this.ingredient_list = data.ingredient_list
                    
                    for (let i = 0; i < this.ingredient_list.length; ++i) {
                        this.ingredient_list[i]["id"] = i+1;
                        if (this.ingredient_list[i].AmountUnit == "unknown") {
                            this.ingredient_list[i].AmountUnit = ""
                        }
                    }
                })
            })
        },
        addIngredient() {
            fetch('http://' + host_info.backend_host + '/insert_ingredient_info', 
                    {
                    method: 'POST',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    },
                    body: JSON.stringify({
                        ingredientName: this.new_ingredient_name,
                        amount: this.new_amount,
                        amountUnit: this.new_amount_unit,
                        cost: this.new_cost
                    })
                })
                .then((response) => {
                    console.log(response)
                    response.json().then((data) => {
                        console.log(data)
                        console.log(data.ingredient_id)
                    })
                })
        },
        editIngredient(edit_ingredient_order_id) {
            this.edit_ingredient_order_id = edit_ingredient_order_id
            this.edit_ingredient_name = this.ingredient_list[edit_ingredient_order_id-1].IngredientName
            this.edit_amount = this.ingredient_list[edit_ingredient_order_id-1].Amount
            this.edit_amount_unit = this.ingredient_list[edit_ingredient_order_id-1].AmountUnit
            this.edit_cost = this.ingredient_list[edit_ingredient_order_id-1].Cost
        },
        deleteIngredient(ingredient_id) {
            fetch('http://' + host_info.backend_host + '/delete_ingredient_info?' + 
                    new URLSearchParams({
                        ingredient_id: ingredient_id
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
                })
        },
        confirmEditIngredient() {
            fetch('http://' + host_info.backend_host + '/update_ingredient_info', 
                    {
                    method: 'PUT',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    },
                    body: JSON.stringify({
                        ingredientID: this.ingredient_list[this.edit_ingredient_order_id-1].IngredientID,
                        ingredientName: this.edit_ingredient_name,
                        amount: this.edit_amount,
                        amountUnit: this.edit_amount_unit,
                        cost: this.edit_cost
                    })
                })
                .then((response) => {
                    console.log(response)
                    // response.json().then((data) => {
                    //     console.log(data)
                    // })
                })
        },
        cancelEditIngredient() {
            this.edit_ingredient_order_id = 0
        }
    },
    beforeMount() {
        this.$store.dispatch('setViewRecipeId', 0)
        this.$store.dispatch('setViewRecipeParentId', 0)
        this.fetchIngredientList()
    }
}
</script>

<style>
</style>