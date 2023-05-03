module "eks" {
  source = "../../modules/eks"
  cluster-name = "ipro-2023-eks-cluster"
  k8s_version = "1.25"

  vpc_id = data.terraform_remote_state.vpc.outputs.vpc.id
  subnet_ids = data.terraform_remote_state.vpc.outputs.subnets.private_subnets

  security_group_ids = []
  
  kms_key_arn = ""
}
