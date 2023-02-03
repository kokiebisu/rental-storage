import React from "react";
import GoogleMapReact from "google-map-react";

const googleMapAPIKey = process.env.GOOGLE_MAP_API_KEY as string;

const AnyReactComponent = ({ text }: { text: string }) => <div>{text}</div>;

export default function SimpleMap() {
  // user's lat and lng
  const defaultProps = {
    center: {
      lat: 49.2827,
      lng: -123.1207,
    },
    zoom: 11,
  };

  return (
    <div style={{ height: "100vh", width: "100%" }} className="">
      <GoogleMapReact
        bootstrapURLKeys={{ key: googleMapAPIKey }}
        defaultCenter={defaultProps.center}
        defaultZoom={defaultProps.zoom}
      >
        <AnyReactComponent text="My Marker" />
      </GoogleMapReact>
    </div>
  );
}
