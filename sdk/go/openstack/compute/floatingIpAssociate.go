// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associate a floating IP to an instance. This can be used instead of the
// `floating_ip` options in `openstack_compute_instance_v2`.
type FloatingIpAssociate struct {
	s *pulumi.ResourceState
}

// NewFloatingIpAssociate registers a new resource with the given unique name, arguments, and options.
func NewFloatingIpAssociate(ctx *pulumi.Context,
	name string, args *FloatingIpAssociateArgs, opts ...pulumi.ResourceOpt) (*FloatingIpAssociate, error) {
	if args == nil || args.FloatingIp == nil {
		return nil, errors.New("missing required argument 'FloatingIp'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["fixedIp"] = nil
		inputs["floatingIp"] = nil
		inputs["instanceId"] = nil
		inputs["region"] = nil
		inputs["waitUntilAssociated"] = nil
	} else {
		inputs["fixedIp"] = args.FixedIp
		inputs["floatingIp"] = args.FloatingIp
		inputs["instanceId"] = args.InstanceId
		inputs["region"] = args.Region
		inputs["waitUntilAssociated"] = args.WaitUntilAssociated
	}
	s, err := ctx.RegisterResource("openstack:compute/floatingIpAssociate:FloatingIpAssociate", name, true, inputs, opts...)
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
		inputs["fixedIp"] = state.FixedIp
		inputs["floatingIp"] = state.FloatingIp
		inputs["instanceId"] = state.InstanceId
		inputs["region"] = state.Region
		inputs["waitUntilAssociated"] = state.WaitUntilAssociated
	}
	s, err := ctx.ReadResource("openstack:compute/floatingIpAssociate:FloatingIpAssociate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FloatingIpAssociate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FloatingIpAssociate) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FloatingIpAssociate) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The specific IP address to direct traffic to.
func (r *FloatingIpAssociate) FixedIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fixedIp"])
}

// The floating IP to associate.
func (r *FloatingIpAssociate) FloatingIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["floatingIp"])
}

// The instance to associte the floating IP with.
func (r *FloatingIpAssociate) InstanceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceId"])
}

// The region in which to obtain the V2 Compute client.
// Keypairs are associated with accounts, but a Compute client is needed to
// create one. If omitted, the `region` argument of the provider is used.
// Changing this creates a new floatingip_associate.
func (r *FloatingIpAssociate) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// In cases where the OpenStack environment
// does not automatically wait until the association has finished, set this
// option to have Terraform poll the instance until the floating IP has been
// associated. Defaults to false.
func (r *FloatingIpAssociate) WaitUntilAssociated() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["waitUntilAssociated"])
}

// Input properties used for looking up and filtering FloatingIpAssociate resources.
type FloatingIpAssociateState struct {
	// The specific IP address to direct traffic to.
	FixedIp interface{}
	// The floating IP to associate.
	FloatingIp interface{}
	// The instance to associte the floating IP with.
	InstanceId interface{}
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region interface{}
	// In cases where the OpenStack environment
	// does not automatically wait until the association has finished, set this
	// option to have Terraform poll the instance until the floating IP has been
	// associated. Defaults to false.
	WaitUntilAssociated interface{}
}

// The set of arguments for constructing a FloatingIpAssociate resource.
type FloatingIpAssociateArgs struct {
	// The specific IP address to direct traffic to.
	FixedIp interface{}
	// The floating IP to associate.
	FloatingIp interface{}
	// The instance to associte the floating IP with.
	InstanceId interface{}
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region interface{}
	// In cases where the OpenStack environment
	// does not automatically wait until the association has finished, set this
	// option to have Terraform poll the instance until the floating IP has been
	// associated. Defaults to false.
	WaitUntilAssociated interface{}
}
