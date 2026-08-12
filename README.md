# AWS CloudWatch Log Alerting Pipeline

A production-style **AWS centralized logging and alerting pipeline** built with **Golang, CloudWatch, Lambda, SNS, Terraform, and GitHub Actions**.

This project demonstrates how to build an **end-to-end observability workflow** that detects application failures and delivers alerts within seconds. The entire infrastructure is provisioned using **Terraform (Infrastructure as Code)** and can be deployed automatically through **GitHub Actions CI/CD**.

---

## 🎯 Project Objectives

The primary goal of this project is to simulate a **real-world production monitoring scenario** and demonstrate modern DevOps practices.

### Key objectives

- 🔍 **Centralized log collection** from an EC2-hosted application
- ⚠️ **Automatic detection of ERROR events** using CloudWatch Subscription Filters
- 🚀 **Event-driven alert processing** with AWS Lambda
- 📧 **Real-time notifications** through Amazon SNS
- 🔐 **Least-privilege IAM design** for secure cloud operations
- 🏗️ **Infrastructure as Code** using Terraform
- 🔄 **CI/CD automation** with GitHub Actions
- ☁️ **Cloud-native observability thinking** from log generation to alert delivery

This repository is intended for:

- DevOps Engineers
- Cloud Engineers
- SRE practitioners
- Engineers learning AWS monitoring and alerting patterns
- Portfolio and interview demonstration projects

---

## 🏗️ Architecture

```text
GitHub Push
    │
    ▼
GitHub Actions
    ├── Terraform Init
    ├── Terraform Validate
    ├── Terraform Plan
    └── Terraform Apply
            │
            ▼
      AWS Infrastructure
            │
            ▼
      EC2 Instance
            │
            ▼
      Golang Gin Application
            │
            ├── stdout
            └── app.log
                    │
                    ▼
         CloudWatch Agent
                    │
                    ▼
         CloudWatch Log Group
                    │
                    ▼
      Subscription Filter
          pattern: ERROR
                    │
                    ▼
         AWS Lambda (Python)
                    │
                    ▼
            Amazon SNS
                    │
                    ▼
           Email Alert
```

---

## ✨ Features

- 🌐 `GET /` → normal INFO traffic
- ⚠️ `GET /warn` → WARNING log generation
- ❌ `GET /error` → ERROR log generation that triggers the alert pipeline
- 📄 Dual logging to **stdout** and **app.log**
- 📡 Real-time log shipping with **CloudWatch Agent**
- 🔔 Automatic Lambda invocation on matching log patterns
- 📧 Email notification through **Amazon SNS**
- 🏗️ Fully automated AWS infrastructure with **Terraform**
- 🔄 Optional **GitHub Actions Terraform CI/CD pipeline**
- 🔐 Support for **GitHub OIDC + AWS IAM Role** authentication (no static AWS keys required)

### Current Terraform Scope

At the current stage of the project, Terraform provisions only the **CloudWatch Log Group** used by the application:

* 📄 **Amazon CloudWatch Log Group** (`/aws/ec2/alerting-app`)
* ⏳ **Log retention policy** (7 days)
* 📤 **Terraform output** exposing the log group name

Planned future Terraform modules will add:

* 🌐 VPC networking
* 🖥️ EC2 instance
* 🔐 Security groups
* 👤 IAM roles and policies
* 📡 CloudWatch Subscription Filter
* ⚡ AWS Lambda function
* 🔔 Amazon SNS topic and email subscription


---

## 🧰 Tech Stack

| Layer | Technology |
|------|-------------|
| Application | Golang 1.25 + Gin |
| Container Runtime | Docker / Docker Compose |
| Compute | Amazon EC2 |
| Log Collection | Amazon CloudWatch Agent |
| Log Storage | Amazon CloudWatch Logs |
| Event Processing | AWS Lambda (Python) |
| Notifications | Amazon SNS |
| Security | AWS IAM |
| IaC | Terraform |
| CI/CD | GitHub Actions |
| Authentication | GitHub OIDC + IAM Role |

---

## 📁 Project Structure

```text
aws-cloudwatch-alerting/
├── app/
│   ├── main.go                 # Gin application
│   ├── go.mod
│   ├── go.sum
│   ├── Dockerfile
│   └── app.log                 # Local log file
│
├── terraform/
│   ├── main.tf                 # Provider and backend
│   ├── variables.tf            # Input variables
│   ├── outputs.tf              # Terraform outputs
│   ├── ec2.tf                  # EC2 instance and security group
│   ├── iam.tf                  # IAM roles and policies
│   ├── cloudwatch.tf           # Log group and subscription filter
│   ├── lambda.tf               # Lambda function and permissions
│   └── sns.tf                  # SNS topic and email subscription
│
├── lambda/
│   └── handler.py              # Lambda alert processor
│
├── .github/
│   └── workflows/
│       └── terraform.yml       # Terraform CI/CD pipeline
│
├── compose.yaml
└── README.md
```

---

## 🚀 Local Development

### Prerequisites

- Docker Desktop
- WSL2 (recommended on Windows)
- Docker Compose

### Run the application

```bash
docker compose up --build
```

The API will be available at:

```text
http://localhost:8080
```

---

## 🧪 API Testing

### INFO endpoint

```bash
curl http://localhost:8080/
```

### WARNING endpoint

```bash
curl http://localhost:8080/warn
```

### ERROR endpoint

