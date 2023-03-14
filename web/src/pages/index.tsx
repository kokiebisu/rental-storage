import Link from "next/link";
import { useDisclosure } from "@mantine/hooks";
import { Card, SignInModal, Header } from "@/components";
import { useQuery } from "@apollo/client";
import { FIND_PROFILE_QUERY } from "@/queries/user";
import { useContext } from "react";
import { CustomApollo } from "@/context/apollo";

const HomePage = () => {
  const [opened, { open, close }] = useDisclosure(false);
  const links = [{ label: "Borrow", link: "/borrow" }];
  const { data, loading, error } = useQuery(FIND_PROFILE_QUERY);

  if (loading) {
    return <div>loading</div>;
  }
  if (error) {
    return <div>error</div>;
  }

  console.log("DATA: ", data);

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
              {Array(5)
                .fill(null)
                .map(() => {
                  return (
                    <Link key={null} href="/spaces/1">
                      <Card />
                    </Link>
                  );
                })}
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
