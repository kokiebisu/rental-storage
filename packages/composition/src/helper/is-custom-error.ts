import { CustomError } from "../error";

export default (err: unknown): err is CustomError => {
  return err instanceof CustomError;
};
