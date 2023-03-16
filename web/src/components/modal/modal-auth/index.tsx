import AuthModalTemplate from "./template";
import useLogic from "./logic";
import ModalWrapper from "../wrapper";

const AuthModal = ({ opened, close }: any) => {
  const logic = useLogic({ close });
  return (
    <ModalWrapper opened={opened} close={close}>
      <AuthModalTemplate {...logic} />
    </ModalWrapper>
  );
};

export default AuthModal;
