import './App.scss';
import {defineComponent, onMounted, ref,} from 'vue';
import {RouterView} from 'vue-router';

export const App = defineComponent({
  setup() {const ok = ref(false);
    onMounted(() => {
      //根据时间判断是否为夜间模式
      const hour = new Date().getHours();
      if (hour >= 20 || hour <= 6) {
        window.document.documentElement.setAttribute('data-theme', 'dark');
      }else {
        window.document.documentElement.setAttribute('data-theme', 'light');
      }
    });
    return () => (
      <div class="page">
        <RouterView/>
      </div>
    );
  }
});
