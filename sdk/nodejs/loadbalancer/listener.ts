// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a V2 listener resource within OpenStack.
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState): Listener {
        return new Listener(name, <any>state, { id });
    }

    /**
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp: pulumi.Output<boolean | undefined>;
    /**
     * The maximum number of connections allowed
     * for the Listener.
     */
    public readonly connectionLimit: pulumi.Output<number>;
    /**
     * The ID of the default pool with which the
     * Listener is associated. Changing this creates a new Listener.
     */
    public readonly defaultPoolId: pulumi.Output<string>;
    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    public readonly defaultTlsContainerRef: pulumi.Output<string | undefined>;
    /**
     * Human-readable description for the Listener.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     */
    public readonly loadbalancerId: pulumi.Output<string>;
    /**
     * Human-readable name for the Listener. Does not have
     * to be unique.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
     * Changing this creates a new Listener.
     */
    public readonly protocol: pulumi.Output<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     */
    public readonly protocolPort: pulumi.Output<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    public readonly sniContainerRefs: pulumi.Output<string[] | undefined>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     */
    public readonly tenantId: pulumi.Output<string>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ListenerState = argsOrState as ListenerState | undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["connectionLimit"] = state ? state.connectionLimit : undefined;
            inputs["defaultPoolId"] = state ? state.defaultPoolId : undefined;
            inputs["defaultTlsContainerRef"] = state ? state.defaultTlsContainerRef : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["loadbalancerId"] = state ? state.loadbalancerId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["protocolPort"] = state ? state.protocolPort : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["sniContainerRefs"] = state ? state.sniContainerRefs : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if (!args || args.loadbalancerId === undefined) {
                throw new Error("Missing required property 'loadbalancerId'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.protocolPort === undefined) {
                throw new Error("Missing required property 'protocolPort'");
            }
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["connectionLimit"] = args ? args.connectionLimit : undefined;
            inputs["defaultPoolId"] = args ? args.defaultPoolId : undefined;
            inputs["defaultTlsContainerRef"] = args ? args.defaultTlsContainerRef : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["loadbalancerId"] = args ? args.loadbalancerId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["protocolPort"] = args ? args.protocolPort : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["sniContainerRefs"] = args ? args.sniContainerRefs : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        super("openstack:loadbalancer/listener:Listener", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The maximum number of connections allowed
     * for the Listener.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * The ID of the default pool with which the
     * Listener is associated. Changing this creates a new Listener.
     */
    readonly defaultPoolId?: pulumi.Input<string>;
    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly defaultTlsContainerRef?: pulumi.Input<string>;
    /**
     * Human-readable description for the Listener.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     */
    readonly loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the Listener. Does not have
     * to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
     * Changing this creates a new Listener.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     */
    readonly protocolPort?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly sniContainerRefs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The maximum number of connections allowed
     * for the Listener.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * The ID of the default pool with which the
     * Listener is associated. Changing this creates a new Listener.
     */
    readonly defaultPoolId?: pulumi.Input<string>;
    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly defaultTlsContainerRef?: pulumi.Input<string>;
    /**
     * Human-readable description for the Listener.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     */
    readonly loadbalancerId: pulumi.Input<string>;
    /**
     * Human-readable name for the Listener. Does not have
     * to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
     * Changing this creates a new Listener.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     */
    readonly protocolPort: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly sniContainerRefs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     */
    readonly tenantId?: pulumi.Input<string>;
}
