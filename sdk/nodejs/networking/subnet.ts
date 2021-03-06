// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a V2 Neutron subnet resource within OpenStack.
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetState): Subnet {
        return new Subnet(name, <any>state, { id });
    }

    /**
     * An array of sub-ranges of CIDR available for
     * dynamic allocation to ports. The allocation_pool object structure is
     * documented below. Changing this creates a new subnet.
     */
    public readonly allocationPools: pulumi.Output<{ end: string, start: string }[]>;
    /**
     * CIDR representing IP range for this subnet, based on IP
     * version. You can omit this option if you are creating a subnet from a
     * subnet pool.
     */
    public readonly cidr: pulumi.Output<string>;
    /**
     * An array of DNS name server names used by hosts
     * in this subnet. Changing this updates the DNS name servers for the existing
     * subnet.
     */
    public readonly dnsNameservers: pulumi.Output<string[] | undefined>;
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value enables or
     * disables the DHCP capabilities of the existing subnet. Defaults to true.
     */
    public readonly enableDhcp: pulumi.Output<boolean | undefined>;
    /**
     * Default gateway used by devices in this subnet.
     * Leaving this blank and not setting `no_gateway` will cause a default
     * gateway of `.1` to be used. Changing this updates the gateway IP of the
     * existing subnet.
     */
    public readonly gatewayIp: pulumi.Output<string>;
    /**
     * An array of routes that should be used by devices
     * with IPs from this subnet (not including local subnet route). The host_route
     * object structure is documented below. Changing this updates the host routes
     * for the existing subnet.
     */
    public readonly hostRoutes: pulumi.Output<{ destinationCidr: string, nextHop: string }[] | undefined>;
    /**
     * IP version, either 4 (default) or 6. Changing this creates a
     * new subnet.
     */
    public readonly ipVersion: pulumi.Output<number | undefined>;
    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    public readonly ipv6AddressMode: pulumi.Output<string>;
    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    public readonly ipv6RaMode: pulumi.Output<string>;
    /**
     * The name of the subnet. Changing this updates the name of
     * the existing subnet.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The UUID of the parent network. Changing this
     * creates a new subnet.
     */
    public readonly networkId: pulumi.Output<string>;
    /**
     * Do not set a gateway IP on this subnet. Changing
     * this removes or adds a default gateway IP of the existing subnet.
     */
    public readonly noGateway: pulumi.Output<boolean | undefined>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnet.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * The ID of the subnetpool associated with the subnet.
     */
    public readonly subnetpoolId: pulumi.Output<string | undefined>;
    /**
     * The owner of the subnet. Required if admin wants to
     * create a subnet for another tenant. Changing this creates a new subnet.
     */
    public readonly tenantId: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetArgs | SubnetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SubnetState = argsOrState as SubnetState | undefined;
            inputs["allocationPools"] = state ? state.allocationPools : undefined;
            inputs["cidr"] = state ? state.cidr : undefined;
            inputs["dnsNameservers"] = state ? state.dnsNameservers : undefined;
            inputs["enableDhcp"] = state ? state.enableDhcp : undefined;
            inputs["gatewayIp"] = state ? state.gatewayIp : undefined;
            inputs["hostRoutes"] = state ? state.hostRoutes : undefined;
            inputs["ipVersion"] = state ? state.ipVersion : undefined;
            inputs["ipv6AddressMode"] = state ? state.ipv6AddressMode : undefined;
            inputs["ipv6RaMode"] = state ? state.ipv6RaMode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkId"] = state ? state.networkId : undefined;
            inputs["noGateway"] = state ? state.noGateway : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["subnetpoolId"] = state ? state.subnetpoolId : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as SubnetArgs | undefined;
            if (!args || args.networkId === undefined) {
                throw new Error("Missing required property 'networkId'");
            }
            inputs["allocationPools"] = args ? args.allocationPools : undefined;
            inputs["cidr"] = args ? args.cidr : undefined;
            inputs["dnsNameservers"] = args ? args.dnsNameservers : undefined;
            inputs["enableDhcp"] = args ? args.enableDhcp : undefined;
            inputs["gatewayIp"] = args ? args.gatewayIp : undefined;
            inputs["hostRoutes"] = args ? args.hostRoutes : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["ipv6AddressMode"] = args ? args.ipv6AddressMode : undefined;
            inputs["ipv6RaMode"] = args ? args.ipv6RaMode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkId"] = args ? args.networkId : undefined;
            inputs["noGateway"] = args ? args.noGateway : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["subnetpoolId"] = args ? args.subnetpoolId : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        super("openstack:networking/subnet:Subnet", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnet resources.
 */
export interface SubnetState {
    /**
     * An array of sub-ranges of CIDR available for
     * dynamic allocation to ports. The allocation_pool object structure is
     * documented below. Changing this creates a new subnet.
     */
    readonly allocationPools?: pulumi.Input<pulumi.Input<{ end: pulumi.Input<string>, start: pulumi.Input<string> }>[]>;
    /**
     * CIDR representing IP range for this subnet, based on IP
     * version. You can omit this option if you are creating a subnet from a
     * subnet pool.
     */
    readonly cidr?: pulumi.Input<string>;
    /**
     * An array of DNS name server names used by hosts
     * in this subnet. Changing this updates the DNS name servers for the existing
     * subnet.
     */
    readonly dnsNameservers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value enables or
     * disables the DHCP capabilities of the existing subnet. Defaults to true.
     */
    readonly enableDhcp?: pulumi.Input<boolean>;
    /**
     * Default gateway used by devices in this subnet.
     * Leaving this blank and not setting `no_gateway` will cause a default
     * gateway of `.1` to be used. Changing this updates the gateway IP of the
     * existing subnet.
     */
    readonly gatewayIp?: pulumi.Input<string>;
    /**
     * An array of routes that should be used by devices
     * with IPs from this subnet (not including local subnet route). The host_route
     * object structure is documented below. Changing this updates the host routes
     * for the existing subnet.
     */
    readonly hostRoutes?: pulumi.Input<pulumi.Input<{ destinationCidr: pulumi.Input<string>, nextHop: pulumi.Input<string> }>[]>;
    /**
     * IP version, either 4 (default) or 6. Changing this creates a
     * new subnet.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    readonly ipv6AddressMode?: pulumi.Input<string>;
    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    readonly ipv6RaMode?: pulumi.Input<string>;
    /**
     * The name of the subnet. Changing this updates the name of
     * the existing subnet.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The UUID of the parent network. Changing this
     * creates a new subnet.
     */
    readonly networkId?: pulumi.Input<string>;
    /**
     * Do not set a gateway IP on this subnet. Changing
     * this removes or adds a default gateway IP of the existing subnet.
     */
    readonly noGateway?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnet.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ID of the subnetpool associated with the subnet.
     */
    readonly subnetpoolId?: pulumi.Input<string>;
    /**
     * The owner of the subnet. Required if admin wants to
     * create a subnet for another tenant. Changing this creates a new subnet.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    /**
     * An array of sub-ranges of CIDR available for
     * dynamic allocation to ports. The allocation_pool object structure is
     * documented below. Changing this creates a new subnet.
     */
    readonly allocationPools?: pulumi.Input<pulumi.Input<{ end: pulumi.Input<string>, start: pulumi.Input<string> }>[]>;
    /**
     * CIDR representing IP range for this subnet, based on IP
     * version. You can omit this option if you are creating a subnet from a
     * subnet pool.
     */
    readonly cidr?: pulumi.Input<string>;
    /**
     * An array of DNS name server names used by hosts
     * in this subnet. Changing this updates the DNS name servers for the existing
     * subnet.
     */
    readonly dnsNameservers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value enables or
     * disables the DHCP capabilities of the existing subnet. Defaults to true.
     */
    readonly enableDhcp?: pulumi.Input<boolean>;
    /**
     * Default gateway used by devices in this subnet.
     * Leaving this blank and not setting `no_gateway` will cause a default
     * gateway of `.1` to be used. Changing this updates the gateway IP of the
     * existing subnet.
     */
    readonly gatewayIp?: pulumi.Input<string>;
    /**
     * An array of routes that should be used by devices
     * with IPs from this subnet (not including local subnet route). The host_route
     * object structure is documented below. Changing this updates the host routes
     * for the existing subnet.
     */
    readonly hostRoutes?: pulumi.Input<pulumi.Input<{ destinationCidr: pulumi.Input<string>, nextHop: pulumi.Input<string> }>[]>;
    /**
     * IP version, either 4 (default) or 6. Changing this creates a
     * new subnet.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    readonly ipv6AddressMode?: pulumi.Input<string>;
    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    readonly ipv6RaMode?: pulumi.Input<string>;
    /**
     * The name of the subnet. Changing this updates the name of
     * the existing subnet.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The UUID of the parent network. Changing this
     * creates a new subnet.
     */
    readonly networkId: pulumi.Input<string>;
    /**
     * Do not set a gateway IP on this subnet. Changing
     * this removes or adds a default gateway IP of the existing subnet.
     */
    readonly noGateway?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnet.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ID of the subnetpool associated with the subnet.
     */
    readonly subnetpoolId?: pulumi.Input<string>;
    /**
     * The owner of the subnet. Required if admin wants to
     * create a subnet for another tenant. Changing this creates a new subnet.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
