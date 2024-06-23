locals {
  name = "study-record-go-api"
  role = "study-record"
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "${local.name}-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["${local.region}a", "${local.region}c"]
  public_subnets  = ["10.0.0.0/24", "10.0.1.0/24"]
  private_subnets = ["10.0.100.0/24", "10.0.101.0/24"]

  enable_nat_gateway = true

  public_subnet_tags = {
    "kubernetes.io/role/elb" = "1"
  }

  private_subnet_tags = {
    "kubernetes.io/role/internal-elb" = "1"
  }
}


resource "aws_ecr_repository" "repository" {
  name                 = local.name
  image_tag_mutability = "MUTABLE"

  force_delete = true

  image_scanning_configuration {
    scan_on_push = true
  }
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_version = "1.28"
  cluster_name    = "${local.name}-eks"
  vpc_id          = module.vpc.vpc_id
  subnet_ids      = module.vpc.private_subnets
  enable_irsa     = true
  eks_managed_node_groups = {
    main = {
      min_size       = 1
      max_size       = 2
      desired_size   = 1
      instance_types = ["t2.nano"]
    }
  }
}

resource "aws_security_group_rule" "allow_ingress_port" {
  security_group_id        = module.eks.node_security_group_id
  type                     = "ingress"
  from_port                = 9443
  to_port                  = 9443
  protocol                 = "tcp"
  source_security_group_id = module.eks.cluster_security_group_id
}
