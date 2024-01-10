import { useMutation, useQueryClient } from "react-query";
import type Todo from "~/types/Todo";
import apiRequest from "~/utils/apiRequest";
import { GET_TODOS_QUERY_KEY } from "~/hooks/useGetTodos";

export const CREATE_TODO_QUERY_KEY = "createTodos";

async function createTodo(todo: Todo) {
  return apiRequest<Todo>("/api/todos", "POST", todo);
}

function useCreateTodo() {
  const queryClient = useQueryClient();

  return useMutation<Todo, Error, Todo>(CREATE_TODO_QUERY_KEY, createTodo, {
    onSuccess: () => {
      void queryClient.invalidateQueries(GET_TODOS_QUERY_KEY);
    },
  });
}

export default useCreateTodo;
