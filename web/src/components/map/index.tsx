import { MapContext } from "@/context/map";
import { Space } from "@/types/interface";
import GoogleMapReact from "google-map-react";
import React, { useContext, useEffect, useState } from "react";
import { useGeolocated } from "react-geolocated";

const googleMapAPIKey = process.env.GOOGLE_MAP_API_KEY as string;

// fetch spaces latitude and longitude
const spaceData: Space[] = [
  {
    id: "1",
    title: "Langara",
    imageUrls: [
      "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
    ],
    location: {
      address: "",
      coordinate: {
        latitude: 49.2244,
        longitude: -123.1089,
      },
    },
  },
  {
    id: "2",
    title: "Second",
    imageUrls: [
      "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
    ],
    location: {
      address: "",
      coordinate: {
        latitude: 50.381832,
        longitude: -120.623177,
      },
    },
  },
  {
    id: "3",
    title: "Third Space",
    imageUrls: [
      "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
    ],
    location: {
      address: "",
      coordinate: {
        latitude: 48.052235,
        longitude: -118.243683,
      },
    },
  },
];

export default function Map() {
  const { spaces, setSpaces, center, setCenter } = useContext(MapContext);
  const { coords, isGeolocationEnabled } = useGeolocated({
    positionOptions: {
      enableHighAccuracy: false,
    },
    userDecisionTimeout: 50000,
    isOptimisticGeolocationEnabled: false,
  });
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    setSpaces(spaceData);
    if (coords) {
      setCenter({ lat: coords?.latitude, lng: coords?.longitude });
      setIsLoading(false);
      // isGeolocationEnabled will be true here
    }
  }, [coords]);

  const renderMarkers = (map: any, maps: any, space: Space) => {
    let marker = new maps.Marker({
      position: {
        lat: space.location.coordinate.longitude,
        lng: space.location.coordinate.longitude,
      },
      map,
    });
    return marker;
  };

  if (isLoading && isGeolocationEnabled)
    return <div>Sorry, We could not find you</div>;

  return (
    <div className="h-screen w-full">
      <GoogleMapReact
        bootstrapURLKeys={{ key: googleMapAPIKey }}
        center={center}
        zoom={10}
        onGoogleApiLoaded={({ map, maps }) => {
          spaces?.map((space) => {
            renderMarkers(map, maps, space);
          });
        }}
      />
    </div>
  );
}
