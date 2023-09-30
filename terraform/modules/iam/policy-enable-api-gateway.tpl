{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "execute-api:Invoke", 
                "execute-api:ManageConnections"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:execute-api:*:*:*"
        }
    ]
}
