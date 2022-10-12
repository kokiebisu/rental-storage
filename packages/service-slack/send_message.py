import logging
import boto3

def handler(event, _):
    logging.info("EVENT: ", event)
    ssm = boto3.client('ssm')
    parameter = ssm.get_parameter(Name='/dev/service-slack/bot-api-key', WithDecryption=True)
    print(parameter['Parameter']['Value'])
    
    return {
        "statusCode": 200,
        "body": "hello"
    }
