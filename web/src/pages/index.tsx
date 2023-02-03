import { useContext, useEffect, useState } from "react";
import { AuthContext } from "@/context/auth";
import { useRouter } from "next/router";
import { Button } from "@/components/button";

const HomePage = () => {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const { signOut } = useContext(AuthContext);

  const handleSignOut = async () => {
    const isSuccess = await signOut();
    if (isSuccess) {
      router.push("/login");
    } else {
      alert("something failed");
    }
  };

  useEffect(() => {
    setIsLoading(true);
    const isSignedIn = !!localStorage.getItem("authorizationToken");
    if (
      !isSignedIn &&
      router.pathname !== "/login" &&
      router.pathname !== "/signup"
    ) {
      router.replace("/login");
    }
    setIsLoading(false);
  }, []);

  if (isLoading) {
    return <div>is loading...</div>;
  }

  return (
    <div>
      <Button label="sign out" onClick={handleSignOut} />
    </div>
  );
};

HomePage.displayName = "HomePage";

export default HomePage;
