import { useMutation, useQueryClient } from "react-query";
import type Todo from "~/types/Todo";
import apiRequest from "~/utils/apiRequest";
import { GET_TODOS_QUERY_KEY } from "~/hooks/useGetTodos";

export const UPDATE_TODO_QUERY_KEY = "updateTodos";

async function updateTodo(todo: Todo) {
  return apiRequest<Todo>(`/api/todos/${todo.id}`, "POST", todo);
}

function useUpdateTodo() {
  const queryClient = useQueryClient();

  return useMutation<Todo, Error, Todo>(UPDATE_TODO_QUERY_KEY, updateTodo, {
    onSuccess: () => {
      void queryClient.invalidateQueries(GET_TODOS_QUERY_KEY);
    },
  });
}

export default useUpdateTodo;
