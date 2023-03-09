export default function SpaceOverviewPage() {
  return (
    <div className="w-full">
      <div
        className="w-3/4 border mx-auto mt-6 flex items-center justify-center"
        style={{ height: 500 }}
      >
        Space Images
      </div>
      <div className="flex w-3/4 mx-auto my-10">
        <div className="w-1/2 h-full">
          <div className="w-full mx-auto h-60 border flex justify-center items-center">
            Space Description
          </div>
          <div
            className="mx-auto border mt-6 flex justify-center items-center"
            style={{ height: 400, width: 500 }}
          >
            Map
          </div>
        </div>
        <div className="w-1/2 h-full">
          <div className="w-full mx-auto h-72 border flex justify-center items-center">
            Calendar
          </div>
          <div className="flex justify-center items-center mt-4">
            <textarea
              placeholder="3 suit cases and 1 backpack"
              cols={60}
              rows={8}
            ></textarea>
          </div>
          <button
            className="flex mx-auto justify-center items-center mt-4 border py-2 px-5"
            onClick={() => alert("add image")}
          >
            Add Images
          </button>
          {/* <div className="flex justify-center mt-4">
            <Button onClick={() => alert("booking")} label="make booking" />
          </div> */}
        </div>
      </div>
    </div>
  );
}
