resource "aws_iam_role" "eks_load_balancer_controller_role" {
  name               = "AmazonEKSLoadBalancerControllerRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
  managed_policy_arns = [aws_iam_policy.eks_iam_policy.arn]
}
 
resource "aws_iam_policy" "eks_iam_policy" {
  name   = "AWSLoadBalancerControllerIAMPolicy"
  policy = file("./iam_policy.json")
}
 
data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    sid     = "EKSClusterAssumeRole"
    actions = ["sts:AssumeRoleWithWebIdentity"]
 
    principals {
      type        = "Federated"
      identifiers = ["${module.eks.oidc_provider_arn}"]
    }
 
    condition {
      test     = "StringEquals"
      variable = "${module.eks.oidc_provider}:sub"
      values = [
        "system:serviceaccount:kube-system:aws-load-balancer-controller"
      ]
    }
 
    condition {
      test     = "StringEquals"
      variable = "${module.eks.oidc_provider}:aud"
      values = [
        "sts.amazonaws.com"
      ]
    }
  }
}