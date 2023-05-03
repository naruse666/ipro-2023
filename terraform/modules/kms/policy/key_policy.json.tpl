{
    Statement = [
      {
        Action = "${action}"
        Effect = "Allow"
        Principal = ${principal_map}

        Resource = "${resource}"
        Sid      = "${sid}"
      },
    ]
    Version = "2012-10-17"
}
