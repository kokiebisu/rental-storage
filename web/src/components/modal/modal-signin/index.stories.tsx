import type { Meta, StoryObj } from "@storybook/react";

import Component from "./template";

const meta: Meta<typeof Component> = {
  title: "Components/Modal",
  component: Component,
  argTypes: {},
};

export default meta;
type Story = StoryObj<typeof Component>;

export const SignIn: Story = {
  args: {
    username: "username",
    password: "password",
    name: "name",
    email: "email",
    active: 0,
    nextStep: () => {},
    prevStep: () => {},
  },
};
