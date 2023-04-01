import { createRouter, createWebHistory } from 'vue-router'
import routes from './routes'

const history = createWebHistory()   //history 相当于 访问路径没有 #
const router = createRouter({
  history,
  routes
})
export default router   //导出router