import {defineComponent, PropType, ref} from 'vue';
import s from './Todo.module.scss';

interface Todo {
  id: number;
  text: string;
  completed: boolean;
}

export const TodoList = defineComponent({
  props: {
    name: {
      type: String as PropType<string>
    }
  },
  setup: (props, context) => {
    const todos = ref<Todo[]>([
      {id: 1, text: '唱跳', completed: false},
      {id: 2, text: 'Rap', completed: false},
      {id: 3, text: '打篮球', completed: false},
    ]);
    const newTodo = ref('');
    let id = todos.value.length + 1;

    function addTodo() {
      // trim()去除首尾空格
      if (newTodo.value.trim()) {
        todos.value.push({
          id: id++,
          text: newTodo.value.trim(),
          completed: false,
        });
        newTodo.value = '';
      }
    }

    function removeTodo(todo: Todo) {
      todos.value = todos.value.filter(t => t !== todo);
    }

    function toggleCompleted(todo: Todo) {
      todo.completed = !todo.completed;
    }
    // 全部removeTodo
    function removeTodoAll() {
      todos.value = [];
    }

    return () => (<>
        <div class={s.wrapper}>
          <h1>Todo List</h1>
          <ul>
            {todos.value.map(todo => (
              <li key={todo.id}>
                <input
                  type="checkbox"
                  checked={todo.completed}
                  onChange={() => toggleCompleted(todo)}
                />
                <span style={{textDecoration: todo.completed ? 'line-through' : 'none'}}>
                {todo.text}
              </span>
                <button onClick={() => removeTodo(todo)}>x</button>
              </li>
            ))}
          </ul>
          <div>
            <input type="text" v-model={newTodo.value}/>
            <div class={s.buttons}>
              <button onClick={addTodo}>Add Todo</button>
              <button onClick={() => removeTodoAll()}>Remove All</button>
            </div>
          </div>
        </div>
      </>
    );
  }
});
export default TodoList;
