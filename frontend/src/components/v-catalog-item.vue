<template>

    <div class="v-catalog-item">
        <p class="v-catalog-item__name">{{product_data.title}}</p>
        <p class="v-catalog-item__price">Price: {{product_data.price}}</p>
        <p class="v-catalog-item__description">{{product_data.description}}</p>
        <p class="v-catalog-item__rating">{{product_data.rating}}</p>
        <button>Add to cart</button>
    </div>
</template>

<script>
    export default {
        name: "v-catalog-item",
        props: {
            product_data: {
                type: Object,
                default() {
                    return {}
                }
            }
        }
    }
</script>

<style lang="scss">
    .v-catalog-item {
        flex-basis: 25%;
        box-shadow: 0 0 8px 0 #e0e0e0;
        padding: $padding * 2;
        margin-bottom: $margin * 2;
    }
</style>