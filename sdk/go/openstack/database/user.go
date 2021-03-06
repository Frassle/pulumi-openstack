// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 DB user resource within OpenStack.
type User struct {
	s *pulumi.ResourceState
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOpt) (*User, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["databases"] = nil
		inputs["host"] = nil
		inputs["instanceId"] = nil
		inputs["name"] = nil
		inputs["password"] = nil
		inputs["region"] = nil
	} else {
		inputs["databases"] = args.Databases
		inputs["host"] = args.Host
		inputs["instanceId"] = args.InstanceId
		inputs["name"] = args.Name
		inputs["password"] = args.Password
		inputs["region"] = args.Region
	}
	s, err := ctx.RegisterResource("openstack:database/user:User", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserState, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["databases"] = state.Databases
		inputs["host"] = state.Host
		inputs["instanceId"] = state.InstanceId
		inputs["name"] = state.Name
		inputs["password"] = state.Password
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("openstack:database/user:User", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *User) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *User) ID() *pulumi.IDOutput {
	return r.s.ID
}

// A list of database user should have access to.
func (r *User) Databases() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["databases"])
}

func (r *User) Host() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["host"])
}

func (r *User) InstanceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceId"])
}

// A unique name for the resource.
func (r *User) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// User's password.
func (r *User) Password() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["password"])
}

// Openstack region resource is created in.
func (r *User) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	// A list of database user should have access to.
	Databases interface{}
	Host interface{}
	InstanceId interface{}
	// A unique name for the resource.
	Name interface{}
	// User's password.
	Password interface{}
	// Openstack region resource is created in.
	Region interface{}
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// A list of database user should have access to.
	Databases interface{}
	Host interface{}
	InstanceId interface{}
	// A unique name for the resource.
	Name interface{}
	// User's password.
	Password interface{}
	// Openstack region resource is created in.
	Region interface{}
}
