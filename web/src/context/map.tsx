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
    lat: 49.2827,
    lng: -123.1207,
  },
  setCenter: () => {},
};

export const MapContext = createContext<UserContextType>(iUserContextState);

export const MapContextProvider = ({ children }: any) => {
  const [spaces, setSpaces] = useState<Space[] | null>([]);
  const [center, setCenter] = useState<Position>({
    lat: 49.2827,
    lng: -123.1207,
  });
  return (
    <MapContext.Provider value={{ spaces, setSpaces, center, setCenter }}>
      {children}
    </MapContext.Provider>
  );
};
