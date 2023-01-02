import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import { svgstore } from './src/vite_plugins/svgstore';
import styleImport, { VantResolve } from 'vite-plugin-style-import';
import Components from 'unplugin-vue-components/vite';
import { VantResolver } from 'unplugin-vue-components/resolvers';
// https://vitejs.dev/config/
export default defineConfig({
  // base:'/catplus/dist/',
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
  server: {				// ← ← ← ← ← ←
    host: '0.0.0.0'	// ← 新增内容 ←
  }
})
