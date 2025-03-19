output "backend_ip" {
  value = module.backend.instance_ip
}

output "frontend_url" {
  value = "http://${module.frontend.bucket_name}.s3-website-us-east-1.amazonaws.com"
}
