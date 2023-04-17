import {defineComponent, PropType, ref, watchEffect} from 'vue';
import s from './PunchIn.module.scss';
import {CountDown} from '../components/punchin/CountDown';
import {usePunchIn} from '../stores/usePunchIn';
import {Weather} from '../components/punchin/Weather';
import {Position} from '../components/punchin/Position';
import {MainLayout} from '../layouts/MainLayout';
import {BackIcon} from '../shared/BackIcon';


export const PunchIn = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const store = usePunchIn();
    const end = ref<number | undefined>();
    watchEffect(() => {
      if (store.gowork) {
        //18:00 end.value
        end.value = new Date(new Date(store.gowork).toLocaleDateString()).getTime() + 18 * 60 * 60 * 1000;
      }
    });
    return () => <>
      <MainLayout class={s.layout}>{{
        title: () => '倒计时',
        icon: () => <BackIcon/>,
        default: () => <>
          <Weather/>
          <Position/>
          <CountDown end={end.value}/>
        </>
      }}</MainLayout>
    </>;
  }
});

export default PunchIn;



