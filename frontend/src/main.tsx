import React from "react";
import ReactDOM from "react-dom/client";
import { QueryClient, QueryClientProvider } from "react-query";
import { UserContextProvider } from "./context/UserContext.tsx";
import { RouterProvider } from "@tanstack/react-router";
import router from "~/routes/router.ts";
import "./index.css";

const queryClient = new QueryClient({
  defaultOptions: { queries: { retry: false } },
});

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <UserContextProvider>
        <RouterProvider router={router} />
      </UserContextProvider>
    </QueryClientProvider>
  </React.StrictMode>,
);
