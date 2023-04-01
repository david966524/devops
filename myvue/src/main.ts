import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
const app = createApp(App)

const pinia = createPinia().use(piniaPluginPersistedstate)

app.use(pinia)
app.use(ElementPlus)
app.use(router)
app.mount('#app')
