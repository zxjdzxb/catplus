import {defineComponent, PropType} from 'vue';
import {MainLayout} from '../layouts/MainLayout';
import s from './PunchIn.module.scss';
import {BackIcon} from '../shared/BackIcon';
import Todo from '../shared/Todo';


export const TodoList = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {

    return () => <>
      <MainLayout class={s.layout}>{{
        title: () => '记事本',
        icon: () => <BackIcon/>,
        default: () => <>
          <Todo/>
        </>
      }}</MainLayout>
    </>;
  }
});
export default TodoList;
