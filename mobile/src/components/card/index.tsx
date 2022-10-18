import * as React from "react";
import { ListingCard, ListingCardProps } from "./card-listing";
import { PostCard, PostCardProps } from "./card-post";

export type CardProps =
  | ({ variant: "post" } & PostCardProps)
  | ({ variant: "listing" } & ListingCardProps);

export const Card = ({ variant, ...props }: CardProps) => {
  switch (variant) {
    case "post":
      return <PostCard {...props} />;
    case "listing":
      return <ListingCard {...props} />;
    default:
      throw new Error("Invalid card variant");
  }
};
