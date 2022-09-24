import {
  AWSRegion,
  BookingInterface,
  DynamoDBRepository,
} from "@rental-storage-project/common";
import { BookingMapper } from "../mapper";

export class BookingRepository extends DynamoDBRepository {
  public static async create(region: AWSRegion): Promise<BookingRepository> {
    return new BookingRepository(region, "BookingRepository");
  }

  public async save(booking: BookingInterface): Promise<void> {
    this._logger.debug(booking, "save()");
    const { id, status, amount, userId, listingId, createdAt, updatedAt } =
      booking;

    const params = {
      Item: {
        id: {
          S: id,
        },
        status: {
          S: status,
        },
        user_id: {
          S: userId,
        },
        listing_id: {
          S: listingId,
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

    await this._client
      .putItem(params as any)
      .promise()
      // .then((_data) => {
      //   return {
      //     id,
      //     status,
      //     userId,
      //     listingId,
      //     amount,
      //     createdAt,
      //     updatedAt,
      //   };
      // })
      .catch((err) => {
        this._logger.error(err, "save()");
      });
  }

  public async delete(id: string): Promise<any> {
    const params = {
      Key: {
        id: {
          S: id,
        },
      },
      TableName: process.env.TABLE_NAME,
    };
    return await this._client
      .deleteItem(params as any)
      .promise()
      .then((_data) => {
        return {
          id,
        };
      })
      .catch((err) => {
        this._logger.error(err, "delete()");
      });
  }

  public async findById(id: string): Promise<any> {
    var params = {
      Key: {
        id: { S: id },
      },
      TableName: process.env.TABLE_NAME,
    };

    const data = await this._client
      .getItem(params as any)
      .promise()
      .then((data: any) => data)
      .catch((err) => {
        this._logger.error(err, "findById()");
      });

    return BookingMapper.toDTOFromRaw(data);
  }

  public async findAll(): Promise<any> {
    const params = {
      TableName: process.env.TODO_TABLE_NAME,
    };
    return this._client
      .scan(params as any)
      .promise()
      .then((data: any) => {
        const todoList = [];
        for (let i = 0; i < data.Items!.length; i++) {
          todoList.push({
            id: data.Items[i].id.S,
            name: data.Items[i].name.S,
            description: data.Items[i].description.S,
          });
        }
        return todoList;
      })
      .catch((err) => {
        this._logger.error(err, "findAll()");
      });
  }

  public async update(
    id: string,
    name: string,
    description: string
  ): Promise<any> {
    const params = {
      ExpressionAttributeNames: {
        "#n": "name",
        "#d": "description",
      },
      ExpressionAttributeValues: {
        ":n": {
          S: name,
        },
        ":d": {
          S: description,
        },
      },
      Key: {
        id: {
          S: id,
        },
      },
      ReturnValues: "ALL_NEW",
      TableName: process.env.TODO_TABLE_NAME,
      UpdateExpression: "SET #n = :n, #d = :d",
    };
    return this._client
      .updateItem(params as any)
      .promise()
      .then((data: any) => {
        const body = data.Attributes;
        return {
          id: body.id.S,
          name: body.name.S,
          description: body.description.S,
        };
      })
      .catch((err) => {
        this._logger.error(err, "update()");
      });
  }
}
