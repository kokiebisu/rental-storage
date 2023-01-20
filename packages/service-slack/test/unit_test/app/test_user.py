from app.user import SlackUserMessageSenderService


def test_success__user_created():
    mock_data = {
        'first_name': "Ben",
        'last_name': 'Parker',
        'email_address': 'random@gmail.com'
    }
    message = SlackUserMessageSenderService\
        .generate_user_account_created_message(
            first_name=mock_data['first_name'],
            last_name=mock_data['last_name'],
            email_address=mock_data['email_address'])
    expected_message = (
        f'User Signed up! :tada:\n'
        f'Name: {mock_data["first_name"]} {mock_data["last_name"]}\n'
        f'Email Address: {mock_data["email_address"]}'
    )
    assert message == expected_message
