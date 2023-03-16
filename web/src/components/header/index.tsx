import { useContext } from "react";

import { AuthContext } from "@/context/auth";
import HeaderTemplate from "./template";
import { useRouter } from "next/router";

interface HeaderSearchProps {
  onSignInClicked: () => void;
}

const Header = ({ onSignInClicked }: HeaderSearchProps) => {
  const router = useRouter();
  const links = [{ label: "Borrow", link: "/borrow" }];
  const { user, signout } = useContext(AuthContext);

  const handleGuestDashboardRedirect = () => {
    router.push("/dashboard/guest");
  };

  const handleLenderDashboardRedirect = () => {
    router.push("/dashboard/lender");
  };

  return (
    <HeaderTemplate
      links={links}
      onSignInClicked={onSignInClicked}
      isAuthenticated={!!user}
      handleSignout={signout}
      handleGuestDashboardRedirect={handleGuestDashboardRedirect}
      handleLenderDashboardRedirect={handleLenderDashboardRedirect}
    />
  );
};

export default Header;
