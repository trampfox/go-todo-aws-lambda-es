provider "aws" {
  region = "${var.region}"
}

resource "aws_vpc" "elasticsearch" {
  cidr_block = "${var.vpc_cidr}"
  enable_dns_hostnames = true

  tags {
    Name = "elasticsearch-vpc"
  }
}

resource "aws_security_group" "es_web" {
  name        = "elasticsearch-sg"
  description = "Allow incoming HTTPS connections & icmp"

  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = -1
    to_port = -1
    protocol = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port = 65535
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  vpc_id = "${aws_vpc.elasticsearch.id}"

  tags {
    Name = "elasticsearch-web-sg"
  }
  
}

resource "aws_iam_service_linked_role" "es" {
  aws_service_name = "es.amazonaws.com"
}

resource "aws_subnet" "elasticsearch_subnet" {
  vpc_id     = "${aws_vpc.elasticsearch.id}"
  cidr_block = "${var.subnet_cidr}"
  availability_zone = "${var.zone}"

  tags {
    Name = "elasticsearch-subnet"
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = "${aws_vpc.elasticsearch.id}"

  tags {
    Name = "main-igw"
  }
}

resource "aws_route_table" "public_rt" {
  vpc_id = "${aws_vpc.elasticsearch.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.gw.id}"
  }

  tags {
    Name = "elasticsearch-public-rt"
  }
}

resource "aws_elasticsearch_domain" "es" {
  domain_name           = "${var.name}"
  elasticsearch_version = "${var.elasticsearch_version}"

  encrypt_at_rest {
    enabled    = "${var.encryption_enabled}"
  }

  cluster_config {
    instance_type            = "${var.instance_type}"
    instance_count           = "${var.instance_number}"
    dedicated_master_enabled = "${var.dedicated_master}"
    dedicated_master_type    = "${var.master_instance_type}"
    dedicated_master_count   = "${var.master_number}"
  }

  access_policies = <<CONFIG
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "es:*",
      "Resource": "arn:aws:es:${var.region}:${var.account_id}:domain/todo-app-logs/*"
    }
  ]
}
  CONFIG

  vpc_options {
    security_group_ids = ["${aws_security_group.es_web.id}"]
    subnet_ids         = ["${aws_subnet.elasticsearch_subnet.id}"]
  }

  ebs_options {
    ebs_enabled = true
    volume_type = "${var.volume_type}"
    volume_size = "${var.volume_size}"
  }

  tags {
    Domain = "${var.name}"
  }
}
