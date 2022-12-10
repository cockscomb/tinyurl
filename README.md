# tinyurl

## Local development

```bash
$ docker compose up --build
```

### Initialize DynamoDB

```bash
$ terraform -chdir=terraform/local init
$ terraform -chdir=terraform/local apply
```
