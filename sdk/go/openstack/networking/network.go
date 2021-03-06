// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 Neutron network resource within OpenStack.
type Network struct {
	s *pulumi.ResourceState
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOpt) (*Network, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adminStateUp"] = nil
		inputs["availabilityZoneHints"] = nil
		inputs["external"] = nil
		inputs["name"] = nil
		inputs["region"] = nil
		inputs["segments"] = nil
		inputs["shared"] = nil
		inputs["tenantId"] = nil
		inputs["valueSpecs"] = nil
	} else {
		inputs["adminStateUp"] = args.AdminStateUp
		inputs["availabilityZoneHints"] = args.AvailabilityZoneHints
		inputs["external"] = args.External
		inputs["name"] = args.Name
		inputs["region"] = args.Region
		inputs["segments"] = args.Segments
		inputs["shared"] = args.Shared
		inputs["tenantId"] = args.TenantId
		inputs["valueSpecs"] = args.ValueSpecs
	}
	s, err := ctx.RegisterResource("openstack:networking/network:Network", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Network{s: s}, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NetworkState, opts ...pulumi.ResourceOpt) (*Network, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adminStateUp"] = state.AdminStateUp
		inputs["availabilityZoneHints"] = state.AvailabilityZoneHints
		inputs["external"] = state.External
		inputs["name"] = state.Name
		inputs["region"] = state.Region
		inputs["segments"] = state.Segments
		inputs["shared"] = state.Shared
		inputs["tenantId"] = state.TenantId
		inputs["valueSpecs"] = state.ValueSpecs
	}
	s, err := ctx.ReadResource("openstack:networking/network:Network", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Network{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Network) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Network) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The administrative state of the network.
// Acceptable values are "true" and "false". Changing this value updates the
// state of the existing network.
func (r *Network) AdminStateUp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["adminStateUp"])
}

// An availability zone is used to make
// network resources highly available. Used for resources with high availability
// so that they are scheduled on different availability zones. Changing this
// creates a new network.
func (r *Network) AvailabilityZoneHints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["availabilityZoneHints"])
}

// Specifies whether the network resource has the
// external routing facility. Valid values are true and false. Defaults to
// false. Changing this updates the external attribute of the existing network.
func (r *Network) External() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["external"])
}

// The name of the network. Changing this updates the name of
// the existing network.
func (r *Network) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron network. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// network.
func (r *Network) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// An array of one or more provider segment objects.
func (r *Network) Segments() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["segments"])
}

// Specifies whether the network resource can be accessed
// by any tenant or not. Changing this updates the sharing capabalities of the
// existing network.
func (r *Network) Shared() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["shared"])
}

// The owner of the network. Required if admin wants to
// create a network for another tenant. Changing this creates a new network.
func (r *Network) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// Map of additional options.
func (r *Network) ValueSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["valueSpecs"])
}

// Input properties used for looking up and filtering Network resources.
type NetworkState struct {
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp interface{}
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints interface{}
	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External interface{}
	// The name of the network. Changing this updates the name of
	// the existing network.
	Name interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// network.
	Region interface{}
	// An array of one or more provider segment objects.
	Segments interface{}
	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabalities of the
	// existing network.
	Shared interface{}
	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp interface{}
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints interface{}
	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External interface{}
	// The name of the network. Changing this updates the name of
	// the existing network.
	Name interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// network.
	Region interface{}
	// An array of one or more provider segment objects.
	Segments interface{}
	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabalities of the
	// existing network.
	Shared interface{}
	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}
