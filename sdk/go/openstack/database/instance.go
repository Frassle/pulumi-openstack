// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 DB instance resource within OpenStack.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.Datastore == nil {
		return nil, errors.New("missing required argument 'Datastore'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["configurationId"] = nil
		inputs["databases"] = nil
		inputs["datastore"] = nil
		inputs["flavorId"] = nil
		inputs["name"] = nil
		inputs["networks"] = nil
		inputs["region"] = nil
		inputs["size"] = nil
		inputs["users"] = nil
	} else {
		inputs["configurationId"] = args.ConfigurationId
		inputs["databases"] = args.Databases
		inputs["datastore"] = args.Datastore
		inputs["flavorId"] = args.FlavorId
		inputs["name"] = args.Name
		inputs["networks"] = args.Networks
		inputs["region"] = args.Region
		inputs["size"] = args.Size
		inputs["users"] = args.Users
	}
	s, err := ctx.RegisterResource("openstack:database/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["configurationId"] = state.ConfigurationId
		inputs["databases"] = state.Databases
		inputs["datastore"] = state.Datastore
		inputs["flavorId"] = state.FlavorId
		inputs["name"] = state.Name
		inputs["networks"] = state.Networks
		inputs["region"] = state.Region
		inputs["size"] = state.Size
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("openstack:database/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Configuration ID to be attached to the instance. Database instance
// will be rebooted when configuration is detached.
func (r *Instance) ConfigurationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["configurationId"])
}

// An array of database name, charset and collate. The database
// object structure is documented below.
func (r *Instance) Databases() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["databases"])
}

// An array of database engine type and version. The datastore
// object structure is documented below. Changing this creates a new instance.
func (r *Instance) Datastore() *pulumi.Output {
	return r.s.State["datastore"]
}

// The flavor ID of the desired flavor for the instance.
// Changing this creates new instance.
func (r *Instance) FlavorId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["flavorId"])
}

// Database to be created on new instance. Changing this creates a
// new instance.
func (r *Instance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// An array of one or more networks to attach to the
// instance. The network object structure is documented below. Changing this
// creates a new instance.
func (r *Instance) Networks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networks"])
}

// The region in which to create the db instance. Changing this
// creates a new instance.
func (r *Instance) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Specifies the volume size in GB. Changing this creates new instance.
func (r *Instance) Size() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["size"])
}

// An array of username, password, host and databases. The user
// object structure is documented below.
func (r *Instance) Users() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// Configuration ID to be attached to the instance. Database instance
	// will be rebooted when configuration is detached.
	ConfigurationId interface{}
	// An array of database name, charset and collate. The database
	// object structure is documented below.
	Databases interface{}
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates a new instance.
	Datastore interface{}
	// The flavor ID of the desired flavor for the instance.
	// Changing this creates new instance.
	FlavorId interface{}
	// Database to be created on new instance. Changing this creates a
	// new instance.
	Name interface{}
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new instance.
	Networks interface{}
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region interface{}
	// Specifies the volume size in GB. Changing this creates new instance.
	Size interface{}
	// An array of username, password, host and databases. The user
	// object structure is documented below.
	Users interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Configuration ID to be attached to the instance. Database instance
	// will be rebooted when configuration is detached.
	ConfigurationId interface{}
	// An array of database name, charset and collate. The database
	// object structure is documented below.
	Databases interface{}
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates a new instance.
	Datastore interface{}
	// The flavor ID of the desired flavor for the instance.
	// Changing this creates new instance.
	FlavorId interface{}
	// Database to be created on new instance. Changing this creates a
	// new instance.
	Name interface{}
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new instance.
	Networks interface{}
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region interface{}
	// Specifies the volume size in GB. Changing this creates new instance.
	Size interface{}
	// An array of username, password, host and databases. The user
	// object structure is documented below.
	Users interface{}
}
