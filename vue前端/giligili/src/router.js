import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/Home.vue'
import Register from "@/views/Register";
import Login from "@/views/Login";
import Order from "@/views/Order";
import Search from "@/views/Search";

Vue.use(Router);

export default new Router({
    routes:[{
            path:'/',
            name:'home',
            component: Home
        },
        {
            path:'/about',
            name:'about',
            component: () =>import('./views/About.vue')
        },
        {
            path:'/register',
            name:'register',
            component:Register
        },
        {
            path:'/index',
            name:'login',
            component:Login
        },
        {
            path:'/order',
            name:'order',
            component:Order
        },
        {
            path:'/search',
            name:'search',
            component:Search
        }]
    }
)
