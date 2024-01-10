import { useMutation, useQueryClient } from "react-query";
import type Todo from "~/types/Todo";
import apiRequest from "~/utils/apiRequest";
import { GET_TODOS_QUERY_KEY } from "~/hooks/useGetTodos";

export const DELETE_TODO_QUERY_KEY = "deleteTodos";

async function deleteTodo(todo: Todo) {
  return apiRequest<Todo>(`/api/todos/${todo.id}`, "DELETE");
}

function useDeleteTodo() {
  const queryClient = useQueryClient();

  return useMutation<Todo, Error, Todo>(DELETE_TODO_QUERY_KEY, deleteTodo, {
    onSuccess: () => {
      void queryClient.invalidateQueries(GET_TODOS_QUERY_KEY);
    },
  });
}

export default useDeleteTodo;
