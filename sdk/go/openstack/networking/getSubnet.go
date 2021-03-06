// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack subnet.
func LookupSubnet(ctx *pulumi.Context, args *GetSubnetArgs) (*GetSubnetResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["cidr"] = args.Cidr
		inputs["dhcpDisabled"] = args.DhcpDisabled
		inputs["dhcpEnabled"] = args.DhcpEnabled
		inputs["gatewayIp"] = args.GatewayIp
		inputs["ipVersion"] = args.IpVersion
		inputs["ipv6AddressMode"] = args.Ipv6AddressMode
		inputs["ipv6RaMode"] = args.Ipv6RaMode
		inputs["name"] = args.Name
		inputs["networkId"] = args.NetworkId
		inputs["region"] = args.Region
		inputs["subnetId"] = args.SubnetId
		inputs["subnetpoolId"] = args.SubnetpoolId
		inputs["tenantId"] = args.TenantId
	}
	outputs, err := ctx.Invoke("openstack:networking/getSubnet:getSubnet", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSubnetResult{
		AllocationPools: outputs["allocationPools"],
		Cidr: outputs["cidr"],
		DnsNameservers: outputs["dnsNameservers"],
		EnableDhcp: outputs["enableDhcp"],
		GatewayIp: outputs["gatewayIp"],
		HostRoutes: outputs["hostRoutes"],
		IpVersion: outputs["ipVersion"],
		Ipv6AddressMode: outputs["ipv6AddressMode"],
		Ipv6RaMode: outputs["ipv6RaMode"],
		Name: outputs["name"],
		NetworkId: outputs["networkId"],
		Region: outputs["region"],
		SubnetId: outputs["subnetId"],
		SubnetpoolId: outputs["subnetpoolId"],
		TenantId: outputs["tenantId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSubnet.
type GetSubnetArgs struct {
	// The CIDR of the subnet.
	Cidr interface{}
	// If the subnet has DHCP disabled.
	DhcpDisabled interface{}
	// If the subnet has DHCP enabled.
	DhcpEnabled interface{}
	// The IP of the subnet's gateway.
	GatewayIp interface{}
	// The IP version of the subnet (either 4 or 6).
	IpVersion interface{}
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode interface{}
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode interface{}
	// The name of the subnet.
	Name interface{}
	// The ID of the network the subnet belongs to.
	NetworkId interface{}
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve subnet ids. If omitted, the
	// `region` argument of the provider is used.
	Region interface{}
	// The ID of the subnet.
	SubnetId interface{}
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId interface{}
	// The owner of the subnet.
	TenantId interface{}
}

// A collection of values returned by getSubnet.
type GetSubnetResult struct {
	// Allocation pools of the subnet.
	AllocationPools interface{}
	Cidr interface{}
	// DNS Nameservers of the subnet.
	DnsNameservers interface{}
	// Whether the subnet has DHCP enabled or not.
	EnableDhcp interface{}
	GatewayIp interface{}
	// Host Routes of the subnet.
	HostRoutes interface{}
	IpVersion interface{}
	Ipv6AddressMode interface{}
	Ipv6RaMode interface{}
	Name interface{}
	NetworkId interface{}
	// See Argument Reference above.
	Region interface{}
	SubnetId interface{}
	SubnetpoolId interface{}
	TenantId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
