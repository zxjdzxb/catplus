//import store from "@/store"
import {defineStore} from 'pinia';
import {computed, ref} from 'vue';

export const usePunchIn = defineStore('counter', () => {
  const gowork = ref<number>();
  const gohometime = ref<number>();
  const hour = ref<number>();
  const minute = ref<number>();
  const second = ref<number>();
  const countdowntime = ref<string>('');
  let timer: any = null;
  const time = computed({
    get: () => {
      if (gowork.value && gohometime.value) {
        return (gohometime.value - gowork.value);
      } else {
        return 0;
      }
    },
    set: (val) => {
      return val;
    }
  });

  function countdown() {
    clearTimeout(timer);
    hour.value = Math.floor(time.value / 1000 / 60 / 60 % 24);
    minute.value = Math.floor(time.value / 1000 / 60 % 60);
    second.value = Math.floor(time.value / 1000 % 60);

    countdowntime.value = `${hour.value}时${minute.value}分${second.value}秒`;
    //@ts-ignore
    gowork.value += 1000;
    timer = setTimeout(() => {
      countdown();
    }, 1000);
  }

  return {gowork, gohometime, time, countdowntime, countdown};
});
