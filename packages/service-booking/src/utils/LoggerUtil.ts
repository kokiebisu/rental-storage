import { createLogger, transports, format, Logger } from "winston";

export class LoggerUtil {
  private _client: Logger;
  private _className: string;
  public constructor(className: string) {
    this._className = className;
    const options = {
      console: {
        format: format.combine(format.splat(), format.json()),
      },
    };
    this._client = createLogger({
      transports: [new transports.Console(options.console)],
      format: format.combine(
        format.timestamp(),
        format.printf(({ timestamp, level, message }) => {
          return `[${timestamp}] ${level}: ${message}`;
        })
      ),
    });
  }

  public debug(message: any, methodName: string) {
    this._client.debug({
      class: this._className,
      methodName,
      message,
    });
  }

  public info(message: any, methodName: string) {
    this._client.info({
      class: this._className,
      methodName,
      message,
    });
  }

  public warn(message: any, methodName: string) {
    this._client.warn({
      class: this._className,
      methodName,
      message,
    });
  }

  public error(error: any, methodName: string) {
    this._client.error({
      class: this._className,
      methodName,
      message: error.message,
    });
  }
}
