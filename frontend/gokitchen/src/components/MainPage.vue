<template>
    <subpage-template>
        <h5 class="mb-1 d-inline"  id="title_p">Welcome!</h5>
        <p class="mb-1" style="white-space: pre-line" id="content_p">Select a recipe or ingredient list from sidebar</p>
    </subpage-template>
</template>

<script>
export default {
    beforeMount() {
        this.$store.dispatch('setViewRecipeId', 0)
        this.$store.dispatch('setViewRecipeParentId', 0)
    }
}
</script>

<style>
</style>