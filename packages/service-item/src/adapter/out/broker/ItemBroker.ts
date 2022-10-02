import {
  SNSClient,
  PublishCommand,
  PublishCommandInput,
} from "@aws-sdk/client-sns";

import { ItemBroker } from "../../../application/port";
import { AWSRegion } from "../../../domain/enum";
import { StorageItemInterface } from "../../../types";
import { LoggerUtil } from "../../../utils";

export class ItemBrokerImpl {
  private _client: SNSClient;
  private _region: AWSRegion;
  private _accountId: string;
  private _logger: LoggerUtil;

  private constructor(region: AWSRegion, accountId: string) {
    this._client = new SNSClient({ region });
    this._region = region;
    this._accountId = accountId;
    this._logger = new LoggerUtil("PublisherService");
  }

  public static async create(region: AWSRegion): Promise<ItemBroker> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new ItemBrokerImpl(region, process.env.ACCOUNT_ID);
  }

  public async dispatchItemSaved(data: StorageItemInterface) {
    this._logger.info(data, "dispatchItemSaved()");
    const stringified = JSON.stringify(data);
    try {
      await this.publish(stringified, this.generateTopicArn("ItemCreated"));
    } catch (err) {
      this._logger.error(err, "dispatchItemSaved()");
    }
  }

  private async publish(message: string, topicArn: string) {
    this._logger.info({ message, topicArn }, "publish()");
    const input: PublishCommandInput = {
      Message: message,
      TopicArn: topicArn,
    };
    const command = new PublishCommand(input);

    try {
      await this._client.send(command);
    } catch (err) {
      this._logger.error(err, "publish()");
    }
  }

  private generateTopicArn(topicName: string) {
    this._logger.info({ topicName }, "generateTopicArn()");
    return `arn:aws:sns:${this._region}:${this._accountId}:rental-storage-item-${process.env.NODE_ENV}-${topicName}`;
  }
}
