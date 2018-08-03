// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack network.
func LookupNetwork(ctx *pulumi.Context, args *GetNetworkArgs) (*GetNetworkResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["external"] = args.External
		inputs["matchingSubnetCidr"] = args.MatchingSubnetCidr
		inputs["name"] = args.Name
		inputs["networkId"] = args.NetworkId
		inputs["region"] = args.Region
		inputs["status"] = args.Status
		inputs["tenantId"] = args.TenantId
	}
	outputs, err := ctx.Invoke("openstack:networking/getNetwork:getNetwork", inputs)
	if err != nil {
		return nil, err
	}
	return &GetNetworkResult{
		AdminStateUp: outputs["adminStateUp"],
		AvailabilityZoneHints: outputs["availabilityZoneHints"],
		Region: outputs["region"],
		Shared: outputs["shared"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getNetwork.
type GetNetworkArgs struct {
	// The external routing facility of the network.
	External interface{}
	// The CIDR of a subnet within the network.
	MatchingSubnetCidr interface{}
	// The name of the network.
	Name interface{}
	// The ID of the network.
	NetworkId interface{}
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve networks ids. If omitted, the
	// `region` argument of the provider is used.
	Region interface{}
	// The status of the network.
	Status interface{}
	// The owner of the network.
	TenantId interface{}
}

// A collection of values returned by getNetwork.
type GetNetworkResult struct {
	// (Optional) The administrative state of the network.
	AdminStateUp interface{}
	// (Optional) The availability zone candidates for the network.
	AvailabilityZoneHints interface{}
	// See Argument Reference above.
	Region interface{}
	// (Optional)  Specifies whether the network resource can be accessed
	// by any tenant or not.
	Shared interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
