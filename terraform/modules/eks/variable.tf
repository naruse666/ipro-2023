variable "cluster-name" {
  type = String
}

variable "version" {
  type = String
}

variable "subnet_ids" {
  type = List
}

variable "security_group_ids" {
  type = List
}

variable "kms_key_arn" {
  type = String
}
