resource "aws_kms_key" "terraform_state" {
  description             = "Creates the KMS key used to encrypt bucket objects"
  deletion_window_in_days = 10
  enable_key_rotation     = true
}

resource "aws_kms_alias" "terraform_state" {
  name          = "alias/terraform-bucket-key"
  target_key_id = aws_kms_key.terraform_state.key_id
}