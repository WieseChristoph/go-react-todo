import { useQuery } from "react-query";
import type Todo from "~/types/Todo";
import apiRequest from "~/utils/apiRequest";

export const GET_TODOS_QUERY_KEY = "getTodos";

async function getTodos() {
  return apiRequest<Todo[]>("/api/todos", "GET");
}

function useGetTodos() {
  return useQuery<Todo[], Error>(GET_TODOS_QUERY_KEY, getTodos);
}

export default useGetTodos;
