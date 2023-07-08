output "name" {
  value = aws_eks_cluster.cluster.name
}

output "role" {
  value = {
    pod_execution_role_arn = aws_iam_role.pod_execution_role.arn
  }
}
