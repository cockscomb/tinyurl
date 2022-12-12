variable "dynamodb_endpoint" {
  default = "http://localhost:8000"
}

provider "aws" {
  region                      = "ap-northeast-1"
  skip_credentials_validation = true
  skip_requesting_account_id  = true
  skip_metadata_api_check     = true
  access_key                  = "dummy"
  secret_key                  = "dummy"
  endpoints {
    dynamodb = var.dynamodb_endpoint
  }
}

module "dynamodb" {
  source         = "../modules/dynamodb"
  url_table_name = "url"
}
