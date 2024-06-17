import "bootstrap/dist/css/bootstrap.css"
/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'
import { faCheckDouble, faPlusCircle, faTrashAlt, faPencilAlt, faPlus, faCheck } from "@fortawesome/free-solid-svg-icons";
import { faPlusSquare } from "@fortawesome/free-regular-svg-icons";
library.add(faCheckDouble, faPlusCircle, faTrashAlt, faPencilAlt, faPlus, faCheck, faPlusSquare)

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
    ]
});

const app = createApp(App)
app.use(router)

app.component('sidebar', Sidebar)
// app.component('note-edit-template', NoteEditTemplate)
app.component('subpage-template', SubpageTemplate)
app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')

import "bootstrap/dist/js/bootstrap.js"