// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack security group.
func LookupSecGroup(ctx *pulumi.Context, args *GetSecGroupArgs) (*GetSecGroupResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["region"] = args.Region
		inputs["secgroupId"] = args.SecgroupId
		inputs["tenantId"] = args.TenantId
	}
	outputs, err := ctx.Invoke("openstack:networking/getSecGroup:getSecGroup", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSecGroupResult{
		Region: outputs["region"],
		TenantId: outputs["tenantId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSecGroup.
type GetSecGroupArgs struct {
	// The name of the security group.
	Name interface{}
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve security groups ids. If omitted, the
	// `region` argument of the provider is used.
	Region interface{}
	// The ID of the security group.
	SecgroupId interface{}
	// The owner of the security group.
	TenantId interface{}
}

// A collection of values returned by getSecGroup.
type GetSecGroupResult struct {
	// See Argument Reference above.
	Region interface{}
	TenantId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
