data "http" "cluster_oidc" {
  url = "${aws_eks_cluster.cluster.identity.0.oidc.0.issuer}/.well-known/openid-configuration"
}
data "tls_certificate" "cluster_oidc" {
  url = jsondecode(data.http.cluster_oidc.response_body).jwks_uri
}
resource "aws_iam_openid_connect_provider" "cluster" {
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = [data.tls_certificate.cluster_oidc.certificates[0].sha1_fingerprint]
  url             = aws_eks_cluster.cluster.identity[0].oidc[0].issuer
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    effect  = "Allow"

    condition {
      test     = "StringEquals"
      variable = "${replace(aws_iam_openid_connect_provider.cluster.url, "https://", "")}:sub"
      values   = ["system:serviceaccount:kube-system:aws-node"]
    }

    principals {
      identifiers = [aws_iam_openid_connect_provider.cluster.arn]
      type        = "Federated"
    }
  }
}

resource "aws_iam_role" "oidc" {
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
  name               = "eks-odic-role"
}
