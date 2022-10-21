import { Item } from "../../domain/model";
import { ItemRepository, ItemService } from "../port";

import { LoggerUtil } from "../../utils";
import { ItemInterface } from "../../types";
import { ItemRepositoryImpl } from "../../adapters/repository";

export class ItemServiceImpl implements ItemService {
  private _itemRepository: ItemRepository;

  private _logger: LoggerUtil;

  private constructor(itemRepository: ItemRepository) {
    this._itemRepository = itemRepository;
    this._logger = new LoggerUtil("ItemServiceImpl");
  }

  public static async create() {
    const itemRepository = await ItemRepositoryImpl.create();
    await itemRepository.setup();
    return new ItemServiceImpl(itemRepository);
  }

  public async addItem(data: ItemInterface): Promise<boolean> {
    this._logger.info(data, "addItem()");
    try {
      const item = new Item({
        name: data.name,
        imageUrls: data.imageUrls,
        ownerId: data.ownerId,
        listingId: data.listingId,
        createdAt: new Date(data.createdAt),
        ...(data.updatedAt && { updatedAt: new Date(data.updatedAt) }),
      });

      await this._itemRepository.save(item);
      return true;
    } catch (err) {
      this._logger.error("Something went wrong", "addItem()");
      return false;
    }
  }
  // MOVE TO items-service
  // public async removeById(id: number): Promise<boolean> {
  //   this._logger.info({ id }, "removeById()");
  //   const userExists = await this.findById(id);
  //   if (!userExists) {
  //     throw new Error(`User with id ${id} doesn't exist`);
  //   }
  //   try {
  //     await this._itemRepository.delete(id);
  //     return true;
  //   } catch (err) {
  //     this._logger.error(err, "removeById()");
  //     return false;
  //   }
  // }

  // MOVE TO items-service
  // public async updateStoredItems(booking: BookingInterface): Promise<boolean> {
  //   this._logger.info({ booking }, "updatedStoredItems()");
  //   try {
  //     await this._itemRepository.updateStoringItem(
  //       booking.userId,
  //       booking.items
  //     );
  //     return true;
  //   } catch (err) {
  //     this._logger.error(err, "updateStoringItem");
  //     return false;
  //   }
  // }
}
