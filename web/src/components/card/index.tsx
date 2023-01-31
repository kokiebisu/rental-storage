import { Space } from "@/types/interface";

export interface SpaceCardProps {
  space: Space;
}

export const SpaceCard = ({ space }: SpaceCardProps) => (
  <div key={space.id} className="group relative">
    <div className="min-h-80 aspect-w-1 aspect-h-1 w-full overflow-hidden rounded-md bg-gray-200 group-hover:opacity-75 lg:aspect-none lg:h-80">
      <img
        src={space.imageSrc}
        alt={space.imageAlt}
        className="h-full w-full object-cover object-center lg:h-full lg:w-full"
      />
    </div>
    <div className="mt-4 flex justify-between">
      <div>
        <h3 className="text-sm text-gray-700">
          <a href={space.href}>
            <span aria-hidden="true" className="absolute inset-0" />
            {space.name}
          </a>
        </h3>
        <p className="mt-1 text-sm text-gray-500">{space.color}</p>
      </div>
      <p className="text-sm font-medium text-gray-900">{space.price}</p>
    </div>
  </div>
);
