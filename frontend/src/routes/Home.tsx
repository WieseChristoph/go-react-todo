import { type FC } from "react";
import useUser from "~/hooks/useUser";

const Home: FC = () => {
  const { user } = useUser();

  return (
    <div className="w-full text-center text-white">
      <h1 className="text-3xl">Welcome to the Todos App</h1>
      <h2>
        {user
          ? "Click 'Todos' in the navigation to manage your Todos"
          : "Log in to create Todos"}
      </h2>
    </div>
  );
};

export default Home;
