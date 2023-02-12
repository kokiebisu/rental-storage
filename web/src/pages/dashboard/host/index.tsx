import { Button } from "@/components/button";
import { Header } from "@/components/header";
import Link from "next/link";

export default function Dashboard() {
  return (
    <>
      <Header />
      <div>
        <div className="mx-auto mt-10 flex justify-center">
          <Link href="/dashboard/host/create">
            <Button onClick={() => {}} label="Add new space" />
          </Link>
        </div>
        <p className="text-center mt-4">My spaces</p>
      </div>
    </>
  );
}
