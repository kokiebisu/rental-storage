import AWS from "aws-sdk";
import { AWSRegion } from "../enum";

export class DynamoDBRepository {
  protected _client: AWS.DynamoDB;
  public constructor(region: AWSRegion) {
    this._client = new AWS.DynamoDB({
      region,
    });
  }
}
