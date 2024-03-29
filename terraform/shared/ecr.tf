locals{
  ecr_repositories = toset([
    aws_ecr_repository.front.name,
    aws_ecr_repository.suicide.name
  ])
  }

data "aws_iam_policy_document" "ecr_policy_document" {
  statement {
    sid    = "ECR policy"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["${data.aws_caller_identity.current.account_id}"]
    }

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:DescribeRepositories",
      "ecr:GetRepositoryPolicy",
      "ecr:ListImages",
      "ecr:DeleteRepository",
      "ecr:BatchDeleteImage",
      "ecr:SetRepositoryPolicy",
      "ecr:DeleteRepositoryPolicy",
    ]
  }
}

resource "aws_ecr_lifecycle_policy" "ecr_lifecycle" {
  for_each = local.ecr_repositories

  repository = each.value

  policy = <<EOF
{
  "rules": [
    {
      "rulePriority": 10,
      "description": "Keep images with the tag dev ",
      "selection": {
        "tagStatus": "tagged",
        "tagPrefixList": [
          "dev"
        ],
        "countType": "sinceImagePushed",
        "countUnit": "days",
        "countNumber": 36500
      },
      "action": {
        "type": "expire"
      }
    },
    {
      "rulePriority": 20,
      "description": "Expire other tags",
      "selection": {
        "tagStatus": "any",
        "countType": "sinceImagePushed",
        "countUnit": "days",
        "countNumber": 3
      },
      "action": {
        "type": "expire"
      }
    }
  ]
}
EOF
}

resource "aws_ecr_repository" "suicide" {
  name                 = "ipro2023/suicide"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  force_delete = true
}

resource "aws_ecr_repository" "front" {
  name                 = "ipro2023/front"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  force_delete = true
}

resource "aws_ecr_repository_policy" "ecr_policy" {
  for_each = local.ecr_repositories
  repository = each.value
  policy     = data.aws_iam_policy_document.ecr_policy_document.json
}
