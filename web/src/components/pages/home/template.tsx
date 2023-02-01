import { SpaceCard } from "@/components/card";
import { Header } from "@/components/header";
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
        <div className="mx-auto max-w-2xl pt-4 pb-16 px-4 sm:pt-4 sm:pb-24 sm:px-6 lg:max-w-7xl lg:px-8">
          <div className="bg-white">
            <div className="py-8 sm:py-8">
              <div className="mt-6 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
                {spaces.map((space) => (
                  <Link key={space.id} href={"/space/1"}>
                    <SpaceCard space={space} />
                  </Link>
                ))}
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </>
);

HomePageTemplate.displayName = "HomePageTemplate";

export default HomePageTemplate;
