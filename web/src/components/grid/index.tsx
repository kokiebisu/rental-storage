import { Skeleton } from "@mantine/core";

type GalleryGridProps = { imageUrls: string[]; title: string };

const GalleryGrid = ({ title, imageUrls }: GalleryGridProps) => (
  <div className="max-w-8xl mx-auto">
    <div className="px-12 2xl:px-0">
      <div className="">
        <div className="pt-7">
          <div className="block">
            <div className="rounded-md overflow-y-hidden">
              <div
                className="relative min-h-[300px] overflow-y-hidden"
                style={{
                  maxHeight: "calc(60vh - 64px)",
                }}
              >
                <div
                  className="h-0 min-h-full min-w-full relative"
                  style={{ paddingTop: "50%" }}
                >
                  <div className="absolute top-0 left-0 w-full h-full">
                    <div className="h-full w-1/2 top-0 left-0 absolute">
                      <div
                        style={{ maxHeight: "calc(60vh - 64px" }}
                        className="min-h-[300px] h-full w-full bg-gray-100"
                      >
                        {imageUrls?.[0] ? (
                          <img
                            alt={title}
                            src={imageUrls[0]}
                            className="h-full w-full bg-gray-100 object-cover"
                          />
                        ) : (
                          <Skeleton animate height={"100%"} />
                        )}
                      </div>
                    </div>
                    <div className="flex flex-col h-full w-1/4 pl-2 left-1/2 top-0 absolute">
                      <div
                        style={{ maxHeight: "calc(60vh - 64px" }}
                        className="min-h-[300px] h-full w-full"
                      >
                        <div className="h-1/2">
                          {imageUrls?.[1] ? (
                            <img
                              alt={title}
                              src={imageUrls[1]}
                              className="h-full w-full bg-gray-100 object-cover"
                            />
                          ) : (
                            <Skeleton animate height={"100%"} />
                          )}
                        </div>
                        <div className="h-1/2 pt-2">
                          {imageUrls?.[2] ? (
                            <img
                              alt={title}
                              src={imageUrls[2]}
                              className="h-full w-full bg-gray-100 object-cover"
                            />
                          ) : (
                            <Skeleton animate height={"100%"} />
                          )}
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col h-full w-1/4 pl-2 right-0 top-0 absolute">
                      <div
                        style={{ maxHeight: "calc(60vh - 64px" }}
                        className="min-h-[300px] h-full w-full"
                      >
                        <div className="h-1/2">
                          {imageUrls?.[3] ? (
                            <img
                              alt={title}
                              src={imageUrls[3]}
                              className="h-full w-full bg-gray-100 object-cover"
                            />
                          ) : (
                            <Skeleton animate height={"100%"} />
                          )}
                        </div>
                        <div className="h-1/2 pt-2">
                          {imageUrls?.[4] ? (
                            <img
                              alt={title}
                              src={imageUrls[4]}
                              className="h-full w-full bg-gray-100 object-cover"
                            />
                          ) : (
                            <Skeleton animate height={"100%"} />
                          )}
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
);

export default GalleryGrid;
