resource "aws_kms_key" "key" {
  deletion_window_in_days = 7
}

resource "aws_kms_key_policy" "key_policy" {
  key_id = aws_kms_key.key.id
  policy = templatefile("${path.module}/policy/key_policy.json.tpl", {
    sid           = var.sid
    action        = var.key_policy_action,
    resource      = var.key_policy_resource,
    principal_map = var.key_policy_principal
  })
}
