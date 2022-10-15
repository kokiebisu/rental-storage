import { AWSRegion, UserEvent } from "../../domain/Enum";
import { LoggerUtil } from "../../utils";
import { UserInterface } from "../../types";
import { UserEventSender } from "../../application/Port";
import { KinesisClient, AddTagsToStreamCommand, PutRecordCommand, PutRecordCommandInput } from "@aws-sdk/client-kinesis";
import { v4 as uuid } from 'uuid'

export class UserKinesisStreamEventSender implements UserEventSender {
  private _client: KinesisClient;
  private _region: AWSRegion;
  private _accountId: string;
  private _logger: LoggerUtil;

  private constructor(region: AWSRegion, accountId: string) {
    this._client = new KinesisClient({ region });
    this._region = region;
    this._accountId = accountId;
    this._logger = new LoggerUtil("PublisherService");
  }
  public static async create(region: AWSRegion): Promise<UserKinesisStreamEventSender> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new UserKinesisStreamEventSender(region, process.env.ACCOUNT_ID);
  }

  public async userCreated(data: UserInterface) {
    const event = {
      sourceEntity: "User",
      eventName: UserEvent.Created,
      data
    }
    await this._publish(JSON.stringify(event))
  }

  private async _publish(message: string) {
    console.log("ENTERED1.2", message)
    
    try {
      const input: PutRecordCommandInput = {
        // Message: message,
        // TopicArn: `arn:aws:sns:${this._region}:${this._accountId}:${process.env.NODE_ENV}-${process.env.NAMESPACE}-user-topic`,
        // MessageAttributes: {
        //   entityType: {
        //     DataType: "String",
        //     StringValue: "user_account",
        //   },
        //   event: {
        //     DataType: "String",
        //     StringValue: eventName,
        //   },
        // },
        
        StreamName: "EventStream",
        PartitionKey: uuid(),
        Data: Uint8Array.from(Array.from(message).map(letter => letter.charCodeAt(0))),
      };
      console.log("ENTERED1.3", input)
      const command = new PutRecordCommand(input);
      await this._client.send(command);
      console.log("ENTERED1.4")
    } catch (err) {
      this._logger.info(err, "_publish()")
      throw err
    }

  }
}
