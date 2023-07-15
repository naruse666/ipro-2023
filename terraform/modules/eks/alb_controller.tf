provider "helm" {
  kubernetes {
    host = aws_eks_cluster.cluster.endpoint
    cluster_ca_certificate = base64decode(aws_eks_cluster.cluster.certificate_authority[0].data)
    token = data.aws_eks_cluster_auth.cluster_auth.token
  }
}

provider "kubernetes" {
    host = aws_eks_cluster.cluster.endpoint
    cluster_ca_certificate = base64decode(aws_eks_cluster.cluster.certificate_authority[0].data)
    token = data.aws_eks_cluster_auth.cluster_auth.token
}

resource "null_resource" "kubeconfig" {
  triggers = {
    cluster_name = aws_eks_cluster.cluster.name
  }
  provisioner "local-exec" {
    command = "aws eks update-kubeconfig --name ${aws_eks_cluster.cluster.name} --region ap-northeast-1"
  }
}

data "aws_iam_policy_document" "alb_controller" {
  statement {
    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.cluster.arn]
    }
    actions = ["sts:AssumeRoleWithWebIdentity"]
    condition {
      test     = "StringEquals"
      variable = "${aws_iam_openid_connect_provider.cluster.url}:aud"
      values   = ["sts.amazonaws.com"]
    }
    condition {
      test     = "StringEquals"
      variable = "${aws_iam_openid_connect_provider.cluster.url}:sub"
      values   = ["system:serviceaccount:kube-system:aws-load-balancer-controller"]
    }
  }
}

resource "aws_iam_role" "alb_controller" {
  name               = "${var.cluster-name}-eks-alb-controller-role"
  assume_role_policy = data.aws_iam_policy_document.alb_controller.json
}

resource "aws_iam_policy" "alb_controller" {
  name   = "AWSLoadBalancerControllerIAMPolicy"
  policy = file("${path.module}/policy/alb_controller_policy.json")
}

resource "aws_iam_role_policy_attachment" "alb_controller" {
  policy_arn = aws_iam_policy.alb_controller.arn
  role       = aws_iam_role.alb_controller.name
}

resource "kubernetes_service_account" "aws_loadbalancer_controller" {
  metadata {
    name      = "aws-load-balancer-controller"
    namespace = "kube-system"
    annotations = {
      "eks.amazonaws.com/role-arn" = aws_iam_role.alb_controller.arn
    }
    labels = {
      "app.kubernetes.io/component"  = "controller"
      "app.kubernetes.io/name"       = "aws-load-balancer-controller"
      "app.kubernetes.io/managed-by" = "terraform"
    }
  }
}

resource "helm_release" "aws-load-balancer-controller" {
    name = "aws-load-balancer-controller"
    repository = "https://aws.github.io/eks-charts"
    chart = "aws-load-balancer-controller"
    version = "1.4.2"
    namespace = "kube-system"
    depends_on = [
      kubernetes_service_account.aws_loadbalancer_controller,
      null_resource.kubeconfig
    ]

  set {
    name  = "clusterName"
    value = aws_eks_cluster.cluster.name
  }

  set {
    name  = "vpcId"
    value = var.vpc_id
  }

  set {
    name  = "serviceAccount.create"
    value = false
  }

  set {
    name  = "serviceAccount.name"
    value = "aws-load-balancer-controller"
  }

  set {
    name  = "image.repository"
    value = "602401143452.dkr.ecr.ap-northeast-1.amazonaws.com/amazon/aws-load-balancer-controller"
  }

  set {
    name  = "image.tag"
    value = "v2.4.4"
  }
}
