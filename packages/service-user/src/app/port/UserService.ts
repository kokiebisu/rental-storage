import { UserInterface } from "../../types";

export interface CreateUserInput {
  emailAddress: string;
  firstName: string;
  lastName: string;
  password: string;
}

export interface UserService {
  createUser(data: CreateUserInput): Promise<UserInterface>;
  removeById(uid: string): Promise<boolean>;
  findById(uid: string): Promise<UserInterface>;
  findByEmail(emailAddress: string): Promise<UserInterface>;
}
