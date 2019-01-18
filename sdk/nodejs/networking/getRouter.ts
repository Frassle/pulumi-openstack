// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack router.
 */
export function getRouter(args?: GetRouterArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterResult> {
    args = args || {};
    return pulumi.runtime.invoke("openstack:networking/getRouter:getRouter", {
        "adminStateUp": args.adminStateUp,
        "description": args.description,
        "distributed": args.distributed,
        "enableSnat": args.enableSnat,
        "name": args.name,
        "region": args.region,
        "routerId": args.routerId,
        "status": args.status,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouter.
 */
export interface GetRouterArgs {
    /**
     * Administrative up/down status for the router (must be "true" or "false" if provided).
     */
    readonly adminStateUp?: boolean;
    /**
     * Human-readable description of the router.
     */
    readonly description?: string;
    /**
     * Indicates whether or not to get a distributed router.
     */
    readonly distributed?: boolean;
    readonly enableSnat?: boolean;
    /**
     * The name of the router.
     */
    readonly name?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve router ids. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The UUID of the router resource.
     */
    readonly routerId?: string;
    /**
     * The status of the router (ACTIVE/DOWN).
     */
    readonly status?: string;
    /**
     * The owner of the router.
     */
    readonly tenantId?: string;
}

/**
 * A collection of values returned by getRouter.
 */
export interface GetRouterResult {
    /**
     * The availability zone that is used to make router resources highly available.
     */
    readonly availabilityZoneHints: string[];
    /**
     * The value that points out if the Source NAT is enabled on the router.
     */
    readonly enableSnat: boolean;
    /**
     * The external fixed IPs of the router.
     */
    readonly externalFixedIps: { ipAddress?: string, subnetId?: string }[];
    /**
     * The network UUID of an external gateway for the router.
     */
    readonly externalNetworkId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}