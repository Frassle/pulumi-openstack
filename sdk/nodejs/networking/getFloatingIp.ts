// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get the ID of an available OpenStack floating IP.
 */
export function getFloatingIp(args?: GetFloatingIpArgs, opts?: pulumi.InvokeOptions): Promise<GetFloatingIpResult> {
    args = args || {};
    return pulumi.runtime.invoke("openstack:networking/getFloatingIp:getFloatingIp", {
        "address": args.address,
        "fixedIp": args.fixedIp,
        "pool": args.pool,
        "portId": args.portId,
        "region": args.region,
        "status": args.status,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFloatingIp.
 */
export interface GetFloatingIpArgs {
    /**
     * The IP address of the floating IP.
     */
    readonly address?: string;
    /**
     * The specific IP address of the internal port which should be associated with the floating IP.
     */
    readonly fixedIp?: string;
    /**
     * The name of the pool from which the floating IP belongs to.
     */
    readonly pool?: string;
    /**
     * The ID of the port the floating IP is attached.
     */
    readonly portId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve floating IP ids. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    readonly status?: string;
    /**
     * The owner of the floating IP.
     */
    readonly tenantId?: string;
}

/**
 * A collection of values returned by getFloatingIp.
 */
export interface GetFloatingIpResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
