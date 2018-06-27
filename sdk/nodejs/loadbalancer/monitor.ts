// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a V2 monitor resource within OpenStack.
 */
export class Monitor extends pulumi.CustomResource {
    /**
     * Get an existing Monitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitorState): Monitor {
        return new Monitor(name, <any>state, { id });
    }

    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp: pulumi.Output<boolean | undefined>;
    /**
     * The time, in seconds, between sending probes to members.
     */
    public readonly delay: pulumi.Output<number>;
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * "200", or a range like "200-202".
     */
    public readonly expectedCodes: pulumi.Output<string>;
    /**
     * Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it
     * defaults to "GET".
     */
    public readonly httpMethod: pulumi.Output<string>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to INACTIVE. Must be a number between 1
     * and 10..
     */
    public readonly maxRetries: pulumi.Output<number>;
    /**
     * The Name of the Monitor.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The id of the pool that this monitor will be assigned to.
     */
    public readonly poolId: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     */
    public readonly tenantId: pulumi.Output<string>;
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     */
    public readonly timeout: pulumi.Output<number>;
    /**
     * The type of probe, which is PING, TCP, HTTP, or HTTPS,
     * that is sent by the load balancer to verify the member state. Changing this
     * creates a new monitor.
     */
    public readonly type: pulumi.Output<string>;
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS.
     */
    public readonly urlPath: pulumi.Output<string>;

    /**
     * Create a Monitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: MonitorArgs | MonitorState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: MonitorState = argsOrState as MonitorState | undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["delay"] = state ? state.delay : undefined;
            inputs["expectedCodes"] = state ? state.expectedCodes : undefined;
            inputs["httpMethod"] = state ? state.httpMethod : undefined;
            inputs["maxRetries"] = state ? state.maxRetries : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["poolId"] = state ? state.poolId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["urlPath"] = state ? state.urlPath : undefined;
        } else {
            const args = argsOrState as MonitorArgs | undefined;
            if (!args || args.delay === undefined) {
                throw new Error("Missing required property 'delay'");
            }
            if (!args || args.maxRetries === undefined) {
                throw new Error("Missing required property 'maxRetries'");
            }
            if (!args || args.poolId === undefined) {
                throw new Error("Missing required property 'poolId'");
            }
            if (!args || args.timeout === undefined) {
                throw new Error("Missing required property 'timeout'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["delay"] = args ? args.delay : undefined;
            inputs["expectedCodes"] = args ? args.expectedCodes : undefined;
            inputs["httpMethod"] = args ? args.httpMethod : undefined;
            inputs["maxRetries"] = args ? args.maxRetries : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["poolId"] = args ? args.poolId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["urlPath"] = args ? args.urlPath : undefined;
        }
        super("openstack:loadbalancer/monitor:Monitor", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Monitor resources.
 */
export interface MonitorState {
    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, between sending probes to members.
     */
    readonly delay?: pulumi.Input<number>;
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * "200", or a range like "200-202".
     */
    readonly expectedCodes?: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it
     * defaults to "GET".
     */
    readonly httpMethod?: pulumi.Input<string>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to INACTIVE. Must be a number between 1
     * and 10..
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The Name of the Monitor.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The id of the pool that this monitor will be assigned to.
     */
    readonly poolId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * The type of probe, which is PING, TCP, HTTP, or HTTPS,
     * that is sent by the load balancer to verify the member state. Changing this
     * creates a new monitor.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS.
     */
    readonly urlPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Monitor resource.
 */
export interface MonitorArgs {
    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, between sending probes to members.
     */
    readonly delay: pulumi.Input<number>;
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * "200", or a range like "200-202".
     */
    readonly expectedCodes?: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it
     * defaults to "GET".
     */
    readonly httpMethod?: pulumi.Input<string>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to INACTIVE. Must be a number between 1
     * and 10..
     */
    readonly maxRetries: pulumi.Input<number>;
    /**
     * The Name of the Monitor.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The id of the pool that this monitor will be assigned to.
     */
    readonly poolId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     */
    readonly timeout: pulumi.Input<number>;
    /**
     * The type of probe, which is PING, TCP, HTTP, or HTTPS,
     * that is sent by the load balancer to verify the member state. Changing this
     * creates a new monitor.
     */
    readonly type: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS.
     */
    readonly urlPath?: pulumi.Input<string>;
}