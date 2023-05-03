resource "aws_eks_cluster" "cluster" {
  name     = var.cluster-name
  version  = var.version
  role_arn = aws_iam_role.cluster.arn

  vpc_config {
    subnet_ids              = var.subnet_ids
    endpoint_private_access = false
    endpoint_public_access  = true
    security_group_ids      = var.security_group_ids
  }

  encryption_config {
    provider {
      key_arn = var.kms_key_arn
    }
    resources = "secrets"
  }

  # Ensure that IAM Role permissions are created before and deleted after EKS Cluster handling.
  # Otherwise, EKS will not be able to properly delete EKS managed EC2 infrastructure such as Security Groups.
  depends_on = [
    aws_iam_role_policy_attachment.AmazonEKSClusterPolicy,
    aws_iam_role_policy_attachment.AmazonEKSVPCResourceController,
  ]
}
