import {
  KinesisClient,
  AddTagsToStreamCommand,
  PutRecordCommand,
  PutRecordCommandInput,
} from "@aws-sdk/client-kinesis";
import { ListingInterface } from "../../types";
import { LoggerUtil } from "../../utils";
import { v4 as uuid } from "uuid";
import { AWSRegion } from "../../domain/enum";
import { ListingEvent } from "../../domain/enum";
import { ListingEventSender } from "../../app/port";

export class ListingKinesisStreamEventSender implements ListingEventSender {
  private _client: KinesisClient;
  private _logger: LoggerUtil;

  private constructor(region: AWSRegion, accountId: string) {
    this._client = new KinesisClient({ region });
    this._logger = new LoggerUtil("PublisherService");
  }
  public static async create(
    region: AWSRegion
  ): Promise<ListingKinesisStreamEventSender> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new ListingKinesisStreamEventSender(region, process.env.ACCOUNT_ID);
  }

  public async listingCreated(data: ListingInterface) {
    const event = {
      sourceEntity: "Listing",
      eventName: ListingEvent.Created,
      data,
    };
    await this._publish(JSON.stringify(event));
  }

  private async _publish(message: string) {
    try {
      const input: PutRecordCommandInput = {
        StreamName: `${process.env.STAGE}-EventStream`,
        PartitionKey: uuid(),
        Data: Uint8Array.from(
          Array.from(message).map((letter) => letter.charCodeAt(0))
        ),
      };
      const command = new PutRecordCommand(input);
      await this._client.send(command);
    } catch (err) {
      this._logger.info(err, "_publish()");
      throw err;
    }
  }
}
