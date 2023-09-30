{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "kinesis:PutRecord", 
                "kinesis:PutRecords"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:kinesis:${region}:${account_id}:stream/*"
        },
        {
            "Action": [
                "kinesis:Get*", 
                "kinesis:List*", 
                "kinesis:Describe*"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:kinesis:${region}:${account_id}:stream/*"
        }
    ]
}