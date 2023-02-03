import { Header } from "@/components/header";
import SimpleMap from "@/components/map";
import { Sidebar } from "@/components/sidebar";

const HomePageTemplate = () => (
  <>
    <div className="min-h-full">
      <Header />
      <main className="relative">
        <div className="">
          <SimpleMap />
        </div>
        <Sidebar />
      </main>
    </div>
  </>
);

HomePageTemplate.displayName = "HomePageTemplate";

export default HomePageTemplate;
