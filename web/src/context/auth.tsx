import { createContext, useEffect, useState } from "react";
import axios from "axios";
import { useLazyQuery } from "@apollo/client";

import { awsLambdaClient } from "@/clients";
import { PROFILE_QUERY } from "@/graphql/queries";
import { User, SignUpParams, SignInParams } from "@/types/interface";

interface AuthContext {
  user?: User | null;
  signup: (data: SignUpParams) => Promise<void>;
  signin: (data: SignInParams) => Promise<void>;
  signout: () => Promise<void>;
}

export const AuthContext = createContext<AuthContext>({
  user: null,
  signup: async () => {},
  signin: async () => {},
  signout: async () => {},
});

export const AuthContextProvider = ({ children }: any) => {
  const [user, setUser] = useState<User | null>(null);
  const [fetch, { data, loading, error }] = useLazyQuery(PROFILE_QUERY, {
    client: awsLambdaClient,
    fetchPolicy: "no-cache",
  });

  useEffect(() => {
    const value = localStorage.getItem("user");
    if (value) {
      setUser(JSON.parse(value));
    }
  }, []);

  useEffect(() => {
    if (!loading && !error && data) {
      if (data.profile) {
        localStorage.setItem("user", JSON.stringify(data.profile));
      }
      setUser(data.profile);
    }
  }, [data, error, loading]);

  const signup = async (params: SignUpParams) => {
    if (
      !params.emailAddress ||
      !params.password ||
      !params.firstName ||
      !params.lastName
    ) {
      throw new Error("input misssing");
    }

    if (!process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT) {
      throw new Error("endpoint misssing");
    }

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT}/auth/signup`,
      {
        emailAddress: params.emailAddress,
        firstName: params.firstName,
        lastName: params.lastName,
        password: params.password,
      }
    );
    if (response.status !== 200) {
      throw new Error("internal server error");
    }
    const { authorizationToken } = response.data;

    localStorage.setItem("bearerToken", authorizationToken);
    fetch();
  };

  const signin = async (params: SignInParams) => {
    if (!params.emailAddress || !params.password) {
      throw new Error("input misssing");
    }

    if (!process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT) {
      throw new Error("endpoint misssing");
    }

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT}/auth/signin`,
      {
        emailAddress: params.emailAddress,
        password: params.password,
      }
    );
    if (response.status !== 200) {
      throw new Error("internal server error");
    }
    const { authorizationToken } = response.data;

    localStorage.setItem("bearerToken", authorizationToken);
    fetch();
  };

  const signout = async () => {
    localStorage.removeItem("bearerToken");
    localStorage.removeItem("user");
    fetch();
  };

  return (
    <AuthContext.Provider value={{ user, signup, signin, signout }}>
      {children}
    </AuthContext.Provider>
  );
};
