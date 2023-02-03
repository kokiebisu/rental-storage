import * as React from "react";
import axios from "axios";
import { createContext, useMemo } from "react";

export const AuthContext = createContext({
  // authState: {
  //   isLoading: false,
  // },
  signUp: (_: any) => Promise<void>,
  signOut: () => Promise<void>,
});

export const AuthContextProvider = ({ children }: any) => {
  // TO BE DETERMINED TO REMOVE OR NOT...
  // const [state, dispatch] = useReducer(
  //   (prevState: any, action: any) => {
  //     switch (action.type) {
  //       case "RESTORE_TOKEN":
  //         return {
  //           ...prevState,
  //           userToken: action.token,
  //           isLoading: false,
  //         };
  //       case "SIGN_IN":
  //         return {
  //           ...prevState,
  //           isSignout: false,
  //           userToken: action.token,
  //         };
  //       case "SIGN_OUT":
  //         return {
  //           ...prevState,
  //           isSignout: true,
  //           userToken: null,
  //         };
  //     }
  //   },
  //   {
  //     isLoading: true,
  //     isSignout: false,
  //     userToken: null,
  //   }
  // );

  //   useEffect(() => {
  //     // fetch the token from storage then navigate to our appropriate place
  //     const bootstrapAsync = async () => {
  //       let userToken;

  //       try {
  //         userToken = await SecureStore.getItemAsync("userToken");
  //       } catch (err) {
  //         // Restoring token failed
  //       }

  //       // after restoring token, we may need to validate it in production apps

  //       // this will switch to the App screen or Auth screen and this loading
  //       // screen will be unmounted and thrown away
  //       dispatch({ type: "RESTORE_TOKEN", token: userToken });
  //     };

  //     bootstrapAsync();
  //   }, []);

  const authContext = useMemo(
    () => ({
      signUp: async (data: any) => {
        if (
          !data.emailAddress ||
          !data.password ||
          !data.firstName ||
          !data.lastName
        ) {
          alert("Input missing");
          return;
        }
        if (!process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT) {
          alert("endpoint missing");
          return;
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
          alert("Something went wrong");
          return;
        }

        const { authorizationToken } = response.data;
        // in a production app, we need to send some data (usually username, password) to server and get a token
        // we will also need to handle errors if sign in faield
        try {
          // after getting token, we need to persist the token using 'SecureStore'
          // in the example, we'll use a dummy token
          await localStorage.setItem("authorizationToken", authorizationToken);
          return authorizationToken;
        } catch (err) {
          console.error(err);
        }
      },
      signOut: async () => {
        try {
          await localStorage.removeItem("authorizationToken");
          return true;
        } catch (err) {
          alert("Somethign went wrong");
          return false;
        }
      },
    }),
    []
  );

  return (
    <AuthContext.Provider value={{ ...authContext } as any}>
      {children}
    </AuthContext.Provider>
  );
};
