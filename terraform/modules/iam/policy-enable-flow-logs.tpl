{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "logs:CreateLogGroup",
            "Effect": "Allow",
            "Resource": "arn:aws:logs:*:*:log-group:flow-logs"
        },
        {
            "Action": "logs:CreateLogStream",
            "Effect": "Allow",
            "Resource": "arn:aws:logs:*:*:log-group:flow-logs:*"
        },
        {
            "Action": "logs:PutLogEvents",
            "Effect": "Allow",
            "Resource": "arn:aws:logs:*:*:log-group:flow-logs:log-stream:*"
        }
    ]
}