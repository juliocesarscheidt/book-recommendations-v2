# Instructions on creating an AWS User to be used on CI to interact with CodeCommit

```bash

cat <<EOF > tmp_policy_code_commit.json
{
  "Version": "2012-10-17",
  "Statement": [
      {
          "Effect": "Allow",
          "Action": [
              "codecommit:CreateBranch",
              "codecommit:ListBranches",
              "codecommit:GitPull",
              "codecommit:ListPullRequests",
              "codecommit:GetReferences",
              "codecommit:GetFile",
              "codecommit:GetCommit",
              "codecommit:GetComment",
              "codecommit:UpdateComment",
              "codecommit:GetCommitHistory",
              "codecommit:GetRepository",
              "codecommit:GetBranch",
              "codecommit:GitPush",
              "codecommit:GetPullRequest"
          ],
          "Resource": "arn:aws:codecommit:${AWS_DEFAULT_REGION}:${AWS_ACCOUNT}:book-recommendations-v2"
      }
  ]
}
EOF

aws iam create-policy --policy-name "book-recommendations-repo-git-pull-push-policy" --policy-document file://tmp_policy_code_commit.json

# "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-repo-git-pull-push-policy"


# create an user to interact with the codecommit repo
aws iam create-user --user-name "book-recommendations-repo-user"

aws iam get-user --user-name "book-recommendations-repo-user"
{
    "User": {
        "Path": "/",
        "UserName": "book-recommendations-repo-user",
        "UserId": "AAAAAAAAAAAAAAAAAAAAA",
        "Arn": "arn:aws:iam::${AWS_ACCOUNT}:user/book-recommendations-repo-user",
        "CreateDate": "2022-05-01T18:41:51+00:00"
    }
}
export CODE_COMMIT_USER_ID=$(aws iam get-user --user-name "book-recommendations-repo-user" | jq -r '.User.UserId')
echo "${CODE_COMMIT_USER_ID}"



aws iam attach-user-policy \
  --user-name "book-recommendations-repo-user" \
  --policy-arn "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-repo-git-pull-push-policy"

aws iam list-attached-user-policies --user-name "book-recommendations-repo-user"
{
    "AttachedPolicies": [
        {
            "PolicyName": "book-recommendations-repo-git-pull-push-policy",
            "PolicyArn": "arn:aws:iam::${AWS_ACCOUNT}:policy/book-recommendations-repo-git-pull-push-policy"
        }
    ]
}


# create a SSH key for usage with codecommit
ssh-keygen -t rsa -b 2048 -f "book-recommendations"
chmod 600 "book-recommendations"
chmod 600 "book-recommendations.pub"


CODE_COMMIT_SSH_KEY_ID=$(aws iam upload-ssh-public-key \
  --user-name "book-recommendations-repo-user" \
  --ssh-public-key-body "file://${PWD}/book-recommendations.pub" | jq -r '.SSHPublicKey.SSHPublicKeyId')
echo "${CODE_COMMIT_SSH_KEY_ID}"


# add codecommit config for usage with git
mkdir -p ~/.ssh

cat > ~/.ssh/config <<EOF
Host git-codecommit.*.amazonaws.com
User ${CODE_COMMIT_SSH_KEY_ID}
IdentityFile ${PWD}/book-recommendations
EOF

chmod 600 ~/.ssh/config
cat ~/.ssh/config


ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts
ssh-keyscan -t rsa git-codecommit.${AWS_REGION}.amazonaws.com >> ~/.ssh/known_hosts


# create a repo on CodeCommit called book-recommendations-v2
aws codecommit create-repository \
  --repository-name "book-recommendations-v2" --region ${AWS_DEFAULT_REGION}

git remote add codecommit "ssh://git-codecommit.${AWS_DEFAULT_REGION}.amazonaws.com/v1/repos/book-recommendations-v2"
git push -f codecommit main


git clone \
  "ssh://git-codecommit.${AWS_DEFAULT_REGION}.amazonaws.com/v1/repos/book-recommendations-v2" \
  "book-recommendations-v2-code-commit"

```
