import {defineComponent, onUnmounted, PropType, ref, watchEffect} from 'vue';
import s from './Clock.module.scss';
import {usePunchIn} from '../../stores/usePunchIn';
import {Button} from '../../shared/Button';

export const Clock = defineComponent({
  props: {
    duration: {
      type: Number as PropType<number>
    }
  },
  setup: (props, context) => {
    const store = usePunchIn();
    const duration = ref<number>(15 * 1000);
    watchEffect(() => {
      duration.value = props.duration ? props.duration : duration.value;
    });

    const elapsed = ref<number>(0);
    let lastTime = performance.now();
    let handle: number;
    const update = () => {
      const time = performance.now();
      elapsed.value += Math.min(time - lastTime, duration.value - elapsed.value);
      lastTime = time;
      handle = requestAnimationFrame(update);
    };
    update();
    onUnmounted(() => {
      cancelAnimationFrame(handle);
    });
    return () => (
      <div class={s.wrapper}>
          <label
          >进度条: <progress value={elapsed.value / duration.value}></progress
          ></label>
          <div>{(elapsed.value / 1000).toFixed(1)}s</div>
        {/*<div>*/}
        {/*  距离时间: <input type="range" v-model={duration.value} min="1" max={30000}/>*/}
        {/*  {(duration.value / 1000).toFixed(1)}s*/}
        {/*</div>*/}
      </div>
    );
  }
});
