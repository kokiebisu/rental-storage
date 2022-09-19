// const AWS = require("aws-sdk");
import AWS from "aws-sdk";

const dynamodb = new AWS.DynamoDB();

export const handler = async (event: any) => {
  const id = event.arguments.id;
  const name = event.arguments.name;
  const description = event.arguments.description;

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

  return dynamodb
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
};
