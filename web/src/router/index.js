import { createRouter, createWebHistory } from 'vue-router';
import Info from '../components/Info.vue';
import Recommend from '../components/Recommend.vue';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import UserPage from '../components/UserPage.vue';
import AdminPage from '../components/AdminPage.vue';
import UserInfo from '../components/UserInfo.vue';

const routes = [
    { path: '/', component: Login }, // 登录页面
    { path: '/register', component: Register }, // 注册页面
    { path: '/user', component: UserPage }, // 用户页面
    { path: '/admin', component: AdminPage }, // 管理员页面
    { path: '/info', component: Info }, // 资讯信息
    { path: '/recommend', component: Recommend }, // 智能推荐
    { path: '/user-info', component: UserInfo }, // 用户信息管理
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// 添加导航守卫
router.beforeEach((to, from, next) => {
    const isAuthenticated = localStorage.getItem('isAuthenticated'); // 检查用户是否已登录
    const userRole = localStorage.getItem('userRole'); // 获取用户角色

    if (to.path === '/admin' && userRole !== '0') {
        next('/'); // 如果不是管理员，重定向到登录页面
    } else if (to.path === '/user' && userRole !== '1') {
        next('/'); // 如果不是普通用户，重定向到登录页面
    } else {
        next(); // 允许访问
    }
});

export default router; 