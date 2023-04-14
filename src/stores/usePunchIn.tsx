//import store from "@/store"
import {defineStore} from 'pinia';
import {computed, ref} from 'vue';

export const usePunchIn = defineStore('counter', () => {
  const gowork = ref<number>();
  const gohometime = ref<number>();
  const hour = ref<number>();
  const minute = ref<number>();
  const second = ref<number>();
  let timer: any = null;
  //function countdown() {
  //  clearInterval(timer);
  //    timer = setInterval(() => {
  //    if (gowork.value && gohometime.value) {
  //      time.value = gohometime.value - gowork.value;
  //      hour.value = Math.floor(time.value / 1000 / 60 / 60 % 24);
  //      minute.value = Math.floor(time.value / 1000 / 60 % 60);
  //      second.value = Math.floor(time.value / 1000 % 60);
  //      //倒计时
  //      gowork.value += 1000;
  //    }
  //  }, 1000);
  //}
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
    if (!gowork.value) { return;}
    hour.value = Math.floor(time.value / 1000 / 60 / 60 % 24);
    minute.value = Math.floor(time.value / 1000 / 60 % 60);
    second.value = Math.floor(time.value / 1000 % 60);
    gowork.value += 1000;
    timer = setTimeout(() => {
      countdown();

    }, 1000);
  }

  return {gowork, gohometime, time, hour, minute, second, countdown};
});
