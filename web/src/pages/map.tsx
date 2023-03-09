import Map from "@/components/map";
import { Sidebar } from "@/components/sidebar";
import { MapContextProvider } from "@/context/map";

export default function MapPage() {
  return (
    <>
      <MapContextProvider>
        <div className="min-h-full">
          <main className="relative">
            <div className="">
              <Map />
            </div>
            <Sidebar />
          </main>
        </div>
      </MapContextProvider>
    </>
  );
}
