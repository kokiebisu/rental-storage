import { useDisclosure } from "@mantine/hooks";

import { Footer, Header, AuthModal } from "@/components";

interface DefaultLayoutProps {
  children: React.ReactNode;
}

const DefaultLayout = ({ children }: DefaultLayoutProps) => {
  const [opened, { open, close }] = useDisclosure(false);
  return (
    <div>
      <Header onSignInClicked={open} />
      <div className="py-6">{children}</div>
      <AuthModal opened={opened} close={close} />
      <Footer />
    </div>
  );
};

export default DefaultLayout;
