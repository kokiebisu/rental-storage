export const consumeEvent = async (event: any) => {
  console.log("CONSUMER LOG EVENT: ", event);
  try {
    for (const record of event.Records) {
      const messageAttributes = record.messageAttributes;
      console.log(
        "Message Attributtes -->  ",
        messageAttributes.AttributeNameHere.stringValue
      );
      console.log("Message Body -->  ", record.body);
      // Do something
    }
  } catch (error) {
    console.log(error);
  }
};
