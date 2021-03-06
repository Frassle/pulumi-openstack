# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Role(pulumi.CustomResource):
    """
    Manages a V3 Role resource within OpenStack Keystone.
    
    Note: You _must_ have admin privileges in your OpenStack cloud to use
    this resource.
    """
    def __init__(__self__, __name__, __opts__=None, domain_id=None, name=None, region=None):
        """Create a Role resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if domain_id and not isinstance(domain_id, basestring):
            raise TypeError('Expected property domain_id to be a basestring')
        __self__.domain_id = domain_id
        """
        The domain the role belongs to.
        """
        __props__['domainId'] = domain_id

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the role.
        """
        __props__['name'] = name

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new Role.
        """
        __props__['region'] = region

        super(Role, __self__).__init__(
            'openstack:identity/role:Role',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'domainId' in outs:
            self.domain_id = outs['domainId']
        if 'name' in outs:
            self.name = outs['name']
        if 'region' in outs:
            self.region = outs['region']
