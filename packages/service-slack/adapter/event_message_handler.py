import json


class SQSEventMessageHandler:
    @staticmethod
    def parse(event):
        messages = []
        for record in event['Records']:
            message = record['Sns']
            messages.append({
                'sourceService': message['MessageAttributes']['sourceService']['value'],
                'entity': json.dumps(message['Message'])
            })
        return messages