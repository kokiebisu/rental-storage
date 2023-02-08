import Router from "next/router";
import { useState } from "react";

export default function Dashboard() {
  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [latitude, setLatitude] = useState<number>();
  const [longitude, setLongitude] = useState<number>();
  const [address, setAddress] = useState<string>("");

  // POST AddListing
  const createNewListing = (e: any): void => {
    e.preventDefault();
    showListingInfo();
    Router.push("/dashboard/host");
  };

  // Modal or Message
  const showListingInfo = (): void => {
    let info: string = "title: " + title;
    info += "\ndescription: " + description;
    info += "\nlat: " + latitude + ", lng: " + longitude;
    info += "\naddress: " + address;
    alert(info);
  };

  return (
    <div className="flex w-full h-screen justify-center items-center">
      <div className="h-full w-full flex flex-col justify-center items-center">
        <form method="POST" className="w-1/3 flex flex-col justify-center">
          <input
            required
            type="text"
            placeholder="title"
            className="mb-4"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <textarea
            required
            rows={10}
            cols={80}
            name="description"
            placeholder="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          ></textarea>

          <label htmlFor="latitude">Latitude</label>
          <input
            required
            type="number"
            id="latitude"
            placeholder="latitude -90 ~ 90"
            className="mb-4 mx-3"
            min={-90}
            max={90}
            value={latitude}
            onChange={(e) => setLatitude(parseFloat(e.target.value))}
          />
          <label htmlFor="longitude">Longitude</label>
          <input
            required
            type="number"
            id="longitude"
            placeholder="longitude -180 ~ 180"
            className="mb-4 mx-3"
            min={-180}
            max={180}
            value={longitude}
            onChange={(e) => setLongitude(parseFloat(e.target.value))}
          />

          <label htmlFor="longitude">Address</label>
          <input
            required
            type="text"
            id="address"
            placeholder="Vancouver, BC"
            className="mb-4"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
          />
          <div className="w-full flex justify-center mt-2">
            <button
              type="submit"
              className="inline-flex items-center px-5 py-2 border border-transparent text-base font-medium rounded-full shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              onClick={createNewListing}
            >
              Create
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
