import {createRouter, createWebHistory} from 'vue-router'
import auth from '@/auth'
import Home from "@/components/Home";
import Login from "@/components/Login";
import Signup from "@/components/Signup";
import SecretQuote from "@/components/SecretQuote";
import UserInfo from "@/components/UserInfo";

function requireAuth (to, from, next) {
    if (!auth.isAuthenticated()) {
        this.$router.replace('/login')
    } else {
        next()
    }
}

const constantRoutes = [
    {
        path: '/',
        component: Home
    },
    {
        path: '/home',
        name: 'home',
        component: Home
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/signup',
        name: 'signup',
        component: Signup
    },
    {
        path: '/secretquote',
        name: 'secretquote',
        component: SecretQuote,
        beforeEnter: requireAuth
    },
    {
        path: '/userinfo',
        name: 'userinfo',
        component: UserInfo,
        beforeEnter: requireAuth
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes:constantRoutes,
})
export default router