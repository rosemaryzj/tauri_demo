import { createApp } from "vue";
import "./styles.css";
import App from "./App.vue";
import router from "./router/router.js";
import "./style/global.scss";
import "../src/assets/iconsmind/iconsmind.css"
import pinia from "./store/store.js";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import PerfectScrollbar from 'vue3-perfect-scrollbar'
import 'vue3-perfect-scrollbar/dist/vue3-perfect-scrollbar.css'

createApp(App)
    .use(router)
    .use(pinia)
    .use(PerfectScrollbar)
    .use(ElementPlus, {
        locale: zhCn
    })
    .mount("#app");
