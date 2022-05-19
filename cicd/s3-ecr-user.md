# Instructions on creating an AWS User to be used on CI to interact with S3 and ECR

```bash

cat <<EOF > tmp_policy_s3_user_s3_policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": ["s3:GetObject"],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::${S3_BUCKET_NAME}/book-recommendations/*"
    }
  ]
}
EOF

aws iam create-policy \
  --policy-name "book-recommendations-s3-get-file" \
  --policy-document file://tmp_policy_s3_user_s3_policy.json

# "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-s3-get-file"


cat <<EOF > tmp_policy_s3_user_ecr_policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ECRPullPolicy",
      "Effect": "Allow",
      "Action": [
        "ecr:CreateRepository",
        "ecr:BatchCheckLayerAvailability",
        "ecr:GetDownloadUrlForLayer",
        "ecr:BatchGetImage",
        "ecr:UploadLayerPart",
        "ecr:PutImage",
        "ecr:InitiateLayerUpload",
        "ecr:CompleteLayerUpload",
        "ecr:DescribeRepositories"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "ECRAuthPolicy",
      "Effect": "Allow",
      "Action": [
        "ecr:GetAuthorizationToken"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
EOF

aws iam create-policy \
  --policy-name "book-recommendations-s3-ecr-cicd" \
  --policy-document file://tmp_policy_s3_user_ecr_policy.json

# "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-s3-ecr-cicd"


# create an user to interact with S3 and ECR
aws iam create-user --user-name "book-recommendations-s3-user"

aws iam get-user --user-name "book-recommendations-s3-user"
{
    "User": {
        "Path": "/",
        "UserName": "book-recommendations-s3-user",
        "UserId": "AAAAAAAAAAAAAAAAAAAAA",
        "Arn": "arn:aws:iam::${AWS_ACCOUNT}:user/book-recommendations-s3-user",
        "CreateDate": "2022-05-01T18:41:51+00:00"
    }
}
export S3_USER_ID=$(aws iam get-user --user-name "book-recommendations-s3-user" | jq -r '.User.UserId')
echo "${S3_USER_ID}"



# create an access key to use on CI
aws iam create-access-key --user-name "book-recommendations-s3-user"
{
    "AccessKey": {
        "UserName": "book-recommendations-s3-user",
        "AccessKeyId": "",
        "Status": "Active",
        "SecretAccessKey": "",
        "CreateDate": "2022-05-01T19:55:17+00:00"
    }
}


# attach policies to user
aws iam attach-user-policy \
  --user-name "book-recommendations-s3-user" \
  --policy-arn "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-s3-get-file"

aws iam attach-user-policy \
  --user-name "book-recommendations-s3-user" \
  --policy-arn "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-s3-ecr-cicd"


aws iam list-attached-user-policies --user-name "book-recommendations-s3-user"
{
    "AttachedPolicies": [
        {
            "PolicyName": "book-recommendations-s3-get-file",
            "PolicyArn": "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-s3-get-file"
        },
        {
            "PolicyName": "book-recommendations-s3-ecr-cicd",
            "PolicyArn": "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-s3-ecr-cicd"
        }
    ]
}

```
