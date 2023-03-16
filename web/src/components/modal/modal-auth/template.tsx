import { SignInParams } from "@/types/interface";
import { Stepper, Group } from "@mantine/core";

import { Button, TextInput, PasswordInput } from "../..";

interface SignUpParams {
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
}

export interface AuthModalTemplateProps {
  isSignInMode: boolean;
  firstNameProps: any;
  lastNameProps: any;
  passwordProps: any;
  emailAddressProps: any;
  active: number;
  handleModeChange: () => void;
  handleSignUp: (userInfo: SignUpParams) => void;
  handleSignIn: (userInfo: SignInParams) => void;
  handleNextStep: () => void;
  handlePrevStep: () => void;
}

const AuthModalTemplate = ({
  isSignInMode,
  firstNameProps,
  lastNameProps,
  passwordProps,
  emailAddressProps,
  active,
  handleModeChange,
  handleSignUp,
  handleSignIn,
  handleNextStep,
  handlePrevStep,
}: AuthModalTemplateProps) => {
  return isSignInMode ? (
    <>
      <TextInput
        label="Email Address"
        placeholder="Email Address"
        {...emailAddressProps}
      />
      <div className="mt-2">
        <PasswordInput {...passwordProps} />
      </div>
      <Group grow>
        <Group position="left" mt="xl">
          <Button
            label={isSignInMode ? "Sign Up" : "Sign In"}
            onClick={handleModeChange}
          />
        </Group>
        <Group position="right" mt="xl">
          <Button
            label="Sign In"
            onClick={() =>
              handleSignIn({
                emailAddress: emailAddressProps.value,
                password: passwordProps.value,
              })
            }
          />
        </Group>
      </Group>
    </>
  ) : (
    <>
      <Stepper active={active} breakpoint="sm">
        <Stepper.Step label="First step" description="Profile settings">
          <TextInput
            label="Email Address"
            placeholder="Email Address"
            {...emailAddressProps}
          />
          <div className="mt-2">
            <PasswordInput {...passwordProps} />
          </div>
        </Stepper.Step>

        <Stepper.Step label="Second step" description="Personal information">
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
      <Group grow>
        <Group position="left" mt="xl">
          <Button
            label={isSignInMode ? "Sign Up" : "Sign In"}
            onClick={handleModeChange}
          />
        </Group>
        <Group position="right" mt="xl">
          {active !== 0 && <Button label="Back" onClick={handlePrevStep} />}
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
            <Button label="Next Step" onClick={handleNextStep} />
          )}
        </Group>
      </Group>
    </>
  );
};

export default AuthModalTemplate;
