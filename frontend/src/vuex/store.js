import Vue from 'vue'
import Vuex from 'vuex'
import commonActions from "./actions/actions";
import apiRequests from "./actions/api-requests";
import mutations from "./mutations/mutations";
import getters from "./getters/getters";

const actions = {...commonActions, ...apiRequests}

Vue.use(Vuex);

let store = new Vuex.Store( {
    state: {
        products: [],
        cart: [],
        registration_status: "",
        auth_text_status: "",
        jwt: {
            username: "",
            token: "",
        },
    },
    mutations,
    actions,
    getters,
})

export default store;