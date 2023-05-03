data "aws_caller_identity" "current" {}

module "eks" {
  source       = "../../modules/eks"
  cluster-name = "ipro-2023-eks-cluster"
  k8s_version  = "1.25"

  vpc_id     = data.terraform_remote_state.vpc.outputs.vpc.id
  subnet_ids = data.terraform_remote_state.vpc.outputs.subnets.private_subnets

  security_group_ids = []

  kms_key_arn = ""
}

module "cluster_encription_key" {
  source               = "../../modules/kms"
  sid                  = "EKS Cluster Encription Key"
  key_policy_action    = "kms:*"
  key_policy_resource  = "*"
  key_policy_principal = data.aws_caller_identity.current.arn
}
