// import { useContext, useEffect, useState } from "react";
// import { AuthContext } from "@/context/auth";
// import { useRouter } from "next/router";
import { Header, SignInModal } from "@/components";
import { useDisclosure } from "@mantine/hooks";

const HomePage = () => {
  const [opened, { open, close }] = useDisclosure(false);
  const links = [{ label: "Borrow", link: "/borrow" }];

  return (
    <main>
      <Header links={links} onSignInClicked={open} />
      <div>
        <h3>Near your area</h3>
        <section></section>
      </div>
      <SignInModal opened={opened} close={close} />
    </main>
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
