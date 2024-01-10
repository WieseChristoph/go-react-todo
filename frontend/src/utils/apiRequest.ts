import type Response from "~/types/Response";

async function apiRequest<T>(
  path: string,
  method: "GET" | "POST" | "PUT" | "DELETE",
  body?: unknown,
) {
  const res = await fetch(path, {
    method,
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify(body),
  });

  const data = (await res.json()) as Response<T>;

  if (data.success) {
    return data.data;
  } else {
    throw new Error(data.message);
  }
}

export default apiRequest;
