import { AWSRegion, UserEvent } from "../../../domain/enum";
import { LoggerUtil } from "../../../utils";
import {
  SNSClient,
  PublishCommand,
  PublishCommandInput,
} from "@aws-sdk/client-sns";

export class UserSNSMessageSender {
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

  private async publish(message: string, eventName: UserEvent) {
    const input: PublishCommandInput = {
      Message: message,
      TopicArn: `arn:aws:sns:${this._region}:${this._accountId}:user`,
      MessageAttributes: {
        sourceService: {
          DataType: "String",
          StringValue: "user",
        },
        event: {
          DataType: "String",
          StringValue: eventName,
        },
      },
    };
    const command = new PublishCommand(input);
    await this._client.send(command);
    this._logger.info(input.TopicArn, "publish()");
  }
}
