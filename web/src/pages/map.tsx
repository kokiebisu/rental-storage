import { Header } from "@/components/header";
import Map from "@/components/map";
import { Sidebar } from "@/components/sidebar";

export default function MapPage() {
  return (
    <>
      <div className="min-h-full">
        <Header />
        <main className="relative">
          <div className="">
            <Map />
          </div>
          <Sidebar />
        </main>
      </div>
    </>
  );
}
