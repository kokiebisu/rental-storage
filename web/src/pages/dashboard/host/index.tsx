import { Button } from "@/components/button";
import Link from "next/link";

export default function Dashboard() {
  return (
    <div className="flex w-full h-full">
      <div className="mx-auto mt-10">
        <Link href="/dashboard/host/create">
          <Button onClick={() => {}} label="Add new listing" />
        </Link>
      </div>
    </div>
  );
}
