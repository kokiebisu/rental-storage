import json

from domain.user import User


class SQSEventMessageHandler:
    @staticmethod
    def parse_event(event):
        messages = []
        for record in event['Records']:
            body = json.loads(record['body'])
            message = body['Message']
            entity_type = body['MessageAttributes']['entityType']['Value']
            event_name = body['MessageAttributes']['event']['Value']
            messages.append({
                'eventName': event_name,
                'entityType': entity_type,
                'data': json.loads(message)
            })
        return messages
