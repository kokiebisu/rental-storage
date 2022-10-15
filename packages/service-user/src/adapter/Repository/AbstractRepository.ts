import { Client } from "pg";

import { LoggerUtil } from "../../utils";

export abstract class AbstractRepositoryImpl<T> {
    protected readonly tableName: string;
    protected _logger: LoggerUtil;

    protected constructor(tableName: string, className: string) {
        this.tableName = tableName;
        this._logger = new LoggerUtil(className);
    }

    protected getDBClient() {
        const client = new Client({
          host: process.env.DB_HOST,
          database: process.env.DB_NAME,
          user: process.env.DB_USERNAME,
          password: process.env.DB_PASSWORD,
          port: 5432,
        });
        return client;
      }

    public abstract setup(): Promise<void>
    public abstract save(data: T): Promise<T>

    protected async startTransaction(operations: (data: T | string) => Promise<any>, client: Client, data: T | string): Promise<any> {
        let result: any
        this._logger.info({}, "startTransaction()");
        client.connect()
        try {
            await client.query('BEGIN')
            result = await operations(data)
            await client.query('COMMIT')
            return result
        } catch (err) {
            this._logger.error(err, "startTransaction()");
            await client.query('ROLLBACK')
            throw err;
        } finally {
            await client.end();
        }
    }
}