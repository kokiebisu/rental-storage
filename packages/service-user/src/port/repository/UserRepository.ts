import { User } from "../../domain/model";

export interface UserRepository {
  setup(): Promise<void>;
  save(data: User): Promise<User>;
  delete(uid: string): Promise<User>;
  findOneById(uid: string): Promise<User>;
  findOneByEmail(emailAddress: string): Promise<User>;
}
