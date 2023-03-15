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
export const Card: Story = {
  render: (args: any) => (
    <div className="w-[380px]">
      <Component {...args} />
    </div>
  ),
  args: {
    title: "Card Title",
    address: "1234 Main St, City, State, 12345",
    imageUrls: ["https://picsum.photos/200"],
  },
};
