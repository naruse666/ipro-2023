# Generated by Terragrunt. Sig: nIlQXj57tbuaRZEa
terraform {
  backend "s3" {
    bucket         = "naruse-tf-backend"
    dynamodb_table = "naruse-tf-backend"
    encrypt        = true
    key            = "ipro-2023/terraform/env/vpc.tfstate"
    region         = "ap-northeast-1"
  }
}
