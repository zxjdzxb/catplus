import { defineConfig, splitVendorChunkPlugin } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import styleImport, { VantResolve } from 'vite-plugin-style-import';
// @ts-nocheck
import { svgstore } from './src/vite_plugins/svgstore';

// https://vitejs.dev/config/
export default defineConfig(({command}) => {
  return {
    base: './',
    define: command === 'build' ? {
      DEBUG: false
    } : {
      DEBUG: true //如果想测试接口改成false
    },
    build: {
      rollupOptions: {
        output: {
          manualChunks(id: any) {
            if (id.includes('echarts')) {
              return 'echarts';
            }
            if (id.includes('mock') || id.includes('faker')) {
              return 'mock';
            }
            if (id.includes('vant')) {
              return 'vant';
            }
            if (id.includes('node_modules')) {
              return 'vendor';
            }
          }
        }
      }
    },
    plugins: [
      vue(),
      vueJsx({
        transformOn: true,
        mergeProps: true
      }),
      svgstore(),
      styleImport({
        resolves: [VantResolve()],
      }),
    ],
    server: {
      host: '0.0.0.0',// ← 新增内容 ←
      proxy: {
        '/api/v1': {
          target: 'http://121.196.236.94:3000/',
        }
      }
    }
  }
})
