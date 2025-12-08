output "vpc_id" {
  description = "The ID of the VPC"
  value       = aws_vpc.vpc-fullcycle.id
}

output "subnet_ids" {
  description = "List of public subnet IDs"
  value       = aws_subnet.subnets[*].id
}