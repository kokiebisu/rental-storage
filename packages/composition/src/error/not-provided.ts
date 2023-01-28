import CustomError from "./base";

export default class NotProvidedError extends CustomError {
  statusCode = 500;
  private field: string;
  constructor(field: string) {
    super(`Field not provided`);
    this.field = field;
    Object.setPrototypeOf(this, NotProvidedError.prototype);
  }
  serializeError() {
    return {
      statusCode: this.statusCode,
      message: "Field not provided",
      field: this.field,
    };
  }
}
