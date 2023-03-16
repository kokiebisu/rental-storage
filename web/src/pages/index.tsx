import { useEffect } from "react";
import { useDisclosure } from "@mantine/hooks";
import { Card, SignInModal, Header } from "@/components";
import { useQuery } from "@apollo/client";
import { FIND_SPACES_QUERY } from "@/queries";
import { apiKeyClient } from "@/apollo";

const HomePage = () => {
  const limit = 6;
  const [opened, { open, close }] = useDisclosure(false);
  const links = [{ label: "Borrow", link: "/borrow" }];
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
          variables: { offset: data.spaces.length, limit },
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
    return <div>loading</div>;
  }
  if (error) {
    return <div>error</div>;
  }

  return (
    <div>
      <Header links={links} onSignInClicked={open} />
      <main>
        <div className="max-w-7xl mx-auto px-5 2xl:px-0 w-full">
          <section>
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
                ? data.spaces.map((space: any) => {
                    return (
                      <div key={space.id}>
                        <Card
                          title={space.title}
                          address={space.location.address}
                          imageUrls={space.imageUrls}
                        />
                      </div>
                    );
                  })
                : null}
            </div>
          </section>
        </div>
        <SignInModal opened={opened} close={close} />
      </main>
    </div>
  );
};

HomePage.displayName = "HomePage";

export default HomePage;
