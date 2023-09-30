{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "sns:CreateTopic", 
                "sns:Publish"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:sns:${region}:${account_id}:*"
        }
    ]
}
