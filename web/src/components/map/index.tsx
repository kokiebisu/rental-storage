import React, { useState } from "react";
import GoogleMapReact from "google-map-react";

const googleMapAPIKey = process.env.GOOGLE_MAP_API_KEY as string;

// fetch listings latitude and longitude
// should use interface later
const markers = [
  {
    id: 1,
    name: "Langara College",
    position: { lat: 49.2244, lng: -123.1089 },
  },
  {
    id: 2,
    name: "Whatever",
    position: { lat: 50.381832, lng: -120.623177 },
  },
  {
    id: 3,
    name: "Los Angeles, California",
    position: { lat: 34.052235, lng: -118.243683 },
  },
  {
    id: 4,
    name: "Hometown",
    position: { lat: 35.52389, lng: 139.69294 },
  },
];

export default function Map() {
  // first focus (user's location)
  const [center, setCenter] = useState({
    lat: 49.2827,
    lng: -123.1207,
  });
  const [zoom, setZoom] = useState(8);

  const renderMarkers = (map: any, maps: any, mark: any) => {
    let marker = new maps.Marker({
      position: mark.position,
      map,
    });
    return marker;
  };

  const handleClick = (): void => {
    setZoom(zoom + 0.5);
  };

  const handleMove = (): void => {
    console.log(center);
    const randomPos =
      markers[Math.floor(Math.random() * markers.length)].position;
    setCenter(randomPos);
  };

  return (
    <div className="h-screen w-full">
      <GoogleMapReact
        bootstrapURLKeys={{ key: googleMapAPIKey }}
        center={center}
        zoom={zoom}
        // put markers on a map
        onGoogleApiLoaded={({ map, maps }) => {
          markers.map((marker) => {
            renderMarkers(map, maps, marker);
          });
        }}
      >
        <button
          onClick={handleClick}
          className="text-red-500 font-bold text-2xl"
        >
          Click ME
        </button>
        <button
          onClick={handleMove}
          className="text-blue-500 font-bold text-2xl mx-5"
        >
          Move Random
        </button>
      </GoogleMapReact>
    </div>
  );
}
