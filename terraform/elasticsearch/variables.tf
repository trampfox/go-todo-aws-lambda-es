variable "name" {
  default     = "todo-app-logs-es"
  description = "Elastic Search Service cluster name."
  type        = "string"
}

variable "instance_type" {
  default     = "t2.micro.elasticsearch"
  description = "Elastic Search Service cluster Ec2 instance type."
  type        = "string"
}

variable "instance_number" {
  default     = 1
  description = "Elastic Search Service cluster Ec2 instance number."
  type        = "string"
}

variable "dedicated_master" {
  default     = false
  description = "Dedicated master nodes enabled/disabled."
  type        = "string"
}

variable "master_instance_type" {
  default     = 0
  description = "Elastic Search Service cluster dedicated master Ec2 instance type."
  type        = "string"
}
variable "master_number" {
  default     = 0
  description = "Elastic Search Service cluster dedicated master Ec2 instance number."
  type        = "string"
}

variable "volume_type" {
  default     = "gp2"
  description = "Default type of the EBS volumes."
  type        = "string"
}

variable "volume_size" {
  default     = "10"
  description = "Default size of the EBS volumes."
  type        = "string"
}

variable "elasticsearch_version" {
  default     = "2.3"
  description = "Elastic Search Service cluster version number. t2.micro free tier supports only 2.3 and 1.5 version."
  type        = "string"
}

variable "encryption_enabled" {
  default     = "false"
  description = "Enable encription in Elastic Search. Free tier doesn't support encryption at rest."
  type        = "string"
}

variable "region" {
  default     = "eu-west-1"
  description = "Indicates cluster region."
  type        = "string"
}

variable "zone" {
  default     = "eu-west-1a"
  description = "Indicates cluster zone."
  type        = "string"
}

variable "account_id" {
  description = "AWS account id."
  type        = "string"
}

variable "vpc_cidr" {
  description = "CIDR for the VPC"
  default = "10.0.0.0/16"
}

variable "subnet_cidr" {
  description = "CIDR for the private subnet"
  default = "10.0.1.0/24"
}
