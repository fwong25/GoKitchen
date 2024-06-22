import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import serverConfig from './../../ipcfg.json'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    host: serverConfig.Configs[serverConfig.frontend_server_idx].Ip,
    port: serverConfig.Configs[serverConfig.frontend_server_idx].Port,
  },
  // server: {
  //   host: '192.168.18.27',
  //   port: 5173,
  // },
})

// module.exports = {
//   devServer: {
//     host: '192.168.18.27',
//     port: 5173, // your desired port number
//   },
// };
