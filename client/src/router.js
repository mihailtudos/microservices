// router.js
import { createRouter, createWebHistory } from 'vue-router'
import Home from "./pages/Home.vue";
import Product from "./pages/Product.vue";
import ManageProduct from "@/pages/ManageProduct.vue";

const routes = [
    { path: '/', component: Home },
    { path: '/products/:id', component: Product, props: true },
    { path: '/admin/products/:id', component: ManageProduct, props: true },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router