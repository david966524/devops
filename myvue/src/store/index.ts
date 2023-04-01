import axios from 'axios'
import { defineStore } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import {UserData, UserInfo} from '../interface/index'
import {loginReq} from "../api/login"

// 创建 store
const useUserStore = defineStore('user', {
    
    state: () => ({
        username: '',
        token: ''
    }),

    // 定义 getters，等同于组件的计算属性
    getters: {
        // getter 函数接收 state 作为参数，推荐使用箭头函数
        hello: state => 'Hello!' + state.username
    },

    // 定义 actions，有同步和异步两种类型
    actions: {
        // 异步 action，一般用来处理异步逻辑
        async login (userData: UserData) {
            const loginData = await loginReq(userData)
            const { data, code } = loginData
            if (code === 200) {
                // action 中修改状态
                      // 将 token 存储到本地存储中
                localStorage.setItem('token', <string>data?.token)
                localStorage.setItem('username', <string>data?.username)
                this.username = <string>data?.username
                this.token = <string>data?.token

            }
        },

        // 同步 action
        logout() {
            this.username = ''
            localStorage.removeItem('token')
            this.token = ''
        }
    }
})

export default useUserStore