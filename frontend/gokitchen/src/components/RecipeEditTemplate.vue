<template>
    <subpage-template>
        <form>
            <div class="input-group mb-2">
                <input type="text" class="form-control" placeholder="Edit mode" readonly>
                <div class="input-group-append text-info">
                    <span class="input-group-text bg-white py-0">
                        <button type="submit" class="btn btn-sm text-info" @click.stop.prevent="$.emit('confirm-action')">
                            <h6>
                                <font-awesome-icon icon="plus-circle" size="lg"/>
                            </h6>
                        </button>
                    </span>
                </div>
                <div class="input-group-append text-info">
                    <span class="input-group-text bg-white py-0">
                        <button type="submit" class="btn btn-sm text-info"  @click.stop.prevent="$.emit('cancel-action')">
                            <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                        </button>
                    </span>
                </div>
            </div>
            <h5 data-placeholder="Insert recipe title here..."class="mb-1" id="title_p" contenteditable="true" ></h5>
            <p data-placeholder="Insert recipe steps here..." class="mb-1" id="steps_p" style="white-space: pre-line" contenteditable="true"></p>
            
        </form>
    </subpage-template>
</template>

<script>
</script>

<style>
    p:empty:not(:focus)::before {
        content: attr(data-placeholder);
        color: grey;
    }
    h5:empty:not(:focus)::before {
        content: attr(data-placeholder);
        color: grey;
    }
</style>