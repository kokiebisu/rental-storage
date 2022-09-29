import { StorageItemInterface } from "@rental-storage-project/common";

export interface ItemService {
  addItem(data: StorageItemInterface): Promise<boolean>;
}
