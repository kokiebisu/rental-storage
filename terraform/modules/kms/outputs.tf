output kms_terraform_bucket_key_arn {
    value = aws_kms_key.terraform_state.arn
}