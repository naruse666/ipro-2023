{
    Statement = [
      {
        Action = "${action}"
        Effect = "Allow"
        Principal = {
          "AWS" = ${principal}"
          }

        Resource = "${resource}"
        Sid      = "${sid}"
      },
    ]
    Version = "2012-10-17"
}
