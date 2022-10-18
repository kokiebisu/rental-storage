import { useQuery } from "@apollo/client";
import { useNavigation } from "@react-navigation/native";
import * as React from "react";
import { useEffect, useState } from "react";
import { View, Text, Button } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { Client } from "../../../../config/appsync";
import { QUERY_FIND_MY_LISTINGS } from "../../../../graphql";

export default () => {
  const navigation = useNavigation();
  const [listings, setListings] = useState([]);
  const { data, loading, error } = useQuery(QUERY_FIND_MY_LISTINGS, {
    client: Client,
  });

  useEffect(() => {
    if (!loading && !error) {
      setListings(data.findMyListings);
    }
  }, [data]);

  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <View>
        <Text>Listings</Text>
      </View>
      <View>
        {listings
          ? listings.map((listing) => {
              return (
                <View key={listing.uid}>
                  <View>
                    <Text>{listing.streetAddress}</Text>
                  </View>
                </View>
              );
            })
          : null}
      </View>
      <View>
        <Button
          title="Create Listing"
          onPress={() => navigation.navigate("Post")}
        />
      </View>
    </SafeAreaView>
  );
};
