// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V3 User resource within OpenStack Keystone.
// 
// Note: You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
type User struct {
	s *pulumi.ResourceState
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultProjectId"] = nil
		inputs["description"] = nil
		inputs["domainId"] = nil
		inputs["enabled"] = nil
		inputs["extra"] = nil
		inputs["ignoreChangePasswordUponFirstUse"] = nil
		inputs["ignoreLockoutFailureAttempts"] = nil
		inputs["ignorePasswordExpiry"] = nil
		inputs["multiFactorAuthEnabled"] = nil
		inputs["multiFactorAuthRules"] = nil
		inputs["name"] = nil
		inputs["password"] = nil
		inputs["region"] = nil
	} else {
		inputs["defaultProjectId"] = args.DefaultProjectId
		inputs["description"] = args.Description
		inputs["domainId"] = args.DomainId
		inputs["enabled"] = args.Enabled
		inputs["extra"] = args.Extra
		inputs["ignoreChangePasswordUponFirstUse"] = args.IgnoreChangePasswordUponFirstUse
		inputs["ignoreLockoutFailureAttempts"] = args.IgnoreLockoutFailureAttempts
		inputs["ignorePasswordExpiry"] = args.IgnorePasswordExpiry
		inputs["multiFactorAuthEnabled"] = args.MultiFactorAuthEnabled
		inputs["multiFactorAuthRules"] = args.MultiFactorAuthRules
		inputs["name"] = args.Name
		inputs["password"] = args.Password
		inputs["region"] = args.Region
	}
	s, err := ctx.RegisterResource("openstack:identity/user:User", name, true, inputs, opts...)
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
		inputs["defaultProjectId"] = state.DefaultProjectId
		inputs["description"] = state.Description
		inputs["domainId"] = state.DomainId
		inputs["enabled"] = state.Enabled
		inputs["extra"] = state.Extra
		inputs["ignoreChangePasswordUponFirstUse"] = state.IgnoreChangePasswordUponFirstUse
		inputs["ignoreLockoutFailureAttempts"] = state.IgnoreLockoutFailureAttempts
		inputs["ignorePasswordExpiry"] = state.IgnorePasswordExpiry
		inputs["multiFactorAuthEnabled"] = state.MultiFactorAuthEnabled
		inputs["multiFactorAuthRules"] = state.MultiFactorAuthRules
		inputs["name"] = state.Name
		inputs["password"] = state.Password
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("openstack:identity/user:User", name, id, inputs, opts...)
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

// The default project this user belongs to.
func (r *User) DefaultProjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultProjectId"])
}

// A description of the user.
func (r *User) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The domain this user belongs to.
func (r *User) DomainId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domainId"])
}

// Whether the user is enabled or disabled. Valid
// values are `true` and `false`.
func (r *User) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// Free-form key/value pairs of extra information.
func (r *User) Extra() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["extra"])
}

// User will not have to
// change their password upon first use. Valid values are `true` and `false`.
func (r *User) IgnoreChangePasswordUponFirstUse() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ignoreChangePasswordUponFirstUse"])
}

// User will not have a failure
// lockout placed on their account. Valid values are `true` and `false`.
func (r *User) IgnoreLockoutFailureAttempts() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ignoreLockoutFailureAttempts"])
}

// User's password will not expire.
// Valid values are `true` and `false`.
func (r *User) IgnorePasswordExpiry() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ignorePasswordExpiry"])
}

// Whether to enable multi-factor
// authentication. Valid values are `true` and `false`.
func (r *User) MultiFactorAuthEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["multiFactorAuthEnabled"])
}

// A multi-factor authentication rule.
// The structure is documented below. Please see the
// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
// for more information on how to use mulit-factor rules.
func (r *User) MultiFactorAuthRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["multiFactorAuthRules"])
}

// The name of the user.
func (r *User) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The password for the user.
func (r *User) Password() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["password"])
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new User.
func (r *User) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	// The default project this user belongs to.
	DefaultProjectId interface{}
	// A description of the user.
	Description interface{}
	// The domain this user belongs to.
	DomainId interface{}
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled interface{}
	// Free-form key/value pairs of extra information.
	Extra interface{}
	// User will not have to
	// change their password upon first use. Valid values are `true` and `false`.
	IgnoreChangePasswordUponFirstUse interface{}
	// User will not have a failure
	// lockout placed on their account. Valid values are `true` and `false`.
	IgnoreLockoutFailureAttempts interface{}
	// User's password will not expire.
	// Valid values are `true` and `false`.
	IgnorePasswordExpiry interface{}
	// Whether to enable multi-factor
	// authentication. Valid values are `true` and `false`.
	MultiFactorAuthEnabled interface{}
	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRules interface{}
	// The name of the user.
	Name interface{}
	// The password for the user.
	Password interface{}
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region interface{}
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The default project this user belongs to.
	DefaultProjectId interface{}
	// A description of the user.
	Description interface{}
	// The domain this user belongs to.
	DomainId interface{}
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled interface{}
	// Free-form key/value pairs of extra information.
	Extra interface{}
	// User will not have to
	// change their password upon first use. Valid values are `true` and `false`.
	IgnoreChangePasswordUponFirstUse interface{}
	// User will not have a failure
	// lockout placed on their account. Valid values are `true` and `false`.
	IgnoreLockoutFailureAttempts interface{}
	// User's password will not expire.
	// Valid values are `true` and `false`.
	IgnorePasswordExpiry interface{}
	// Whether to enable multi-factor
	// authentication. Valid values are `true` and `false`.
	MultiFactorAuthEnabled interface{}
	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRules interface{}
	// The name of the user.
	Name interface{}
	// The password for the user.
	Password interface{}
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region interface{}
}