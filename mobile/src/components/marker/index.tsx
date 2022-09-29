import { Marker } from "react-native-maps";
import { Text, View } from "react-native";

export const CustomMarker = ({
  onPress,
  latitude,
  longitude,
  price,
  isSelected,
}) => {
  return (
    <Marker onPress={onPress} coordinate={{ latitude, longitude }}>
      <View
        style={{
          padding: 5,
          backgroundColor: isSelected ? "black" : "white",
          borderRadius: 20,
          borderColor: "grey",
          borderWidth: 1,
        }}
      >
        <Text
          style={{ fontWeight: "700", color: isSelected ? "white" : "black" }}
        >
          ${price}
        </Text>
      </View>
    </Marker>
  );
};
