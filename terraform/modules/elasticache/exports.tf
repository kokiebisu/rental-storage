resource "aws_ssm_parameter" "elasticache_host" {
    name = "/terraform/elasticache/elasticache-host"
    type = "String"
    value = aws_elasticache_cluster.elasticache.cache_nodes.0.address
}

resource "aws_ssm_parameter" "elasticache_port" {
    name = "/terraform/elasticache/elasticache-port"
    type = "String"
    value = aws_elasticache_cluster.elasticache.cache_nodes.0.port 
}
