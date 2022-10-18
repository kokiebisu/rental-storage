import { AWSRegion, UserEvent } from "../../Domain/Enum";
import { LoggerUtil } from "../../Utils";
import {
  SNSClient,
  PublishCommand,
  PublishCommandInput,
} from "@aws-sdk/client-sns";
import { UserInterface } from "../../Types";
import { UserMessageSender } from "../../Application/Port";

export class UserSNSMessageSender implements UserMessageSender {
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
  public static async create(region: AWSRegion): Promise<UserSNSMessageSender> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new UserSNSMessageSender(region, process.env.ACCOUNT_ID);
  }

  public async userCreated(data: UserInterface) {
    const message = JSON.stringify(data);
    await this._publish(message, UserEvent.Created);
  }

  private async _publish(message: string, eventName: UserEvent) {
    try {
      const input: PublishCommandInput = {
        Message: message,
        TopicArn: `arn:aws:sns:${this._region}:${this._accountId}:${process.env.NODE_ENV}-${process.env.NAMESPACE}-user-topic`,
        MessageAttributes: {
          entityType: {
            DataType: "String",
            StringValue: "user_account",
          },
          event: {
            DataType: "String",
            StringValue: eventName,
          },
        },
      };
      const command = new PublishCommand(input);
      await this._client.send(command);
    } catch (err) {
      this._logger.info(err, "_publish()");
      throw err;
    }
  }
}
