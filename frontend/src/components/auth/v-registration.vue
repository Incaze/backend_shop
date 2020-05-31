<template>
    <div class="v-registration">
        <v-links/>
        <h2>Регистрация</h2>
        <div class="v-registration__form">
            <form class="forms">
                <input id="username" placeholder="Логин" required
                       type="text" v-model="body.username">
                <input id="password" placeholder="Пароль" required
                       type="password" v-model="body.password">
                <div>
                    <input @click="Registration" type="submit" value="Зарегистрироваться">
                    Уже зарегистрированы? <router-link :to="{name: 'sign-in'}">Авторизуйтесь</router-link>
                </div>
            </form>
        </div>
        <v-auth-info/>
    </div>
</template>

<script>
    import { mapActions, mapGetters } from 'vuex'
    import router from '@/router/router'
    import VLinks from "../utils/v-links";
    import VAuthInfo from "../utils/v-auth-info";

    export default {
        name: "v-registration",
        components: {VAuthInfo, VLinks},
        props: {},
        data() {
            return {
                body: {
                    "username": "",
                    "password": ""
                }
            }
        },
        computed: {
            ...mapGetters([
                'REGISTRATION_STATUS',
            ]),
        },
        mounted() {
            this.AUTH_TEXT_STATUS = ""
        },
        methods: {
            ...mapActions([
                'REGISTRATION',
                'AUTH_UPDATE_STATUS'
            ]),
            Registration() {
                this.REGISTRATION(this.body).then(() => {
                    if (this.REGISTRATION_STATUS === 200) {
                        router.push('/');
                    } else {
                        this.AUTH_UPDATE_STATUS("Произошла ошибка регистрации")
                    }
                })
            }
        }
    }
</script>

<style lang="scss">

</style>