export const handler = async (event: any) => {
  console.log("CONSUMER LOG EVENT: ", event);
  try {
    for (const record of event.Records) {
      const { Message: message } = JSON.parse(record.body);
      console.log("message: ", message);
    }
  } catch (error) {
    console.log(error);
  }
};
