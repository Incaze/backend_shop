<template>
    <div class="v-cart">
            <v-links/>
    <div> <h1>Корзина</h1></div>
        <p v-if="!cart_data.length">Ваша корзина пуста</p>
        <v-cart-item
            v-for="(item, index) in cart_data"
            :key="item.id"
            :cart_item_data="item"
            @deleteFromCart="deleteFromCart(index)"
            @increment="increment(index)"
            @decrement="decrement(index)"
        />
        <div class="v-cart__total info_bg">
            <p class="info_text">Всего: {{cartTotalCost}} р.</p>
        </div>
    </div>
</template>

<script>
    import vCartItem from "./v-cart-item";
    import {mapActions} from 'vuex';
    import VLinks from "../utils/v-links";

    export default {
        name: "v-cart",
        components:{
            VLinks,
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

</style>