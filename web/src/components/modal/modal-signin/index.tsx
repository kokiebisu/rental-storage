import SignInModalTemplate from "./template";
import useLogic from "./logic";
import ModalWrapper from "../wrapper";

const SignInModal = ({ opened, close }: any) => {
  const logic = useLogic({ close });
  return (
    <ModalWrapper opened={opened} close={close}>
      <SignInModalTemplate {...logic} />
    </ModalWrapper>
  );
};

export default SignInModal;
