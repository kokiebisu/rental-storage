import { Button } from "@/components/button";
import { Avatar } from "@/components/avatar";

export const Header = () => {
  return (
    <header className="bg-white shadow">
      <div className="flex justify-between items-center mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 className="text-3xl font-bold tracking-tight text-gray-900">
          Rental Storage
        </h1>
        <div className="flex items-center">
          <div className="mr-8">
            <Button
              label="Add my space"
              onClick={() => alert("adding my place")}
            />
          </div>
          <div>
            <Avatar />
          </div>
        </div>
      </div>
    </header>
  );
};
