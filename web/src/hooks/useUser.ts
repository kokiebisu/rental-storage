import { useContext, useState } from "react";
import { AuthContext } from "../context/auth";
import { useLocalStorage } from "./useLocalStorage";

export interface User {
  id: string;
  name: string;
  email: string;
  authToken?: string;
}

export interface AddUserParam {}

export const useUser = () => {
  const [user, setUser] = useState<User | null>();
  const { setItem } = useLocalStorage();

  const addUser = (data: AddUserParam) => {
    setUser(user);
    setItem("user", JSON.stringify(user));
  };

  const removeUser = () => {
    setUser(null);
    setItem("user", "");
  };

  return { user, addUser, removeUser };
};
