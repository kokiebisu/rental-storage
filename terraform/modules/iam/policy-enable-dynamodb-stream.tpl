{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "dynamodb:ListStreams"
            ],
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Action": [
                "dynamodb:GetRecords",
                "dynamodb:DescribeStream",
                "dynamodb:GetShardIterator"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:dynamodb:${region}:${account_id}:table/*/stream/*"
        }
    ]
}
