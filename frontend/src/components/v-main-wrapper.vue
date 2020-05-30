<template>
    <div class="v-main-wrapper">
        <keep-alive>
            <router-view/>
        </keep-alive>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex';

    export default {
        name: 'v-main-wrapper',
        data() {
            return {}
        },
        computed: {
            ...mapGetters([
                'CART'
            ])
        }
    }
</script>

<style>
    .v-main-wrapper {
        margin: 0 auto;
    }
</style>