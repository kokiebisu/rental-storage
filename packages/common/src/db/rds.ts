import { LoggerService } from "../services";

export class RDSRepository {
  public readonly tableName: string;
  protected _client: any;
  protected _logger: LoggerService;
  public constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerService(className);
  }
}
