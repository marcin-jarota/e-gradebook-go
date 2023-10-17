import { fileURLToPath, URL } from 'node:url'
import path from 'path'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// import htmlPurge from 'vite-plugin-html-purgecss'
import htmlPurge from 'vite-plugin-purgecss'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [htmlPurge(), vue()],
  css: {
    preprocessorOptions: {
      scss: {
        // additionalData: '@import "bootstrap/scss/bootstrap";'
      }
    }
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '~bootstrap': path.resolve(__dirname, 'node_modules/bootstrap')
    }
  }
})
