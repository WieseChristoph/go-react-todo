import useUser from "~/hooks/useUser";

import DiscordLoginButton from "./DiscordLoginButton";
import LogoutButton from "./LogoutButton";
import Avatar from "../Utils/Avatar";
import { Link } from "@tanstack/react-router";

const Navbar = () => {
  const { user, logout } = useUser();

  return (
    <nav className="flex w-full flex-row items-center gap-5 border-b border-gray-600 p-3">
      <Link to="/" className="text-3xl font-bold text-white">
        Todos
      </Link>
      <ul className="mr-auto flex flex-row gap-3">
        <Link to="/" className="text-gray-300 underline hover:text-white">
          Home
        </Link>
        {user && (
          <>
            <Link
              to="/todos"
              className="text-gray-300 underline hover:text-white"
            >
              Todos
            </Link>
            <Link
              to="/profile"
              className="text-gray-300 underline hover:text-white"
            >
              Profile
            </Link>
          </>
        )}
      </ul>
      {user ? (
        <>
          <Link to="/profile">
            <Avatar
              className="rounded-full"
              userId={user.id}
              avatar={user.avatar}
              size={48}
            />
          </Link>
          <LogoutButton onClick={() => logout()} />
        </>
      ) : (
        <DiscordLoginButton href="/api/auth/discord/login" />
      )}
    </nav>
  );
};

export default Navbar;
