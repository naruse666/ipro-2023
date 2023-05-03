data "terraform_remote_state" "vpc" {
  backend = "s3"

  config = {
    bucket = "naruse-tf-backend"
    key    = "ipro-2023/terraform/env/vpc.tfstate"
    region = "ap-northeast-1"
  }
}
