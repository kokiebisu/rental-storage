import { Stepper, Group } from "@mantine/core";

import { SignUpParams } from "@/hooks/useAuth";

import { Button, TextInput, PasswordInput } from "../..";

export interface SignInModalTemplateProps {
  firstNameProps: any;
  lastNameProps: any;
  passwordProps: any;
  emailAddressProps: any;
  active: number;
  handleSignUp: (userInfo: SignUpParams) => void;
  nextStep: () => void;
  prevStep: () => void;
}

const SignInModalTemplate = ({
  firstNameProps,
  lastNameProps,
  passwordProps,
  emailAddressProps,
  active,
  handleSignUp,
  nextStep,
  prevStep,
}: SignInModalTemplateProps) => {
  return (
    <>
      <Stepper active={active} breakpoint="sm">
        <Stepper.Step label="First step" description="Profile settings">
          <TextInput
            label="FirstName"
            placeholder="First Name"
            {...firstNameProps}
          />
          <div className="mt-2">
            <TextInput
              label="Last Name"
              placeholder="Last Name"
              {...lastNameProps}
            />
          </div>
        </Stepper.Step>

        <Stepper.Step label="Second step" description="Personal information">
          <TextInput
            label="Email Address"
            placeholder="Email Address"
            {...emailAddressProps}
          />
          <div className="mt-2">
            <PasswordInput {...passwordProps} />
          </div>
        </Stepper.Step>

        {/* <Stepper.Completed>
          Completed!
          <Code block mt="xl">
            {JSON.stringify(
              {
                firstName: firstNameProps.value,
                lastName: lastNameProps.value,
                emailAddress: emailAddressProps.value,
                password: passwordProps.value,
              },
              null,
              2
            )}
          </Code>
        </Stepper.Completed> */}
      </Stepper>

      <Group position="right" mt="xl">
        {active !== 0 && <Button label="Back" onClick={prevStep} />}
        {active == 1 && (
          <Button
            label="Sign Up"
            onClick={() =>
              handleSignUp({
                firstName: firstNameProps.value,
                lastName: lastNameProps.value,
                emailAddress: emailAddressProps.value,
                password: passwordProps.value,
              })
            }
          />
        )}
        {active != 1 && active != 2 && (
          <Button label="Next Step" onClick={nextStep} />
        )}
      </Group>
    </>
  );
};

export default SignInModalTemplate;
