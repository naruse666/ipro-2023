module "eks" {
  source       = "../modules/eks"
  cluster-name = "${local.name}-eks-cluster"
  k8s_version  = local.k8s_version

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  security_group_ids = []

  kms_key_arn = module.cluster_encription_key.key.arn
}

module "cluster_encription_key" {
  source               = "../modules/kms"
  sid                  = "EKS Cluster Encription Key"
  key_policy_action    = "kms:*"
  key_policy_resource  = "*"
  key_policy_principal = data.aws_caller_identity.current.arn
}
