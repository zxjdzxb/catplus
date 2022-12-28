import "./App.scss"
import { defineComponent, Transition, VNode } from "vue";
import { RouteLocationNormalizedLoaded, RouterView } from "vue-router";

export const App = defineComponent({
  setup() {
    return () => (
      <div class="page">
        <RouterView />
      </div>
    )
  }
})
