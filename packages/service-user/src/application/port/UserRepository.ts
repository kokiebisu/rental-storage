import { Payment, User } from "../../domain/model";

export interface UserRepository {
  setup(): Promise<void>;
  save(data: User): Promise<User>;
  delete(uid: string): Promise<void>;
  findOneById(uid: string): Promise<User | null>;
  findOneByEmail(emailAddress: string): Promise<User | null>;
}
