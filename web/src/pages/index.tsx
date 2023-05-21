import { useEffect } from "react";
import Link from "next/link";
import { useQuery } from "@apollo/client";

import { Card } from "@/components";
import { FIND_SPACES_QUERY } from "@/graphql/queries";
import { apiKeyClient } from "@/clients";
import { DefaultLayout } from "@/layout";
import { Spinner } from "@/components/spinner";
import { Space } from "@/types/interface";

export default function LandingPage() {
  const limit = 6;
  const { data, loading, error, fetchMore } = useQuery(FIND_SPACES_QUERY, {
    client: apiKeyClient,
    variables: {
      filter: {
        offset: 0,
        limit,
      },
    },
  });

  useEffect(() => {
    const handleScroll = () => {
      const scrollTop = document.documentElement.scrollTop;
      const windowHeight = window.innerHeight;
      const scrollHeight = document.documentElement.scrollHeight;

      if (scrollTop + windowHeight >= scrollHeight) {
        fetchMore({
          variables: { filter: { offset: data.spaces.length, limit } },
          updateQuery: (prev, { fetchMoreResult }) => {
            if (!fetchMoreResult) return prev;
            return { spaces: [...prev.spaces, ...fetchMoreResult.spaces] };
          },
        });
      }
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, [data, fetchMore]);

  if (loading) {
    return (
      <div className="absolute top-1/2 left-1/2">
        <Spinner />
      </div>
    );
  }
  if (error) {
    return <div>error</div>;
  }

  return (
    <DefaultLayout>
      <main>
        <div className="py-12 max-w-7xl mx-auto px-5 2xl:px-0 w-full">
          <section className="pb-12">
            <h3 className="text-3xl font-bold">Near your area</h3>
            <div className="mt-4">
              <h5>
                Refine your Canada real estate search by price, bedroom, or type
                (house, townhouse, or condo). View up-to-date MLS® listings in
                Canada.
              </h5>
            </div>
            <div className="mt-4">
              <h5>1 storage available in Canada</h5>
            </div>
            <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-5 mt-6">
              {data.spaces.length > 0
                ? data.spaces.map((space: Space, index: number) => {
                    return (
                      <Link href={`/spaces/${space.id}`} key={index}>
                        <Card
                          title={space.title}
                          address={space.location.address}
                          imageUrls={space.imageUrls}
                        />
                      </Link>
                    );
                  })
                : null}
            </div>
          </section>
        </div>
      </main>
    </DefaultLayout>
  );
}
