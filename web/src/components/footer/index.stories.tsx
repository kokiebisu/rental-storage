import type { Meta, StoryObj } from "@storybook/react";

import Component from ".";

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Component> = {
  title: "Components",
  component: Component,
  argTypes: {},
};

export default meta;
type Story = StoryObj<typeof Component>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Footer: Story = {
  args: {
    data: [
      {
        title: "Company",
        links: [
          {
            link: "/",
            label: "Borrow",
          },
        ],
      },
    ],
  },
};
