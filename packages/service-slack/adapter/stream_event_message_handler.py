import base64
import json


class KinesisEventStreamHandler:
    @staticmethod
    def parse_event(event):
        messages = []
        for record in event['Records']:
            decoded = base64.b64decode(record["kinesis"]["data"]).decode("utf-8")
            event = json.loads(decoded)
            messages.append({
                'eventEntity': event['sourceEntity'],
                'eventName': event['eventName'],
                'data': event['data']
            })
            
            # body = json.loads(record['body'])
            # message = body['Message']
            # entity_type = body['MessageAttributes']['entityType']['Value']
            # event_name = body['MessageAttributes']['event']['Value']
            # messages.append({
            #     'eventName': event_name,
            #     'entityType': entity_type,
            #     'data': json.loads(message)
            # })
        print(messages)
        return messages





    