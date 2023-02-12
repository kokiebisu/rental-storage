resource "aws_sqs_queue_policy" "booking" {
    queue_url = aws_sqs_queue.booking.id

    policy = <<POLICY
    {
        "Version": "2012-10-17",
        "Id": "sqspolicy",
        "Statement": [
            {
                "Sid": "allow-sns-messages",
                "Effect": "Allow",
                "Principal": "*",
                "Resource": "${aws_sqs_queue.booking.arn}",
                "Condition": {
                    "ArnEquals": {
                        "aws:SourceArn": [
                            "${var.payment_topic_arn}",
                            "${var.user_topic_arn}"
                        ]
                    }
                }
            }
        ]
    }
    POLICY
} 

resource "aws_sqs_queue_policy" "space" {
    queue_url = aws_sqs_queue.space.id

    policy = <<POLICY
    {
        "Version": "2012-10-17",
        "Id": "sqspolicy",
        "Statement": [
            {
                "Sid": "allow-sns-messages",
                "Effect": "Allow",
                "Principal": "*",
                "Resource": "${aws_sqs_queue.space.arn}",
                "Condition": {
                    "ArnEquals": {
                        "aws:SourceArn": [
                            "${var.space_topic_arn}"
                        ]
                    }
                }
            }
        ]
    }
    POLICY
} 


resource "aws_sqs_queue_policy" "payment" {
    queue_url = aws_sqs_queue.payment.id

    policy = <<POLICY
    {
        "Version": "2012-10-17",
        "Id": "sqspolicy",
        "Statement": [
            {
                "Sid": "allow-sns-messages",
                "Effect": "Allow",
                "Principal": "*",
                "Resource": "${aws_sqs_queue.payment.arn}",
                "Condition": {
                    "ArnEquals": {
                        "aws:SourceArn": [
                            "${var.payment_topic_arn}"
                        ]
                    }
                }
            }
        ]
    }
    POLICY
} 

# resource "aws_sqs_queue_policy" "user" {
#     queue_url = aws_sqs_queue.user.id

#     policy = <<POLICY
#     {
#         "Version": "2012-10-17",
#         "Id": "sqspolicy",
#         "Statement": [
#             {
#                 "Sid": "allow-sns-messages",
#                 "Effect": "Allow",
#                 "Principal": "*",
#                 "Resource": "${aws_sqs_queue.user.arn}",
#                 "Condition": {
#                     "ArnEquals": {
#                         "aws:SourceArn": [
#                             "${var.item_topic_arn}"
#                         ]
#                     }
#                 }
#             }
#         ]
#     }
#     POLICY
# } 