```bash
curl http://localhost:8080/error
```

Check the generated logs:

```bash
cat app/app.log
```

Example output:

```text
2026/08/13 21:10:01 INFO: normal request received
2026/08/13 21:10:05 WARNING: suspicious activity detected
2026/08/13 21:10:09 ERROR: application failure occurred
```

---

## ☁️ AWS Deployment

### 1. Clone the repository

```bash
git clone https://github.com/ArashMHD91/aws-cloudwatch-alerting.git
cd aws-cloudwatch-alerting
```

### 2. Initialize Terraform

```bash
cd terraform
terraform init
```

### 3. Validate the configuration

```bash
terraform fmt -check
terraform validate
```

### 4. Review the execution plan

```bash
terraform plan -out=tfplan
```

### 5. Apply the infrastructure

```bash
terraform apply tfplan
```


## 🔐 GitHub Actions + OIDC Authentication

This project supports **GitHub OpenID Connect (OIDC)** authentication, which is the recommended modern approach for GitHub Actions on AWS.

### Why OIDC?

✅ No long-lived AWS access keys stored in GitHub
✅ Temporary credentials issued by AWS STS
✅ Repository and branch-level access control
✅ Enterprise-grade security model

### Required AWS resources

#### Create the OIDC provider

```bash
aws iam create-open-id-connect-provider \
  --url https://token.actions.githubusercontent.com \
  --client-id-list sts.amazonaws.com
```

#### Create the IAM role

Trust policy example:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::123456789012:oidc-provider/token.actions.githubusercontent.com"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
        },
        "StringLike": {
          "token.actions.githubusercontent.com:sub": "repo:ArashMHD91/aws-cloudwatch-alerting:*"
        }
      }
    }
  ]
}
```

---

## 🔄 GitHub Actions Workflow

```yaml
name: Terraform

on:
  push:
    branches: [main]
    paths:
      - 'terraform/**'

permissions:
  id-token: write
  contents: read

jobs:
  terraform:
    runs-on: ubuntu-latest

    defaults:
      run:
        working-directory: terraform

    steps:
      - uses: actions/checkout@v4

      - uses: hashicorp/setup-terraform@v3

      - uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: arn:aws:iam::123456789012:role/github-actions-terraform
          aws-region: eu-west-2

      - run: terraform init
      - run: terraform fmt -check
      - run: terraform validate
      - run: terraform plan -out=tfplan
      - run: terraform apply -auto-approve tfplan
```

---

## 📬 How the Alert Pipeline Works

### Step 1 — Application generates an ERROR log

```text
ERROR: application failure occurred
```

### Step 2 — CloudWatch Agent ships the log

The agent continuously monitors `app.log` and forwards new entries to the configured **CloudWatch Log Group**.

### Step 3 — Subscription Filter matches the pattern

```hcl
pattern = "ERROR"
```

### Step 4 — Lambda processes the event

CloudWatch Logs invokes the Lambda function and sends the log payload as a **base64 + gzip encoded event**.

### Step 5 — SNS sends the notification

The Lambda function publishes a formatted message to **Amazon SNS**, which delivers an email alert to the subscribed address.

---

## 📊 Example Alert Email

```text
Subject: AWS CloudWatch Alert

ERROR detected in application logs

Log Group: /aws/ec2/alerting-app
Message: ERROR: application failure occurred
Timestamp: 2026-08-13T21:10:09Z
```

---

## 🛡️ Security Considerations

This project follows several AWS security best practices:

- 🔐 **Least-privilege IAM policies**
- 🚫 **No hardcoded AWS credentials**
- ⏳ **Temporary STS credentials via OIDC**
- 📦 Isolated Lambda execution role
- 📡 Restricted CloudWatch Logs invocation permissions
- 🌍 Region-specific CloudWatch Logs principals

---

## 📚 Key Learning Outcomes

During this project I explored:

- How **CloudWatch Subscription Filters** trigger Lambda functions
- Why Lambda permissions require a **region-specific CloudWatch Logs principal** (`logs.<region>.amazonaws.com`)
- Designing **least-privilege IAM roles** for EC2, Lambda, and CloudWatch
- Building **event-driven AWS architectures**
- Implementing **Infrastructure CI/CD** with GitHub Actions
- Using **GitHub OIDC federation with AWS STS** for secure automation
- Thinking about **observability as a complete workflow** rather than isolated services

---

## 🧹 Cleanup

To remove all AWS resources and avoid unnecessary charges:

```bash
cd terraform
terraform destroy
```

Confirm with:

```text
yes
```

---

## 🔮 Future Improvements

- [ ] Add **structured JSON logging**
- [ ] Send alerts to **Slack or Microsoft Teams**
- [ ] Store Terraform remote state in **S3 + DynamoDB locking**
- [ ] Add **Prometheus and Grafana dashboards**
- [ ] Package the application as a **multi-stage Docker image**
- [ ] Deploy the container to **Amazon ECS Fargate**
- [ ] Add **automated integration tests** in GitHub Actions

---

## 👨‍💻 Author

**Ali Fateh**

- LinkedIn: linkedin.com/in/mohammad-ali-fatehchehr-47425589/

---

## ⭐ Support

If this project helped you learn something about **AWS, Terraform, CloudWatch, or DevOps observability**, consider giving it a **⭐ star** on GitHub. It helps others discover the repository and supports continued open-source learning and experimentation.
