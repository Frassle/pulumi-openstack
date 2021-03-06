# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Policy(pulumi.CustomResource):
    """
    Manages a v1 firewall policy resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, audited=None, description=None, name=None, region=None, rules=None, shared=None, tenant_id=None, value_specs=None):
        """Create a Policy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if audited and not isinstance(audited, bool):
            raise TypeError('Expected property audited to be a bool')
        __self__.audited = audited
        """
        Audit status of the firewall policy
        (must be "true" or "false" if provided - defaults to "false").
        This status is set to "false" whenever the firewall policy or any of its
        rules are changed. Changing this updates the `audited` status of an existing
        firewall policy.
        """
        __props__['audited'] = audited

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        A description for the firewall policy. Changing
        this updates the `description` of an existing firewall policy.
        """
        __props__['description'] = description

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        A name for the firewall policy. Changing this
        updates the `name` of an existing firewall policy.
        """
        __props__['name'] = name

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the v1 networking client.
        A networking client is needed to create a firewall policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall policy.
        """
        __props__['region'] = region

        if rules and not isinstance(rules, list):
            raise TypeError('Expected property rules to be a list')
        __self__.rules = rules
        """
        An array of one or more firewall rules that comprise
        the policy. Changing this results in adding/removing rules from the
        existing firewall policy.
        """
        __props__['rules'] = rules

        if shared and not isinstance(shared, bool):
            raise TypeError('Expected property shared to be a bool')
        __self__.shared = shared
        """
        Sharing status of the firewall policy (must be "true"
        or "false" if provided). If this is "true" the policy is visible to, and
        can be used in, firewalls in other tenants. Changing this updates the
        `shared` status of an existing firewall policy. Only administrative users
        can specify if the policy should be shared.
        """
        __props__['shared'] = shared

        if tenant_id and not isinstance(tenant_id, basestring):
            raise TypeError('Expected property tenant_id to be a basestring')
        __self__.tenant_id = tenant_id
        __props__['tenantId'] = tenant_id

        if value_specs and not isinstance(value_specs, dict):
            raise TypeError('Expected property value_specs to be a dict')
        __self__.value_specs = value_specs
        """
        Map of additional options.
        """
        __props__['valueSpecs'] = value_specs

        super(Policy, __self__).__init__(
            'openstack:firewall/policy:Policy',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'audited' in outs:
            self.audited = outs['audited']
        if 'description' in outs:
            self.description = outs['description']
        if 'name' in outs:
            self.name = outs['name']
        if 'region' in outs:
            self.region = outs['region']
        if 'rules' in outs:
            self.rules = outs['rules']
        if 'shared' in outs:
            self.shared = outs['shared']
        if 'tenantId' in outs:
            self.tenant_id = outs['tenantId']
        if 'valueSpecs' in outs:
            self.value_specs = outs['valueSpecs']
