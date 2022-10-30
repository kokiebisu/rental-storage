import { createContext, useEffect, useReducer } from "react";
import { useQuery } from "@apollo/client";

import { QUERY_FIND_ME } from "../graphql";

export const ProfileContext = createContext(null);

export const ProfileContextProvider = ({ children }) => {
  const { data, error, loading } = useQuery(QUERY_FIND_ME);
  const [state, dispatch] = useReducer(
    (prevState, action) => {
      switch (action.type) {
        case "SET":
          return {
            ...prevState,
            uid: action.profile.uid,
            emailAddress: action.profile.emailAddress,
            firstName: action.profile.firstName,
            lastName: action.profile.lastName,
          };
        case "FLUSH":
          return {
            ...prevState,
            uid: null,
            emailAddress: null,
            firstName: null,
            lastName: null,
          };
      }
    },
    {
      uid: null,
      emailAddress: null,
      firstName: null,
      lastName: null,
    }
  );

  useEffect(() => {
    if (!error && !loading) {
      dispatch({
        type: "SET",
        profile: data.findMe,
      });
    }
  }, [data]);

  if (
    !error &&
    !loading &&
    (!state.uid || !state.emailAddress || !state.firstName || !state.lastName)
  ) {
    dispatch({
      type: "SET",
      profile: data.findMe,
    });
  }

  return (
    <ProfileContext.Provider
      value={{ profileState: state, profileDispatch: dispatch }}
    >
      {children}
    </ProfileContext.Provider>
  );
};
