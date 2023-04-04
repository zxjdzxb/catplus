//import store from "@/store"
import {defineStore} from 'pinia';
import {computed, ref, watchEffect} from 'vue';

export const usePunchIn = defineStore('counter', () => {
  const gowork = ref<number>();
  const gohometime = ref<number>();
  const time = ref<number>();
  const hour = ref<number>();
  const minute = ref<number>();
  const second = ref<number>();
  function countdown() {
    setInterval(() => {
      if (gowork.value && gohometime.value) {
        time.value = gohometime.value - gowork.value;
        hour.value = Math.floor(time.value / 1000 / 60 / 60 % 24);
        minute.value = Math.floor(time.value / 1000 / 60 % 60);
        second.value = Math.floor(time.value / 1000 % 60);
        //倒计时
        gowork.value += 1000;
      }
    }, 1000);
  }
  return { gowork, gohometime, time, hour, minute, second , countdown}
})
