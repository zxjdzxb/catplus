import {defineComponent, onUnmounted, PropType, ref} from 'vue';
import s from './Clock.module.scss';

export const Clock = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const duration = ref<number>(15 * 1000)
    const elapsed = ref<number>(0)
    let lastTime = performance.now()
    let handle: number
    const update = () => {
      const time = performance.now()
      elapsed.value += Math.min(time - lastTime, duration.value - elapsed.value)
      lastTime = time
      handle = requestAnimationFrame(update)
    }
    update()
    onUnmounted(() => {
      cancelAnimationFrame(handle)
    })
    return () => (
      <div class={s.wrapper}>
      <div>
        <label
        >Elapsed Time: <progress value={elapsed.value / duration.value}></progress
        ></label>
      </div>
        <div>{ (elapsed.value / 1000).toFixed(1) }s</div>
        <div>
          Duration: <input type="range" v-model={duration.value} min="1" max={duration.value*5} />
          { (duration.value / 1000).toFixed(1) }s
        </div>
        <button onClick ={()=>elapsed.value = 0}>Reset</button>
      </div>
    );
  }
});
