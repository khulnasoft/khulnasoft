import khulnasoft
import khulnasoft_aws as aws

db_cluster = aws.rds.Cluster("dbCluster", master_password=khulnasoft.Output.secret("foobar"))
