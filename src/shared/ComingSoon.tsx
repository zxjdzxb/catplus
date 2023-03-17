import { defineComponent, PropType } from 'vue'
import { Center } from './Center'
import s from './ComingSoon.module.scss'
import { Icon } from './Icon'
import {useRouter} from 'vue-router';
import {Button} from './Button';
import {BackIcon} from './BackIcon';
import {Tab, Tabs} from './Tabs';
import {Tags} from '../components/item/Tags';
import {InputPad} from '../components/item/InputPad';
import {MainLayout} from '../layouts/MainLayout';
export const ComingSoon = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const router = useRouter()
    const onClick = () => {
      router.back()
    }
    return () => (
      <MainLayout class={s.layout}>{{
        title: () => '敬请期待',
        icon: () => <BackIcon/>,
        default: () => <>
          <div>
            <Center class={s.pig_wrapper}>
              <Icon name="forward" class={s.pig} />
            </Center>
            <p class={s.text}>开发ing...</p>
          </div>
        </>
      }}</MainLayout>

    )
  }
})
