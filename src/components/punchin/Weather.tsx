import {defineComponent, onMounted, PropType, reactive, ref} from 'vue';
import s from './Weather.module.scss';
import axios from 'axios';
import {Icon} from '../../shared/Icon';

type Weather = {
  adcode: string
  city: string
  humidity: string
  humidity_float: string
  province: string
  reporttime: string
  temperature: string
  temperature_float: string
  weather: string
  winddirection: string
  windpower: string
}
export const Weather = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    let weather = reactive<Weather>({
      adcode: '',
      city: '',
      humidity: '',
      humidity_float: '',
      province: '',
      reporttime: '',
      temperature: '',
      temperature_float: '',
      weather: '',
      winddirection: '',
      windpower: ''
    });
    const key = ref('91fe89880ecb5819480f9e5f16d09e4a');
    const iconCondition = ref<string>('');
    const getLocationInfo = async () => {
      const params = {
        key: key.value,
      };
      const {data} = await axios.get('https://restapi.amap.com/v3/ip', {params});
      // data.adcode值为获取天气需要的city值
      getWeather(data.adcode);
    };
    const getWeather = async (adcode: string) => {
      const params = {
        key: key.value,
        city: adcode
      };
      const {data} = await axios.get(`https://restapi.amap.com/v3/weather/weatherInfo`, {params});
      Object.assign(weather, data.lives[0]);
      iconCondition.value = setWeatherIcon(weather.weather);
    };
    onMounted(() => {
      getLocationInfo();
    });

    function setWeatherIcon(weather: string) {
      // 只处理了基础的天气，可以继续精细化处理
      if (weather === '晴') {
        return 'clear-day';
      } else if (weather.includes('云')) {
        return 'partly-cloudy-day';
      } else if (weather.includes('风')) {
        return 'wind';
      } else if (weather.includes('雨')) {
        return 'rain';
      } else if (weather.includes('雪')) {
        return 'snow';
      } else if (weather.includes('雾')) {
        return 'fog';
      } else if (weather.includes('霾')) {
        return 'haze';
      }
      return 'cloudy';
    }

    return () => (
      <div class={s.wrapper}>
        <div class={s.flexcenter}>
          <Icon class={s.icons} name={iconCondition.value}/>
          <div class={s.font22}>
            <span>{weather.city}</span>
            <span>{weather.temperature}℃ </span>
            <span>{weather.weather}</span>
          </div>
        </div>
      </div>
    );
  }
});
