import { DefaultLayout } from "@/layout";

export default function Dashboard() {
  return (
    <DefaultLayout>
      <div className="w-full h-full">
        <div>
          <p className="text-center">My bookings (pending, confirmed)</p>
        </div>
      </div>
    </DefaultLayout>
  );
}
