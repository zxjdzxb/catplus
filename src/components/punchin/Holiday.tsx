import {defineComponent, onMounted, PropType, reactive, ref} from 'vue';
import axios from 'axios';

export const Holiday = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const nearestHoliday = ref<JSONValue>(null);
    let typeDes = ref<Array<string>>([]);
    let aaa = ref<Array<number | string>>([]);
    let Holiday = reactive({
      typeDes: [],
      aaa: []
    });
    const getHolidayData = async () => {
      const res = await axios.get(`https://www.mxnzp.com/api/holiday/list/year/${new Date().getFullYear()}`, {
        params: {
          app_id: 'qoirmtnytsifwkqo',  // 请替换成自己申请的app_id
          app_secret: 'd0pJeTVWY2lOTTdJZEt1YmN4anRhUT09',  // 请替换成自己申请的app_secret
        }
      });
      const data = res.data.data;
      // 获取当前月份
      const month = new Date().getMonth() + 1;
      for (let i = new Date().getMonth(); i < data.length; i++) {
        nearestHoliday.value = data[i];
        for (let j = 0; j < nearestHoliday.value?.days.length; j++) {
          if (nearestHoliday.value.days[j].type === 2) {
            let today: Date = new Date();
            const holiday: Date = new Date((nearestHoliday.value?.days[j].date).replace(/-/g, '/'));
            //@ts-ignore
            let days: number = Math.floor((holiday - today) / (1000 * 60 * 60 * 24));
            if (days > 0) {
              typeDes.value.push(nearestHoliday.value?.days[j].typeDes);
              aaa.value.push(days);
              Object.assign(Holiday, {typeDes: typeDes.value, aaa: aaa.value});
            }
            break;
          }
        }
      }

    };

    onMounted(getHolidayData);

    return () => (
      <>
        <div>
          <p>距离</p>
          {
            Holiday.typeDes.map((item, index) => {
              return <li key={index}>{item}还有{Holiday.aaa[index]}天</li>;
            })
          }
        </div>
      </>
    );
  }
});
export default Holiday;
