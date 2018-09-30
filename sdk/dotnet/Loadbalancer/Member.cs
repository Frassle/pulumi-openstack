// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Loadbalancer {

	/// <summary>
	/// The set of arguments for constructing a Member resource.
	/// </summary>
	public struct MemberArgs {
		/// <summary>
		/// The IP address of the member to receive traffic from
		/// the load balancer. Changing this creates a new member.
		/// </summary>
		public Pulumi.IO<string> Address { get; set; }

		/// <summary>
		/// The administrative state of the member.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// Human-readable name for the member.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The id of the pool that this member will be
		/// assigned to.
		/// </summary>
		public Pulumi.IO<string> PoolId { get; set; }

		/// <summary>
		/// The port on which to listen for client traffic.
		/// Changing this creates a new member.
		/// </summary>
		public Pulumi.IO<int> ProtocolPort { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// member.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// The subnet in which to access the member
		/// </summary>
		public Pulumi.IO<string> SubnetId { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the member.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new member.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// A positive integer value that indicates the relative
		/// portion of traffic that this member should receive from the pool. For
		/// example, a member with a weight of 10 receives five times as much traffic
		/// as a member with a weight of 2.
		/// </summary>
		public Pulumi.IO<int> Weight { get; set; }

	} // MemberArgs

	/// <summary>
	/// Manages a V2 member resource within OpenStack.
	/// </summary>
	public class Member : Pulumi.CustomResource {
		/// <summary>
		/// The IP address of the member to receive traffic from
		/// the load balancer. Changing this creates a new member.
		/// </summary>
		public Pulumi.IO<string> Address { get; set; }

		/// <summary>
		/// The administrative state of the member.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// Human-readable name for the member.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The id of the pool that this member will be
		/// assigned to.
		/// </summary>
		public Pulumi.IO<string> PoolId { get; set; }

		/// <summary>
		/// The port on which to listen for client traffic.
		/// Changing this creates a new member.
		/// </summary>
		public Pulumi.IO<int> ProtocolPort { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// member.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// The subnet in which to access the member
		/// </summary>
		public Pulumi.IO<string> SubnetId { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the member.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new member.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// A positive integer value that indicates the relative
		/// portion of traffic that this member should receive from the pool. For
		/// example, a member with a weight of 10 receives five times as much traffic
		/// as a member with a weight of 2.
		/// </summary>
		public Pulumi.IO<int> Weight { get; set; }

		public Member(string name, MemberArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:loadbalancer/member:Member", name, SerialiseArgs(args), opts) {
			Address = Outputs["address"].Select(item => Protobuf.ToString(item));
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			PoolId = Outputs["poolId"].Select(item => Protobuf.ToString(item));
			ProtocolPort = Outputs["protocolPort"].Select(item => Protobuf.ToInt(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			SubnetId = Outputs["subnetId"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			Weight = Outputs["weight"].Select(item => Protobuf.ToInt(item));
			Address = Outputs["address"].Select(item => Protobuf.ToString(item));
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			PoolId = Outputs["poolId"].Select(item => Protobuf.ToString(item));
			ProtocolPort = Outputs["protocolPort"].Select(item => Protobuf.ToInt(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			SubnetId = Outputs["subnetId"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			Weight = Outputs["weight"].Select(item => Protobuf.ToInt(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(MemberArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["address"] = Protobuf.ToProtobuf(args.Address);
			props["adminStateUp"] = Protobuf.ToProtobuf(args.AdminStateUp);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["poolId"] = Protobuf.ToProtobuf(args.PoolId);
			props["protocolPort"] = Protobuf.ToProtobuf(args.ProtocolPort);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["subnetId"] = Protobuf.ToProtobuf(args.SubnetId);
			props["tenantId"] = Protobuf.ToProtobuf(args.TenantId);
			props["weight"] = Protobuf.ToProtobuf(args.Weight);
			return props;
		} // SerialiseArgs

	} // Member
} // Pulumi.Openstack.Loadbalancer
