import pino from "pino";

export default class {
  private client: pino.Logger;

  constructor() {
    this.client = pino({
      serializers: {
        req: pino.stdSerializers.req,
        res: pino.stdSerializers.res,
        err: pino.stdSerializers.err,
        // add any additional serializers you need
        // for example:
        payload: (obj) => obj, // just returns the object as is
      },
    });
  }

  public info(payload: unknown, filepath: string, line: number): void {
    this.client.info({ payload }, filepath, line);
  }

  public error(payload: unknown, filepath: string, line: number): void {
    this.client.error({ payload }, filepath, line);
  }

  public warn(payload: unknown, filepath: string, line: number): void {
    this.client.warn({ payload }, filepath, line);
  }

  public debug(payload: unknown, filepath: string, line: number): void {
    this.client.debug({ payload }, filepath, line);
  }
}
