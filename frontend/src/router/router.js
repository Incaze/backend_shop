import Vue from 'vue'
import Router from 'vue-router'

import vCatalog from '../components/catalog/v-catalog'
import vCart from '../components/cart/v-cart'
import vRegistration from '../components/auth/v-registration'
import vSignIn from '../components/auth/v-sign-in'

Vue.use(Router)

let router = new Router({
   routes: [
       {
           path: '/',
           name: 'catalog',
           component: vCatalog
       },
       {
           path: '/cart',
           name: 'cart',
           component: vCart,
           props: true
       },
       {
           path: '/registration',
           name: 'registration',
           component: vRegistration,
       },
       {
           path: '/sign_in',
           name: 'sign-in',
           component: vSignIn,
       },
   ]
})

export default router