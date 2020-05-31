<template>
    <div class="v-auth">
        <v-links/>
        <h2>Авторизация</h2>
        <div class="v-sign-in__form">
            <form class="forms">
                <input id="username" placeholder="Логин" required
                       type="text" v-model="body.username"/>
                <input id="password" placeholder="Пароль" required
                       type="password" v-model="body.password">
                <div>
                    <input @click="SignIn" type="submit" value="Авторизоваться">
                    Нет аккаунта? <router-link :to="{name: 'registration'}">Зарегистрируйтесь</router-link>
                </div>
            </form>
        </div>
        <v-auth-info/>
    </div>
</template>

<script>
    import {mapActions, mapGetters} from 'vuex'
    import router from '@/router/router'
    import VLinks from "../utils/v-links";
    import VAuthInfo from "../utils/v-auth-info";

    export default {
        name: "v-auth",
        components: {VAuthInfo, VLinks},
        props: {},
        data() {
            return {
                body: {
                    "username": "",
                    "password": ""
                },
            }
        },

        computed: {
            ...mapGetters([
                'JWT_AUTH',
            ]),

        },
        mounted() {
            this.AUTH_TEXT_STATUS = ""
        },
        methods: {
            ...mapActions([
                'SIGN_IN',
                'AUTH_UPDATE_STATUS'
            ]),
            SignIn() {
                this.SIGN_IN(this.body).then(() => {
                    if (this.JWT_AUTH.token) {
                        router.push('/');
                    } else {
                        this.AUTH_UPDATE_STATUS("Произошла ошибка регистрации")
                    }
                })
            },
        }
    }
</script>

<style lang="scss">

</style>