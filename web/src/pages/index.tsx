// import { useContext, useEffect, useState } from "react";
// import { AuthContext } from "@/context/auth";
// import { useRouter } from "next/router";
// import { Header, SignInModal } from "@/components";
import { Card, SignInModal } from "@/components";
import { Header } from "@/components";
import { useDisclosure } from "@mantine/hooks";
import Link from "next/link";

const HomePage = () => {
  const [opened, { open, close }] = useDisclosure(false);
  const links = [{ label: "Borrow", link: "/borrow" }];

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
                .fill()
                .map(() => {
                  return (
                    <Link href="/spaces/1">
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
  // const router = useRouter();
  // const [isLoading, setIsLoading] = useState(false);
  // const { logout, isAuthenticated } = useContext(AuthContext);
  // const handleSignOut = async () => {
  //   try {
  //     await logout();
  //     router.push("/login");
  //   } catch (err) {
  //     alert("something failed");
  //   }
  // };
  // useEffect(() => {
  //   setIsLoading(true);
  //   if (
  //     !isAuthenticated &&
  //     router.pathname !== "/login" &&
  //     router.pathname !== "/signup"
  //   ) {
  //     router.replace("/login");
  //   }
  //   setIsLoading(false);
  // }, []);
  // if (isLoading) {
  //   return <div>is loading...</div>;
  // }
  // return (
  //   <div>
  //     <HeaderMenu />
  //     <Button label="sign out" onClick={handleSignOut} />
  //   </div>
  // );
};

HomePage.displayName = "HomePage";

export default HomePage;
