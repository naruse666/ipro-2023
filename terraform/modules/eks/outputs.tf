output "role" {
  value = {
    pod_execution_role_arn = aws_iam_role.pod_execution_role.arn
  }
}

output "cluster" {
  value = {
    name     = aws_eks_cluster.cluster.name
    endpoint = aws_eks_cluster.cluster.endpoint
    token    = data.aws_eks_cluster_auth.cluster_auth.token
    ca       = base64decode(aws_eks_cluster.cluster.certificate_authority[0].data)
  }
}
