from app.space import SlackSpaceMessageSenderService


def test_success__space_created():
    mock_data = {
        'lender_id': '123',
        'street_address': 'Random Street'
    }
    message = SlackSpaceMessageSenderService.\
        generate_space_created_message(
            lender_id=mock_data['lender_id'],
            street_address=mock_data['street_address'])
    expected_message = (
        f'Space was posted! :tada:\n'
        f'Lender Id: {mock_data["lender_id"]}\n'
        f'Street Address: {mock_data["street_address"]}\n'
    )
    assert message == expected_message


def test_success__space_deleted():
    mock_data = {
        'lender_id': '123',
        'street_address': 'Random Street'
    }
    message = SlackSpaceMessageSenderService.\
        generate_space_deleted_message(
            lender_id=mock_data['lender_id'],
            street_address=mock_data['street_address'])
    expected_message = (
        f'Space was deleted! :tada:\n'
        f'Lender Id: {mock_data["lender_id"]}\n'
        f'Street Address: {mock_data["street_address"]}\n'
    )
    assert message == expected_message
