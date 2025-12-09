variable "prefix" {
  description = "Prefixo para os recursos"
  type        = string
  default     = "fullcycle"
}

variable "retention_days" {
  description = "Número de dias para retenção dos logs"
  type        = number
  default     = 7
}

variable "node_desired_size" {
  description = "Número desejado de nós no cluster"
  type        = number
  default     = 1
}

variable "node_max_size" {
  description = "Número máximo de nós no cluster"
  type        = number
  default     = 2
}

variable "node_min_size" {
  description = "Número mínimo de nós no cluster"
  type        = number
  default     = 1
}

variable "vpc_cidr_block" {
  description = "Bloco CIDR para a VPC"
  type        = string
  default     = "10.0.0.0/16" 
}