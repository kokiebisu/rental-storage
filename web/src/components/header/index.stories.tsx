import type { Meta, StoryObj } from "@storybook/react";

import Component from "./template";

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Component> = {
  title: "Components",
  component: Component,
  argTypes: {},
};

export default meta;
type Story = StoryObj<typeof Component>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Header: Story = {
  args: {
    links: [
      {
        link: "/",
        label: "Borrow",
      },
    ],
    onSignInClicked: () => alert("Sign in clicked!"),
    onMapScreenNavigate: () => alert("Map screen navigate!"),
  },
};
