import khulnasoft
import khulnasoft_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[{
    "protocol": "tcp",
    "from_port": 0,
    "to_port": 0,
    "cidr_blocks": ["0.0.0.0/0"],
}])
# Get the ID for the latest Amazon Linux AMI.
ami = aws.get_ami(filters=[{
        "name": "name",
        "values": ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",
    tags={
        "Name": "web-server-www",
    },
    instance_type=aws.ec2.InstanceType.T2_MICRO,
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")
khulnasoft.export("publicIp", server.public_ip)
khulnasoft.export("publicHostName", server.public_dns)
