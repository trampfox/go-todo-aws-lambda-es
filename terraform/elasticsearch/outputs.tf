output "subnet_id" {
  value = "${aws_subnet.elasticsearch_subnet.id}"
}

output "security_group" {
    value = "${aws_security_group.es_web.id}"
}

output "es_endpoint" {
  value = "${aws_elasticsearch_domain.es.endpoint}"
}
