import axios from "axios";

const API_URL = 'http://localhost:80/'

export default {
    GET_PRODUCTS_FROM_API({commit}) {
        return axios(API_URL + 'products', {
            method: "GET"
        })
            .then((products) => {
                commit('SET_PRODUCTS_TO_STATE', products.data);
                return products;
            })
            .catch((error) => {
                return error;
            })
    },
    REGISTRATION({commit}, body) {
        return axios(API_URL +'register', {
            method: "POST",
            data: body
        })
            .then(registration => {
                commit('SET_REGISTRATION_STATUS_TO_STATE', registration.status);
                return registration;
            })
            .catch((error) => {
                return error;
            })
    },
    SIGN_IN({commit}, body) {
        return axios(API_URL +'get_token', {
            method: "POST",
            data: body
        })
            .then(authentication => {
                commit('SET_AUTH_DATA_TO_STATE', authentication.data);
                return authentication;
            })
            .catch((error) => {
                return error;
            })
    }
}