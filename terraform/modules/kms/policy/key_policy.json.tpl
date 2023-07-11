{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "${action}",
      "Effect": "Allow",
      "Principal": {
        "AWS": "${principal}"
      },

      "Resource": "${resource}",
      "Sid": "${sid}"
    },
    {
      "Sid": "Allow use by github actioins",
      "Effect": "Allow",
      "Principal": { 
        "AWS": "${github_actions_role}" 
       },
      "Action": [
      "kms:DescribeKey",
      "kms:Get*",
      "kms:ListResourceTags"
      ], 
      "Resource": "*"
    }
  ]
}
