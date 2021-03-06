# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class LoadBalancer(pulumi.CustomResource):
    """
    Manages a V2 loadbalancer resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, admin_state_up=None, description=None, flavor=None, loadbalancer_provider=None, name=None, region=None, security_group_ids=None, tenant_id=None, vip_address=None, vip_subnet_id=None):
        """Create a LoadBalancer resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError('Expected property admin_state_up to be a bool')
        __self__.admin_state_up = admin_state_up
        """
        The administrative state of the Loadbalancer.
        A valid value is true (UP) or false (DOWN).
        """
        __props__['adminStateUp'] = admin_state_up

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        Human-readable description for the Loadbalancer.
        """
        __props__['description'] = description

        if flavor and not isinstance(flavor, basestring):
            raise TypeError('Expected property flavor to be a basestring')
        __self__.flavor = flavor
        """
        The UUID of a flavor. Changing this creates a new
        loadbalancer.
        """
        __props__['flavor'] = flavor

        if loadbalancer_provider and not isinstance(loadbalancer_provider, basestring):
            raise TypeError('Expected property loadbalancer_provider to be a basestring')
        __self__.loadbalancer_provider = loadbalancer_provider
        """
        The name of the provider. Changing this
        creates a new loadbalancer.
        """
        __props__['loadbalancerProvider'] = loadbalancer_provider

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Human-readable name for the Loadbalancer. Does not have
        to be unique.
        """
        __props__['name'] = name

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB member. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB member.
        """
        __props__['region'] = region

        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError('Expected property security_group_ids to be a list')
        __self__.security_group_ids = security_group_ids
        """
        A list of security group IDs to apply to the
        loadbalancer. The security groups must be specified by ID and not name (as
        opposed to how they are configured with the Compute Instance).
        """
        __props__['securityGroupIds'] = security_group_ids

        if tenant_id and not isinstance(tenant_id, basestring):
            raise TypeError('Expected property tenant_id to be a basestring')
        __self__.tenant_id = tenant_id
        """
        Required for admins. The UUID of the tenant who owns
        the Loadbalancer.  Only administrative users can specify a tenant UUID
        other than their own.  Changing this creates a new loadbalancer.
        """
        __props__['tenantId'] = tenant_id

        if vip_address and not isinstance(vip_address, basestring):
            raise TypeError('Expected property vip_address to be a basestring')
        __self__.vip_address = vip_address
        """
        The ip address of the load balancer.
        Changing this creates a new loadbalancer.
        """
        __props__['vipAddress'] = vip_address

        if not vip_subnet_id:
            raise TypeError('Missing required property vip_subnet_id')
        elif not isinstance(vip_subnet_id, basestring):
            raise TypeError('Expected property vip_subnet_id to be a basestring')
        __self__.vip_subnet_id = vip_subnet_id
        """
        The network on which to allocate the
        Loadbalancer's address. A tenant can only create Loadbalancers on networks
        authorized by policy (e.g. networks that belong to them or networks that
        are shared).  Changing this creates a new loadbalancer.
        """
        __props__['vipSubnetId'] = vip_subnet_id

        __self__.vip_port_id = pulumi.runtime.UNKNOWN
        """
        The Port ID of the Load Balancer IP.
        """

        super(LoadBalancer, __self__).__init__(
            'openstack:loadbalancer/loadBalancer:LoadBalancer',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'adminStateUp' in outs:
            self.admin_state_up = outs['adminStateUp']
        if 'description' in outs:
            self.description = outs['description']
        if 'flavor' in outs:
            self.flavor = outs['flavor']
        if 'loadbalancerProvider' in outs:
            self.loadbalancer_provider = outs['loadbalancerProvider']
        if 'name' in outs:
            self.name = outs['name']
        if 'region' in outs:
            self.region = outs['region']
        if 'securityGroupIds' in outs:
            self.security_group_ids = outs['securityGroupIds']
        if 'tenantId' in outs:
            self.tenant_id = outs['tenantId']
        if 'vipAddress' in outs:
            self.vip_address = outs['vipAddress']
        if 'vipPortId' in outs:
            self.vip_port_id = outs['vipPortId']
        if 'vipSubnetId' in outs:
            self.vip_subnet_id = outs['vipSubnetId']
