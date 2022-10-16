import * as React from "react";
import axios from "axios";
import { createContext, useEffect, useMemo, useReducer } from "react";
import * as SecureStore from "expo-secure-store";

export const AuthContext = createContext({});

export const AuthContextProvider = ({ children }) => {
  const [state, dispatch] = useReducer(
    (prevState, action) => {
      switch (action.type) {
        case "RESTORE_TOKEN":
          return {
            ...prevState,
            userToken: action.token,
            isLoading: false,
          };
        case "SIGN_IN":
          return {
            ...prevState,
            isSignout: false,
            userToken: action.token,
          };
        case "SIGN_OUT":
          return {
            ...prevState,
            isSignout: true,
            userToken: null,
          };
      }
    },
    {
      isLoading: true,
      isSignout: false,
      userToken: null,
    }
  );

  useEffect(() => {
    // fetch the token from storage then navigate to our appropriate place
    const bootstrapAsync = async () => {
      let userToken;

      try {
        userToken = await SecureStore.getItemAsync("userToken");
        console.log("USER TOKEN: ", userToken);
      } catch (err) {
        // Restoring token failed
      }

      // after restoring token, we may need to validate it in production apps

      // this will switch to the App screen or Auth screen and this loading
      // screen will be unmounted and thrown away
      dispatch({ type: "RESTORE_TOKEN", token: userToken });
    };

    bootstrapAsync();
  }, []);

  const authContext = useMemo(
    () => ({
      signIn: async (data) => {
        if (!data.emailAddress || !data.password) {
          alert("Input missing!");
          return;
        }
        const response = await axios.post(
          "https://17gm42qnnb.execute-api.us-east-1.amazonaws.com/dev/auth/signin",
          {
            emailAddress: data.emailAddress,
            password: data.password,
          }
        );

        console.log("RESPONSE: ", response);

        if (response.status !== 200) {
          alert("something went wrong");
          return;
        }

        const { authorizationToken } = response.data;

        console.log("AUTHORIZATION TOKEN: ", authorizationToken);
        // in a production app, we need to send some data (usually username, password) to server and get a token
        // we will also need to handle errors if sign in faield
        try {
          // after getting token, we need to persist the token using 'SecureStore'
          // in the example, we'll use a dummy token
          await SecureStore.setItemAsync("userToken", authorizationToken);
          dispatch({ type: "SIGN_IN", token: authorizationToken });
        } catch (err) {
          console.error(err);
        }
      },
      signOut: async () => {
        try {
          await SecureStore.deleteItemAsync("userToken");
          dispatch({ type: "SIGN_OUT" });
        } catch (err) {
          console.error(err);
        }
      },
      signUp: async (data) => {
        if (
          !data.emailAddress ||
          !data.password ||
          !data.firstName ||
          !data.lastName
        ) {
          alert("Input missing");
          return;
        }
        const response = await axios.post(
          "https://17gm42qnnb.execute-api.us-east-1.amazonaws.com/dev/auth/signup",
          {
            emailAddress: data.emailAddress,
            firstName: data.firstName,
            lastName: data.lastName,
            password: data.password,
          }
        );

        if (response.status !== 200) {
          alert("Something went wrong");
          return;
        }

        const { authorizationToken } = response.data;
        console.log("RESPONSE: ", response.data);
        // in a production app, we need to send some data (usually username, password) to server and get a token
        // we will also need to handle errors if sign in faield
        try {
          // after getting token, we need to persist the token using 'SecureStore'
          // in the example, we'll use a dummy token
          await SecureStore.setItemAsync("userToken", authorizationToken);
          dispatch({ type: "SIGN_IN", token: authorizationToken });
        } catch (err) {
          console.error(err);
        }
      },
    }),
    []
  );

  return (
    <AuthContext.Provider value={{ ...authContext, authState: state }}>
      {children}
    </AuthContext.Provider>
  );
};
