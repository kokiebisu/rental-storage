import { Guest } from "../entity";

export const isGuest = (data: any): data is Guest => {
  return (
    (data as Guest).firstName !== undefined &&
    (data as Guest).lastName !== undefined
  );
};
