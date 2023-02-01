from app.booking import SlackBookingMessageSenderService


def test_success__user_booking():
    mock_data = {
        'first_name': "Ben",
        'last_name': 'Parker',
        'email_address': 'random@gmail.com',
        'user_id': 'random123'
    }
    message = SlackBookingMessageSenderService\
        .generate_booking_created_message(
            first_name=mock_data['first_name'],
            last_name=mock_data['last_name'],
            email_address=mock_data['email_address'],
            user_id=mock_data['user_id'])
    expected_message = (
        f'User booking created! :tada:\n'
        f'Name: {mock_data["first_name"]} {mock_data["last_name"]}\n'
        f'Email Address: {mock_data["email_address"]}\n'
        f'User Id: {mock_data["user_id"]}\n'
    )
    assert message == expected_message
