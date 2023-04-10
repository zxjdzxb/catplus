import {defineComponent, PropType, ref, watchEffect} from 'vue';
import s from './PunchIn.module.scss';
import {RouterView} from 'vue-router';
import {CountDown} from '../components/punchin/CountDown';
import {usePunchIn} from '../stores/usePunchIn';
import {Clock} from '../components/punchin/Clock';
import {Weather} from '../components/punchin/Weather';


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
      <Weather/>
      <CountDown end={end.value}/>
    </>;
  }
});

export default PunchIn;



