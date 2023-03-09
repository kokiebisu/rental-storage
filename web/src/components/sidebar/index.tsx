import { SearchBar } from "./searchBar";

import { useContext } from "react";
import { MapContext } from "@/context/map";

export const Sidebar = () => {
  return (
    <nav className="absolute top-0 h-full w-1/4 bg-white">
      <SearchBar />
    </nav>
  );
};
