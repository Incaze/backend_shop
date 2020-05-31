export default {
    ADD_TO_CART({commit}, product){
        commit('SET_CART', product)
    },
    DELETE_FROM_CART({commit}, index){
        commit('REMOVE_FROM_CART', index)
    },
    INCREMENT_CART_ITEM({commit}, index){
        commit('INCREMENT', index)
    },
    DECREMENT_CART_ITEM({commit}, index){
        commit('DECREMENT', index)
    },
    AUTH_UPDATE_STATUS({commit}, text){
        commit('AUTH_UPDATE', text)
    },
    SIGN_OUT({commit}){
        commit('REMOVE_JWT_TOKEN')
    }


}

