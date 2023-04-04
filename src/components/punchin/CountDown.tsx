import {computed, defineComponent, PropType, ref, watchEffect} from 'vue';
import s from './CountDown.module.scss';
import {DatetimePicker, Popup} from 'vant';
import {usePunchIn} from '../../stores/usePunchIn';

export const CountDown = defineComponent({
  props: {
    happenAt: {
      type: String as PropType<string>
    },
    end: { // 传入到期时间（时间戳）
      type: Number as PropType<number>,
    }
  },
  setup: (props, context) => {
    // computed当前时间
    const refDatePickerVisible = ref(false);
    const store = usePunchIn();
    const showDatePicker = () => refDatePickerVisible.value = true;
    const hideDatePicker = () => refDatePickerVisible.value = false;
    const setDate = (date: Date) => {
      store.gowork = date.getTime();
      localStorage.setItem('gowork', store.gowork.toString());
      setTimeout(() => {
        store.gohometime = props.end;
      }, 1000);
      hideDatePicker();
    };
    const time = ref<string | Date>();
    const Prodate = () => {
      watchEffect(() => {
        if (store.gowork) {
          const gowork = localStorage.getItem('gowork');
          let date = new Date(Number(gowork));
          let year = date.getFullYear();
          let month: number | string = date.getMonth() + 1;
          month = month < 10 ? ('0' + month) : month;
          let day: number | string = date.getDate();
          day = day < 10 ? ('0' + day) : day;
          let hours: number | string = date.getHours();
          hours = hours < 10 ? ('0' + hours) : hours;
          let minutes: number | string = date.getMinutes();
          minutes = minutes < 10 ? ('0' + minutes) : minutes;
          let seconds: number | string = date.getSeconds();
          seconds = seconds < 10 ? ('0' + seconds) : seconds;
          time.value = year + '-' + month + '-' + day + ' ' + hours + ':' + minutes + ':' + seconds;
        }

      });
    };
    store.countdown();


    return () => (
      <div class={s.wrapper}>
        <div>
          <span>
            <span onClick={showDatePicker}>签到<input onClick={Prodate} placeholder="查看签到时间" value={store.gowork ? time.value : null}/>
            </span>
            <Popup position="bottom" v-model:show={refDatePickerVisible.value}>
              <DatetimePicker modelValue={store.gowork ? new Date(store.gowork) : new Date()}
                              type="datetime" title="选择年月日"
                              onConfirm={setDate} onCancel={hideDatePicker}
                              min-date={new Date(2020, 0, 1)}
              />
            </Popup>
          </span>
        </div>
        <div>
          <span>距离下班还有： </span>
          {store.gowork ? <span>{store.hour}时{store.minute}分{store.second}秒</span> : <span>请选择时间</span>}
        </div>
      </div>
    );
  }
});

