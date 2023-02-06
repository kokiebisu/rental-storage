import { useEffect } from "react";
import { useUser } from "./useUser";
import { useLocalStorage } from "./useLocalStorage";
import axios from "axios";

export interface SignUpParams {
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
}

export interface LoginParams {
  emailAddress: string;
  password: string;
}

export const useAuth = () => {
  const { user, addUser, removeUser } = useUser();
  const { getItem, removeItem, setItem } = useLocalStorage();

  useEffect(() => {
    const user = getItem("user");
    if (user) {
      addUser(JSON.parse(user));
    }
  }, []);

  const signup = async (data: SignUpParams) => {
    if (
      !data.emailAddress ||
      !data.password ||
      !data.firstName ||
      !data.lastName
    ) {
      throw new Error("input misssing");
    }

    if (!process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT) {
      throw new Error("endpoint misssing");
    }

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT}/auth/signup`,
      {
        emailAddress: data.emailAddress,
        firstName: data.firstName,
        lastName: data.lastName,
        password: data.password,
      }
    );
    if (response.status !== 200) {
      throw new Error("internal server error");
    }
    const { authenticationToken } = response.data;

    // in a production app, we need to send some data (usually username, password) to server and get a token
    // we will also need to handle errors if sign in faield

    setItem("authenticationToken", authenticationToken);
  };

  const isAuthenticated = () => {
    const authenticationToken = getItem("authenticationToken");
    return !!authenticationToken;
  };

  const login = async (data: LoginParams) => {
    if (!data.emailAddress || !data.password) {
      throw new Error("input missing");
    }

    if (!process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT) {
      throw new Error("endpoint missing");
    }

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT}/auth/signin`,
      {
        emailAddress: data.emailAddress,
        password: data.password,
      }
    );
    if (response.status !== 200) {
      throw new Error("something went wrong");
    }
    const { authorizationToken } = response.data;

    // in a production app, we need to send some data (usually username, password) to server and get a token
    // we will also need to handle errors if sign in faield

    setItem("authenticationToken", authorizationToken);
  };

  const logout = () => {
    removeItem("authenticationToken");
    removeUser();
  };

  return { user, signup, login, logout, isAuthenticated };
};