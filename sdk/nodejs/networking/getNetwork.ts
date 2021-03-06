// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get the ID of an available OpenStack network.
 */
export function getNetwork(args?: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {
    args = args || {};
    return pulumi.runtime.invoke("openstack:networking/getNetwork:getNetwork", {
        "external": args.external,
        "matchingSubnetCidr": args.matchingSubnetCidr,
        "name": args.name,
        "networkId": args.networkId,
        "region": args.region,
        "status": args.status,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * The external routing facility of the network.
     */
    readonly external?: boolean;
    /**
     * The CIDR of a subnet within the network.
     */
    readonly matchingSubnetCidr?: string;
    /**
     * The name of the network.
     */
    readonly name?: string;
    /**
     * The ID of the network.
     */
    readonly networkId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve networks ids. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The status of the network.
     */
    readonly status?: string;
    /**
     * The owner of the network.
     */
    readonly tenantId?: string;
}

/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    /**
     * (Optional) The administrative state of the network.
     */
    readonly adminStateUp: string;
    /**
     * (Optional) The availability zone candidates for the network.
     */
    readonly availabilityZoneHints: string[];
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * (Optional)  Specifies whether the network resource can be accessed
     * by any tenant or not.
     */
    readonly shared: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
