import {
  AWSRegion,
  PublisherService,
  StorageItemInterface,
} from "@rental-storage-project/common";

export class ItemPublisherService extends PublisherService {
  public static async create(region: AWSRegion): Promise<ItemPublisherService> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new ItemPublisherService(region, process.env.ACCOUNT_ID);
  }

  public async dispatchItemSaved(data: StorageItemInterface) {
    this._logger.info(data, "dispatchItemSaved()");
    const stringified = JSON.stringify(data);
    try {
      await this.publish(stringified, "item-saved");
    } catch (err) {
      this._logger.error(err, "dispatchItemSaved()");
    }
  }
}
