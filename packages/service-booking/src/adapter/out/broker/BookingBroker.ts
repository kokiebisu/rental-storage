import { AWSRegion } from "../../../domain/enum";
import { LoggerUtil } from "../../../utils";
import AWS from "aws-sdk";

export class BookingPublisherService {
  private _client: AWS.SNS;
  private _region: AWSRegion;
  private _accountId: string;
  private _logger: LoggerUtil;

  private constructor(region: AWSRegion, accountId: string) {
    this._client = new AWS.SNS({ region });
    this._region = region;
    this._accountId = accountId;
    this._logger = new LoggerUtil("PublisherService");
  }
  public static async create(
    region: AWSRegion
  ): Promise<BookingPublisherService> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new BookingPublisherService(region, process.env.ACCOUNT_ID);
  }

  private async publish(message: string, topicName: string) {
    const params = {
      Message: message,
      TopicArn: `arn:aws:sns:${this._region}:${this._accountId}:${topicName}`,
    };

    await this._client.publish(params).promise();
    this._logger.info(params.TopicArn, "publish()");
  }
}
