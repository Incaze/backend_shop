<template>
    <div class="v-links">
        <router-link :to="{name: 'catalog'}">
            <div class="v-links__link_to_catalog links">Каталог</div>
        </router-link>
        <router-link v-if="!JWT_AUTH.token.length" :to="{name: 'registration'}">
            <div class="v-links__link_to_registration links">Регистрация</div>
        </router-link>
        <router-link v-if="!JWT_AUTH.token.length" :to="{name: 'sign-in'}">
            <div class="v-links__link_to_sign_in links">Авторизация</div>
        </router-link>

        <div
                v-if="JWT_AUTH.token.length"
                class="v-links__link_to_sign_out links"
                @click="SignOut"
        >Выход</div>

        <p v-if="JWT_AUTH.token.length">Добрый день, {{JWT_AUTH.username}}!</p>
        <router-link v-if="JWT_AUTH.token.length" :to="{name: 'cart', params: {cart_data: CART}}">
            <div class="v-links__link_to_cart links">Корзина: {{CART.length}}</div>
        </router-link>
    </div>
</template>

<script>
    import {mapActions, mapGetters} from "vuex";

    export default {
        name: "v-links",
        computed: {
            ...mapGetters([
                'CART',
                'JWT_AUTH'
            ])
        },
        methods: {
            ...mapActions([
                'SIGN_OUT'
            ]),
            SignOut() {
                this.SIGN_OUT;
                window.location.reload();
            }
        }
    }
</script>

<style lang="scss">
    .v-links {
        &__link_to_cart {
            right: 110px;
        }
        &__link_to_catalog {
            right: 10px;
        }
        &__link_to_sign_in {
            left: 10px;
        }
        &__link_to_registration {
            left: 150px;
        }
    }
</style>