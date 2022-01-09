resource "aws_vpc" "main" {
  #  cidr_block         = "from resource ${cidr_block}"
  cidr_block         = var.cidr_block
  enable_dns_support = var.enable_dns
  ipv6_cidr_block    = var.ma_super_variable
  instance_tenancy = local.super_var_2
}
variable "enable_dns" {
  default     = "3"
  description = "My awesome variable"
}

variable "cidr_block" {
  description = "my cidr block"
}

variable "ma_super_variable" {
  description = "my cidr block"
}
locals {
  super_var_2 = "add_spice_2"
}