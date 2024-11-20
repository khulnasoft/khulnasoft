package main

import (
	"encoding/json"
	"fmt"

	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/ec2"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/eks"
	"github.com/khulnasoft/khulnasoft-aws/sdk/v5/go/aws/iam"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// VPC
		eksVpc, err := ec2.NewVpc(ctx, "eksVpc", &ec2.VpcArgs{
			CidrBlock:          khulnasoft.String("10.100.0.0/16"),
			InstanceTenancy:    khulnasoft.String("default"),
			EnableDnsHostnames: khulnasoft.Bool(true),
			EnableDnsSupport:   khulnasoft.Bool(true),
			Tags: khulnasoft.StringMap{
				"Name": khulnasoft.String("khulnasoft-eks-vpc"),
			},
		})
		if err != nil {
			return err
		}
		eksIgw, err := ec2.NewInternetGateway(ctx, "eksIgw", &ec2.InternetGatewayArgs{
			VpcId: eksVpc.ID(),
			Tags: khulnasoft.StringMap{
				"Name": khulnasoft.String("khulnasoft-vpc-ig"),
			},
		})
		if err != nil {
			return err
		}
		eksRouteTable, err := ec2.NewRouteTable(ctx, "eksRouteTable", &ec2.RouteTableArgs{
			VpcId: eksVpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: khulnasoft.String("0.0.0.0/0"),
					GatewayId: eksIgw.ID(),
				},
			},
			Tags: khulnasoft.StringMap{
				"Name": khulnasoft.String("khulnasoft-vpc-rt"),
			},
		})
		if err != nil {
			return err
		}
		zones, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{}, nil)
		if err != nil {
			return err
		}
		var vpcSubnet []*ec2.Subnet
		for key0, val0 := range zones.Names {
			__res, err := ec2.NewSubnet(ctx, fmt.Sprintf("vpcSubnet-%v", key0), &ec2.SubnetArgs{
				AssignIpv6AddressOnCreation: khulnasoft.Bool(false),
				VpcId:                       eksVpc.ID(),
				MapPublicIpOnLaunch:         khulnasoft.Bool(true),
				CidrBlock:                   khulnasoft.Sprintf("10.100.%v.0/24", key0),
				AvailabilityZone:            khulnasoft.String(val0),
				Tags: khulnasoft.StringMap{
					"Name": khulnasoft.Sprintf("khulnasoft-sn-%v", val0),
				},
			})
			if err != nil {
				return err
			}
			vpcSubnet = append(vpcSubnet, __res)
		}
		var rta []*ec2.RouteTableAssociation
		for key0, _ := range zones.Names {
			__res, err := ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("rta-%v", key0), &ec2.RouteTableAssociationArgs{
				RouteTableId: eksRouteTable.ID(),
				SubnetId:     vpcSubnet[key0].ID(),
			})
			if err != nil {
				return err
			}
			rta = append(rta, __res)
		}
		var splat0 khulnasoft.StringArray
		for _, val0 := range vpcSubnet {
			splat0 = append(splat0, val0.ID())
		}
		subnetIds := splat0
		eksSecurityGroup, err := ec2.NewSecurityGroup(ctx, "eksSecurityGroup", &ec2.SecurityGroupArgs{
			VpcId:       eksVpc.ID(),
			Description: khulnasoft.String("Allow all HTTP(s) traffic to EKS Cluster"),
			Tags: khulnasoft.StringMap{
				"Name": khulnasoft.String("khulnasoft-cluster-sg"),
			},
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					CidrBlocks: khulnasoft.StringArray{
						khulnasoft.String("0.0.0.0/0"),
					},
					FromPort:    khulnasoft.Int(443),
					ToPort:      khulnasoft.Int(443),
					Protocol:    khulnasoft.String("tcp"),
					Description: khulnasoft.String("Allow pods to communicate with the cluster API Server."),
				},
				&ec2.SecurityGroupIngressArgs{
					CidrBlocks: khulnasoft.StringArray{
						khulnasoft.String("0.0.0.0/0"),
					},
					FromPort:    khulnasoft.Int(80),
					ToPort:      khulnasoft.Int(80),
					Protocol:    khulnasoft.String("tcp"),
					Description: khulnasoft.String("Allow internet access to pods"),
				},
			},
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Principal": map[string]interface{}{
						"Service": "eks.amazonaws.com",
					},
					"Effect": "Allow",
					"Sid":    "",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		eksRole, err := iam.NewRole(ctx, "eksRole", &iam.RoleArgs{
			AssumeRolePolicy: khulnasoft.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "servicePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      eksRole.ID(),
			PolicyArn: khulnasoft.String("arn:aws:iam::aws:policy/AmazonEKSServicePolicy"),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "clusterPolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      eksRole.ID(),
			PolicyArn: khulnasoft.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
		})
		if err != nil {
			return err
		}
		tmpJSON1, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Principal": map[string]interface{}{
						"Service": "ec2.amazonaws.com",
					},
					"Effect": "Allow",
					"Sid":    "",
				},
			},
		})
		if err != nil {
			return err
		}
		json1 := string(tmpJSON1)
		ec2Role, err := iam.NewRole(ctx, "ec2Role", &iam.RoleArgs{
			AssumeRolePolicy: khulnasoft.String(json1),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "workerNodePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      ec2Role.ID(),
			PolicyArn: khulnasoft.String("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "cniPolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      ec2Role.ID(),
			PolicyArn: khulnasoft.String("arn:aws:iam::aws:policy/AmazonEKSCNIPolicy"),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "registryPolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      ec2Role.ID(),
			PolicyArn: khulnasoft.String("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"),
		})
		if err != nil {
			return err
		}
		eksCluster, err := eks.NewCluster(ctx, "eksCluster", &eks.ClusterArgs{
			RoleArn: eksRole.Arn,
			Tags: khulnasoft.StringMap{
				"Name": khulnasoft.String("khulnasoft-eks-cluster"),
			},
			VpcConfig: &eks.ClusterVpcConfigArgs{
				PublicAccessCidrs: khulnasoft.StringArray{
					khulnasoft.String("0.0.0.0/0"),
				},
				SecurityGroupIds: khulnasoft.StringArray{
					eksSecurityGroup.ID(),
				},
				SubnetIds: subnetIds,
			},
		})
		if err != nil {
			return err
		}
		_, err = eks.NewNodeGroup(ctx, "nodeGroup", &eks.NodeGroupArgs{
			ClusterName:   eksCluster.Name,
			NodeGroupName: khulnasoft.String("khulnasoft-eks-nodegroup"),
			NodeRoleArn:   ec2Role.Arn,
			SubnetIds:     subnetIds,
			Tags: khulnasoft.StringMap{
				"Name": khulnasoft.String("khulnasoft-cluster-nodeGroup"),
			},
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: khulnasoft.Int(2),
				MaxSize:     khulnasoft.Int(2),
				MinSize:     khulnasoft.Int(1),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("clusterName", eksCluster.Name)
		ctx.Export("kubeconfig", khulnasoft.All(eksCluster.Endpoint, eksCluster.CertificateAuthority, eksCluster.Name).ApplyT(func(_args []interface{}) (string, error) {
			endpoint := _args[0].(string)
			certificateAuthority := _args[1].(eks.ClusterCertificateAuthority)
			name := _args[2].(string)
			var _zero string
			tmpJSON2, err := json.Marshal(map[string]interface{}{
				"apiVersion": "v1",
				"clusters": []map[string]interface{}{
					map[string]interface{}{
						"cluster": map[string]interface{}{
							"server":                     endpoint,
							"certificate-authority-data": certificateAuthority.Data,
						},
						"name": "kubernetes",
					},
				},
				"contexts": []map[string]interface{}{
					map[string]interface{}{
						"contest": map[string]interface{}{
							"cluster": "kubernetes",
							"user":    "aws",
						},
					},
				},
				"current-context": "aws",
				"kind":            "Config",
				"users": []map[string]interface{}{
					map[string]interface{}{
						"name": "aws",
						"user": map[string]interface{}{
							"exec": map[string]interface{}{
								"apiVersion": "client.authentication.k8s.io/v1alpha1",
								"command":    "aws-iam-authenticator",
							},
							"args": []string{
								"token",
								"-i",
								name,
							},
						},
					},
				},
			})
			if err != nil {
				return _zero, err
			}
			json2 := string(tmpJSON2)
			return json2, nil
		}).(khulnasoft.StringOutput))
		return nil
	})
}
