export const appsyncConfig = {
  GRAPHQL_ENDPOINT: process.env.NEXT_PUBLIC_APPSYNC_ENDPOINT || "",
  REGION: process.env.NEXT_PUBLIC_APPSYNC_REGION || "us-east-1",
  AUTHENTICATION_TYPE: "AWS_LAMBDA",
};
