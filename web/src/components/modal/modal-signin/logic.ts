import { useForm } from "@mantine/form";
import { useState } from "react";

const useSignInModal = () => {
  const [active, setActive] = useState(0);

  const form = useForm({
    initialValues: {
      username: "",
      password: "",
      name: "",
      email: "",
      website: "",
      github: "",
    },

    validate: (values) => {
      if (active === 0) {
        return {
          username:
            values.username.trim().length < 6
              ? "Username must include at least 6 characters"
              : null,
          password:
            values.password.length < 6
              ? "Password must include at least 6 characters"
              : null,
        };
      }

      if (active === 1) {
        return {
          name:
            values.name.trim().length < 2
              ? "Name must include at least 2 characters"
              : null,
          email: /^\S+@\S+$/.test(values.email) ? null : "Invalid email",
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
    usernameProps: form.getInputProps("username"),
    passwordProps: form.getInputProps("password"),
    nameProps: form.getInputProps("name"),
    emailProps: form.getInputProps("email"),
    active,
    nextStep,
    prevStep,
  };
};

export default useSignInModal;
