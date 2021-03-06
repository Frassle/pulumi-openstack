# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetAuthScopeResult(object):
    """
    A collection of values returned by getAuthScope.
    """
    def __init__(__self__, project_domain_id=None, project_domain_name=None, project_id=None, project_name=None, region=None, roles=None, user_domain_id=None, user_domain_name=None, user_id=None, user_name=None, id=None):
        if project_domain_id and not isinstance(project_domain_id, basestring):
            raise TypeError('Expected argument project_domain_id to be a basestring')
        __self__.project_domain_id = project_domain_id
        """
        The domain ID of the project.
        """
        if project_domain_name and not isinstance(project_domain_name, basestring):
            raise TypeError('Expected argument project_domain_name to be a basestring')
        __self__.project_domain_name = project_domain_name
        """
        The domain name of the project.
        """
        if project_id and not isinstance(project_id, basestring):
            raise TypeError('Expected argument project_id to be a basestring')
        __self__.project_id = project_id
        """
        The project ID of the scope.
        """
        if project_name and not isinstance(project_name, basestring):
            raise TypeError('Expected argument project_name to be a basestring')
        __self__.project_name = project_name
        """
        The project name of the scope.
        """
        if region and not isinstance(region, basestring):
            raise TypeError('Expected argument region to be a basestring')
        __self__.region = region
        if roles and not isinstance(roles, list):
            raise TypeError('Expected argument roles to be a list')
        __self__.roles = roles
        """
        A list of roles in the current scope. See reference below.
        """
        if user_domain_id and not isinstance(user_domain_id, basestring):
            raise TypeError('Expected argument user_domain_id to be a basestring')
        __self__.user_domain_id = user_domain_id
        """
        The domain ID of the user.
        """
        if user_domain_name and not isinstance(user_domain_name, basestring):
            raise TypeError('Expected argument user_domain_name to be a basestring')
        __self__.user_domain_name = user_domain_name
        """
        The domain name of the user.
        """
        if user_id and not isinstance(user_id, basestring):
            raise TypeError('Expected argument user_id to be a basestring')
        __self__.user_id = user_id
        """
        The user ID the of the scope.
        """
        if user_name and not isinstance(user_name, basestring):
            raise TypeError('Expected argument user_name to be a basestring')
        __self__.user_name = user_name
        """
        The username of the scope.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_auth_scope(name=None, region=None):
    """
    Use this data source to get authentication information about the current
    auth scope in use. This can be used as self-discovery or introspection of
    the username or project name currently in use.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['region'] = region
    __ret__ = pulumi.runtime.invoke('openstack:identity/getAuthScope:getAuthScope', __args__)

    return GetAuthScopeResult(
        project_domain_id=__ret__.get('projectDomainId'),
        project_domain_name=__ret__.get('projectDomainName'),
        project_id=__ret__.get('projectId'),
        project_name=__ret__.get('projectName'),
        region=__ret__.get('region'),
        roles=__ret__.get('roles'),
        user_domain_id=__ret__.get('userDomainId'),
        user_domain_name=__ret__.get('userDomainName'),
        user_id=__ret__.get('userId'),
        user_name=__ret__.get('userName'),
        id=__ret__.get('id'))
