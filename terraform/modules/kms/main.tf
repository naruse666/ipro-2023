data "aws_iam_role" "github_actions" {
  name = "AllowAccessFromIpro2023"
}
resource "aws_kms_key" "key" {
  deletion_window_in_days = 7
}

resource "aws_kms_key_policy" "key_policy" {
  key_id = aws_kms_key.key.id
  policy = templatefile("${path.module}/policy/key_policy.json.tpl", {
    sid                 = var.sid
    action              = var.key_policy_action,
    resource            = var.key_policy_resource,
    principal           = var.key_policy_principal,
    github_actions_role = data.aws_iam_role.github_actions.arn
  })
}
