// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 router interface resource within OpenStack.
type RouterInterface struct {
	s *pulumi.ResourceState
}

// NewRouterInterface registers a new resource with the given unique name, arguments, and options.
func NewRouterInterface(ctx *pulumi.Context,
	name string, args *RouterInterfaceArgs, opts ...pulumi.ResourceOpt) (*RouterInterface, error) {
	if args == nil || args.RouterId == nil {
		return nil, errors.New("missing required argument 'RouterId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["portId"] = nil
		inputs["region"] = nil
		inputs["routerId"] = nil
		inputs["subnetId"] = nil
	} else {
		inputs["portId"] = args.PortId
		inputs["region"] = args.Region
		inputs["routerId"] = args.RouterId
		inputs["subnetId"] = args.SubnetId
	}
	s, err := ctx.RegisterResource("openstack:networking/routerInterface:RouterInterface", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouterInterface{s: s}, nil
}

// GetRouterInterface gets an existing RouterInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterInterface(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouterInterfaceState, opts ...pulumi.ResourceOpt) (*RouterInterface, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["portId"] = state.PortId
		inputs["region"] = state.Region
		inputs["routerId"] = state.RouterId
		inputs["subnetId"] = state.SubnetId
	}
	s, err := ctx.ReadResource("openstack:networking/routerInterface:RouterInterface", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouterInterface{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RouterInterface) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RouterInterface) ID() *pulumi.IDOutput {
	return r.s.ID
}

// ID of the port this interface connects to. Changing
// this creates a new router interface.
func (r *RouterInterface) PortId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["portId"])
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to create a router. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// router interface.
func (r *RouterInterface) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// ID of the router this interface belongs to. Changing
// this creates a new router interface.
func (r *RouterInterface) RouterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["routerId"])
}

// ID of the subnet this interface connects to. Changing
// this creates a new router interface.
func (r *RouterInterface) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// Input properties used for looking up and filtering RouterInterface resources.
type RouterInterfaceState struct {
	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortId interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router interface.
	Region interface{}
	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterId interface{}
	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetId interface{}
}

// The set of arguments for constructing a RouterInterface resource.
type RouterInterfaceArgs struct {
	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortId interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router interface.
	Region interface{}
	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterId interface{}
	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetId interface{}
}
