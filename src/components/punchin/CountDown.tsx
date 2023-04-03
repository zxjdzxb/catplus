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

      hideDatePicker();
    };
    watchEffect(() => {
      //计算时间差 转换成时分秒
      if (store.gowork) {
        store.gohometime = new Date(store.gowork).setHours(18, 0, 0, 0);
        store.time = store.gohometime - store.gowork;
        store.hour = Math.floor(store.time / 1000 / 60 / 60 % 24);
        store.minute = Math.floor(store.time / 1000 / 60 % 60);
        store.second = Math.floor(store.time / 1000 % 60);
      }
    });
    store.increment();


    return () => (
      <div class={s.wrapper}>
        <div>
          <span>
            <span onClick={showDatePicker}>当前时间<input value={store.gowork ? new Date(store.gowork) : new Date()}/></span>
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

