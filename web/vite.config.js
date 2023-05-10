import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import Components from 'unplugin-vue-components/vite'
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'




// https://vitejs.dev/config/
export default ({ mode }) => defineConfig({
  base: loadEnv(mode, process.cwd()).VITE_PUBLIC_PATH,
  plugins: [vue(),
    Components({
      dts: true,
      include: [/[\\/].vite[\\/]/],
      resolvers: [
        AntDesignVueResolver(),
      ],
    })],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src')
    }
  },
  build: {
    terserOptions: {
      compress: {
        // Used to delete console in production environment
        drop_console: true,
      },
    }
  },
  server: {
    open: true,
    port: 8080,
    proxy: {

    }
  }
})
