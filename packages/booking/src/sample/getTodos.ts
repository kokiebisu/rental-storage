// const AWS = require("aws-sdk");
import AWS from "aws-sdk";

const dynamodb = new AWS.DynamoDB();

export const handler = async (event: any) => {
  const params = {
    TableName: process.env.TODO_TABLE_NAME,
  };

  return dynamodb
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
};
