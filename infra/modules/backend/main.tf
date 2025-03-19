resource "aws_instance" "backend" {
  ami           = "ami-0c55b159cbfafe1f0" # Ubuntu 22.04
  instance_type = "t2.micro"
  subnet_id     = var.subnet_id
  security_groups = [var.security_group_id]
  key_name      = var.key_name

  user_data = <<-EOF
    #!/bin/bash
    sudo apt update -y
    sudo apt install docker.io -y
    sudo systemctl start docker
    sudo systemctl enable docker
    docker run -d -p 8080:8080 mi-imagen-backend
  EOF

  tags = { Name = "backend-server" }
}
