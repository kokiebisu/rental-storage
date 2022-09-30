import { StorageItemInterface } from "../../../types";

export interface StorageItemRepository {
  setup(): Promise<void>;
  save(data: StorageItemInterface): Promise<{ insertId: number } | null>;
  delete(id: string): Promise<StorageItemInterface>;
  findOneById(id: string): Promise<StorageItemInterface>;
}
