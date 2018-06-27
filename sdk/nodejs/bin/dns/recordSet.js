"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
/**
 * Manages a DNS record set in the OpenStack DNS Service.
 */
class RecordSet extends pulumi.CustomResource {
    /**
     * Get an existing RecordSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name, id, state) {
        return new RecordSet(name, state, { id });
    }
    constructor(name, argsOrState, opts) {
        let inputs = {};
        if (opts && opts.id) {
            const state = argsOrState;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["records"] = state ? state.records : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        }
        else {
            const args = argsOrState;
            if (!args || args.zoneId === undefined) {
                throw new Error("Missing required property 'zoneId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["records"] = args ? args.records : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
        }
        super("openstack:dns/recordSet:RecordSet", name, inputs, opts);
    }
}
exports.RecordSet = RecordSet;
//# sourceMappingURL=recordSet.js.map