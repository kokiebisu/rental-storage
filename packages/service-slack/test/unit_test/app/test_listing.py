from app.listing import SlackListingMessageSenderService

def test_success__listing_created():
    mock_adapter = {}
    mock = {
        "lender_id": "123",
        "street_address": "Random Street"
    }
    service = SlackListingMessageSenderService(bot_adapter=mock_adapter)
    message = service._generate_listing_created_message(lender_id="123", street_address="Random Street")
    expected_message = (
        f'Listing was posted! :tada:\n'
        f'Lender Id: {mock["lender_id"]}\n'
        f'Street Address: {mock["street_address"]}\n'
    )
    assert message == expected_message