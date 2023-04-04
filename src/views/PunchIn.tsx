import {defineComponent, PropType, ref, watchEffect} from 'vue';
import s from './PunchIn.module.scss';
import {RouterView} from 'vue-router';
import {CountDown} from '../components/punchin/CountDown';
import {usePunchIn} from '../stores/usePunchIn';


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
        end.value = store.gowork + 9 * 60 * 60 * 1000;
      }
    });
    return () => <>
      <CountDown end={end.value}/>
    </>;
  }
});

export default PunchIn;



