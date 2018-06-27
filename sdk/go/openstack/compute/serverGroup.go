// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 Server Group resource within OpenStack.
type ServerGroup struct {
	s *pulumi.ResourceState
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOpt) (*ServerGroup, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["policies"] = nil
		inputs["region"] = nil
		inputs["valueSpecs"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["policies"] = args.Policies
		inputs["region"] = args.Region
		inputs["valueSpecs"] = args.ValueSpecs
	}
	inputs["members"] = nil
	s, err := ctx.RegisterResource("openstack:compute/serverGroup:ServerGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServerGroup{s: s}, nil
}

// GetServerGroup gets an existing ServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServerGroupState, opts ...pulumi.ResourceOpt) (*ServerGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["members"] = state.Members
		inputs["name"] = state.Name
		inputs["policies"] = state.Policies
		inputs["region"] = state.Region
		inputs["valueSpecs"] = state.ValueSpecs
	}
	s, err := ctx.ReadResource("openstack:compute/serverGroup:ServerGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServerGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ServerGroup) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ServerGroup) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The instances that are part of this server group.
func (r *ServerGroup) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// A unique name for the server group. Changing this creates
// a new server group.
func (r *ServerGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The set of policies for the server group. Only two
// two policies are available right now, and both are mutually exclusive. See
// the Policies section for more information. Changing this creates a new
// server group.
func (r *ServerGroup) Policies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policies"])
}

// The region in which to obtain the V2 Compute client.
// If omitted, the `region` argument of the provider is used. Changing
// this creates a new server group.
func (r *ServerGroup) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Map of additional options.
func (r *ServerGroup) ValueSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["valueSpecs"])
}

// Input properties used for looking up and filtering ServerGroup resources.
type ServerGroupState struct {
	// The instances that are part of this server group.
	Members interface{}
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name interface{}
	// The set of policies for the server group. Only two
	// two policies are available right now, and both are mutually exclusive. See
	// the Policies section for more information. Changing this creates a new
	// server group.
	Policies interface{}
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region interface{}
	// Map of additional options.
	ValueSpecs interface{}
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name interface{}
	// The set of policies for the server group. Only two
	// two policies are available right now, and both are mutually exclusive. See
	// the Policies section for more information. Changing this creates a new
	// server group.
	Policies interface{}
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region interface{}
	// Map of additional options.
	ValueSpecs interface{}
}