import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 Neutron subnetpool resource within OpenStack.
 */
export declare class SubnetPool extends pulumi.CustomResource {
    /**
     * Get an existing SubnetPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetPoolState): SubnetPool;
    /**
     * The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     */
    readonly addressScopeId: pulumi.Output<string | undefined>;
    /**
     * The time at which subnetpool was created.
     */
    readonly createdAt: pulumi.Output<string>;
    /**
     * The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     */
    readonly defaultPrefixlen: pulumi.Output<number>;
    /**
     * The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     */
    readonly defaultQuota: pulumi.Output<number | undefined>;
    /**
     * The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * The IP protocol version.
     */
    readonly ipVersion: pulumi.Output<number>;
    /**
     * Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     */
    readonly isDefault: pulumi.Output<boolean | undefined>;
    /**
     * The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     */
    readonly maxPrefixlen: pulumi.Output<number>;
    /**
     * The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     */
    readonly minPrefixlen: pulumi.Output<number>;
    /**
     * The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     */
    readonly name: pulumi.Output<string>;
    /**
     * A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     */
    readonly prefixes: pulumi.Output<string[]>;
    /**
     * The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     */
    readonly projectId: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The revision number of the subnetpool.
     */
    readonly revisionNumber: pulumi.Output<number>;
    /**
     * Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     */
    readonly shared: pulumi.Output<boolean | undefined>;
    /**
     * The time at which subnetpool was created.
     */
    readonly updatedAt: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a SubnetPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetPoolArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering SubnetPool resources.
 */
export interface SubnetPoolState {
    /**
     * The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     */
    readonly addressScopeId?: pulumi.Input<string>;
    /**
     * The time at which subnetpool was created.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     */
    readonly defaultPrefixlen?: pulumi.Input<number>;
    /**
     * The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     */
    readonly defaultQuota?: pulumi.Input<number>;
    /**
     * The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The IP protocol version.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     */
    readonly isDefault?: pulumi.Input<boolean>;
    /**
     * The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     */
    readonly maxPrefixlen?: pulumi.Input<number>;
    /**
     * The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     */
    readonly minPrefixlen?: pulumi.Input<number>;
    /**
     * The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     */
    readonly prefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The revision number of the subnetpool.
     */
    readonly revisionNumber?: pulumi.Input<number>;
    /**
     * Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     */
    readonly shared?: pulumi.Input<boolean>;
    /**
     * The time at which subnetpool was created.
     */
    readonly updatedAt?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a SubnetPool resource.
 */
export interface SubnetPoolArgs {
    /**
     * The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     */
    readonly addressScopeId?: pulumi.Input<string>;
    /**
     * The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     */
    readonly defaultPrefixlen?: pulumi.Input<number>;
    /**
     * The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     */
    readonly defaultQuota?: pulumi.Input<number>;
    /**
     * The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The IP protocol version.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     */
    readonly isDefault?: pulumi.Input<boolean>;
    /**
     * The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     */
    readonly maxPrefixlen?: pulumi.Input<number>;
    /**
     * The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     */
    readonly minPrefixlen?: pulumi.Input<number>;
    /**
     * The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     */
    readonly prefixes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     */
    readonly shared?: pulumi.Input<boolean>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}