import { DynamoDBRepository } from "@rental-storage-project/common";

export class BookingRepository extends DynamoDBRepository {
  public static async create(): Promise<BookingRepository> {
    return new BookingRepository();
  }

  public async save(
    id: string,
    name: string,
    description: string
  ): Promise<any> {
    const params = {
      Item: {
        id: {
          S: id,
        },
        name: {
          S: name,
        },
        description: {
          S: description,
        },
      },
      ReturnConsumedCapacity: "TOTAL",
      TableName: process.env.TODO_TABLE_NAME,
    };

    return this._client
      .putItem(params as any)
      .promise()
      .then((data) => {
        return {
          id,
          name,
          description,
        };
      })
      .catch((err) => {
        console.log(err);
      });
  }

  public async delete(id: string): Promise<any> {
    const params = {
      Key: {
        id: {
          S: id,
        },
      },
      TableName: process.env.TODO_TABLE_NAME,
    };
    return this._client
      .deleteItem(params as any)
      .promise()
      .then((data) => {
        return {
          id,
        };
      })
      .catch((err) => {
        console.log(err);
      });
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
        console.log(err);
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
        console.log(err);
      });
  }
}
