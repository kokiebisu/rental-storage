import * as React from "react";
import { View, Text, Image, Pressable } from "react-native";
import styles from "./styles";

const Post = ({ item, onPress, width }) => {
  return (
    <Pressable
      onPress={onPress}
      style={[styles.container, { width: width - 30 }]}
    >
      <View style={styles.innerContainer}>
        <Image
          style={styles.image}
          source={{
            uri: "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Garage__wide_shot_2.max-1300x1300.png",
          }}
        />
        <View
          style={{ flex: 1, marginHorizontal: 10, justifyContent: "center" }}
        >
          <View>
            <Text style={styles.description} numberOfLines={2}>
              {item.title}
            </Text>

            <Text style={styles.prices}>
              <Text style={styles.price}>${item.price} </Text>/ month
            </Text>
          </View>
        </View>
      </View>
    </Pressable>
  );
};

export default Post;
