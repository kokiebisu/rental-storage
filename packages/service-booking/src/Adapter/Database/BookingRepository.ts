import {
  DynamoDBClient,
  PutItemCommand,
  PutItemCommandInput,
  DeleteItemCommand,
  DeleteItemCommandInput,
  GetItemCommand,
  GetItemCommandInput,
} from "@aws-sdk/client-dynamodb";

import { BookingMapper } from "../mapper";
import { BookingRepository } from "../../app/port";
import { LoggerUtil } from "../../utils";
import { AWSRegion } from "../../domain/enum";
import { BookingInterface } from "../../types";

export class BookingRepositoryImpl implements BookingRepository {
  private _client: DynamoDBClient;
  private _logger: LoggerUtil;

  private constructor(region: AWSRegion, repositoryName: string) {
    this._client = new DynamoDBClient({
      region,
    });
    this._logger = new LoggerUtil(repositoryName);
  }

  public static async create(
    region: AWSRegion
  ): Promise<BookingRepositoryImpl> {
    return new BookingRepositoryImpl(region, "BookingRepository");
  }

  public async save(booking: BookingInterface): Promise<void> {
    this._logger.info(booking, "save()");
    const {
      id,
      status,
      amount,
      ownerId,
      listingId,
      createdAt,
      updatedAt,
      items,
    } = booking;
    const itemsStringified = items.map((item: any) => {
      return {
        S: JSON.stringify(item),
      };
    });

    const input: PutItemCommandInput = {
      Item: {
        id: {
          S: id,
        },
        status: {
          S: status,
        },
        owner_id: {
          N: ownerId.toString(),
        },
        listing_id: {
          N: listingId.toString(),
        },
        amount: {
          M: {
            value: {
              N: amount.value.toString(),
            },
            currency: {
              S: amount.currency,
            },
          },
        },
        items: {
          L: itemsStringified,
        },
        created_at: {
          S: createdAt,
        },
        ...(updatedAt && {
          updated_at: {
            S: updatedAt,
          },
        }),
      },
      ReturnConsumedCapacity: "TOTAL",
      TableName: process.env.TABLE_NAME,
    };
    const command = new PutItemCommand(input);

    try {
      await this._client.send(command);
    } catch (err) {
      this._logger.error(err, "save()");
    }
  }

  public async delete(id: string): Promise<any> {
    const input: DeleteItemCommandInput = {
      Key: {
        id: {
          S: id,
        },
      },
      TableName: process.env.TABLE_NAME,
    };
    const command = new DeleteItemCommand(input);

    try {
      await this._client.send(command);
    } catch (err) {
      this._logger.error(err, "delete()");
    }
  }

  public async findById(id: string): Promise<any> {
    const input: GetItemCommandInput = {
      Key: {
        id: { S: id },
      },
      TableName: process.env.TABLE_NAME,
    };
    const command = new GetItemCommand(input);

    try {
      const data = await this._client.send(command);
      return BookingMapper.toDTOFromRaw(data);
    } catch (err) {
      this._logger.error(err, "findById()");
    }
  }

  // public async findAll(): Promise<any> {
  //   const params = {
  //     TableName: process.env.TODO_TABLE_NAME,
  //   };
  //   return this._client
  //     .scan(params as any)
  //     .promise()
  //     .then((data: any) => {
  //       const todoList = [];
  //       for (let i = 0; i < data.Items!.length; i++) {
  //         todoList.push({
  //           id: data.Items[i].id.S,
  //           name: data.Items[i].name.S,
  //           description: data.Items[i].description.S,
  //         });
  //       }
  //       return todoList;
  //     })
  //     .catch((err) => {
  //       this._logger.error(err, "findAll()");
  //     });
  // }

  // public async update(
  //   id: string,
  //   name: string,
  //   description: string
  // ): Promise<any> {
  //   const params = {
  //     ExpressionAttributeNames: {
  //       "#n": "name",
  //       "#d": "description",
  //     },
  //     ExpressionAttributeValues: {
  //       ":n": {
  //         S: name,
  //       },
  //       ":d": {
  //         S: description,
  //       },
  //     },
  //     Key: {
  //       id: {
  //         S: id,
  //       },
  //     },
  //     ReturnValues: "ALL_NEW",
  //     TableName: process.env.TODO_TABLE_NAME,
  //     UpdateExpression: "SET #n = :n, #d = :d",
  //   };
  //   return this._client
  //     .updateItem(params as any)
  //     .promise()
  //     .then((data: any) => {
  //       const body = data.Attributes;
  //       return {
  //         id: body.id.S,
  //         name: body.name.S,
  //         description: body.description.S,
  //       };
  //     })
  //     .catch((err) => {
  //       this._logger.error(err, "update()");
  //     });
  // }
}
