# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Service(pulumi.CustomResource):
    """
    Manages a V2 Neutron VPN service resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, admin_state_up=None, description=None, name=None, region=None, router_id=None, subnet_id=None, tenant_id=None, value_specs=None):
        """Create a Service resource with the given unique name, props, and options."""
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
        The administrative state of the resource. Can either be up(true) or down(false).
        Changing this updates the administrative state of the existing service.
        """
        __props__['adminStateUp'] = admin_state_up

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        The human-readable description for the service.
        Changing this updates the description of the existing service.
        """
        __props__['description'] = description

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the service. Changing this updates the name of
        the existing service.
        """
        __props__['name'] = name

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a VPN service. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        service.
        """
        __props__['region'] = region

        if not router_id:
            raise TypeError('Missing required property router_id')
        elif not isinstance(router_id, basestring):
            raise TypeError('Expected property router_id to be a basestring')
        __self__.router_id = router_id
        """
        The ID of the router. Changing this creates a new service.
        """
        __props__['routerId'] = router_id

        if subnet_id and not isinstance(subnet_id, basestring):
            raise TypeError('Expected property subnet_id to be a basestring')
        __self__.subnet_id = subnet_id
        """
        SubnetID is the ID of the subnet. Default is null.
        """
        __props__['subnetId'] = subnet_id

        if tenant_id and not isinstance(tenant_id, basestring):
            raise TypeError('Expected property tenant_id to be a basestring')
        __self__.tenant_id = tenant_id
        """
        The owner of the service. Required if admin wants to
        create a service for another project. Changing this creates a new service.
        """
        __props__['tenantId'] = tenant_id

        if value_specs and not isinstance(value_specs, dict):
            raise TypeError('Expected property value_specs to be a dict')
        __self__.value_specs = value_specs
        """
        Map of additional options.
        """
        __props__['valueSpecs'] = value_specs

        __self__.external_v4_ip = pulumi.runtime.UNKNOWN
        """
        The read-only external (public) IPv4 address that is used for the VPN service.
        """
        __self__.external_v6_ip = pulumi.runtime.UNKNOWN
        """
        The read-only external (public) IPv6 address that is used for the VPN service.
        """
        __self__.status = pulumi.runtime.UNKNOWN
        """
        Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
        """

        super(Service, __self__).__init__(
            'openstack:vpnaas/service:Service',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'adminStateUp' in outs:
            self.admin_state_up = outs['adminStateUp']
        if 'description' in outs:
            self.description = outs['description']
        if 'externalV4Ip' in outs:
            self.external_v4_ip = outs['externalV4Ip']
        if 'externalV6Ip' in outs:
            self.external_v6_ip = outs['externalV6Ip']
        if 'name' in outs:
            self.name = outs['name']
        if 'region' in outs:
            self.region = outs['region']
        if 'routerId' in outs:
            self.router_id = outs['routerId']
        if 'status' in outs:
            self.status = outs['status']
        if 'subnetId' in outs:
            self.subnet_id = outs['subnetId']
        if 'tenantId' in outs:
            self.tenant_id = outs['tenantId']
        if 'valueSpecs' in outs:
            self.value_specs = outs['valueSpecs']
