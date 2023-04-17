import type { RouteRecordRaw } from 'vue-router'
import Login from '../views/login.vue'
import test from '../views/test.vue'
import Couldflare from '../views/cloudflare.vue'
import Home from "../views/home.vue"
import User from "../views/user.vue"
import Airborne from "../views/airborne.vue"
import Im from "../views/im.vue"
import Boce from "../components/dnstools/boce.vue"
import Itdog from "../components/dnstools/itdog.vue"
import Asiacloud from "../views/asiacloud.vue"
import Aliyun from "../views/aliyun.vue"
import Jenkins from "../components/dnstools/jenkins.vue"
// 使用路由项目类型声明一个路由数组
const routes: Array<RouteRecordRaw> = [
    {
        path: '/' ,
        name: 'login',
        component: Login
      },
    //   {
    //     path: '/cloudflare',
    //     name: 'Couldflare',
    //     component: Couldflare
    //   },
      {
        path: '/test',
        name: 'test',
        component: test
      },
      {
        path: '/home',
        name: 'home',
        component: Home,
        children: [
            {path: '/cloudflare', component: Couldflare},
            {path: '/user', component: User},
            {path: '/airborne' , component: Airborne},
            {path: '/im' , component: Im},
            {path: '/boce',component: Boce},
            {path: '/itdog',component: Itdog},
            {path: '/asiacloud',component: Asiacloud},
            {path: '/aliyun',component: Aliyun},
            {path: '/jenkins',component: Jenkins}
        ]
      },
]

export default routes