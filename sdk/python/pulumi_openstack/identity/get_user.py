# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetUserResult(object):
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, default_project_id=None, domain_id=None, region=None, id=None):
        if default_project_id and not isinstance(default_project_id, basestring):
            raise TypeError('Expected argument default_project_id to be a basestring')
        __self__.default_project_id = default_project_id
        """
        See Argument Reference above.
        """
        if domain_id and not isinstance(domain_id, basestring):
            raise TypeError('Expected argument domain_id to be a basestring')
        __self__.domain_id = domain_id
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, basestring):
            raise TypeError('Expected argument region to be a basestring')
        __self__.region = region
        """
        The region the user is located in.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_user(domain_id=None, enabled=None, idp_id=None, name=None, password_expires_at=None, protocol_id=None, region=None, unique_id=None):
    """
    Use this data source to get the ID of an OpenStack user.
    """
    __args__ = dict()

    __args__['domainId'] = domain_id
    __args__['enabled'] = enabled
    __args__['idpId'] = idp_id
    __args__['name'] = name
    __args__['passwordExpiresAt'] = password_expires_at
    __args__['protocolId'] = protocol_id
    __args__['region'] = region
    __args__['uniqueId'] = unique_id
    __ret__ = pulumi.runtime.invoke('openstack:identity/getUser:getUser', __args__)

    return GetUserResult(
        default_project_id=__ret__.get('defaultProjectId'),
        domain_id=__ret__.get('domainId'),
        region=__ret__.get('region'),
        id=__ret__.get('id'))
