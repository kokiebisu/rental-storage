import { useContext } from "react";
import { AuthContext } from "@/context/auth";
import { useRouter } from "next/router";
import { Button } from "@/components/button";

const HomePage = () => {
  const router = useRouter();
  const { signOut } = useContext(AuthContext);

  const handleSignOut = async () => {
    const isSuccess = await signOut();
    if (isSuccess) {
      router.push("/login");
    } else {
      alert("something failed");
    }
  };
  return (
    <div>
      <Button label="sign out" onClick={handleSignOut} />
    </div>
  );
};

HomePage.displayName = "HomePage";

export default HomePage;
