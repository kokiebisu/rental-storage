import { AWSRegion, UserEvent } from "../../domain/enum";
import { LoggerUtil } from "../../utils";
import { UserInterface } from "../../domain/types";
import { UserEventSender } from "../../port";
import {
  KinesisClient,
  AddTagsToStreamCommand,
  PutRecordCommand,
  PutRecordCommandInput,
} from "@aws-sdk/client-kinesis";
import { v4 as uuid } from "uuid";

export class UserKinesisStreamEventSender implements UserEventSender {
  private _client: KinesisClient;
  private _logger: LoggerUtil;

  private constructor(region: AWSRegion, accountId: string) {
    this._client = new KinesisClient({ region });
    this._logger = new LoggerUtil("PublisherService");
  }
  public static async create(
    region: AWSRegion
  ): Promise<UserKinesisStreamEventSender> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new UserKinesisStreamEventSender(region, process.env.ACCOUNT_ID);
  }

  public async userCreated(data: UserInterface) {
    const event = {
      sourceEntity: "User",
      eventName: UserEvent.Created,
      data,
    };
    await this._publish(JSON.stringify(event));
  }

  private async _publish(message: string) {
    this._logger.info({ env: process.env.STAGE }, "YOYO");
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
