<script>
export default {
  mounted() {
        document.body.classList.add('bg-light')
    },
    destroyed() {
        document.body.classList.remove('bg-light')
    }
}
</script>

<template>
  <router-view :key="$route.fullPath"></router-view>
</template>

<style>
  .astext {
      background:none;
      border:none;
      margin:0;
      padding:0;
      cursor: pointer;
  }
  .header-font {
    color: rgb(78, 121, 174);
  }
</style>
