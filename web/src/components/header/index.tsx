import { Avatar } from "../";
import Link from "next/link";

export const Header = () => {
  return (
    <header className="bg-white shadow">
      <div className="flex justify-between items-center mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <Link href="/map">
          <h1 className="text-3xl font-bold tracking-tight text-gray-900">
            Rental Storage
          </h1>
        </Link>
        <div className="flex items-center">
          <div className="mr-8">
            <Link href="/dashboard/host">
              <button className="inline-flex items-center px-5 py-2 border border-transparent text-base font-medium rounded-full shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                Host Dashboard
              </button>
            </Link>
          </div>
          <div className="mr-8">
            <Link href="/dashboard/me">
              <button className="inline-flex items-center px-5 py-2 border border-transparent text-base font-medium rounded-full shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                User Dashboard
              </button>
            </Link>
          </div>
          <div>
            <Avatar />
          </div>
        </div>
      </div>
    </header>
  );
};
