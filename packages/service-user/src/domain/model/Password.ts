import { genSalt, hash } from "bcryptjs";

export class Password {
  public readonly value: string;
  private constructor(value: string) {
    this.value = value;
  }

  public static async create(value: string) {
    this.validatePassword(value);
    const encryptedPassword = await this.encryptPassword(value);
    return new Password(encryptedPassword);
  }

  private static async encryptPassword(
    value: string,
    salt = 10
  ): Promise<string> {
    const generatedSalt = await genSalt(salt);
    return await hash(value, generatedSalt);
  }

  private static validatePassword(value: string) {
    if (!value) {
      throw new Error("Password was not provided");
    }
    if (value.length === 0) {
      throw new Error("Password cannot be empty");
    }
  }
}
