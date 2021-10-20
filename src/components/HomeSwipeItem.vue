<template>
    <div class="home_swipe_item">
        <img v-lazy="data.picUrl" alt>
        <div class="label">
            {{ data.label }}
        </div>
    </div>
</template>

<script>

export default {
    name: 'HomeSwipeItem',
    props: {
        data: {
            type: Object,
            default: () => ({})
        }
    }
}
</script>

<style lang="scss" scoped>
.home_swipe_item {
    width: 100%;
    height: 100%;
    img {
        width: 100%;
        height: 100%;
    }
    .label {
        color: white;
        background: rgba($color: black, $alpha: 0.5);
        position: absolute;
        bottom: 0;
        left: 0;
        right: 0;
        height: 30px;
        line-height: 30px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        padding: 0 10px;
    }
}
</style>