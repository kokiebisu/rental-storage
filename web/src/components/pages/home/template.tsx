import { SpaceCard } from "@/components/card";
import { Header } from "@/components/header";
import SimpleMap from "@/components/map";
import { Space } from "@/types/interface";
import Link from "next/link";

export interface HomePageTemplateProps {
  spaces: Space[];
}

const HomePageTemplate = ({ spaces }: HomePageTemplateProps) => (
  <>
    <div className="min-h-full">
      <Header />
      <main>
        <div className="">
          <SimpleMap />
        </div>
      </main>
    </div>
  </>
);

HomePageTemplate.displayName = "HomePageTemplate";

export default HomePageTemplate;
