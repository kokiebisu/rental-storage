import AWS from "aws-sdk";
import { AWSRegion } from "../enum";

export class PublisherService {
  protected _client: AWS.SNS;
  protected _region: AWSRegion;
  protected _accountId: string;

  protected constructor(region: AWSRegion, accountId: string) {
    this._client = new AWS.SNS({ region });
    this._region = region;
    this._accountId = accountId;
  }

  protected async publish(message: string, topicName: string) {
    const params = {
      Message: message,
      TopicArn: `arn:aws:sns:${this._region}:${this._accountId}:${topicName}`,
    };
    console.debug("PUBLISH PARAMS: ", params);
    try {
      await this._client.publish(params).promise();
      console.debug(`Successfully published to topicArn: ${params.TopicArn}`);
    } catch (err) {
      console.error(err);
    }
  }
}
