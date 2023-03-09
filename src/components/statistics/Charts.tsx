import {computed, defineComponent, onMounted, PropType, ref, watch} from 'vue';
import s from './Charts.module.scss';
import {FormItem} from '../../shared/Form';
import {PieChart} from './PieChart';
import {LineChart} from './LineChart';
import {Bars} from './Bars';
import {http} from '../../shared/Http';


type Data1Item = { happen_at: string, amount: number }
type Data1 = Data1Item[]
type Data2Item = { tag_id: number; tag: Tag; amount: number }
type Data2 = Data2Item[]
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
      console.log(array);
      return array as [string, number][];

    });

    const fetchData1 = async () => {
      const response = await http.get<{ groups: Data1, summary: number }>('/items/summary', {
        happen_after: props.startDate,
        happen_before: props.endDate,
        kind: kind.value,
        group_by: 'happen_at',
      }, {
        _mock: 'itemSummary',
        _autoLoading: true
      });
      data1.value = response.data.groups;

    };
    onMounted(fetchData1)
    watch(() => kind.value, fetchData1)
    const data2 = ref<Data2>([])
    const betterData2 = computed<{ name: string; value: number }[]>(() =>
      data2.value.map((item) => ({
        name: item.tag.name,
        value: item.amount
      }))
    )

    const fetchData2 = async ()=>{
      const response = await http.get<{ groups: Data2; summary: number }>('/items/summary', {
        happen_after: props.startDate,
        happen_before: props.endDate,
        kind: kind.value,
        group_by: 'tag_id',
      }, {
        _mock: 'itemSummary'
      })
      data2.value = response.data.groups
    }
    onMounted(fetchData2)
    watch(() => kind.value, fetchData2)
    const betterData3 = computed<{tag:Tag, amount:number, percent: number}[]>(()=>{
      const total = data2.value.reduce((sum, item) => sum + item.amount, 0)
      return data2.value.map(item => ({
        ...item,
        percent: Math.round(item.amount / total * 100)
      }))
    })
    return () => (
      <div class={s.wrapper}>
        <FormItem label="类型" type="select" options={[
          {value: 'expenses', text: '支出'},
          {value: 'income', text: '收入'}
        ]} v-model={kind.value}/>
        <LineChart data={betterData1.value}/>
        <PieChart data={betterData2.value} />
        <Bars data={betterData3.value}/>
      </div>
    );
  }
});
