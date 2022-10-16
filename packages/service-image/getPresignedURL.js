var AWS = require('aws-sdk');
var s3 = new AWS.S3({
  signatureVersion: 'v4',
});


exports.handler = (event, context) => {
  console.log('event:'+ JSON.stringify(event));
  const {filename} = event.arguments
    
  try {
  const url = s3.getSignedUrl('putObject', {
    Bucket: 'dev-rental-a-locker-listing-profile',
    ACL: 'public-read',
    Key: filename,
  });

  console.log('presigned url: ', url);  

  console.log("REUSLT: ", {
    filename,
    url
  })

  return {
    filename,
    url
  }
  } catch (err) {
    console.error(err)
    throw err
  }




};