<template>
    <div class="container-fluid">
        <div class="row flex-nowrap">
            <sidebar id="sidebar" v-show="sidebar_visible"/>
            <div class="col py-3">
                <div class="container">
                    <div class="row mt-1">
                        <div class="col-md-10 offset-md-1">
                            <div class="card">
                                <div class="card-header shadow-sm bg-white">
                                    <h1 class="display-5 header-font">
                                        <button type="button" class="header-font astext" @click.stop.prevent="toggleSidebar()">
                                            <font-awesome-icon icon="check-double" /> 
                                        </button>
                                        <button id="test" type="submit" class="header-font astext" @click.stop.prevent="backToHome()">
                                            &nbsp;GoKitchen    
                                        </button>
                                    </h1>
                                </div>
                                <div class="card-body">
                                    <slot />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            sidebar_visible: true
        }
    },
    methods: {
        backToHome() {
            this.$router.push({ path: '/' });
        },
        toggleSidebar() {
            this.sidebar_visible = !this.sidebar_visible
        }
    }
}
</script>

<style>
</style>