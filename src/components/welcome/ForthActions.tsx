import s from './welcome.module.scss';
import { RouterLink } from 'vue-router';
const onClick = () => {
  localStorage.setItem('skipFeatures', 'yes')
}
export const ForthActions = () => (
  <div class={s.actions}>
    <RouterLink class={s.fake} to="/start" >跳过</RouterLink>
    <span onClick={onClick}>
      <RouterLink to="/items">完成</RouterLink>
    </span>
    <RouterLink class={s.fake} to="/start" >跳过</RouterLink>
  </div>
)

ForthActions.displayName = 'ForthActions'
