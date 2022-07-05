import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

const srcPath = path.resolve(__dirname, 'src')

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: [
      { find: '@', replacement: srcPath },
    ]
  },
  server: {
    proxy: {
      '^/user/.+': 'http://localhost:8888',
      '^/role/.+': 'http://localhost:8888',
      '^/object/.+': 'http://localhost:8888'
    }
  }
})
