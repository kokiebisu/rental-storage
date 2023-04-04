resource "aws_elasticache_cluster" "elasticache" {
  cluster_id           = "elasticache-cluster"
  engine               = "redis"
  node_type            = "cache.t2.micro"
  num_cache_nodes      = 1
  parameter_group_name = "default.redis7"
  port                 = 6379
  security_group_ids   = [var.elasticache_security_group_id]
  subnet_group_name    = var.elasticache_subnet_group_name
  engine_version       = "7.0"
}
