import { UserInterface, StorageItemInterface } from "../../types";

export interface UserRepository {
  setup(): Promise<void>;
  save(data: UserInterface): Promise<{ insertId: number } | undefined>;
  delete(id: number): Promise<UserInterface>;
  findOneById(id: number): Promise<UserInterface | null>;
  findOneByEmail(emailAddress: string): Promise<UserInterface | null>;
}
