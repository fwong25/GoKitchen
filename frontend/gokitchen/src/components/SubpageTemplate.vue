<template>
    <div class="container-fluid">
        <div class="row flex-nowrap">
            <sidebar />
            <div class="col py-3">
                <div class="container">
                    <div class="row mt-1">
                        <div class="col-md-10 offset-md-1">
                            <div class="card">
                                <div class="card-header shadow-sm bg-white">
                                    <h1 class="display-5 header-font"><font-awesome-icon icon="check-double" /> GoKitchen</h1>
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
</script>

<style>
</style>