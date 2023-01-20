from app.listing import SlackListingMessageSenderService

def test_success__listing_created():
    mock_data = {
        'lender_id': '123',
        'street_address': 'Random Street'
    }
    message = SlackListingMessageSenderService.generate_listing_created_message(lender_id=mock['lender_id'], street_address=mock['street_address'])
    expected_message = (
        f'Listing was posted! :tada:\n'
        f'Lender Id: {mock_data["lender_id"]}\n'
        f'Street Address: {mock_data["street_address"]}\n'
    )
    assert message == expected_message
