import {computed, defineComponent, onMounted, PropType, ref} from 'vue';
import s from './Charts.module.scss';
import {FormItem} from '../../shared/Form';
import {PieChart} from './PieChart';
import {LineChart} from './LineChart';
import {Bars} from './Bars';
import {http} from '../../shared/Http';


type Data1Item = { happen_at: string, amount: number }
type Data1 = Data1Item[]
export const Charts = defineComponent({
  props: {
    startDate: {
      type: String as PropType<string>,
      required: false
    },
    endDate: {
      type: String as PropType<string>,
      required: false
    }
  },
  setup: (props, context) => {
    const kind = ref('expenses');
    const data1 = ref<Data1>([]);
    const betterData1 = computed(() => {
      if (!props.startDate || !props.endDate) {
        return [];
      }
      const array = [];
      const n = new Date(props.endDate).getDate() - new Date(props.startDate).getDate() + 1;
      let j = 0;
      for (let i = 0; i < n; i++) {
        const date = new Date(props.startDate + 'T16:00:00.000Z');
        date.setDate(date.getDate() + i);
        const dateStr = date.toISOString();
        if (data1.value[j] && data1.value[j].happen_at === dateStr) {
          array.push([dateStr, data1.value[j].amount]);
          j += 1;
        } else {
          array.push([dateStr, 0]);
        }
      }
      return array as [string, number][];

    });

    onMounted(async () => {
      const response = await http.get<{ groups: Data1, summary: number }>('/items/summary', {
        happen_after: props.startDate,
        happen_before: props.endDate,
        kind: kind.value,
        _mock: 'itemSummary'
      });
      data1.value = response.data.groups;
    });

    return () => (
      <div class={s.wrapper}>
        <FormItem label="类型" type="select" options={[
          {value: 'expenses', text: '支出'},
          {value: 'income', text: '收入'}
        ]} v-model={kind.value}/>
        <LineChart data={betterData1.value}/>
        <PieChart/>
        <Bars/>
      </div>
    );
  }
});
