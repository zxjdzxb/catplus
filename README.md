# Vue 3 + Typescript + Vite

# 本地运行
pnpm run dev
# build 
pnpm run build
# oss
bin/ossutil cp -r dist oss://catplus
# dev 接口测试备注vite.config.ts

define: command === 'build' ? {
     DEBUG: false
 } : {
     DEBUG: true //如果想测试接口改成false
},
