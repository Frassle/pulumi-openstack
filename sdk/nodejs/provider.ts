// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * The provider type for the openstack package
 */
export class Provider extends pulumi.ProviderResource {

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            inputs["authUrl"] = args ? args.authUrl : undefined;
            inputs["cacertFile"] = args ? args.cacertFile : undefined;
            inputs["cert"] = args ? args.cert : undefined;
            inputs["cloud"] = args ? args.cloud : undefined;
            inputs["defaultDomain"] = args ? args.defaultDomain : undefined;
            inputs["domainId"] = args ? args.domainId : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["endpointType"] = args ? args.endpointType : undefined;
            inputs["insecure"] = args ? args.insecure : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["projectDomainId"] = args ? args.projectDomainId : undefined;
            inputs["projectDomainName"] = args ? args.projectDomainName : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["swauth"] = args ? args.swauth : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["tenantName"] = args ? args.tenantName : undefined;
            inputs["token"] = args ? args.token : undefined;
            inputs["useOctavia"] = args ? args.useOctavia : undefined;
            inputs["userDomainId"] = args ? args.userDomainId : undefined;
            inputs["userDomainName"] = args ? args.userDomainName : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["userName"] = args ? args.userName : undefined;
        }
        super("openstack", name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The Identity authentication URL.
     */
    readonly authUrl?: pulumi.Input<string>;
    /**
     * A Custom CA certificate.
     */
    readonly cacertFile?: pulumi.Input<string>;
    /**
     * A client certificate to authenticate with.
     */
    readonly cert?: pulumi.Input<string>;
    /**
     * An entry in a `clouds.yaml` file to use.
     */
    readonly cloud?: pulumi.Input<string>;
    /**
     * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
     */
    readonly defaultDomain?: pulumi.Input<string>;
    /**
     * The ID of the Domain to scope to (Identity v3).
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * The name of the Domain to scope to (Identity v3).
     */
    readonly domainName?: pulumi.Input<string>;
    readonly endpointType?: pulumi.Input<string>;
    /**
     * Trust self-signed certificates.
     */
    readonly insecure?: pulumi.Input<boolean>;
    /**
     * A client private key to authenticate with.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * Password to login with.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The ID of the domain where the proejct resides (Identity v3).
     */
    readonly projectDomainId?: pulumi.Input<string>;
    /**
     * The name of the domain where the project resides (Identity v3).
     */
    readonly projectDomainName?: pulumi.Input<string>;
    /**
     * The OpenStack region to connect to.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
     */
    readonly swauth?: pulumi.Input<boolean>;
    /**
     * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
     */
    readonly tenantName?: pulumi.Input<string>;
    /**
     * Authentication token to use as an alternative to username/password.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service
     * (Neutron).
     */
    readonly useOctavia?: pulumi.Input<boolean>;
    /**
     * The ID of the domain where the user resides (Identity v3).
     */
    readonly userDomainId?: pulumi.Input<string>;
    /**
     * The name of the domain where the user resides (Identity v3).
     */
    readonly userDomainName?: pulumi.Input<string>;
    /**
     * Username to login with.
     */
    readonly userId?: pulumi.Input<string>;
    /**
     * Username to login with.
     */
    readonly userName?: pulumi.Input<string>;
}
