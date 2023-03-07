import {defineComponent, onMounted, PropType} from 'vue';
import s from './Navbar.module.scss';
export const Navbar = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const {slots} = context

    const onClick = () => {
      const button_bg = getComputedStyle(document.documentElement).getPropertyValue(`--button-bg`);
      document.documentElement.style.setProperty(`--button-bg`, button_bg === 'red' ? '#5c33be' : 'red');
    }
    return () => (
      <div class={s.navbar}>
        <span class={s.icon_wrapper}>
          {slots.icon?.()}
        </span>
        <span class={s.title_wrapper}>
          {slots.default?.()}
        </span>
        <button onClick={onClick}>切换</button>
      </div>
    )
  }
})
