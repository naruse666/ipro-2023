output "vpc" {
  value = {
    id = module.vpc.vpc_id
  }
}

output "subnets" {
  value = {
    private_subnets      = module.vpc.private_subnets
    private_subnets_cidr = module.vpc.private_subnets_cidr_blocks

    public_subnets      = module.vpc.public_subnets
    public_subnets_cidr = module.vpc.public_subnets_cidr_blocks
  }
}
