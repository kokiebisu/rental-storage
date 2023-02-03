import React from "react";
import GoogleMapReact from "google-map-react";

const googleMapAPIKey = process.env.GOOGLE_MAP_API_KEY as string;

// fetch listings latitude and longitude
// should use interface later
const markers = [
  {
    id: 0,
    name: "Langara College",
    position: { lat: 49.2244, lng: -123.1089 },
  },
  {
    id: 1,
    name: "Whatever",
    position: { lat: 50.381832, lng: -120.623177 },
  },
  {
    id: 2,
    name: "Denver, Colorado",
    position: { lat: 39.739235, lng: -104.99025 },
  },
  {
    id: 3,
    name: "Los Angeles, California",
    position: { lat: 34.052235, lng: -118.243683 },
  },
  {
    id: 4,
    name: "New York, New York",
    position: { lat: 40.712776, lng: -74.005974 },
  },
  {
    id: 5,
    name: "Hometown",
    position: { lat: 35.52389, lng: 139.69294 },
  },
];

export default function SimpleMap() {
  // first focus (user's location)
  const defaultProps = {
    center: {
      lat: 49.2827,
      lng: -123.1207,
    },
    zoom: 8,
  };

  const renderMarkers = (map: any, maps: any, mark: any) => {
    let marker = new maps.Marker({
      position: mark.position,
      map,
    });
    return marker;
  };

  return (
    <div style={{ height: "100vh", width: "100%" }} className="">
      <GoogleMapReact
        bootstrapURLKeys={{ key: googleMapAPIKey }}
        defaultCenter={defaultProps.center}
        defaultZoom={defaultProps.zoom}
        // put markers on a map
        onGoogleApiLoaded={({ map, maps }) => {
          markers.map((marker) => {
            renderMarkers(map, maps, marker);
          });
        }}
      ></GoogleMapReact>
    </div>
  );
}
