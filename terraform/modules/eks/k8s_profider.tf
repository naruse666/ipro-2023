provider "helm" {
  kubernetes {
    host                   = aws_eks_cluster.cluster.endpoint
    cluster_ca_certificate = base64decode(aws_eks_cluster.cluster.certificate_authority[0].data)
    token                  = data.aws_eks_cluster_auth.cluster_auth.token
  }
}

provider "kubernetes" {
  host                   = aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(aws_eks_cluster.cluster.certificate_authority[0].data)
  token                  = data.aws_eks_cluster_auth.cluster_auth.token
}

resource "null_resource" "kubeconfig" {
  triggers = {
    cluster_name = aws_eks_cluster.cluster.name
  }
  provisioner "local-exec" {
    command = "aws eks update-kubeconfig --name ${aws_eks_cluster.cluster.name} --region ap-northeast-1"
  }
}

