export interface CreateUserInput {
  emailAddress: string;
  firstName: string;
  lastName: string;
  password: string;
}

export interface UserService {
  createUser(data: CreateUserInput): Promise<{ uid: string } | null>;
}
