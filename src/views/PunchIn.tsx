import {defineComponent, PropType} from 'vue';
import s from './PunchIn.module.scss';
import {RouterView} from 'vue-router';
import {CountDown} from '../components/punchin/CountDown';

export const PunchIn = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    return () => <>
      <CountDown />
   </>;
  }
});

export default PunchIn;



