import { Header } from "@/components/header";

export default function Dashboard() {
  return (
    <div className="w-full h-full">
      <Header />
      <div>
        <p className="text-center">My bookings (pending, confirmed)</p>
      </div>
    </div>
  );
}
