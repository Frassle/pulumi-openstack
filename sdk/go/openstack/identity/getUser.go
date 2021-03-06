// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an OpenStack user.
func LookupUser(ctx *pulumi.Context, args *GetUserArgs) (*GetUserResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["domainId"] = args.DomainId
		inputs["enabled"] = args.Enabled
		inputs["idpId"] = args.IdpId
		inputs["name"] = args.Name
		inputs["passwordExpiresAt"] = args.PasswordExpiresAt
		inputs["protocolId"] = args.ProtocolId
		inputs["region"] = args.Region
		inputs["uniqueId"] = args.UniqueId
	}
	outputs, err := ctx.Invoke("openstack:identity/getUser:getUser", inputs)
	if err != nil {
		return nil, err
	}
	return &GetUserResult{
		DefaultProjectId: outputs["defaultProjectId"],
		DomainId: outputs["domainId"],
		Region: outputs["region"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// The domain this user belongs to.
	DomainId interface{}
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled interface{}
	// The identity provider ID of the user.
	IdpId interface{}
	// The name of the user.
	Name interface{}
	// Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
	PasswordExpiresAt interface{}
	// The protocol ID of the user.
	ProtocolId interface{}
	Region interface{}
	// The unique ID of the user.
	UniqueId interface{}
}

// A collection of values returned by getUser.
type GetUserResult struct {
	// See Argument Reference above.
	DefaultProjectId interface{}
	// See Argument Reference above.
	DomainId interface{}
	// The region the user is located in.
	Region interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
