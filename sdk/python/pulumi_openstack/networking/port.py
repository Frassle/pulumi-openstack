# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Port(pulumi.CustomResource):
    """
    Manages a V2 port resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, admin_state_up=None, allowed_address_pairs=None, device_id=None, device_owner=None, fixed_ips=None, mac_address=None, name=None, network_id=None, no_security_groups=None, region=None, security_group_ids=None, tenant_id=None, value_specs=None):
        """Create a Port resource with the given unique name, props, and options."""
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
        Administrative up/down status for the port
        (must be "true" or "false" if provided). Changing this updates the
        `admin_state_up` of an existing port.
        """
        __props__['adminStateUp'] = admin_state_up

        if allowed_address_pairs and not isinstance(allowed_address_pairs, list):
            raise TypeError('Expected property allowed_address_pairs to be a list')
        __self__.allowed_address_pairs = allowed_address_pairs
        """
        An IP/MAC Address pair of additional IP
        addresses that can be active on this port. The structure is described
        below.
        """
        __props__['allowedAddressPairs'] = allowed_address_pairs

        if device_id and not isinstance(device_id, basestring):
            raise TypeError('Expected property device_id to be a basestring')
        __self__.device_id = device_id
        """
        The ID of the device attached to the port. Changing this
        creates a new port.
        """
        __props__['deviceId'] = device_id

        if device_owner and not isinstance(device_owner, basestring):
            raise TypeError('Expected property device_owner to be a basestring')
        __self__.device_owner = device_owner
        """
        The device owner of the Port. Changing this creates
        a new port.
        """
        __props__['deviceOwner'] = device_owner

        if fixed_ips and not isinstance(fixed_ips, list):
            raise TypeError('Expected property fixed_ips to be a list')
        __self__.fixed_ips = fixed_ips
        """
        An array of desired IPs for this port. The structure is
        described below.
        """
        __props__['fixedIps'] = fixed_ips

        if mac_address and not isinstance(mac_address, basestring):
            raise TypeError('Expected property mac_address to be a basestring')
        __self__.mac_address = mac_address
        """
        The additional MAC address.
        """
        __props__['macAddress'] = mac_address

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        A unique name for the port. Changing this
        updates the `name` of an existing port.
        """
        __props__['name'] = name

        if not network_id:
            raise TypeError('Missing required property network_id')
        elif not isinstance(network_id, basestring):
            raise TypeError('Expected property network_id to be a basestring')
        __self__.network_id = network_id
        """
        The ID of the network to attach the port to. Changing
        this creates a new port.
        """
        __props__['networkId'] = network_id

        if no_security_groups and not isinstance(no_security_groups, bool):
            raise TypeError('Expected property no_security_groups to be a bool')
        __self__.no_security_groups = no_security_groups
        """
        If set to
        `true`, then no security groups are applied to the port. If set to `false` and
        no `security_group_ids` are specified, then the Port will yield to the default
        behavior of the Networking service, which is to usually apply the "default"
        security group.
        """
        __props__['noSecurityGroups'] = no_security_groups

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a port. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        port.
        """
        __props__['region'] = region

        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError('Expected property security_group_ids to be a list')
        __self__.security_group_ids = security_group_ids
        """
        A list
        of security group IDs to apply to the port. The security groups must be
        specified by ID and not name (as opposed to how they are configured with
        the Compute Instance).
        """
        __props__['securityGroupIds'] = security_group_ids

        if tenant_id and not isinstance(tenant_id, basestring):
            raise TypeError('Expected property tenant_id to be a basestring')
        __self__.tenant_id = tenant_id
        """
        The owner of the Port. Required if admin wants
        to create a port for another tenant. Changing this creates a new port.
        """
        __props__['tenantId'] = tenant_id

        if value_specs and not isinstance(value_specs, dict):
            raise TypeError('Expected property value_specs to be a dict')
        __self__.value_specs = value_specs
        """
        Map of additional options.
        """
        __props__['valueSpecs'] = value_specs

        __self__.all_fixed_ips = pulumi.runtime.UNKNOWN
        """
        The collection of Fixed IP addresses on the port in the
        order returned by the Network v2 API.
        """
        __self__.all_security_group_ids = pulumi.runtime.UNKNOWN
        """
        The collection of Security Group IDs on the port
        which have been explicitly and implicitly added.
        """

        super(Port, __self__).__init__(
            'openstack:networking/port:Port',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'adminStateUp' in outs:
            self.admin_state_up = outs['adminStateUp']
        if 'allFixedIps' in outs:
            self.all_fixed_ips = outs['allFixedIps']
        if 'allSecurityGroupIds' in outs:
            self.all_security_group_ids = outs['allSecurityGroupIds']
        if 'allowedAddressPairs' in outs:
            self.allowed_address_pairs = outs['allowedAddressPairs']
        if 'deviceId' in outs:
            self.device_id = outs['deviceId']
        if 'deviceOwner' in outs:
            self.device_owner = outs['deviceOwner']
        if 'fixedIps' in outs:
            self.fixed_ips = outs['fixedIps']
        if 'macAddress' in outs:
            self.mac_address = outs['macAddress']
        if 'name' in outs:
            self.name = outs['name']
        if 'networkId' in outs:
            self.network_id = outs['networkId']
        if 'noSecurityGroups' in outs:
            self.no_security_groups = outs['noSecurityGroups']
        if 'region' in outs:
            self.region = outs['region']
        if 'securityGroupIds' in outs:
            self.security_group_ids = outs['securityGroupIds']
        if 'tenantId' in outs:
            self.tenant_id = outs['tenantId']
        if 'valueSpecs' in outs:
            self.value_specs = outs['valueSpecs']
