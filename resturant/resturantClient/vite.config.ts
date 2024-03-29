import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import * as path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 4000,
    open: "http://localhost:4000"
  },
  resolve: {
    alias: {
      "@src": path.resolve(__dirname, "./src"),
    }
  }
})
