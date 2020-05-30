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
    }
}