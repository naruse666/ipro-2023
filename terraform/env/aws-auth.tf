provider "kubernetes" {
  host                   = module.eks.cluster.endpoint
  cluster_ca_certificate = module.eks.cluster.ca
  token                  = module.eks.cluster.token
}


resource "kubernetes_config_map" "aws_auth" {
  metadata {
    name      = "aws-auth"
    namespace = "kube-system"
  }

  data = {
    "mapRoles" = <<EOT
- rolearn = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/*"
  username: system:node:{{EC2PrivateDNSName}}
  groups:
    - system:bootstrappers
    - system:nodes
EOT
  }
}
