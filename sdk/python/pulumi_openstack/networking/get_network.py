# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetNetworkResult(object):
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, admin_state_up=None, availability_zone_hints=None, region=None, shared=None, id=None):
        if admin_state_up and not isinstance(admin_state_up, basestring):
            raise TypeError('Expected argument admin_state_up to be a basestring')
        __self__.admin_state_up = admin_state_up
        """
        (Optional) The administrative state of the network.
        """
        if availability_zone_hints and not isinstance(availability_zone_hints, list):
            raise TypeError('Expected argument availability_zone_hints to be a list')
        __self__.availability_zone_hints = availability_zone_hints
        """
        (Optional) The availability zone candidates for the network.
        """
        if region and not isinstance(region, basestring):
            raise TypeError('Expected argument region to be a basestring')
        __self__.region = region
        """
        See Argument Reference above.
        """
        if shared and not isinstance(shared, basestring):
            raise TypeError('Expected argument shared to be a basestring')
        __self__.shared = shared
        """
        (Optional)  Specifies whether the network resource can be accessed
        by any tenant or not.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_network(external=None, matching_subnet_cidr=None, name=None, network_id=None, region=None, status=None, tenant_id=None):
    """
    Use this data source to get the ID of an available OpenStack network.
    """
    __args__ = dict()

    __args__['external'] = external
    __args__['matchingSubnetCidr'] = matching_subnet_cidr
    __args__['name'] = name
    __args__['networkId'] = network_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['tenantId'] = tenant_id
    __ret__ = pulumi.runtime.invoke('openstack:networking/getNetwork:getNetwork', __args__)

    return GetNetworkResult(
        admin_state_up=__ret__.get('adminStateUp'),
        availability_zone_hints=__ret__.get('availabilityZoneHints'),
        region=__ret__.get('region'),
        shared=__ret__.get('shared'),
        id=__ret__.get('id'))
