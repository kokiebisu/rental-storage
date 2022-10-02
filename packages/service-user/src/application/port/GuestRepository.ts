import { GuestInterface, StorageItemInterface } from "../../types";

export interface GuestRepository {
  setup(): Promise<void>;
  save(data: GuestInterface): Promise<{ insertId: number } | undefined>;
  delete(id: number): Promise<GuestInterface>;
  findOneById(id: number): Promise<GuestInterface>;
  findAllItemIdsByGuestId(id: number): Promise<any>;
  updateStoringItem(id: number, items: StorageItemInterface[]): Promise<void>;
}
