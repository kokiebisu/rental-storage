import CustomError from "./base";

export default class InternalServerError extends CustomError {
  statusCode = 500;
  constructor() {
    super("Internal server error");
    Object.setPrototypeOf(this, InternalServerError.prototype);
  }
  serializeError() {
    return {
      statusCode: this.statusCode,
      message: "something went wrong",
    };
  }
}
