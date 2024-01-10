import { type FC } from "react";
import useUser from "~/hooks/useUser";

import TodosList from "~/components/Todos/TodosList";

const Todos: FC = () => {
  const { user } = useUser();

  if (user) return <TodosList />;
  else
    return (
      <div className="w-full text-center text-xl text-red-500">
        You need to login to see your todos
      </div>
    );
};

export default Todos;
