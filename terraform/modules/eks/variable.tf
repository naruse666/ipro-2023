variable "cluster-name" {
  type = string
}

variable "k8s_version" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(any)
}

variable "kms_key_arn" {
  type = string
}
