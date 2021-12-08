
data "aws_eks_cluster" "eks" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "eks" {
  name = module.eks.cluster_id
}

data "aws_vpcs" "myvpcs" {
}

output "myvpcs" {
  value = "${data.aws_vpcs.myvpcs.ids}"
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.eks.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.eks.certificate_authority[0].data)
  token                  = data.aws_eks_cluster_auth.eks.token
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"

  cluster_version = "1.21"
  cluster_name    = "test-cluster"
  vpc_id          = "vpc-09d70d4727e49dde6"
  subnets         = ["subnet-0213058ab09c17cca", "subnet-003c36f6f9c457619", "subnet-00a17755b8471dbd5"]

  worker_groups = [
    {
      instance_type = "m4.large"
      asg_max_size  = 2
    }
  ]
}