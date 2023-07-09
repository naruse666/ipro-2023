- rolearn: "${pod-role}"
  username: "system:node:{{SessionName}}"
  groups:
    - system:bootstrappers
    - system:nodes
    - system:node-proxier
