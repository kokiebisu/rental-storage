import { useQuery } from "@apollo/client";
import { createContext, useEffect, useState } from "react";
import { Client } from "../config/appsync";
import { QUERY_FIND_ME } from "../graphql";

export const ProfileContext = createContext(null);

export const ProfileContextProvider = ({ children }) => {
  const [profile, setProfile] = useState(null);
  const { data, error, loading } = useQuery(QUERY_FIND_ME, {
    client: Client,
  });

  useEffect(() => {
    if (!error && !loading && !profile) {
      setProfile(data.findMe);
    }
  }, [data]);

  return (
    <ProfileContext.Provider value={{ profile }}>
      {children}
    </ProfileContext.Provider>
  );
};
