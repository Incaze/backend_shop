<template>
    <div class="v-catalog">
        <v-links/>
        <h1>Каталог</h1>
        <div class="v-catalog__list">
            <v-catalog-item
                    v-for="product in PRODUCTS"
                    :key="product.id"
                    :product_data = "product"
                    @addToCart = "addToCart"
            />
        </div>

    </div>
</template>

<script>
    import vCatalogItem from "./v-catalog-item";
    import {mapActions, mapGetters} from 'vuex'
    import VLinks from "../utils/v-links";

    export default {
        name: "v-catalog",
        components: {
            VLinks,
            vCatalogItem
        },
        data() {
            return {}
        },
        methods: {
            ...mapActions([
               'GET_PRODUCTS_FROM_API',
                'ADD_TO_CART'
            ]),
            addToCart(data){
                this.ADD_TO_CART(data)
            }
        },
        mounted() {
            this.GET_PRODUCTS_FROM_API()
        },
        computed: {
            ...mapGetters([
                'PRODUCTS',
            ])
        }

    }
</script>

<style lang="scss">
    .v-catalog {
        &__list {
            display: flex;
            flex-wrap: wrap;
            justify-content: space-between;
            align-items: center;
        }
    }
</style>