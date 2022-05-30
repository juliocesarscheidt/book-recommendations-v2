resource "aws_s3_bucket" "api_gw_s3_bucket" {
  bucket = local.api_gw_bucket_name
  acl    = "private"
  # with versioning enabled
  versioning {
    enabled = true
  }
  lifecycle_rule {
    id      = "old-files"
    enabled = true
    prefix  = "book-recommendations-files/"
    # move to one zone infrequent access
    transition {
      days          = 30
      storage_class = "ONEZONE_IA"
    }
    # remove older versions of objects
    noncurrent_version_expiration {
      days = 90
    }
  }
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256" # SSE-S3
      }
    }
  }
  tags = {
    Name = local.api_gw_bucket_name
  }
}

resource "aws_s3_bucket_policy" "api_gw_s3_bucket_policy" {
  bucket = aws_s3_bucket.api_gw_s3_bucket.id
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
        {
            "Sid": "DenyIncorrectEncryptionHeader",
            "Effect": "Deny",
            "Principal": "*",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::${local.api_gw_bucket_name}/*",
            "Condition": {
                "StringNotEquals": {
                    "s3:x-amz-server-side-encryption": "AES256"
                }
            }
        },
        {
            "Sid": "DenyUnencryptedObjectUploads",
            "Effect": "Deny",
            "Principal": "*",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::${local.api_gw_bucket_name}/*",
            "Condition": {
                "Null": {
                    "s3:x-amz-server-side-encryption": "true"
                }
            }
        }
    ]
}
EOF

  depends_on = [
    aws_s3_bucket.api_gw_s3_bucket,
  ]
}
