provider "kubernetes" {
  host                   = module.eks.cluster.endpoint
  cluster_ca_certificate = module.eks.cluster.ca
  token                  = module.eks.cluster.token
}

resource "kubernetes_config_map_v1_data" "aws_auth" {
  force = true

  metadata {
    name      = "aws-auth"
    namespace = "kube-system"
  }

  data = {
    "mapRoles" = templatefile("${path.module}/configmaps/map-roles.yaml.tpl", {
      pod-role = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/ipro-2023-pod-execution-role"
    })

    "mapUsers" = templatefile("${path.module}/configmaps/map-users.yaml.tpl", {
      naruse = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/naruse",
      admin  = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/cli-user"
    })

    "mapAccounts" = templatefile("${path.module}/configmaps/map-account.yaml.tpl", {
      account_id = "${data.aws_caller_identity.current.account_id}"
    })
  }
}

# Cluster Role
resource "kubernetes_cluster_role" "viewer" {
  metadata {
    name = "ViewerClusterRole"
  }

  rule {
    api_groups = [""]
    resources  = ["nodes", "services", "namespaces", "pods", "endpoints", "persistentvolumeclaims", "persistentvolumes"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = ["apps"]
    resources  = ["deployments", "replicasets", "statefulsets", "daemonsets"]
    verbs      = ["get", "list", "watch"]
  }
}

# Cluster Role Binding
resource "kubernetes_cluster_role_binding" "viewer" {
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
    name      = "ViewerClusterRole"
  }
}
