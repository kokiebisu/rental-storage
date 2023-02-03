import { SearchBar } from "./searchBar";
import spacesArr from "../pages/home/hook";
import { SpaceCard } from "../card";
import { Space } from "@/types/interface";

const spaces: Space[] = spacesArr().spaces;
export const Sidebar = () => {
  return (
    <nav className="absolute top-0 h-full w-1/4 bg-white">
      <SearchBar />
      <div className="">
        {spaces.map((space: Space) => (
          <SpaceCard space={space} />
        ))}
      </div>
    </nav>
  );
};
