"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
/**
 * Manages a V2 router interface resource within OpenStack.
 */
class RouterInterface extends pulumi.CustomResource {
    /**
     * Get an existing RouterInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name, id, state) {
        return new RouterInterface(name, state, { id });
    }
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["portId"] = state ? state.portId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["routerId"] = state ? state.routerId : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        }
        else {
            const args = argsOrState;
            if (!args || args.routerId === undefined) {
                throw new Error("Missing required property 'routerId'");
            }
            inputs["portId"] = args ? args.portId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["routerId"] = args ? args.routerId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
        }
        super("openstack:networking/routerInterface:RouterInterface", name, inputs, opts);
    }
}
exports.RouterInterface = RouterInterface;
//# sourceMappingURL=routerInterface.js.map