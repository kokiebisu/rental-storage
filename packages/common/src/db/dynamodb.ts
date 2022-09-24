import AWS from "aws-sdk";
import { AWSRegion } from "../enum";
import { LoggerService } from "../services";

export class DynamoDBRepository {
  protected _client: AWS.DynamoDB;
  protected _logger: LoggerService;
  public constructor(region: AWSRegion, repositoryName: string) {
    this._client = new AWS.DynamoDB({
      region,
    });
    this._logger = new LoggerService(repositoryName);
  }
}
