<template>
    <div class="v-cart">
        <router-link :to="{name: 'catalog'}">
            <div class="v-cart__link_to_catalog">Каталог</div>
        </router-link>
        <h1>Корзина</h1>
        <p v-if="!cart_data.length">Ваша корзина пуста</p>
        <v-cart-item
            v-for="(item, index) in cart_data"
            :key="item.id"
            :cart_item_data="item"
            @deleteFromCart="deleteFromCart(index)"
            @increment="increment(index)"
            @decrement="decrement(index)"
        />
        <div class="v-cart__total">
            <p class="total_name">Всего: {{cartTotalCost}} р.</p>
        </div>
    </div>
</template>

<script>
    import vCartItem from "./v-cart-item";
    import {mapActions} from 'vuex';

    export default {
        name: "v-cart",
        components:{
            vCartItem
        },
        props:{
            cart_data:{
                type: Array,
                default() {
                    return [];
                }
            }
        },
        methods:{
            ...mapActions([
                'DELETE_FROM_CART',
                'INCREMENT_CART_ITEM',
                'DECREMENT_CART_ITEM',
            ]),
            increment(index){
                this.INCREMENT_CART_ITEM(index)
            },
            decrement(index){
                this.DECREMENT_CART_ITEM(index)
            },
            deleteFromCart(index){
                this.DELETE_FROM_CART(index)
            }
        },
        computed:{
            cartTotalCost(){
                return this.cart_data.reduce((res, item) => res + item.price * item.quantity, 0)
            }
        }
    }
</script>

<style lang="scss">
    .v-cart{
        margin-bottom: 100px;
        &__link_to_catalog {
            position: absolute;
            top: 10px;
            right: 10px;
            padding: $padding * 2;
            border: solid 1px #aeaeae;
            background: #ffffff;
        }
        &__total {
            position: fixed;
            bottom: 0;
            right: 0;
            left: 0;
            padding: $padding*2 $padding*3;
            display: flex;
            justify-content: center;
            background: $blue-bg;
            color: #ffffff;
            font-size: 20px;
        }
        .total__name {
            margin-right: $margin * 2;
        }
    }
</style>