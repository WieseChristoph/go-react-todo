import { type FC } from "react";
import useGetTodos from "~/hooks/useGetTodos";
import useDeleteTodo from "~/hooks/useDeleteTodo";
import useCreateTodo from "~/hooks/useCreateTodo";
import useUpdateTodo from "~/hooks/useUpdateTodo";

import TodoCard from "./TodoCard";
import CreateTodoForm from "./CreateTodoForm";

const TodosList: FC = () => {
  const { data: todos, isLoading, error } = useGetTodos();
  const { mutate: createTodo } = useCreateTodo();
  const { mutate: updateTodo } = useUpdateTodo();
  const { mutate: deleteTodo } = useDeleteTodo();

  if (isLoading) return <div>Loading...</div>;

  if (error) return <div>{error.message}</div>;

  return (
    <div className="flex flex-row flex-wrap gap-6">
      <CreateTodoForm onSubmit={(todo) => createTodo(todo)} />
      {todos &&
        !!todos.length &&
        todos.map((todo) => (
          <TodoCard
            key={todo.id}
            todo={todo}
            onClick={() => updateTodo({ ...todo, status: !todo.status })}
            onDelete={() => deleteTodo(todo)}
          />
        ))}
    </div>
  );
};

export default TodosList;
