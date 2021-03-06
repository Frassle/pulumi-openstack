// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 router resource within OpenStack.
type Router struct {
	s *pulumi.ResourceState
}

// NewRouter registers a new resource with the given unique name, arguments, and options.
func NewRouter(ctx *pulumi.Context,
	name string, args *RouterArgs, opts ...pulumi.ResourceOpt) (*Router, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adminStateUp"] = nil
		inputs["availabilityZoneHints"] = nil
		inputs["distributed"] = nil
		inputs["enableSnat"] = nil
		inputs["externalFixedIps"] = nil
		inputs["externalGateway"] = nil
		inputs["externalNetworkId"] = nil
		inputs["name"] = nil
		inputs["region"] = nil
		inputs["tenantId"] = nil
		inputs["valueSpecs"] = nil
		inputs["vendorOptions"] = nil
	} else {
		inputs["adminStateUp"] = args.AdminStateUp
		inputs["availabilityZoneHints"] = args.AvailabilityZoneHints
		inputs["distributed"] = args.Distributed
		inputs["enableSnat"] = args.EnableSnat
		inputs["externalFixedIps"] = args.ExternalFixedIps
		inputs["externalGateway"] = args.ExternalGateway
		inputs["externalNetworkId"] = args.ExternalNetworkId
		inputs["name"] = args.Name
		inputs["region"] = args.Region
		inputs["tenantId"] = args.TenantId
		inputs["valueSpecs"] = args.ValueSpecs
		inputs["vendorOptions"] = args.VendorOptions
	}
	s, err := ctx.RegisterResource("openstack:networking/router:Router", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Router{s: s}, nil
}

// GetRouter gets an existing Router resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouterState, opts ...pulumi.ResourceOpt) (*Router, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adminStateUp"] = state.AdminStateUp
		inputs["availabilityZoneHints"] = state.AvailabilityZoneHints
		inputs["distributed"] = state.Distributed
		inputs["enableSnat"] = state.EnableSnat
		inputs["externalFixedIps"] = state.ExternalFixedIps
		inputs["externalGateway"] = state.ExternalGateway
		inputs["externalNetworkId"] = state.ExternalNetworkId
		inputs["name"] = state.Name
		inputs["region"] = state.Region
		inputs["tenantId"] = state.TenantId
		inputs["valueSpecs"] = state.ValueSpecs
		inputs["vendorOptions"] = state.VendorOptions
	}
	s, err := ctx.ReadResource("openstack:networking/router:Router", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Router{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Router) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Router) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Administrative up/down status for the router
// (must be "true" or "false" if provided). Changing this updates the
// `admin_state_up` of an existing router.
func (r *Router) AdminStateUp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["adminStateUp"])
}

// An availability zone is used to make 
// network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
// this creates a new router.
func (r *Router) AvailabilityZoneHints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["availabilityZoneHints"])
}

// Indicates whether or not to create a
// distributed router. The default policy setting in Neutron restricts
// usage of this property to administrative users only.
func (r *Router) Distributed() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["distributed"])
}

// Enable Source NAT for the router. Valid values are
// "true" or "false". An `external_network_id` has to be set in order to
// set this property. Changing this updates the `enable_snat` of the router.
func (r *Router) EnableSnat() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableSnat"])
}

// An external fixed IP for the router. This
// can be repeated. The structure is described below. An `external_network_id`
// has to be set in order to set this property. Changing this updates the
// external fixed IPs of the router.
func (r *Router) ExternalFixedIps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["externalFixedIps"])
}

// The
// network UUID of an external gateway for the router. A router with an
// external gateway is required if any compute instances or load balancers
// will be using floating IPs. Changing this updates the external gateway
// of an existing router.
func (r *Router) ExternalGateway() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["externalGateway"])
}

// The network UUID of an external gateway
// for the router. A router with an external gateway is required if any
// compute instances or load balancers will be using floating IPs. Changing
// this updates the external gateway of the router.
func (r *Router) ExternalNetworkId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["externalNetworkId"])
}

// A unique name for the router. Changing this
// updates the `name` of an existing router.
func (r *Router) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to create a router. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// router.
func (r *Router) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The owner of the floating IP. Required if admin wants
// to create a router for another tenant. Changing this creates a new router.
func (r *Router) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// Map of additional driver-specific options.
func (r *Router) ValueSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["valueSpecs"])
}

// Map of additional vendor-specific options.
// Supported options are described below.
func (r *Router) VendorOptions() *pulumi.Output {
	return r.s.State["vendorOptions"]
}

// Input properties used for looking up and filtering Router resources.
type RouterState struct {
	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// `admin_state_up` of an existing router.
	AdminStateUp interface{}
	// An availability zone is used to make 
	// network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
	// this creates a new router.
	AvailabilityZoneHints interface{}
	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed interface{}
	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An `external_network_id` has to be set in order to
	// set this property. Changing this updates the `enable_snat` of the router.
	EnableSnat interface{}
	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An `external_network_id`
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIps interface{}
	// The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	ExternalGateway interface{}
	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkId interface{}
	// A unique name for the router. Changing this
	// updates the `name` of an existing router.
	Name interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router.
	Region interface{}
	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantId interface{}
	// Map of additional driver-specific options.
	ValueSpecs interface{}
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions interface{}
}

// The set of arguments for constructing a Router resource.
type RouterArgs struct {
	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// `admin_state_up` of an existing router.
	AdminStateUp interface{}
	// An availability zone is used to make 
	// network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
	// this creates a new router.
	AvailabilityZoneHints interface{}
	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed interface{}
	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An `external_network_id` has to be set in order to
	// set this property. Changing this updates the `enable_snat` of the router.
	EnableSnat interface{}
	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An `external_network_id`
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIps interface{}
	// The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	ExternalGateway interface{}
	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkId interface{}
	// A unique name for the router. Changing this
	// updates the `name` of an existing router.
	Name interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router.
	Region interface{}
	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantId interface{}
	// Map of additional driver-specific options.
	ValueSpecs interface{}
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions interface{}
}
