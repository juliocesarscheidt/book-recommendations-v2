resource "aws_iam_role" "s3_role_api_gw" {
  name               = "s3-role-api-gw-${var.env}"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": [
          "ecs.amazonaws.com",
          "ecs-tasks.amazonaws.com"
        ]
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "s3_policy_api_gw" {
  name   = "s3-policy-api-gw-${var.env}"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::${local.api_gw_bucket_name}/book-recommendations-files/*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_role_policy_api_gw" {
  role       = aws_iam_role.s3_role_api_gw.name
  policy_arn = aws_iam_policy.s3_policy_api_gw.arn
  depends_on = [
    aws_iam_role.s3_role_api_gw,
    aws_iam_policy.s3_policy_api_gw,
  ]
}
