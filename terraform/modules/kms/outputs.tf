output "key" {
  value = {
    arn = aws_kms_key.key.arn
    id  = aws_kms_key.key.key_id
  }
}
