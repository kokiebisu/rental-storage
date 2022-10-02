import { StorageItem } from "../../domain/model";
import { StorageItemMapper } from "../../adapter/in/mapper";

import { ItemBroker, ItemService, StorageItemRepository } from "../port";
import { StorageItemRepositoryImpl } from "../../adapter/out/db";
import { LoggerUtil } from "../../utils";
import { StorageItemInterface } from "../../types";
import { AWSRegion } from "../../domain/enum";
import { ItemBrokerImpl } from "../../adapter/out/broker";

export class ItemServiceImpl implements ItemService {
  private _storageItemRepository: StorageItemRepository;
  private _logger: LoggerUtil;
  private _broker: ItemBroker;

  private constructor(
    storageItemRepository: StorageItemRepository,
    broker: ItemBroker
  ) {
    this._storageItemRepository = storageItemRepository;
    this._logger = new LoggerUtil("UserServiceImpl");
    this._broker = broker;
  }

  public static async create() {
    const storageItemRepository = await StorageItemRepositoryImpl.create();
    const publisher = await ItemBrokerImpl.create(AWSRegion.US_EAST_1);

    await storageItemRepository.setup();
    return new ItemServiceImpl(storageItemRepository, publisher);
  }

  public async addItem(data: StorageItemInterface): Promise<boolean> {
    this._logger.info(data, "addItem()");
    try {
      const entity = new StorageItem({
        name: data.name,
        imageUrls: data.imageUrls,
        guestId: data.guestId,
        listingId: data.listingId,
        createdAt: new Date(data.createdAt),
        ...(data.updatedAt && { updatedAt: new Date(data.updatedAt) }),
      });
      const dto = StorageItemMapper.toDTOFromEntity(entity);
      const result = await this._storageItemRepository.save(dto);
      console.log("INSERTED RESULT: ", result);
      if (result?.insertId) {
        await this._broker.dispatchItemSaved({
          ...dto,
          id: result.insertId,
        });
      }

      return true;
    } catch (err) {
      this._logger.error("Something went wrong", "addItem()");
      return false;
    }
  }
}
