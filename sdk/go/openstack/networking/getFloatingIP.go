// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack floating IP.
func LookupFloatingIP(ctx *pulumi.Context, args *GetFloatingIPArgs) (*GetFloatingIPResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["address"] = args.Address
		inputs["fixedIp"] = args.FixedIp
		inputs["pool"] = args.Pool
		inputs["portId"] = args.PortId
		inputs["region"] = args.Region
		inputs["status"] = args.Status
		inputs["tenantId"] = args.TenantId
	}
	outputs, err := ctx.Invoke("openstack:networking/getFloatingIP:getFloatingIP", inputs)
	if err != nil {
		return nil, err
	}
	return &GetFloatingIPResult{
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getFloatingIP.
type GetFloatingIPArgs struct {
	// The IP address of the floating IP.
	Address interface{}
	// The specific IP address of the internal port which should be associated with the floating IP.
	FixedIp interface{}
	// The name of the pool from which the floating IP belongs to.
	Pool interface{}
	// The ID of the port the floating IP is attached.
	PortId interface{}
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve floating IP ids. If omitted, the
	// `region` argument of the provider is used.
	Region interface{}
	Status interface{}
	// The owner of the floating IP.
	TenantId interface{}
}

// A collection of values returned by getFloatingIP.
type GetFloatingIPResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}