import { createRouter, createWebHashHistory} from "vue-router";
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
const routes = [
    {
        path: "",
        component: ()=>import("../layout/index.vue"),
        redirect: "/login",
        children: [
            {
                path: "/home",
                name: "home",
                component: ()=>import("../views/Home.vue")
            }
        ]
    },
    {
        path: "/login",
        name: "login",
        component: ()=>import("../views/SignIn.vue")
    },
    {
        path: "/sign_up",
        name: "sign_up",
        component: ()=>import("../views/SignUp.vue")
    }
]

const router = createRouter({
    history: createWebHashHistory(''),
    routes
})
// 页面渲染成功之后，展示进度条（实际效果：Mac的Chrome就是在页面顶部有条2px左右的进度条）
router.beforeEach(async (to, from) => {
    NProgress.start();
})

// 页面加载成功之后，关闭进度条
router.afterEach(to => {
    NProgress.done()
})

export default router;