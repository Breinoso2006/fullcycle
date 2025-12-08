module "new-vpc" {
  source = "./modules/vpc"
  prefix = var.prefix
  vpc_cidr_block = var.vpc_cidr_block
}

module "eks" {
  source = "./modules/eks"
  vpc_id            = module.new-vpc.vpc_id
  prefix = var.prefix
  node_desired_size = var.node_desired_size
  node_max_size     = var.node_max_size
  node_min_size     = var.node_min_size
  subnet_ids        = module.new-vpc.subnet_ids
  retention_days = var.retention_days
}