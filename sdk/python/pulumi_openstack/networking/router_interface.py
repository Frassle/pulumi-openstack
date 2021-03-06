# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class RouterInterface(pulumi.CustomResource):
    """
    Manages a V2 router interface resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, port_id=None, region=None, router_id=None, subnet_id=None):
        """Create a RouterInterface resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if port_id and not isinstance(port_id, basestring):
            raise TypeError('Expected property port_id to be a basestring')
        __self__.port_id = port_id
        """
        ID of the port this interface connects to. Changing
        this creates a new router interface.
        """
        __props__['portId'] = port_id

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a router. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        router interface.
        """
        __props__['region'] = region

        if not router_id:
            raise TypeError('Missing required property router_id')
        elif not isinstance(router_id, basestring):
            raise TypeError('Expected property router_id to be a basestring')
        __self__.router_id = router_id
        """
        ID of the router this interface belongs to. Changing
        this creates a new router interface.
        """
        __props__['routerId'] = router_id

        if subnet_id and not isinstance(subnet_id, basestring):
            raise TypeError('Expected property subnet_id to be a basestring')
        __self__.subnet_id = subnet_id
        """
        ID of the subnet this interface connects to. Changing
        this creates a new router interface.
        """
        __props__['subnetId'] = subnet_id

        super(RouterInterface, __self__).__init__(
            'openstack:networking/routerInterface:RouterInterface',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'portId' in outs:
            self.port_id = outs['portId']
        if 'region' in outs:
            self.region = outs['region']
        if 'routerId' in outs:
            self.router_id = outs['routerId']
        if 'subnetId' in outs:
            self.subnet_id = outs['subnetId']
