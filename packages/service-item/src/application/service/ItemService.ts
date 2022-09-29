import { StorageItem } from "../../domain/model";
import { StorageItemMapper } from "../../adapter/in/mapper";
import { ItemPublisherService } from "../../adapter/out/broker/ItemBroker";
import { ItemService, StorageItemRepository } from "../port";
import { StorageItemRepositoryImpl } from "../../adapter/out/db";
import { LoggerUtil } from "../../utils";
import { StorageItemInterface } from "../../types";
import { AWSRegion } from "../../domain/enum";

export class ItemServiceImpl implements ItemService {
  private _storageItemRepository: StorageItemRepository;
  private _logger: LoggerUtil;
  private _publisher: ItemPublisherService;

  private constructor(
    storageItemRepository: StorageItemRepository,
    publisher: ItemPublisherService
  ) {
    this._storageItemRepository = storageItemRepository;
    this._logger = new LoggerUtil("UserServiceImpl");
    this._publisher = publisher;
  }

  public static async create() {
    const storageItemRepository = await StorageItemRepositoryImpl.create();
    const publisher = await ItemPublisherService.create(AWSRegion.US_EAST_1);

    await storageItemRepository.setup();
    return new ItemServiceImpl(storageItemRepository, publisher);
  }

  public async addItem(data: StorageItemInterface): Promise<boolean> {
    this._logger.info(data, "addItem()");
    try {
      const entity = new StorageItem({
        name: data.name,
        imageUrls: data.imageUrls,
        userId: data.userId,
        listingId: data.listingId,
        createdAt: data.createdAt,
        ...(data.updatedAt && { updatedAt: data.updatedAt }),
      });
      const dto = StorageItemMapper.toDTOFromEntity(entity);
      const result = await this._storageItemRepository.save(dto);
      console.log("INSERTED RESULT: ", result);
      await this._publisher.dispatchItemSaved({
        ...dto,
        id: result.insertId.toString(),
      });

      return true;
    } catch (err) {
      this._logger.error("Something went wrong", "addItem()");
      return false;
    }
  }
}
