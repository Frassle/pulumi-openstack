# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class SiteConnection(pulumi.CustomResource):
    """
    Manages a V2 Neutron IPSec site connection resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, admin_state_up=None, description=None, dpds=None, ikepolicy_id=None, initiator=None, ipsecpolicy_id=None, local_ep_group_id=None, local_id=None, mtu=None, name=None, peer_address=None, peer_cidrs=None, peer_ep_group_id=None, peer_id=None, psk=None, region=None, tenant_id=None, value_specs=None, vpnservice_id=None):
        """Create a SiteConnection resource with the given unique name, props, and options."""
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
        Changing this updates the administrative state of the existing connection.
        """
        __props__['adminStateUp'] = admin_state_up

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        The human-readable description for the connection.
        Changing this updates the description of the existing connection.
        """
        __props__['description'] = description

        if dpds and not isinstance(dpds, list):
            raise TypeError('Expected property dpds to be a list')
        __self__.dpds = dpds
        """
        A dictionary with dead peer detection (DPD) protocol controls.
        - `action` - (Optional) The dead peer detection (DPD) action.
        A valid value is clear, hold, restart, disabled, or restart-by-peer.
        Default value is hold.
        """
        __props__['dpds'] = dpds

        if not ikepolicy_id:
            raise TypeError('Missing required property ikepolicy_id')
        elif not isinstance(ikepolicy_id, basestring):
            raise TypeError('Expected property ikepolicy_id to be a basestring')
        __self__.ikepolicy_id = ikepolicy_id
        """
        The ID of the IKE policy. Changing this creates a new connection.
        """
        __props__['ikepolicyId'] = ikepolicy_id

        if initiator and not isinstance(initiator, basestring):
            raise TypeError('Expected property initiator to be a basestring')
        __self__.initiator = initiator
        """
        A valid value is response-only or bi-directional. Default is bi-directional.
        """
        __props__['initiator'] = initiator

        if not ipsecpolicy_id:
            raise TypeError('Missing required property ipsecpolicy_id')
        elif not isinstance(ipsecpolicy_id, basestring):
            raise TypeError('Expected property ipsecpolicy_id to be a basestring')
        __self__.ipsecpolicy_id = ipsecpolicy_id
        """
        The ID of the IPsec policy. Changing this creates a new connection.
        """
        __props__['ipsecpolicyId'] = ipsecpolicy_id

        if local_ep_group_id and not isinstance(local_ep_group_id, basestring):
            raise TypeError('Expected property local_ep_group_id to be a basestring')
        __self__.local_ep_group_id = local_ep_group_id
        """
        The ID for the endpoint group that contains private subnets for the local side of the connection.
        You must specify this parameter with the peer_ep_group_id parameter unless
        in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
        Changing this updates the existing connection.
        """
        __props__['localEpGroupId'] = local_ep_group_id

        if local_id and not isinstance(local_id, basestring):
            raise TypeError('Expected property local_id to be a basestring')
        __self__.local_id = local_id
        """
        An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
        Most often, local ID would be domain name, email address, etc.
        If this is not configured then the external IP address will be used as the ID.
        """
        __props__['localId'] = local_id

        if mtu and not isinstance(mtu, int):
            raise TypeError('Expected property mtu to be a int')
        __self__.mtu = mtu
        """
        The maximum transmission unit (MTU) value to address fragmentation.
        Minimum value is 68 for IPv4, and 1280 for IPv6.
        """
        __props__['mtu'] = mtu

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the connection. Changing this updates the name of
        the existing connection.
        """
        __props__['name'] = name

        if not peer_address:
            raise TypeError('Missing required property peer_address')
        elif not isinstance(peer_address, basestring):
            raise TypeError('Expected property peer_address to be a basestring')
        __self__.peer_address = peer_address
        """
        The peer gateway public IPv4 or IPv6 address or FQDN.
        """
        __props__['peerAddress'] = peer_address

        if peer_cidrs and not isinstance(peer_cidrs, list):
            raise TypeError('Expected property peer_cidrs to be a list')
        __self__.peer_cidrs = peer_cidrs
        """
        Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
        """
        __props__['peerCidrs'] = peer_cidrs

        if peer_ep_group_id and not isinstance(peer_ep_group_id, basestring):
            raise TypeError('Expected property peer_ep_group_id to be a basestring')
        __self__.peer_ep_group_id = peer_ep_group_id
        """
        The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
        You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
        where peer_cidrs is provided with a subnet_id for the VPN service.
        """
        __props__['peerEpGroupId'] = peer_ep_group_id

        if not peer_id:
            raise TypeError('Missing required property peer_id')
        elif not isinstance(peer_id, basestring):
            raise TypeError('Expected property peer_id to be a basestring')
        __self__.peer_id = peer_id
        """
        The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
        Typically, this value matches the peer_address value.
        Changing this updates the existing policy.
        """
        __props__['peerId'] = peer_id

        if not psk:
            raise TypeError('Missing required property psk')
        elif not isinstance(psk, basestring):
            raise TypeError('Expected property psk to be a basestring')
        __self__.psk = psk
        """
        The pre-shared key. A valid value is any string.
        """
        __props__['psk'] = psk

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an IPSec site connection. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        site connection.
        """
        __props__['region'] = region

        if tenant_id and not isinstance(tenant_id, basestring):
            raise TypeError('Expected property tenant_id to be a basestring')
        __self__.tenant_id = tenant_id
        """
        The owner of the connection. Required if admin wants to
        create a connection for another project. Changing this creates a new connection.
        """
        __props__['tenantId'] = tenant_id

        if value_specs and not isinstance(value_specs, dict):
            raise TypeError('Expected property value_specs to be a dict')
        __self__.value_specs = value_specs
        """
        Map of additional options.
        """
        __props__['valueSpecs'] = value_specs

        if not vpnservice_id:
            raise TypeError('Missing required property vpnservice_id')
        elif not isinstance(vpnservice_id, basestring):
            raise TypeError('Expected property vpnservice_id to be a basestring')
        __self__.vpnservice_id = vpnservice_id
        """
        The ID of the VPN service. Changing this creates a new connection.
        """
        __props__['vpnserviceId'] = vpnservice_id

        super(SiteConnection, __self__).__init__(
            'openstack:vpnaas/siteConnection:SiteConnection',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'adminStateUp' in outs:
            self.admin_state_up = outs['adminStateUp']
        if 'description' in outs:
            self.description = outs['description']
        if 'dpds' in outs:
            self.dpds = outs['dpds']
        if 'ikepolicyId' in outs:
            self.ikepolicy_id = outs['ikepolicyId']
        if 'initiator' in outs:
            self.initiator = outs['initiator']
        if 'ipsecpolicyId' in outs:
            self.ipsecpolicy_id = outs['ipsecpolicyId']
        if 'localEpGroupId' in outs:
            self.local_ep_group_id = outs['localEpGroupId']
        if 'localId' in outs:
            self.local_id = outs['localId']
        if 'mtu' in outs:
            self.mtu = outs['mtu']
        if 'name' in outs:
            self.name = outs['name']
        if 'peerAddress' in outs:
            self.peer_address = outs['peerAddress']
        if 'peerCidrs' in outs:
            self.peer_cidrs = outs['peerCidrs']
        if 'peerEpGroupId' in outs:
            self.peer_ep_group_id = outs['peerEpGroupId']
        if 'peerId' in outs:
            self.peer_id = outs['peerId']
        if 'psk' in outs:
            self.psk = outs['psk']
        if 'region' in outs:
            self.region = outs['region']
        if 'tenantId' in outs:
            self.tenant_id = outs['tenantId']
        if 'valueSpecs' in outs:
            self.value_specs = outs['valueSpecs']
        if 'vpnserviceId' in outs:
            self.vpnservice_id = outs['vpnserviceId']
