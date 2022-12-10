variable "url_table_name" {
  default = "url"
}
resource "aws_dynamodb_table" "url" {
  name         = var.url_table_name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }
}
