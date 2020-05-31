<template>
    <div class="v-catalog-item">
        <img class="v-catalog-item__image" :src=" require('@/assets/images/' + product_data.image) " alt="img">
        <p class="v-catalog-item__name">{{product_data.title}}</p>
        <p class="v-catalog-item__price">Цена: {{product_data.price}} р</p>
        <p class="v-catalog-item__description">{{product_data.description}}</p>
        <p class="v-catalog-item__rating">Оценка: {{product_data.rating}}/5.0</p>
        <button
                v-if="JWT_AUTH.token.length"
                class="v-catalog-item__add_to_cart_btn btn"
                @click="addToCart"
        >Добавить в корзину
        </button>
    </div>
</template>

<script>
    import {mapGetters} from "vuex";

    export default {
        name: "v-catalog-item",
        props: {
            product_data: {
                type: Object,
                default() {
                    return {}
                }
            }
        },
        data(){
            return {}
        },
        computed: {
            ...mapGetters([
                'JWT_AUTH'
            ])
        },
        methods: {
            addToCart(){
                this.$emit('addToCart',this.product_data)
            }
        },
        mounted() {
            this.$set(this.product_data, 'quantity', 1)
        }
    }
</script>

<style lang="scss">
    .v-catalog-item {
        flex-basis: 25%;
        box-shadow: 0 0 8px 0 #e0e0e0;
        padding: $padding * 2;
        margin-bottom: $margin * 2;

        &__image {
            width: 120px;
        }
        &__description {
            text-align: justify;
        }
    }
</style>