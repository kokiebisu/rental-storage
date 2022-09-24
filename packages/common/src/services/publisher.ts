import AWS from "aws-sdk";
import { LoggerService } from "..";
import { AWSRegion } from "../enum";

export class PublisherService {
  protected _client: AWS.SNS;
  protected _region: AWSRegion;
  protected _accountId: string;
  protected _logger: LoggerService;

  protected constructor(region: AWSRegion, accountId: string) {
    this._client = new AWS.SNS({ region });
    this._region = region;
    this._accountId = accountId;
    this._logger = new LoggerService("PublisherService");
  }

  protected async publish(message: string, topicName: string) {
    const params = {
      Message: message,
      TopicArn: `arn:aws:sns:${this._region}:${this._accountId}:${topicName}`,
    };

    await this._client.publish(params).promise();
    this._logger.info(params.TopicArn, "publish()");
  }
}
