// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Networking {

	public sealed class PortArgsAllowedAddressPair : Pulumi.IIOProtobuf {
		public Pulumi.IO<string> IpAddress { get; set; }
		public Pulumi.IO<string> MacAddress { get; set; }

		public Pulumi.IO<Google.Protobuf.WellKnownTypes.Value> ToProtobuf() {
			return Pulumi.Protobuf.ToProtobuf(
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("ip_address", Protobuf.ToProtobuf(IpAddress)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("mac_address", Protobuf.ToProtobuf(MacAddress)));
		} // ToProtobuf

	} // PortArgsAllowedAddressPair

	public sealed class PortArgsFixedIp : Pulumi.IIOProtobuf {
		public Pulumi.IO<string> IpAddress { get; set; }
		public Pulumi.IO<string> SubnetId { get; set; }

		public Pulumi.IO<Google.Protobuf.WellKnownTypes.Value> ToProtobuf() {
			return Pulumi.Protobuf.ToProtobuf(
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("ip_address", Protobuf.ToProtobuf(IpAddress)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("subnet_id", Protobuf.ToProtobuf(SubnetId)));
		} // ToProtobuf

	} // PortArgsFixedIp

	public sealed class PortAllowedAddressPair : Pulumi.IIOProtobuf {
		public Pulumi.IO<string> IpAddress { get; set; }
		public Pulumi.IO<string> MacAddress { get; set; }

		public Pulumi.IO<Google.Protobuf.WellKnownTypes.Value> ToProtobuf() {
			return Pulumi.Protobuf.ToProtobuf(
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("ip_address", Protobuf.ToProtobuf(IpAddress)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("mac_address", Protobuf.ToProtobuf(MacAddress)));
		} // ToProtobuf

		public static PortAllowedAddressPair FromProtobuf(Google.Protobuf.WellKnownTypes.Value value) {
			var obj = value.StructValue;
			return new PortAllowedAddressPair() {
				IpAddress = Protobuf.ToString(obj.Fields["ip_address"]),
				MacAddress = Protobuf.ToString(obj.Fields["mac_address"]),
			};
		} // FromProtobuf

	} // PortAllowedAddressPair

	public sealed class PortFixedIp : Pulumi.IIOProtobuf {
		public Pulumi.IO<string> IpAddress { get; set; }
		public Pulumi.IO<string> SubnetId { get; set; }

		public Pulumi.IO<Google.Protobuf.WellKnownTypes.Value> ToProtobuf() {
			return Pulumi.Protobuf.ToProtobuf(
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("ip_address", Protobuf.ToProtobuf(IpAddress)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("subnet_id", Protobuf.ToProtobuf(SubnetId)));
		} // ToProtobuf

		public static PortFixedIp FromProtobuf(Google.Protobuf.WellKnownTypes.Value value) {
			var obj = value.StructValue;
			return new PortFixedIp() {
				IpAddress = Protobuf.ToString(obj.Fields["ip_address"]),
				SubnetId = Protobuf.ToString(obj.Fields["subnet_id"]),
			};
		} // FromProtobuf

	} // PortFixedIp

	/// <summary>
	/// The set of arguments for constructing a Port resource.
	/// </summary>
	public struct PortArgs {
		/// <summary>
		/// Administrative up/down status for the port
		/// (must be "true" or "false" if provided). Changing this updates the
		/// `admin_state_up` of an existing port.
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// An IP/MAC Address pair of additional IP
		/// addresses that can be active on this port. The structure is described
		/// below.
		/// </summary>
		public Pulumi.IO<Pulumi.IO<PortArgsAllowedAddressPair>[]> AllowedAddressPairs { get; set; }

		/// <summary>
		/// The ID of the device attached to the port. Changing this
		/// creates a new port.
		/// </summary>
		public Pulumi.IO<string> DeviceId { get; set; }

		/// <summary>
		/// The device owner of the Port. Changing this creates
		/// a new port.
		/// </summary>
		public Pulumi.IO<string> DeviceOwner { get; set; }

		/// <summary>
		/// An array of desired IPs for this port. The structure is
		/// described below.
		/// </summary>
		public Pulumi.IO<Pulumi.IO<PortArgsFixedIp>[]> FixedIps { get; set; }

		/// <summary>
		/// Specify a specific MAC address for the port. Changing
		/// this creates a new port.
		/// </summary>
		public Pulumi.IO<string> MacAddress { get; set; }

		/// <summary>
		/// A unique name for the port. Changing this
		/// updates the `name` of an existing port.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The ID of the network to attach the port to. Changing
		/// this creates a new port.
		/// </summary>
		public Pulumi.IO<string> NetworkId { get; set; }

		/// <summary>
		/// If set to
		/// `true`, then no security groups are applied to the port. If set to `false` and
		/// no `security_group_ids` are specified, then the Port will yield to the default
		/// behavior of the Networking service, which is to usually apply the "default"
		/// security group.
		/// </summary>
		public Pulumi.IO<bool> NoSecurityGroups { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 networking client.
		/// A networking client is needed to create a port. If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// port.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// A list
		/// of security group IDs to apply to the port. The security groups must be
		/// specified by ID and not name (as opposed to how they are configured with
		/// the Compute Instance).
		/// </summary>
		public Pulumi.IO<Pulumi.IO<string>[]> SecurityGroupIds { get; set; }

		/// <summary>
		/// The owner of the Port. Required if admin wants
		/// to create a port for another tenant. Changing this creates a new port.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// Map of additional options.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> ValueSpecs { get; set; }

	} // PortArgs

	/// <summary>
	/// Manages a V2 port resource within OpenStack.
	/// </summary>
	public class Port : Pulumi.CustomResource {
		/// <summary>
		/// Administrative up/down status for the port
		/// (must be "true" or "false" if provided). Changing this updates the
		/// `admin_state_up` of an existing port.
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// The collection of Fixed IP addresses on the port in the
		/// order returned by the Network v2 API.
		/// </summary>
		public Pulumi.IO<string[]> AllFixedIps { get; set; }

		/// <summary>
		/// The collection of Security Group IDs on the port
		/// which have been explicitly and implicitly added.
		/// </summary>
		public Pulumi.IO<string[]> AllSecurityGroupIds { get; set; }

		/// <summary>
		/// An IP/MAC Address pair of additional IP
		/// addresses that can be active on this port. The structure is described
		/// below.
		/// </summary>
		public Pulumi.IO<PortAllowedAddressPair[]> AllowedAddressPairs { get; set; }

		/// <summary>
		/// The ID of the device attached to the port. Changing this
		/// creates a new port.
		/// </summary>
		public Pulumi.IO<string> DeviceId { get; set; }

		/// <summary>
		/// The device owner of the Port. Changing this creates
		/// a new port.
		/// </summary>
		public Pulumi.IO<string> DeviceOwner { get; set; }

		/// <summary>
		/// An array of desired IPs for this port. The structure is
		/// described below.
		/// </summary>
		public Pulumi.IO<PortFixedIp[]> FixedIps { get; set; }

		/// <summary>
		/// Specify a specific MAC address for the port. Changing
		/// this creates a new port.
		/// </summary>
		public Pulumi.IO<string> MacAddress { get; set; }

		/// <summary>
		/// A unique name for the port. Changing this
		/// updates the `name` of an existing port.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The ID of the network to attach the port to. Changing
		/// this creates a new port.
		/// </summary>
		public Pulumi.IO<string> NetworkId { get; set; }

		/// <summary>
		/// If set to
		/// `true`, then no security groups are applied to the port. If set to `false` and
		/// no `security_group_ids` are specified, then the Port will yield to the default
		/// behavior of the Networking service, which is to usually apply the "default"
		/// security group.
		/// </summary>
		public Pulumi.IO<bool> NoSecurityGroups { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 networking client.
		/// A networking client is needed to create a port. If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// port.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// A list
		/// of security group IDs to apply to the port. The security groups must be
		/// specified by ID and not name (as opposed to how they are configured with
		/// the Compute Instance).
		/// </summary>
		public Pulumi.IO<string[]> SecurityGroupIds { get; set; }

		/// <summary>
		/// The owner of the Port. Required if admin wants
		/// to create a port for another tenant. Changing this creates a new port.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// Map of additional options.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> ValueSpecs { get; set; }

		public Port(string name, PortArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:networking/port:Port", name, SerialiseArgs(args), opts) {
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			AllowedAddressPairs = Outputs["allowedAddressPairs"].Select(item => Protobuf.ToList(item, item1 => PortAllowedAddressPair.FromProtobuf(item1)));
			DeviceId = Outputs["deviceId"].Select(item => Protobuf.ToString(item));
			DeviceOwner = Outputs["deviceOwner"].Select(item => Protobuf.ToString(item));
			FixedIps = Outputs["fixedIps"].Select(item => Protobuf.ToList(item, item1 => PortFixedIp.FromProtobuf(item1)));
			MacAddress = Outputs["macAddress"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			NetworkId = Outputs["networkId"].Select(item => Protobuf.ToString(item));
			NoSecurityGroups = Outputs["noSecurityGroups"].Select(item => Protobuf.ToBool(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			SecurityGroupIds = Outputs["securityGroupIds"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			ValueSpecs = Outputs["valueSpecs"].Select(item => Protobuf.ToMap(item));
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			AllFixedIps = Outputs["allFixedIps"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			AllSecurityGroupIds = Outputs["allSecurityGroupIds"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			AllowedAddressPairs = Outputs["allowedAddressPairs"].Select(item => Protobuf.ToList(item, item1 => PortAllowedAddressPair.FromProtobuf(item1)));
			DeviceId = Outputs["deviceId"].Select(item => Protobuf.ToString(item));
			DeviceOwner = Outputs["deviceOwner"].Select(item => Protobuf.ToString(item));
			FixedIps = Outputs["fixedIps"].Select(item => Protobuf.ToList(item, item1 => PortFixedIp.FromProtobuf(item1)));
			MacAddress = Outputs["macAddress"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			NetworkId = Outputs["networkId"].Select(item => Protobuf.ToString(item));
			NoSecurityGroups = Outputs["noSecurityGroups"].Select(item => Protobuf.ToBool(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			SecurityGroupIds = Outputs["securityGroupIds"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			ValueSpecs = Outputs["valueSpecs"].Select(item => Protobuf.ToMap(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(PortArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["adminStateUp"] = Protobuf.ToProtobuf(args.AdminStateUp);
			props["allowedAddressPairs"] = Protobuf.ToProtobuf(args.AllowedAddressPairs, item => Protobuf.ToProtobuf(item));
			props["deviceId"] = Protobuf.ToProtobuf(args.DeviceId);
			props["deviceOwner"] = Protobuf.ToProtobuf(args.DeviceOwner);
			props["fixedIps"] = Protobuf.ToProtobuf(args.FixedIps, item => Protobuf.ToProtobuf(item));
			props["macAddress"] = Protobuf.ToProtobuf(args.MacAddress);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["networkId"] = Protobuf.ToProtobuf(args.NetworkId);
			props["noSecurityGroups"] = Protobuf.ToProtobuf(args.NoSecurityGroups);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["securityGroupIds"] = Protobuf.ToProtobuf(args.SecurityGroupIds, item => Protobuf.ToProtobuf(item));
			props["tenantId"] = Protobuf.ToProtobuf(args.TenantId);
			props["valueSpecs"] = Protobuf.ToProtobuf(args.ValueSpecs);
			props["allFixedIps"] = null; //out
			props["allSecurityGroupIds"] = null; //out
			return props;
		} // SerialiseArgs

	} // Port
} // Pulumi.Openstack.Networking
