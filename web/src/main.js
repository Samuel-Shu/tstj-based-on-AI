import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import 'bootstrap/dist/css/bootstrap.css';
import 'element-plus/theme-chalk/index.css';
import ElementPlus from 'element-plus';

const app = createApp(App);
app.use(router);
app.use(ElementPlus);
app.mount('#app'); 