module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = local.name

  cidr = local.cidr

  azs             = local.azs
  private_subnets = local.private_subnets
  public_subnets  = local.public_subnets

  enable_nat_gateway = true
  enable_vpn_gateway = true
  single_nat_gateway = true

  enable_dns_hostnames = true
  enable_dns_support   = true

  private_subnet_tags = {
    "kubernetes.io/role/internal-elb" = "1"
  }

  public_subnet_tags = {
    "kubernetes.io/role/elb" = "1"
  }
}
