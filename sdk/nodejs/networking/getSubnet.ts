// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get the ID of an available OpenStack subnet.
 */
export function getSubnet(args?: GetSubnetArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetResult> {
    args = args || {};
    return pulumi.runtime.invoke("openstack:networking/getSubnet:getSubnet", {
        "cidr": args.cidr,
        "dhcpDisabled": args.dhcpDisabled,
        "dhcpEnabled": args.dhcpEnabled,
        "gatewayIp": args.gatewayIp,
        "ipVersion": args.ipVersion,
        "ipv6AddressMode": args.ipv6AddressMode,
        "ipv6RaMode": args.ipv6RaMode,
        "name": args.name,
        "networkId": args.networkId,
        "region": args.region,
        "subnetId": args.subnetId,
        "subnetpoolId": args.subnetpoolId,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnet.
 */
export interface GetSubnetArgs {
    /**
     * The CIDR of the subnet.
     */
    readonly cidr?: string;
    /**
     * If the subnet has DHCP disabled.
     */
    readonly dhcpDisabled?: boolean;
    /**
     * If the subnet has DHCP enabled.
     */
    readonly dhcpEnabled?: boolean;
    /**
     * The IP of the subnet's gateway.
     */
    readonly gatewayIp?: string;
    /**
     * The IP version of the subnet (either 4 or 6).
     */
    readonly ipVersion?: number;
    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    readonly ipv6AddressMode?: string;
    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    readonly ipv6RaMode?: string;
    /**
     * The name of the subnet.
     */
    readonly name?: string;
    /**
     * The ID of the network the subnet belongs to.
     */
    readonly networkId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve subnet ids. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The ID of the subnet.
     */
    readonly subnetId?: string;
    /**
     * The ID of the subnetpool associated with the subnet.
     */
    readonly subnetpoolId?: string;
    /**
     * The owner of the subnet.
     */
    readonly tenantId?: string;
}

/**
 * A collection of values returned by getSubnet.
 */
export interface GetSubnetResult {
    /**
     * Allocation pools of the subnet.
     */
    readonly allocationPools: { end: string, start: string }[];
    readonly cidr: string;
    /**
     * DNS Nameservers of the subnet.
     */
    readonly dnsNameservers: string[];
    /**
     * Whether the subnet has DHCP enabled or not.
     */
    readonly enableDhcp: boolean;
    readonly gatewayIp: string;
    /**
     * Host Routes of the subnet.
     */
    readonly hostRoutes: { destinationCidr: string, nextHop: string }[];
    readonly ipVersion: number;
    readonly ipv6AddressMode: string;
    readonly ipv6RaMode: string;
    readonly name: string;
    readonly networkId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    readonly subnetId: string;
    readonly subnetpoolId: string;
    readonly tenantId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
