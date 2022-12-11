variable "dynamodb_endpoint" {
  default = "http://localhost:8000"
}

provider "aws" {
  region = "ap-northeast-1"
  endpoints {
    dynamodb = var.dynamodb_endpoint
  }
}

module "dynamodb" {
  source         = "../modules/dynamodb"
  url_table_name = "url"
}
