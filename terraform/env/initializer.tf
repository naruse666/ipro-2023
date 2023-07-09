resource "null_resource" "patch-coredns"{
  triggers = {
    endpoint = module.eks.cluster.endpoint
    ca_crt   = module.eks.cluster.ca
    token    = module.eks.cluster.token
  }

  provisioner "local-exec" {
    command = <<EOH
cat >/tmp/ca.crt <<EOF
${module.eks.cluster.ca}
EOF
kubectl patch deployment coredns \
  --server="${module.eks.cluster.endpoint}" \
  --certificate_authority=/tmp/ca.crt \
  --token="${module.eks.cluster.token}" \
  -n kube-system \
  --type json \
  -p='[{"op": "remove", "path": "/spec/template/metadata/annotations/eks.amazonaws.com~1compute-type"}]';

kubectl rollout restart -n kube-system deployment coredns \
  --server="${module.eks.cluster.endpoint}" \
  --certificate_authority=/tmp/ca.crt \
  --token="${module.eks.cluster.token}";
EOH
  }
}
