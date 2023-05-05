import {defineComponent, onMounted, PropType, reactive} from 'vue';
import s from './Position.module.scss';
import AMapLoader from '@amap/amap-jsapi-loader';

type Position = {
  path: string[]
  current_position: string[]
};
export const Position = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const state = reactive<Position>({
      path: [],
      current_position: [],
    });


    function initMap() {
      AMapLoader.load({
        key: '641de6dcc99dbf8afb4725649c15949f', // 申请好的Web端开发者Key，首次调用 load 时必填
        version: '2.0', // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
      })
        .then((AMap) => {
          const map = new AMap.Map('container', {
            //设置地图容器id
            viewMode: '2D', //是否为3D地图模式
            zoom: 16, //初始化地图级别
          });
          //添加插件
          AMap.plugin('AMap.Geolocation', function () {
            var geolocation = new AMap.Geolocation({
              enableHighAccuracy: true,//是否使用高精度定位，默认:true
              timeout: 20000,          //超过10秒后停止定位，默认：5s
              position: 'RB',    //定位按钮的停靠位置
              offset: [10, 20], //定位按钮与设置的停靠位置的偏移量，默认：[10, 20]
              zoomToAccuracy: true,   //定位成功后是否自动调整地图视野到定位点

            });
            map.addControl(geolocation);
            geolocation.getCurrentPosition(function (status: string, result: object) {
              if (status == 'complete') {
                onComplete(result);
              } else {
                onError(result);
              }
            });

            function onComplete(data: any) {
              // data是具体的定位信息
              console.log(data);
            }

            function onError(data: any) {
              // 定位出错
              console.error(data);
            }
          });
          // 单击


        })
        .catch((e) => {
          console.log(e);
        });
    }

    onMounted(() => {
      initMap();
    });
    return () => (
      <div class={s.wrapper}>
        <div id="container" class={s.container}></div>
      </div>
    );
  }
});
