package main

import (
	"fmt"
	"github.com/pulumi/pulumi-awsx/sdk/go/awsx/ec2"
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	_ "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		// Create a new VPC, subnets, and associated infrastructure
		vpcNetworkCidr := "10.0.0.0/16"
		eksVpc, err := ec2.NewVpc(ctx, "eks-vpc", &ec2.VpcArgs{
			EnableDnsHostnames: pulumi.Bool(true),
			CidrBlock:          &vpcNetworkCidr,
		})
		if err != nil {
			return err
		}

		// Create a new EKS cluster
		eksCluster, err := eks.NewCluster(ctx, "eks-cluster", &eks.ClusterArgs{
			VpcId:                        eksVpc.VpcId,
			PublicSubnetIds:              eksVpc.PublicSubnetIds,
			PrivateSubnetIds:             eksVpc.PrivateSubnetIds,
			InstanceType:                 pulumi.String("t2.medium"),
			DesiredCapacity:              pulumi.Int(3),
			MinSize:                      pulumi.Int(3),
			MaxSize:                      pulumi.Int(6),
			NodeAssociatePublicIpAddress: pulumi.BoolRef(false),
		})
		if err != nil {
			return err
		}

		// Export some values in case they are needed elsewhere
		ctx.Export("kubeconfig", eksCluster.Kubeconfig)
		ctx.Export("vpcId", eksVpc.VpcId)
		return nil
	})
	if err != nil {
		fmt.Printf("Boot failed: %v", err)
	}
}
