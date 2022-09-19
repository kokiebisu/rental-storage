export class RDSRepository {
  public readonly tableName: string;
  protected _client: any;
  public constructor(client: any, tableName: string) {
    this._client = client;
    this.tableName = tableName;
  }
}
