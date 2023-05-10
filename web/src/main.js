import { createApp } from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue'
import { setupRouter } from './router'
import 'ant-design-vue/dist/antd.css';
import store from './store';

const app = createApp(App)

setupRouter(app)

app.use(Antd).use(store).mount('#app')
