module "eks" {
  source       = "../modules/eks"
  cluster-name = local.name
  k8s_version  = local.k8s_version

  vpc_id     = module.vpc.vpc_id
  subnet_ids = concat(module.vpc.private_subnets, module.vpc.public_subnets)

  kms_key_arn = module.cluster_encription_key.key.arn
}

module "cluster_encription_key" {
  source               = "../modules/kms"
  sid                  = "EKS Cluster Encription Key"
  key_policy_action    = "kms:*"
  key_policy_resource  = "*"
  key_policy_principal = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
}

resource "aws_eks_fargate_profile" "kube-system_eks_node" {
  cluster_name           = module.eks.cluster.name
  fargate_profile_name   = "${local.name}-kube-system-fargate-profile"
  pod_execution_role_arn = module.eks.role.pod_execution_role_arn
  subnet_ids             = module.vpc.private_subnets

  selector {
    namespace = "kube-system"
  }
}

resource "aws_eks_fargate_profile" "app_eks_node" {
  cluster_name           = module.eks.cluster.name
  fargate_profile_name   = "${local.name}-app-fargate-profile"
  pod_execution_role_arn = module.eks.role.pod_execution_role_arn
  subnet_ids             = module.vpc.private_subnets

  selector {
    namespace = "app"
  }
}
