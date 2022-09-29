import { StorageItemInterface } from "../../../types";

export interface ItemBroker {
  dispatchItemSaved(data: StorageItemInterface): Promise<void>;
}
