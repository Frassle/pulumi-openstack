"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
/**
 * Manages a V2 neutron security group resource within OpenStack.
 * Unlike Nova security groups, neutron separates the group from the rules
 * and also allows an admin to target a specific tenant_id.
 */
class SecGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name, id, state) {
        return new SecGroup(name, state, { id });
    }
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["deleteDefaultRules"] = state ? state.deleteDefaultRules : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        }
        else {
            const args = argsOrState;
            inputs["deleteDefaultRules"] = args ? args.deleteDefaultRules : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        super("openstack:networking/secGroup:SecGroup", name, inputs, opts);
    }
}
exports.SecGroup = SecGroup;
//# sourceMappingURL=secGroup.js.map