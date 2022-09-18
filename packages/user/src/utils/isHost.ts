import { Host } from "../entity";

export const isHost = (data: any): data is Host => {
  return (
    (data as Host).firstName !== undefined &&
    (data as Host).lastName !== undefined
  );
};
