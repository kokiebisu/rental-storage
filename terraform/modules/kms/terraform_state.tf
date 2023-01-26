resource "aws_kms_key" "backend" {
  description             = "Creates the KMS key used to encrypt bucket objects"
  deletion_window_in_days = 10
  enable_key_rotation     = true
}

resource "aws_kms_alias" "backend" {
  name          = "alias/terraform-bucket-key"
  target_key_id = aws_kms_key.backend.key_id
}