import json from './../../../ipcfg.json'
console.log(json.Configs[json.frontend_server_idx])
export const host_info = {
    frontend_host: json.Configs[json.frontend_server_idx].Ip + ':' + json.Configs[json.frontend_server_idx].Port,
    frontend_ip: json.Configs[json.frontend_server_idx].Ip,
    fontend_port: json.Configs[json.frontend_server_idx].Port,
    backend_host: json.Configs[json.backend_server_idx].Ip + ':' + json.Configs[json.backend_server_idx].Port,
}
console.log(host_info)

import "bootstrap/dist/css/bootstrap.css"
/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'
import { faCheckDouble, faPlusCircle, faTrashAlt, faPencilAlt, faPlus, faCheck, faCaretDown } from "@fortawesome/free-solid-svg-icons";
import { faPlusSquare } from "@fortawesome/free-regular-svg-icons";
library.add(faCheckDouble, faPlusCircle, faTrashAlt, faPencilAlt, faPlus, faCheck, faPlusSquare, faCaretDown)

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import App from './App.vue'
import MainPage from './components/MainPage.vue'
import RecipeView from './components/RecipeView.vue'
import IngredientInfoView from './components/IngredientInfoView.vue'
import RecipeAddNew from './components/RecipeAddNew.vue'
import RecipeEdit from './components/RecipeEdit.vue'
import Sidebar from './components/Sidebar.vue'
// import NoteEditTemplate from './components/NoteEditTemplate.vue'
import SubpageTemplate from './components/SubpageTemplate.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '', component: MainPage },
        { path: '/recipe_add_new', component: RecipeAddNew },
        { path: '/recipe_view', component: RecipeView },
        { path: '/ingredient_info_view', component: IngredientInfoView },
        { path: '/recipe_edit', component: RecipeEdit },
        { path: '/sidebar', component: Sidebar },
    ]
});

const app = createApp(App)
app.use(router)

// export var subrecipe_list_id = 0
import { createStore } from 'vuex'
const store = createStore({
    state () {
      return {
        subrecipe_list_id: 0
      }
    },
    getters: {
        subrecipeListId: function (state) {
            console.log('get subrecipe_list_id')
            console.log(state.subrecipe_list_id)
            return state.subrecipe_list_id
        }
    },
    actions: {
      setSubrecipeListId (context, payload) {
        const subrecipe_list_id = payload
        context.commit('UPDATE_SUBRECIPE_LIST_ID', subrecipe_list_id)
      }
    },
    mutations: {
        UPDATE_SUBRECIPE_LIST_ID(state, payload) {
            state.subrecipe_list_id = payload
        }
    }
})
app.use(store)

app.component('sidebar', Sidebar)
// app.component('note-edit-template', NoteEditTemplate)
app.component('subpage-template', SubpageTemplate)
app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')

import "bootstrap/dist/js/bootstrap.js"