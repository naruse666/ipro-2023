data "aws_route53_zone" "ipro2023" {
  name         = local.domain_name
  private_zone = false
}

resource "aws_acm_certificate" "certificate" {
  domain_name       = local.domain_name
  validation_method = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "cert_dns" {
  allow_overwrite = true
  name            = tolist(aws_acm_certificate.certificate.domain_validation_options)[0].resource_record_name
  records         = [tolist(aws_acm_certificate.certificate.domain_validation_options)[0].resource_record_value]
  type            = tolist(aws_acm_certificate.certificate.domain_validation_options)[0].resource_record_type
  zone_id         = data.aws_route53_zone.ipro2023.zone_id
  ttl             = 60
}

resource "aws_acm_certificate_validation" "cert_validate" {
  certificate_arn = aws_acm_certificate.certificate.arn
  #validation_record_fqdns = [for record in aws_route53_record.cert_dns: record.fqdn]
  validation_record_fqdns = [aws_route53_record.cert_dns.fqdn]
}
