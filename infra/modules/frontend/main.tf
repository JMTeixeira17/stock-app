resource "aws_s3_bucket" "frontend" {
  bucket = var.bucket_name
  acl    = "public-read"
}

resource "aws_s3_bucket_website_configuration" "frontend" {
  bucket = aws_s3_bucket.frontend.id

  index_document { suffix = "index.html" }
}
