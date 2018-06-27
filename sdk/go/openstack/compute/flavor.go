// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 flavor resource within OpenStack.
type Flavor struct {
	s *pulumi.ResourceState
}

// NewFlavor registers a new resource with the given unique name, arguments, and options.
func NewFlavor(ctx *pulumi.Context,
	name string, args *FlavorArgs, opts ...pulumi.ResourceOpt) (*Flavor, error) {
	if args == nil || args.Disk == nil {
		return nil, errors.New("missing required argument 'Disk'")
	}
	if args == nil || args.Ram == nil {
		return nil, errors.New("missing required argument 'Ram'")
	}
	if args == nil || args.Vcpus == nil {
		return nil, errors.New("missing required argument 'Vcpus'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["disk"] = nil
		inputs["ephemeral"] = nil
		inputs["extraSpecs"] = nil
		inputs["isPublic"] = nil
		inputs["name"] = nil
		inputs["ram"] = nil
		inputs["region"] = nil
		inputs["rxTxFactor"] = nil
		inputs["swap"] = nil
		inputs["vcpus"] = nil
	} else {
		inputs["disk"] = args.Disk
		inputs["ephemeral"] = args.Ephemeral
		inputs["extraSpecs"] = args.ExtraSpecs
		inputs["isPublic"] = args.IsPublic
		inputs["name"] = args.Name
		inputs["ram"] = args.Ram
		inputs["region"] = args.Region
		inputs["rxTxFactor"] = args.RxTxFactor
		inputs["swap"] = args.Swap
		inputs["vcpus"] = args.Vcpus
	}
	s, err := ctx.RegisterResource("openstack:compute/flavor:Flavor", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Flavor{s: s}, nil
}

// GetFlavor gets an existing Flavor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlavor(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FlavorState, opts ...pulumi.ResourceOpt) (*Flavor, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["disk"] = state.Disk
		inputs["ephemeral"] = state.Ephemeral
		inputs["extraSpecs"] = state.ExtraSpecs
		inputs["isPublic"] = state.IsPublic
		inputs["name"] = state.Name
		inputs["ram"] = state.Ram
		inputs["region"] = state.Region
		inputs["rxTxFactor"] = state.RxTxFactor
		inputs["swap"] = state.Swap
		inputs["vcpus"] = state.Vcpus
	}
	s, err := ctx.ReadResource("openstack:compute/flavor:Flavor", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Flavor{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Flavor) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Flavor) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The amount of disk space in gigabytes to use for the root
// (/) partition. Changing this creates a new flavor.
func (r *Flavor) Disk() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["disk"])
}

func (r *Flavor) Ephemeral() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ephemeral"])
}

// Key/Value pairs of metadata for the flavor.
func (r *Flavor) ExtraSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["extraSpecs"])
}

// Whether the flavor is public. Changing this creates
// a new flavor.
func (r *Flavor) IsPublic() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["isPublic"])
}

// A unique name for the flavor. Changing this creates a new
// flavor.
func (r *Flavor) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The amount of RAM to use, in megabytes. Changing this
// creates a new flavor.
func (r *Flavor) Ram() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ram"])
}

// The region in which to obtain the V2 Compute client.
// Flavors are associated with accounts, but a Compute client is needed to
// create one. If omitted, the `region` argument of the provider is used.
// Changing this creates a new flavor.
func (r *Flavor) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// RX/TX bandwith factor. The default is 1. Changing
// this creates a new flavor.
func (r *Flavor) RxTxFactor() *pulumi.Float64Output {
	return (*pulumi.Float64Output)(r.s.State["rxTxFactor"])
}

// The amount of disk space in megabytes to use. If
// unspecified, the default is 0. Changing this creates a new flavor.
func (r *Flavor) Swap() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["swap"])
}

// The number of virtual CPUs to use. Changing this creates
// a new flavor.
func (r *Flavor) Vcpus() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["vcpus"])
}

// Input properties used for looking up and filtering Flavor resources.
type FlavorState struct {
	// The amount of disk space in gigabytes to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk interface{}
	Ephemeral interface{}
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs interface{}
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic interface{}
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name interface{}
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram interface{}
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region interface{}
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor interface{}
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap interface{}
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus interface{}
}

// The set of arguments for constructing a Flavor resource.
type FlavorArgs struct {
	// The amount of disk space in gigabytes to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk interface{}
	Ephemeral interface{}
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs interface{}
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic interface{}
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name interface{}
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram interface{}
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region interface{}
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor interface{}
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap interface{}
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus interface{}
}