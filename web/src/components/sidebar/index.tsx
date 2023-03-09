import { SearchBar } from "./searchBar";
import { SpaceCard } from "../card";
import { Space } from "@/types/interface";
import { useContext } from "react";
import { MapContext } from "@/context/map";

export const Sidebar = () => {
  const { spaces } = useContext(MapContext);
  return (
    <nav className="absolute top-0 h-full w-1/4 bg-white">
      <SearchBar />
    </nav>
  );
};
