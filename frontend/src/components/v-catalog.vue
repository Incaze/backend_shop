<template>
    <div class="v-catalog">
        <h1>Catalog</h1>
        <div class="v-catalog__list">
            <v-catalog-item
                    v-for="product in products"
                    :key="product.id"
                    :product_data = "product"
            />
        </div>

    </div>
</template>

<script>
    import vCatalogItem from "./v-catalog-item";
    export default {
        name: "v-catalog",
        components: {
            vCatalogItem
        },
        data() {
            {
                return {
                    products: [
                        {
                            id: 0,
                            title: "T-shirt 0",
                            image: "0.jpg",
                            description: "SADSX dsaD SADasd",
                            price: 2100.234234234,
                            rating: 3,
                            category: "Мужские"
                        },
                        {
                            id: 1,
                            title: "T-shirt 1",
                            image: "1.jpg",
                            description: "TEXT 1",
                            price: 4200.51524,
                            rating: 5,
                            category: "Мужские"
                        },
                        {
                            id: 2,
                            title: "T-shirt 2",
                            image: "2.jpg",
                            description: "TEXT 2",
                            price: 5300.1245512,
                            rating: 4,
                            category: "Мужские"
                        },
                        {
                            id: 3,
                            title: "T-shirt 3",
                            image: "3.jpg",
                            description: "TEXT 3",
                            price: 6500.3522314,
                            rating: 2,
                            category: "Мужские"
                        },
                        {
                            id: 4,
                            title: "T-shirt 4",
                            image: "4.jpg",
                            description: "TEXT 4",
                            price: 8700.4124123,
                            rating: 4,
                            category: "Мужские"
                        },
                        {
                            id: 5,
                            title: "T-shirt 5",
                            image: "5.jpg",
                            description: "TEXT 5",
                            price: 2340.15,
                            rating: 3,
                            category: "Мужские"
                        },
                    ]
                }
            }
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
        &__link_to_cart {
            position: fixed;
            top: 80px;
            right: 10px;
            padding: $padding*2;
            border: solid 1px #aeaeae;
            background: #ffffff;
        }
    }

</style>