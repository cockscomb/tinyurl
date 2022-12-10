provider "aws" {
  region     = "ap-northeast-1"
  endpoints {
    dynamodb = "http://localhost:8000"
  }
}

module "dynamodb" {
  source         = "../modules/dynamodb"
  url_table_name = "url"
}
