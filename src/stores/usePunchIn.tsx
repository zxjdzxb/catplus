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
