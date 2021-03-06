// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V3 Role resource within OpenStack Keystone.
// 
// Note: You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
type Role struct {
	s *pulumi.ResourceState
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOpt) (*Role, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["domainId"] = nil
		inputs["name"] = nil
		inputs["region"] = nil
	} else {
		inputs["domainId"] = args.DomainId
		inputs["name"] = args.Name
		inputs["region"] = args.Region
	}
	s, err := ctx.RegisterResource("openstack:identity/role:Role", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Role{s: s}, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RoleState, opts ...pulumi.ResourceOpt) (*Role, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["domainId"] = state.DomainId
		inputs["name"] = state.Name
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("openstack:identity/role:Role", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Role{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Role) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Role) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The domain the role belongs to.
func (r *Role) DomainId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domainId"])
}

// The name of the role.
func (r *Role) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new Role.
func (r *Role) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering Role resources.
type RoleState struct {
	// The domain the role belongs to.
	DomainId interface{}
	// The name of the role.
	Name interface{}
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new Role.
	Region interface{}
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// The domain the role belongs to.
	DomainId interface{}
	// The name of the role.
	Name interface{}
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new Role.
	Region interface{}
}
