export class RDSRepository {
  public readonly tableName: string;
  protected _client: any;
  public constructor(client: any, tableName: string) {
    this._client = client;
    this.tableName = tableName;
  }

  public async findById(id: string, filter: any): Promise<any> {
    var response = await this._client.execute(
      `SELECT * name FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    if (response.length == 0) {
      throw new Error(`id not found in ${this.tableName}`);
    }
    return response;
  }
}
