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

  function increment() {
    setInterval(() => {
      if (gowork.value) {
        gowork.value = new Date().getTime();
      }
    }, 1000);
  }
  return { gowork, gohometime, time, hour, minute, second ,increment}
})
