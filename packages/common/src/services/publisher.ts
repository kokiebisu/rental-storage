import AWS from "aws-sdk";
import { AWSRegion } from "aws-sdk/clients/cur";

export class PublisherService {
  protected _client: AWS.SNS;

  protected constructor(region: AWSRegion) {
    this._client = new AWS.SNS({ region });
  }

  protected async publish(message: string, topicArn: string) {
    const params = {
      Message: message,
      topicArn,
    };
    try {
      await this._client.publish(params).promise();
      console.debug(`Successfully published to topicArn: ${params.topicArn}`);
    } catch (err) {
      console.error(err);
    }
  }
}
