// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associates a floating IP to a port. This is useful for situations
// where you have a pre-allocated floating IP or are unable to use the
// `openstack_networking_floatingip_v2` resource to create a floating IP.
type FloatingIpAssociate struct {
	s *pulumi.ResourceState
}

// NewFloatingIpAssociate registers a new resource with the given unique name, arguments, and options.
func NewFloatingIpAssociate(ctx *pulumi.Context,
	name string, args *FloatingIpAssociateArgs, opts ...pulumi.ResourceOpt) (*FloatingIpAssociate, error) {
	if args == nil || args.FloatingIp == nil {
		return nil, errors.New("missing required argument 'FloatingIp'")
	}
	if args == nil || args.PortId == nil {
		return nil, errors.New("missing required argument 'PortId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["floatingIp"] = nil
		inputs["portId"] = nil
		inputs["region"] = nil
	} else {
		inputs["floatingIp"] = args.FloatingIp
		inputs["portId"] = args.PortId
		inputs["region"] = args.Region
	}
	s, err := ctx.RegisterResource("openstack:networking/floatingIpAssociate:FloatingIpAssociate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FloatingIpAssociate{s: s}, nil
}

// GetFloatingIpAssociate gets an existing FloatingIpAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIpAssociate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FloatingIpAssociateState, opts ...pulumi.ResourceOpt) (*FloatingIpAssociate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["floatingIp"] = state.FloatingIp
		inputs["portId"] = state.PortId
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("openstack:networking/floatingIpAssociate:FloatingIpAssociate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FloatingIpAssociate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FloatingIpAssociate) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FloatingIpAssociate) ID() *pulumi.IDOutput {
	return r.s.ID
}

// IP Address of an existing floating IP.
func (r *FloatingIpAssociate) FloatingIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["floatingIp"])
}

// ID of an existing port with at least one IP address to
// associate with this floating IP.
func (r *FloatingIpAssociate) PortId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["portId"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a floating IP that can be used with
// another networking resource, such as a load balancer. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// floating IP (which may or may not have a different address).
func (r *FloatingIpAssociate) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering FloatingIpAssociate resources.
type FloatingIpAssociateState struct {
	// IP Address of an existing floating IP.
	FloatingIp interface{}
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region interface{}
}

// The set of arguments for constructing a FloatingIpAssociate resource.
type FloatingIpAssociateArgs struct {
	// IP Address of an existing floating IP.
	FloatingIp interface{}
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region interface{}
}
