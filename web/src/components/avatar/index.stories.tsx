import type { Meta, StoryObj } from "@storybook/react";

import { Avatar } from ".";

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Avatar> = {
  title: "Components/Avatar",
  component: Avatar,
  argTypes: {},
};

export default meta;
type Story = StoryObj<typeof Avatar>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const NoProfilePic: Story = {
  args: {},
};
