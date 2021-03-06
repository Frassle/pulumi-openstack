// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 load balancer monitor resource within OpenStack.
type MonitorV1 struct {
	s *pulumi.ResourceState
}

// NewMonitorV1 registers a new resource with the given unique name, arguments, and options.
func NewMonitorV1(ctx *pulumi.Context,
	name string, args *MonitorV1Args, opts ...pulumi.ResourceOpt) (*MonitorV1, error) {
	if args == nil || args.Delay == nil {
		return nil, errors.New("missing required argument 'Delay'")
	}
	if args == nil || args.MaxRetries == nil {
		return nil, errors.New("missing required argument 'MaxRetries'")
	}
	if args == nil || args.Timeout == nil {
		return nil, errors.New("missing required argument 'Timeout'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adminStateUp"] = nil
		inputs["delay"] = nil
		inputs["expectedCodes"] = nil
		inputs["httpMethod"] = nil
		inputs["maxRetries"] = nil
		inputs["region"] = nil
		inputs["tenantId"] = nil
		inputs["timeout"] = nil
		inputs["type"] = nil
		inputs["urlPath"] = nil
	} else {
		inputs["adminStateUp"] = args.AdminStateUp
		inputs["delay"] = args.Delay
		inputs["expectedCodes"] = args.ExpectedCodes
		inputs["httpMethod"] = args.HttpMethod
		inputs["maxRetries"] = args.MaxRetries
		inputs["region"] = args.Region
		inputs["tenantId"] = args.TenantId
		inputs["timeout"] = args.Timeout
		inputs["type"] = args.Type
		inputs["urlPath"] = args.UrlPath
	}
	s, err := ctx.RegisterResource("openstack:loadbalancer/monitorV1:MonitorV1", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MonitorV1{s: s}, nil
}

// GetMonitorV1 gets an existing MonitorV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitorV1(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MonitorV1State, opts ...pulumi.ResourceOpt) (*MonitorV1, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adminStateUp"] = state.AdminStateUp
		inputs["delay"] = state.Delay
		inputs["expectedCodes"] = state.ExpectedCodes
		inputs["httpMethod"] = state.HttpMethod
		inputs["maxRetries"] = state.MaxRetries
		inputs["region"] = state.Region
		inputs["tenantId"] = state.TenantId
		inputs["timeout"] = state.Timeout
		inputs["type"] = state.Type
		inputs["urlPath"] = state.UrlPath
	}
	s, err := ctx.ReadResource("openstack:loadbalancer/monitorV1:MonitorV1", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MonitorV1{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MonitorV1) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MonitorV1) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The administrative state of the monitor.
// Acceptable values are "true" and "false". Changing this value updates the
// state of the existing monitor.
func (r *MonitorV1) AdminStateUp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["adminStateUp"])
}

// The time, in seconds, between sending probes to members.
// Changing this creates a new monitor.
func (r *MonitorV1) Delay() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["delay"])
}

// equired for HTTP(S) types. Expected HTTP codes
// for a passing HTTP(S) monitor. You can either specify a single status like
// "200", or a range like "200-202". Changing this updates the expected_codes
// of the existing monitor.
func (r *MonitorV1) ExpectedCodes() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["expectedCodes"])
}

// Required for HTTP(S) types. The HTTP method used
// for requests by the monitor. If this attribute is not specified, it defaults
// to "GET". Changing this updates the http_method of the existing monitor.
func (r *MonitorV1) HttpMethod() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["httpMethod"])
}

// Number of permissible ping failures before changing
// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
// this updates the max_retries of the existing monitor.
func (r *MonitorV1) MaxRetries() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxRetries"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an LB monitor. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// LB monitor.
func (r *MonitorV1) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The owner of the monitor. Required if admin wants to
// create a monitor for another tenant. Changing this creates a new monitor.
func (r *MonitorV1) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// Maximum number of seconds for a monitor to wait for a
// ping reply before it times out. The value must be less than the delay value.
// Changing this updates the timeout of the existing monitor.
func (r *MonitorV1) Timeout() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["timeout"])
}

// The type of probe, which is PING, TCP, HTTP, or HTTPS,
// that is sent by the monitor to verify the member state. Changing this
// creates a new monitor.
func (r *MonitorV1) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Required for HTTP(S) types. URI path that will be
// accessed if monitor type is HTTP or HTTPS. Changing this updates the
// url_path of the existing monitor.
func (r *MonitorV1) UrlPath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["urlPath"])
}

// Input properties used for looking up and filtering MonitorV1 resources.
type MonitorV1State struct {
	// The administrative state of the monitor.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing monitor.
	AdminStateUp interface{}
	// The time, in seconds, between sending probes to members.
	// Changing this creates a new monitor.
	Delay interface{}
	// equired for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202". Changing this updates the expected_codes
	// of the existing monitor.
	ExpectedCodes interface{}
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it defaults
	// to "GET". Changing this updates the http_method of the existing monitor.
	HttpMethod interface{}
	// Number of permissible ping failures before changing
	// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
	// this updates the max_retries of the existing monitor.
	MaxRetries interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB monitor. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB monitor.
	Region interface{}
	// The owner of the monitor. Required if admin wants to
	// create a monitor for another tenant. Changing this creates a new monitor.
	TenantId interface{}
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// Changing this updates the timeout of the existing monitor.
	Timeout interface{}
	// The type of probe, which is PING, TCP, HTTP, or HTTPS,
	// that is sent by the monitor to verify the member state. Changing this
	// creates a new monitor.
	Type interface{}
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS. Changing this updates the
	// url_path of the existing monitor.
	UrlPath interface{}
}

// The set of arguments for constructing a MonitorV1 resource.
type MonitorV1Args struct {
	// The administrative state of the monitor.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing monitor.
	AdminStateUp interface{}
	// The time, in seconds, between sending probes to members.
	// Changing this creates a new monitor.
	Delay interface{}
	// equired for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202". Changing this updates the expected_codes
	// of the existing monitor.
	ExpectedCodes interface{}
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it defaults
	// to "GET". Changing this updates the http_method of the existing monitor.
	HttpMethod interface{}
	// Number of permissible ping failures before changing
	// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
	// this updates the max_retries of the existing monitor.
	MaxRetries interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB monitor. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB monitor.
	Region interface{}
	// The owner of the monitor. Required if admin wants to
	// create a monitor for another tenant. Changing this creates a new monitor.
	TenantId interface{}
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// Changing this updates the timeout of the existing monitor.
	Timeout interface{}
	// The type of probe, which is PING, TCP, HTTP, or HTTPS,
	// that is sent by the monitor to verify the member state. Changing this
	// creates a new monitor.
	Type interface{}
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS. Changing this updates the
	// url_path of the existing monitor.
	UrlPath interface{}
}
