{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "ssm:GetParameter",
                "ssm:GetParameters",
                "ssm:GetParametersByPath",
                "ssm:DescribeParameters",
                "ssm:PutParameter"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:ssm:${region}:${account_id}:*"
        }
    ]
}
