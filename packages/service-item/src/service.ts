import {
  StorageItemInterface,
  LoggerService,
} from "@rental-storage-project/common";
import { StorageItem } from "./entity";
import { StorageItemMapper } from "./mapper";
import { StorageItemRepository } from "./repository";

interface ItemService {
  addItem(data: StorageItemInterface): Promise<boolean>;
}

export class ItemServiceImpl implements ItemService {
  private _storageItemRepository: StorageItemRepository;
  private _logger: LoggerService;

  private constructor(storageItemRepository: StorageItemRepository) {
    this._storageItemRepository = storageItemRepository;
    this._logger = new LoggerService("UserServiceImpl");
  }

  public static async create() {
    const storageItemRepository = await StorageItemRepository.create();

    await storageItemRepository.setup();
    return new ItemServiceImpl(storageItemRepository);
  }

  public async addItem(data: StorageItemInterface): Promise<boolean> {
    try {
      const entity = new StorageItem({
        name: data.name,
        imageUrls: data.imageUrls,
        userId: data.userId,
        listingId: data.listingId,
      });
      await this._storageItemRepository.save(
        StorageItemMapper.toDTOFromEntity(entity)
      );
      return true;
    } catch (err) {
      this._logger.error("Something went wrong", "addItem()");
      return false;
    }
  }
}
