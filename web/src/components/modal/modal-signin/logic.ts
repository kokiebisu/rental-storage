import { useContext, useState } from "react";
import { useForm } from "@mantine/form";
import { useRouter } from "next/router";

import { AuthContext } from "@/context/auth";
import { SignUpParams } from "@/hooks/useAuth";

const useSignInModal = () => {
  const router = useRouter();
  const { signup } = useContext(AuthContext);
  const [active, setActive] = useState(0);

  const handleSignUp = async (userInfo: SignUpParams) => {
    try {
      await signup(userInfo);
      router.reload();
    } catch (err) {
      alert(err);
    }
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

      if (active === 1) {
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

      return {};
    },
  });

  const nextStep = () => {
    setActive((current) => {
      if (form.validate().hasErrors) {
        return current;
      }
      return current < 3 ? current + 1 : current;
    });
  };

  const prevStep = () => {
    setActive((current) => (current > 0 ? current - 1 : current));
  };

  return {
    firstNameProps: form.getInputProps("firstName"),
    lastNameProps: form.getInputProps("lastName"),
    passwordProps: form.getInputProps("password"),
    emailAddressProps: form.getInputProps("emailAddress"),
    handleSignUp,
    active,
    nextStep,
    prevStep,
  };
};

export default useSignInModal;
