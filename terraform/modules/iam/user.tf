resource "aws_iam_user" "terraform" {
  name = "terraform"
}

resource "aws_iam_access_key" "terraform" {
  user = aws_iam_user.terraform.name
}

resource "aws_iam_policy_attachment" "test-attach" {
  name = "terraform-policy-attachment"
  users      = [aws_iam_user.terraform.name]
  policy_arn = aws_iam_policy.enable_s3.arn
}
