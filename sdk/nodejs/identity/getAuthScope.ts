// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get authentication information about the current
 * auth scope in use. This can be used as self-discovery or introspection of
 * the username or project name currently in use.
 */
export function getAuthScope(args: GetAuthScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthScopeResult> {
    return pulumi.runtime.invoke("openstack:identity/getAuthScope:getAuthScope", {
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthScope.
 */
export interface GetAuthScopeArgs {
    /**
     * The name of the scope. This is an arbitrary name which is
     * only used as a unique identifier so an actual token isn't used as the ID.
     */
    readonly name: string;
    /**
     * The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getAuthScope.
 */
export interface GetAuthScopeResult {
    /**
     * The domain ID of the project.
     */
    readonly projectDomainId: string;
    /**
     * The domain name of the project.
     */
    readonly projectDomainName: string;
    /**
     * The project ID of the scope.
     */
    readonly projectId: string;
    /**
     * The project name of the scope.
     */
    readonly projectName: string;
    readonly region: string;
    /**
     * A list of roles in the current scope. See reference below.
     */
    readonly roles: { roleId: string, roleName: string }[];
    /**
     * The domain ID of the user.
     */
    readonly userDomainId: string;
    /**
     * The domain name of the user.
     */
    readonly userDomainName: string;
    /**
     * The user ID the of the scope.
     */
    readonly userId: string;
    /**
     * The username of the scope.
     */
    readonly userName: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
