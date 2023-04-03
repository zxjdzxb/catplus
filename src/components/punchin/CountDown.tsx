import {computed, defineComponent, PropType, ref, watchEffect} from 'vue';
import s from './CountDown.module.scss';
import {DatetimePicker, Popup} from 'vant';
import {usePunchIn} from '../../stores/usePunchIn';

export const CountDown = defineComponent({
  props: {
    happenAt: {
      type: String as PropType<string>
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
      store.gohometime = store.gowork+ 8 * 60 * 60 * 1000
      hideDatePicker();
    };
    store.increment();
    return () => (
      <div class={s.wrapper}>
        <div>
          <span>
            <span onClick={showDatePicker}>签到时间<input value={store.gowork ? new Date(store.gowork) : null}/></span>
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

