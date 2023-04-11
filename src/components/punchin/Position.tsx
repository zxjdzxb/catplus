import {defineComponent, onMounted, PropType, reactive, ref} from 'vue';
import s from './Position.module.scss';
import AMapLoader from '@amap/amap-jsapi-loader';
import axios from 'axios';

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
    const circle = ref<number[]>([]);
    const key = ref('91fe89880ecb5819480f9e5f16d09e4a');
    const getLocationInfo = async () => {
      const params = {
        key: key.value,
        address: '北京市朝阳区朝阳公园'
      };
      const {data} = await axios.get('https://restapi.amap.com/v3/geocode/geo?parameters', {params});
      circle.value.push(Number(data.geocodes[0].location.split(',')[0]));
      circle.value.push(Number(data.geocodes[0].location.split(',')[1]));
    };

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
            center: circle.value, //初始化地图中心点位置
          });
          //添加插件
          AMap.plugin(['AMap.ToolBar', 'AMap.Scale', 'AMap.HawkEye'], function () {
            //异步同时加载多个插件
            map.addControl(new AMap.HawkEye()); //显示缩略图
            map.addControl(new AMap.Scale()); //显示当前地图中心的比例尺
          });
          AMap.plugin('AMap.Geolocation', function () {
            var geolocation = new AMap.Geolocation({
              enableHighAccuracy: true,//是否使用高精度定位，默认:true
              timeout: 10000,          //超过10秒后停止定位，默认：5s
              buttonPosition: 'RB',    //定位按钮的停靠位置
              buttonOffset: new AMap.Pixel(10, 20),//定位按钮与设置的停靠位置的偏移量，默认：Pixel(10, 20)
              zoomToAccuracy: true,   //定位成功后是否自动调整地图视野到定位点

            });
            map.addControl(geolocation);

          });
          // 单击
          map.on('click', (e: any) => {
            // console.log(e);
            state.current_position = [e.lnglat.KL, e.lnglat.kT];
            //@ts-ignore
            state.path.push([e.lnglat.KL, e.lnglat.kT]);
            addMarker();
            addPolyLine();
            // 地图按照适合展示图层内数据的缩放等级展示
            // map.setFitView();
          });

          function addMarker() {
            const marker = new AMap.Marker({
              icon: '//a.amap.com/jsapi_demos/static/demo-center/icons/poi-marker-default.png',
              position: state.current_position,
              offset: new AMap.Pixel(-26, -54),
            });
            marker.setMap(map);
          }

          // 折线
          function addPolyLine() {
            const polyline = new AMap.Polyline({
              path: state.path,
              isOutline: true,
              outlineColor: '#ffeeff',
              borderWeight: 1,
              strokeColor: '#3366FF',
              strokeOpacity: 0.6,
              strokeWeight: 5,
              // 折线样式还支持 'dashed'
              strokeStyle: 'solid',
              // strokeStyle是dashed时有效
              // strokeDasharray: [10, 5],
              lineJoin: 'round',
              lineCap: 'round',
              zIndex: 50,
            });
            map.add([polyline]);
          }
        })
        .catch((e) => {
          console.log(e);
        });
    }

    onMounted(() => {
      getLocationInfo().then(r => {
        initMap();
      });
    });
    return () => (
      <div class={s.wrapper}>
        <div id="container" class={s.container}></div>
      </div>
    );
  }
});
