import {defineComponent, onMounted, PropType, ref} from 'vue';
import s from './Navbar.module.scss';
import {Icon} from './Icon';

export const Navbar = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const {slots} = context;
    const ok = ref(false);
    onMounted(() => {
      //根据时间判断是否为夜间模式
      const hour = new Date().getHours();
      if (hour >= 20 || hour <= 6) {
        window.document.documentElement.setAttribute('data-theme', 'dark');
        ok.value = true;
      }else {
        window.document.documentElement.setAttribute('data-theme', 'light');
        ok.value = false;
      }
    });
    const themeSwitch = () => {
      const currentTheme = window.document.documentElement.getAttribute('data-theme');
      ok.value = !ok.value;
      if (currentTheme === 'dark') {
        window.document.documentElement.setAttribute('data-theme', 'light');
      } else {
        window.document.documentElement.setAttribute('data-theme', 'dark');
      }

    };
    return () => (
      <div class={s.navbar}>
        <span class={s.icon_wrapper}>
          {slots.icon?.()}
        </span>
        <span class={s.title_wrapper}>
          {slots.default?.()}
        </span>
        <span onClick={themeSwitch} class={s.status}>
         {!ok.value ? <Icon name="daytime" class={s.icon_wrapper}/> : <Icon name="night" class={s.icon_wrapper}/>}
        </span>
      </div>
    );
  }
});
