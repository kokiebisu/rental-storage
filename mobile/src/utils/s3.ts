const generateEndpoint = async ({ filename }) => {
  try {
    const {
      data: { url },
    } = await UserPoolClient.query({
      query: USER_POOL_GET_PRESIGNED_URL_S3,
      variables: {
        filename,
        filepath: "receipts",
        bucket: ImageBuckets.support,
      },
    });

    return url.split("url=")[1].slice(0, -1);
  } catch (error) {
    consoleUtil({
      error,
      category: "apollo",
      level: "Error",
      message: `[UPLOAD IMAGE ERROR]: ${error}`,
    });
  }
};
