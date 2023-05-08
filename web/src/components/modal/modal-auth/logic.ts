import { useContext, useState } from "react";
import { useForm } from "@mantine/form";

import { AuthContext } from "@/context/auth";
import { SignInParams, SignUpParams } from "@/types/interface";

const useAuthModal = ({ close }: any) => {
  const { signup, signin } = useContext(AuthContext);
  const [isSignInMode, setIsSignInMode] = useState(true);
  const [active, setActive] = useState(0);

  const handleSignUp = async (userInfo: SignUpParams) => {
    try {
      await signup(userInfo);
      close();
    } catch (err) {
      alert(err);
    }
  };

  const handleSignIn = async (userInfo: SignInParams) => {
    try {
      await signin(userInfo);
      close();
    } catch (err) {
      alert(err);
    }
  };

  const handleModeChange = () => {
    setIsSignInMode((current) => !current);
  };

  const form = useForm({
    initialValues: {
      firstName: "",
      lastName: "",
      password: "",
      emailAddress: "",
    },

    validate: (values) => {
      if (active === 0) {
        return {
          password:
            values.password.length < 6
              ? "Password must include at least 6 characters"
              : null,
          emailAddress: /^\S+@\S+$/.test(values.emailAddress)
            ? null
            : "Invalid email",
        };
      }

      if (active === 1) {
        return {
          firstName:
            values.firstName.trim().length < 2
              ? "First name must include at least 2 characters"
              : null,
          lastName:
            values.lastName.trim().length < 2
              ? "Last name must include at least 2 characters"
              : null,
        };
      }

      return {};
    },
  });

  const handleNextStep = () => {
    setActive((current) => {
      if (form.validate().hasErrors) {
        return current;
      }
      return current < 3 ? current + 1 : current;
    });
  };

  const handlePrevStep = () => {
    setActive((current) => (current > 0 ? current - 1 : current));
  };

  return {
    firstNameProps: form.getInputProps("firstName"),
    lastNameProps: form.getInputProps("lastName"),
    passwordProps: form.getInputProps("password"),
    emailAddressProps: form.getInputProps("emailAddress"),
    isSignInMode,
    active,
    handleModeChange,
    handleSignUp,
    handleSignIn,
    handleNextStep,
    handlePrevStep,
  };
};

export default useAuthModal;
