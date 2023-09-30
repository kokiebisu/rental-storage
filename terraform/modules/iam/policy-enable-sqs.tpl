{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "sqs:ReceiveMessage",
                "sqs:DeleteMessage",
                "sqs:GetQueueAttributes"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:sqs:${region}:${account_id}:*"
        }
    ]
}
