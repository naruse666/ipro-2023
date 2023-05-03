variable "sid" {
  type = string
}

variable "key_policy_action" {
  type = string
}

variable "key_policy_resource" {
  type = string
}

variable "key_policy_principal" {
  type = map(any)
}
