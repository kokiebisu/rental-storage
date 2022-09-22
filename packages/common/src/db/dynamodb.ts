import AWS from "aws-sdk";

export class DynamoDBRepository {
  protected _client: AWS.DynamoDB;
  public constructor() {
    this._client = new AWS.DynamoDB();
  }
}
