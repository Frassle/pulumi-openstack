// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpnaas

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 Neutron Endpoint Group resource within OpenStack.
type EndpointGroup struct {
	s *pulumi.ResourceState
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOpt) (*EndpointGroup, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["endpoints"] = nil
		inputs["name"] = nil
		inputs["region"] = nil
		inputs["tenantId"] = nil
		inputs["type"] = nil
		inputs["valueSpecs"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["endpoints"] = args.Endpoints
		inputs["name"] = args.Name
		inputs["region"] = args.Region
		inputs["tenantId"] = args.TenantId
		inputs["type"] = args.Type
		inputs["valueSpecs"] = args.ValueSpecs
	}
	s, err := ctx.RegisterResource("openstack:vpnaas/endpointGroup:EndpointGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EndpointGroup{s: s}, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EndpointGroupState, opts ...pulumi.ResourceOpt) (*EndpointGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["endpoints"] = state.Endpoints
		inputs["name"] = state.Name
		inputs["region"] = state.Region
		inputs["tenantId"] = state.TenantId
		inputs["type"] = state.Type
		inputs["valueSpecs"] = state.ValueSpecs
	}
	s, err := ctx.ReadResource("openstack:vpnaas/endpointGroup:EndpointGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EndpointGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EndpointGroup) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EndpointGroup) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The human-readable description for the group.
// Changing this updates the description of the existing group.
func (r *EndpointGroup) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
// Changing this creates a new group.
func (r *EndpointGroup) Endpoints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["endpoints"])
}

// The name of the group. Changing this updates the name of
// the existing group.
func (r *EndpointGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an endpoint group. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// group.
func (r *EndpointGroup) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The owner of the group. Required if admin wants to
// create an endpoint group for another project. Changing this creates a new group.
func (r *EndpointGroup) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
// Changing this creates a new group.
func (r *EndpointGroup) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Map of additional options.
func (r *EndpointGroup) ValueSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["valueSpecs"])
}

// Input properties used for looking up and filtering EndpointGroup resources.
type EndpointGroupState struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description interface{}
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints interface{}
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region interface{}
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId interface{}
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type interface{}
	// Map of additional options.
	ValueSpecs interface{}
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description interface{}
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints interface{}
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region interface{}
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId interface{}
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type interface{}
	// Map of additional options.
	ValueSpecs interface{}
}