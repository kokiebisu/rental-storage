import { Button } from "@/components";
import { DefaultLayout } from "@/layout";
import Link from "next/link";

export default function Dashboard() {
  return (
    <DefaultLayout>
      <div>
        <div className="mx-auto mt-10 flex justify-center">
          <Link href="/dashboard/lender/create">
            <Button onClick={() => {}} label="Add new space" />
          </Link>
        </div>
        <p className="text-center mt-4">My spaces</p>
      </div>
    </DefaultLayout>
  );
}
