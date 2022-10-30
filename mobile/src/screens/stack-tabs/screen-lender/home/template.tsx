import * as React from "react";
import { useEffect, useState } from "react";
import { View, Button, FlatList, Pressable } from "react-native";
import { useQuery } from "@apollo/client";
import { useNavigation } from "@react-navigation/native";
import { SafeAreaView } from "react-native-safe-area-context";

import { QUERY_FIND_MY_LISTINGS } from "../../../../graphql";
import { Card } from "../../../../components/card";
import { Typography } from "../../../../components/typography";
import { Spacing } from "../../../../components/spacing";

export default () => {
  const navigation = useNavigation();
  const [listings, setListings] = useState([]);
  const { data, loading, error } = useQuery(QUERY_FIND_MY_LISTINGS);

  useEffect(() => {
    if (!loading && !error) {
      setListings(data.findMyListings);
    }
  }, [data]);

  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <View>
        <Spacing variant="lg">
          <Typography variant="h2">Listing</Typography>
        </Spacing>
      </View>
      <FlatList
        showsVerticalScrollIndicator={false}
        data={listings}
        renderItem={({ item }) => {
          return (
            <View style={{ marginBottom: 16 }}>
              <Pressable onPress={() => alert("pressed")}>
                <Card
                  variant="listing"
                  streetAddress={item.streetAddress}
                  imageUrls={item.imageUrls}
                />
              </Pressable>
            </View>
          );
        }}
      />
      <View>
        <Button
          title="Create Listing"
          onPress={() => navigation.navigate("Post")}
        />
      </View>
    </SafeAreaView>
  );
};
