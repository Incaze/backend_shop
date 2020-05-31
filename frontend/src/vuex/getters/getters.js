export default {
    PRODUCTS(state){
        return state.products;
    },
    CART(state){
        return state.cart;
    },
    REGISTRATION_STATUS(state){
        return state.registration_status;
    },
    AUTH_TEXT_STATUS(state){
        return state.auth_text_status;
    },
    JWT_AUTH(state){
        return state.jwt;
    }
}