from adapter.bot import SlackBotAdapter


class SlackUserMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot_adapter: SlackBotAdapter):
        self.bot = bot_adapter

    def alert_user_sign_up(self, data):
        print("ENTERED3", data)
        payload = {
            'channel_name': 'C0464PCNZH8',
            'message': 'Congrats! A User signed up'
        }
        self.bot.send_message(channel_name=payload['channel_name'], message=payload['message'])
