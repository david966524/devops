import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target: 'http://192.168.96.128:8080',
        changeOrigin: true,  // 允许跨域
        rewrite: (path) => path.replace(/^\/api/, ''),
      }
    }
  }
})
