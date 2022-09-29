import { StorageItemInterface } from "../../types";

export interface ItemService {
  addItem(data: StorageItemInterface): Promise<boolean>;
}
