<template>
    <div class="v-main-wrapper">
        <v-catalog/>
        <v-cart/>
    </div>
</template>

<script>
    import vCatalog from "./v-catalog";
    import vCart from "./v-cart";

    export default {
        name: 'v-main-wrapper',
        components: {
            vCatalog,
            vCart,
        },
        props: {},
        data() {
            return {

            }
        },
        computed: {},
        methods: {},
        watch: {},
        mounted() {
        }
    }
</script>

<style>
    .v-main-wrapper {
        margin: 0 auto;
    }
</style>