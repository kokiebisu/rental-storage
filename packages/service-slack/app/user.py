from adapter.bot import SlackBotAdapter


class SlackUserMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot):
        self.bot = bot

    def alert_user_sign_up(self):
        self.bot.send_message_to_channel("Congrats! A User signed up")
