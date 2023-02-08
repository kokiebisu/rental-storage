import { MapContext } from "@/context/map";
import { Space } from "@/types/interface";
import Link from "next/link";
import { useContext } from "react";

export interface SpaceCardProps {
  space: Space;
}

export const SpaceCard = ({ space }: SpaceCardProps) => {
  const { setCenter } = useContext(MapContext);
  return (
    <Link
      onMouseEnter={() => setCenter({ lat: space.lat, lng: space.lng })}
      href={`/listings/${space.id}`}
      className="mx-auto w-full my-2 block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
    >
      <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
        {space.name}
      </h5>
      <p className="font-normal text-gray-700 dark:text-gray-400">
        Here are the biggest enterprise technology acquisitions of 2021 so far,
        in reverse chronological order.
      </p>
    </Link>
  );
};
