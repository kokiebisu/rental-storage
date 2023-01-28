import CustomError from "./base";

export default class ValidationError extends CustomError {
  statusCode = 500;
  private field: string;
  constructor(field: string) {
    super(`Validation failed`);
    this.field = field;
    Object.setPrototypeOf(this, ValidationError.prototype);
  }
  serializeError() {
    return {
      statusCode: this.statusCode,
      message: "Validation is wrong",
      field: this.field,
    };
  }
}
