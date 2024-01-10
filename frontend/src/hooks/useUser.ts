import { useContext } from "react";
import UserContext from "~/context/UserContext";
import apiRequest from "~/utils/apiRequest";

function useUser() {
  const { user, setUser } = useContext(UserContext);

  const logout = () =>
    apiRequest<undefined>("/api/auth/logout", "POST")
      .then(() => setUser(undefined))
      .catch((e) => console.error(e));

  return { user, logout };
}

export default useUser;
