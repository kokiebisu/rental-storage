import { Space } from "@/types/interface";
import * as React from "react";
import { createContext, useState } from "react";

interface Position {
  lat: number;
  lng: number;
}

type UserContextType = {
  spaces: Space[] | null;
  setSpaces: React.Dispatch<React.SetStateAction<Space[] | null>>;
  center: Position;
  setCenter: React.Dispatch<React.SetStateAction<Position>>;
};

const iUserContextState = {
  spaces: [],
  setSpaces: () => {},
  center: {
    lat: 38.8977,
    lng: -77.0365,
  },
  setCenter: () => {},
};

export const MapContext = createContext<UserContextType>(iUserContextState);

export const MapContextProvider = ({ children }: any) => {
  const [spaces, setSpaces] = useState<Space[] | null>([]);
  const [center, setCenter] = useState<Position>({
    // This default position is used when a user's does not allow location access
    lat: 38.8977,
    lng: -77.0365,
  });
  return (
    <MapContext.Provider value={{ spaces, setSpaces, center, setCenter }}>
      {children}
    </MapContext.Provider>
  );
};
