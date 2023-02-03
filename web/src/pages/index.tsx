import { useContext, useEffect, useState } from "react";
import { AuthContext } from "@/context/auth";
import { useRouter } from "next/router";
import { Button } from "@/components/button";
import { useQuery } from "@apollo/client";
import { QUERY_FIND_ME } from "@/graphql";

const HomePage = () => {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const { logout, isAuthenticated } = useContext(AuthContext);

  const handleSignOut = async () => {
    try {
      await logout();
      router.push("/login");
    } catch (err) {
      alert("something failed");
    }
  };

  useEffect(() => {
    setIsLoading(true);
    if (
      !isAuthenticated &&
      router.pathname !== "/login" &&
      router.pathname !== "/signup"
    ) {
      router.replace("/login");
    }
    setIsLoading(false);
  }, []);

  const { data, error, loading } = useQuery(QUERY_FIND_ME);

  if (!loading) {
    return <div>loading...</div>;
  }

  if (error) {
    return <div>error</div>;
  }

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
