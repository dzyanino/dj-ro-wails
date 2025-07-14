import { createRouter, createWebHashHistory, type Router } from 'vue-router';
import Index from '@/pages/index.vue';
import Data from '@/pages/Data.vue';
import Help from '@/pages/Help.vue';

const routes = [
    { path: '/', component: Index },
    { path: '/data', component: Data },
    { path: '/aide', component: Help }
]

const router: Router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router