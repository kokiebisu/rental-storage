import { Payment, User } from "../../domain/model";

export interface UserRepository {
  setup(): Promise<void>;
  save(data: User): Promise<User>;
  delete(id: number): Promise<void>;
  findOneById(id: number): Promise<User | null>;
  findOneByEmail(emailAddress: string): Promise<User | null>;
}
