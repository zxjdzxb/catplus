import {Dialog, Toast} from 'vant';
import {defineComponent, onMounted, PropType, ref} from 'vue';
import {RouterLink, useRoute} from 'vue-router';
import {useMeStore} from '../stores/useMeStore';
import {Icon} from './Icon';
import s from './Overlay.module.scss';

export const Overlay = defineComponent({
  props: {
    onClose: {
      type: Function as PropType<() => void>,
    },
  },
  setup: (props) => {
    const meStore = useMeStore();
    const close = () => {
      props.onClose?.();
    };
    const route = useRoute();
    const me = ref<User>();
    onMounted(async () => {
      const response = await meStore.mePromise;
      me.value = response?.data.resource;
    });
    const onSignOut = async () => {
      await Dialog.confirm({
        title: '确认',
        message: '你真的要退出登录吗？',
      });
      localStorage.removeItem('jwt');
      window.location.reload();
    };
    const linkArray = ref([
      {to: '/items', iconName: 'mangosteen', text: '欢迎使用'},
      {to: '/items/create', iconName: 'pig', text: '记一笔账'},
      {to: '/statistics', iconName: 'charts', text: '统计图表'},
      {to: '/punchin', iconName: 'punchin', text: '打卡'},
      {to: '/todolist', iconName: 'todolist', text: '记事本'},
      {to: '/export', iconName: 'export', text: '导出数据'},
      {to: '/notify', iconName: 'notify', text: '记账提醒'},

    ]);

    const refSelected = ref<number>(0);
    onMounted(() => {
      const path = route.path;
      switch (path) {
        case '/items':
          refSelected.value = 0;
          break;
        case '/items/create':
          refSelected.value = 1;
          break;
        case '/statistics':
          refSelected.value = 2;
          break;
        case '/export':
          refSelected.value = 3;
          break;
        case '/notify':
          refSelected.value = 4;
          break;
      }
    });
    const onIndexRepeat = (index: number, text: string) => {
      if (index === refSelected.value) {
        Toast({
          message: `当前已是${text}`,
          icon: 'success',
        });
      }
    };
    return () => (
      <>
        <div class={s.mask} onClick={close}></div>
        <div class={s.overlay}>
          <section class={s.currentUser}>
            {me.value ? (
              <div>
                <h2 class={s.email}>{me.value.email}</h2>
                <p onClick={onSignOut}>点击这里退出登录</p>
              </div>
            ) : (
              <RouterLink to={`/sign_in?return_to=${route.fullPath}`}>
                <h2>未登录用户</h2>
                <p>点击这里登录</p>
              </RouterLink>
            )}
          </section>
          <nav>
            <ul class={s.action_list}>
              {linkArray.value.map((item, i) => (
                <li onClick={() => onIndexRepeat(i, item.text)}>
                  <RouterLink to={item.to} class={s.action}>
                    <Icon name={item.iconName} class={s.icon}/>
                    <span>{item.text}</span>
                  </RouterLink>
                </li>
              ))}
            </ul>
          </nav>
        </div>
      </>
    );
  },
});

export const OverlayIcon = defineComponent({
  setup: () => {
    const refOverlayVisible = ref(false);
    const onClickMenu = () => {
      refOverlayVisible.value = !refOverlayVisible.value;
    };
    return () => (
      <>
        <Icon name="menu" class={s.icon} onClick={onClickMenu}/>
        {refOverlayVisible.value && (
          <Overlay onClose={() => (refOverlayVisible.value = false)}/>
        )}
      </>
    );
  },
});
