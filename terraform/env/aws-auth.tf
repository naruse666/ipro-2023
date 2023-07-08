provider "kubernetes" {
  host                   = module.eks.cluster.endpoint
  cluster_ca_certificate = module.eks.cluster.ca
  token                  = module.eks.cluster.token
}


resource "kubernetes_config_map" "aws_auth" {
  metadata {
    name      = "aws-auth"
    namespace = "kube-system"
  }

  data = {
    "mapUsers" = <<EOT
- userarn = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/cli-user"
  username: admin
  groups:
    - system:bootstrappers
    - system:nodes
    - eks-console-dashboard-full-access-binding"
- userarn = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/naruse"
  username: admin
  groups:
    - system:bootstrappers
    - system:nodes
    - eks-console-dashboard-full-access-binding"
EOT
  }
}

# Role binding
resource "kubernetes_role_binding" "aws-auth" {
  metadata {
    name      = "eks-console-dashboard-restricted-access-role-binding"
    namespace = "default"
  }

  role_ref {
    kind      = "Role"
    api_group = "rbac.authorization.k8s.io"
    name      = "eks-console-dashboard-restricted-access-role"
  }

  subject {
    kind      = "Group"
    api_group = "rbac.authorization.k8s.io"
    name      = "eks-console-dashboard-restricted-access-group"
  }
}

# Cluster Role Binding
resource "kubernetes_cluster_role_binding" "aws-auth" {
  metadata {
    name = "eks-console-dashboard-full-access-binding"
  }

  subject {
    kind      = "Group"
    api_group = "rbac.authorization.k8s.io"
    name      = "eks-console-dashboard-full-access-group"
  }

  role_ref {
    kind      = "ClusterRole"
    api_group = "rbac.authorization.k8s.io"
    name      = "eks-console-dashboard-full-access-clusterrole"
  }
}

# Validation webhook
resource "kubernetes_validating_webhook_configuration_v1" "aws-auth" {
  metadata {
    name = "eks-aws-auth-configmap-validation-webhook"
  }

  webhook {
    name = "eks-aws-auth-configmap-validation-webhook.amazonaws.com"

    admission_review_versions = ["v1"]

    client_config {
      ca_bundle = module.eks.cluster.ca
      url = "https://127.0.0.1:21375/validate"
    }

    rule {
      api_groups   = [""]
      api_versions = ["v1"]
      operations   = ["UPDATE"]
      resources    = ["configmaps"]
      scope        = "*"
    }

    namespace_selector  {
      match_labels = {"kubernetes.io/metadata.name" = "kube-system"}
    }

    side_effects = "None"
  }
}