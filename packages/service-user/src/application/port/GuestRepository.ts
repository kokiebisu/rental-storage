import { GuestInterface } from "../../types";
import { StorageItemInterface } from "@rental-storage-project/item";

export interface GuestRepository {
  setup(): Promise<void>;
  save(data: GuestInterface): Promise<GuestInterface>;
  delete(id: string): Promise<GuestInterface>;
  findOneById(id: string): Promise<GuestInterface>;
  findAllItemIdsByUserId(guestId: string): Promise<any>;
  updateStoringItem(
    userId: string,
    items: StorageItemInterface[]
  ): Promise<void>;
}
