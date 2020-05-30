<template>
    <div class="v-cart-item">
        <img class="v-cart-item__image" :src=" require('@/assets/images/' + cart_item_data.image) " alt="img">
        <div class="v-cart-item-info">
            <p>{{cart_item_data.title}}</p>
            <p>{{cart_item_data.price}} р.</p>
        </div>

        <div class="v-cart-item__quantity">
            <p>Кол-во</p>
            <span class="quantity__tools">
                <span class="quantity__btn span" @click="incItem">+</span>
                {{cart_item_data.quantity}}
                <span class="quantity__btn span" @click="decItem">-</span>
            </span>
        </div>
            <button class = "btn_remove" @click="deleteFromCart">Удалить</button>
    </div>
</template>

<script>
    export default {
        name: "v-cart-item",
        props:{
            cart_item_data:{
                type: Object,
                default(){
                    return {}
                }
            }
        },
        methods:{
            deleteFromCart(){
                this.$emit('deleteFromCart')
            },
            incItem(){
                this.$emit('increment')
            },
            decItem(){
                this.$emit('decrement')
            }
        },
        mounted() {

        }
    }
</script>

<style lang="scss">
    .v-cart-item {
        display: flex;
        flex-wrap: nowrap;
        justify-content: space-between;
        align-items: center;
        box-shadow: 0 0 8px 0 #e0e0e0;
        padding: $padding*2;
        margin-bottom: $margin*2;
        &__image {
            max-width: 120px;
        }
        .quantity__btn {
            cursor: pointer;
        }
        .quantity__tools {
            user-select: none;
        }
    }
</style>