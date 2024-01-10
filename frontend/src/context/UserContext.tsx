import {
  type FC,
  createContext,
  type ReactNode,
  useState,
  type Dispatch,
  type SetStateAction,
  useEffect,
} from "react";
import useGetUser from "~/hooks/useGetUser";
import type User from "~/types/User";

const UserContext = createContext<{
  user?: User;
  setUser: Dispatch<SetStateAction<User | undefined>>;
}>({
  user: undefined,
  setUser: () => undefined,
});

export const UserContextProvider: FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [user, setUser] = useState<User | undefined>(undefined);
  const { data } = useGetUser();

  useEffect(() => {
    if (!data) return;

    setUser(data);
  }, [data]);

  return (
    <UserContext.Provider value={{ user: user, setUser: setUser }}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext;
