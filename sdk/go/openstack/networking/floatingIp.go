// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 floating IP resource within OpenStack Neutron (networking)
// that can be used for load balancers.
// These are similar to Nova (compute) floating IP resources,
// but only compute floating IPs can be used with compute instances.
type FloatingIp struct {
	s *pulumi.ResourceState
}

// NewFloatingIp registers a new resource with the given unique name, arguments, and options.
func NewFloatingIp(ctx *pulumi.Context,
	name string, args *FloatingIpArgs, opts ...pulumi.ResourceOpt) (*FloatingIp, error) {
	if args == nil || args.Pool == nil {
		return nil, errors.New("missing required argument 'Pool'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["address"] = nil
		inputs["fixedIp"] = nil
		inputs["pool"] = nil
		inputs["portId"] = nil
		inputs["region"] = nil
		inputs["subnetId"] = nil
		inputs["tenantId"] = nil
		inputs["valueSpecs"] = nil
	} else {
		inputs["address"] = args.Address
		inputs["fixedIp"] = args.FixedIp
		inputs["pool"] = args.Pool
		inputs["portId"] = args.PortId
		inputs["region"] = args.Region
		inputs["subnetId"] = args.SubnetId
		inputs["tenantId"] = args.TenantId
		inputs["valueSpecs"] = args.ValueSpecs
	}
	s, err := ctx.RegisterResource("openstack:networking/floatingIp:FloatingIp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FloatingIp{s: s}, nil
}

// GetFloatingIp gets an existing FloatingIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FloatingIpState, opts ...pulumi.ResourceOpt) (*FloatingIp, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["address"] = state.Address
		inputs["fixedIp"] = state.FixedIp
		inputs["pool"] = state.Pool
		inputs["portId"] = state.PortId
		inputs["region"] = state.Region
		inputs["subnetId"] = state.SubnetId
		inputs["tenantId"] = state.TenantId
		inputs["valueSpecs"] = state.ValueSpecs
	}
	s, err := ctx.ReadResource("openstack:networking/floatingIp:FloatingIp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FloatingIp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FloatingIp) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FloatingIp) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The actual/specific floating IP to obtain. By default,
// non-admin users are not able to specify a floating IP, so you must either be
// an admin user or have had a custom policy or role applied to your OpenStack
// user or project.
func (r *FloatingIp) Address() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["address"])
}

// Fixed IP of the port to associate with this floating IP. Required if
// the port has multiple fixed IPs.
func (r *FloatingIp) FixedIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fixedIp"])
}

// The name of the pool from which to obtain the floating
// IP. Changing this creates a new floating IP.
func (r *FloatingIp) Pool() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["pool"])
}

// ID of an existing port with at least one IP address to
// associate with this floating IP.
func (r *FloatingIp) PortId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["portId"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a floating IP that can be used with
// another networking resource, such as a load balancer. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// floating IP (which may or may not have a different address).
func (r *FloatingIp) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The subnet ID of the floating IP pool. Specify this if
// the floating IP network has multiple subnets.
func (r *FloatingIp) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// The target tenant ID in which to allocate the floating
// IP, if you specify this together with a port_id, make sure the target port
// belongs to the same tenant. Changing this creates a new floating IP (which
// may or may not have a different address)
func (r *FloatingIp) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// Map of additional options.
func (r *FloatingIp) ValueSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["valueSpecs"])
}

// Input properties used for looking up and filtering FloatingIp resources.
type FloatingIpState struct {
	// The actual/specific floating IP to obtain. By default,
	// non-admin users are not able to specify a floating IP, so you must either be
	// an admin user or have had a custom policy or role applied to your OpenStack
	// user or project.
	Address interface{}
	// Fixed IP of the port to associate with this floating IP. Required if
	// the port has multiple fixed IPs.
	FixedIp interface{}
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool interface{}
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region interface{}
	// The subnet ID of the floating IP pool. Specify this if
	// the floating IP network has multiple subnets.
	SubnetId interface{}
	// The target tenant ID in which to allocate the floating
	// IP, if you specify this together with a port_id, make sure the target port
	// belongs to the same tenant. Changing this creates a new floating IP (which
	// may or may not have a different address)
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}

// The set of arguments for constructing a FloatingIp resource.
type FloatingIpArgs struct {
	// The actual/specific floating IP to obtain. By default,
	// non-admin users are not able to specify a floating IP, so you must either be
	// an admin user or have had a custom policy or role applied to your OpenStack
	// user or project.
	Address interface{}
	// Fixed IP of the port to associate with this floating IP. Required if
	// the port has multiple fixed IPs.
	FixedIp interface{}
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool interface{}
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region interface{}
	// The subnet ID of the floating IP pool. Specify this if
	// the floating IP network has multiple subnets.
	SubnetId interface{}
	// The target tenant ID in which to allocate the floating
	// IP, if you specify this together with a port_id, make sure the target port
	// belongs to the same tenant. Changing this creates a new floating IP (which
	// may or may not have a different address)
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}
