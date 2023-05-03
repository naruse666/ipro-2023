resource "aws_security_group" "cluster_shared" {
  name   = "${var.cluster-name}-cluster-shared"
  vpc_id = var.vpc_id

  tags = {
    Name = "${var.cluster-name}-cluster-shared"
  }
}

resource "aws_security_group_rule" "default_cluster_to_node" {
  security_group_id        = aws_security_group.cluster_shared.id
  source_security_group_id = aws_eks_cluster.cluster.vpc_config[0].cluster_security_group_id

  type      = "ingress"
  from_port = 0
  to_port   = 65535
  protocol  = "-1"
}

resource "aws_security_group_rule" "inter_node" {
  security_group_id        = aws_security_group.cluster_shared.id
  source_security_group_id = aws_security_group.cluster_shared.id

  type      = "ingress"
  from_port = 0
  to_port   = 65535
  protocol  = "-1"
}

resource "aws_security_group_rule" "node_to_default_cluster" {
  security_group_id        = aws_eks_cluster.cluster.vpc_config[0].cluster_security_group_id
  source_security_group_id = aws_security_group.cluster_shared.id

  type      = "ingress"
  from_port = 0
  to_port   = 65535
  protocol  = "-1"
}
