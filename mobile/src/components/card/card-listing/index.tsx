import * as React from "react";
import { View, Text, Image } from "react-native";

export interface ListingCardProps {
  imageUrls: string[];
  streetAddress: string;
}

export const ListingCard = ({ streetAddress, imageUrls }: ListingCardProps) => {
  return (
    <View>
      <View style={{ height: 300 }}>
        <Image
          style={{
            // height: "100%",
            width: "100%",
            aspectRatio: 5 / 4,
            resizeMode: "cover",
          }}
          source={{
            uri: imageUrls[0],
          }}
        />
      </View>
      <Text>{streetAddress}</Text>
    </View>
  );
};
