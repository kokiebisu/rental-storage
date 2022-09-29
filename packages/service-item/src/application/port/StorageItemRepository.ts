import { StorageItemInterface } from "@rental-storage-project/common";

export interface StorageItemRepository {
  setup(): Promise<void>;
  save(data: StorageItemInterface): Promise<{ insertId: number }>;
  delete(id: string): Promise<StorageItemInterface>;
  findOneById(id: string): Promise<StorageItemInterface>;
}
