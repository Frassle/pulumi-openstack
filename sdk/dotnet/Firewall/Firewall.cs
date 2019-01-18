// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Firewall {

	/// <summary>
	/// The set of arguments for constructing a Firewall resource.
	/// </summary>
	public struct FirewallArgs {
		/// <summary>
		/// Administrative up/down status for the firewall
		/// (must be "true" or "false" if provided - defaults to "true").
		/// Changing this updates the `admin_state_up` of an existing firewall.
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// Router(s) to associate this firewall instance
		/// with. Must be a list of strings. Changing this updates the associated routers
		/// of an existing firewall. Conflicts with `no_routers`.
		/// </summary>
		public Pulumi.IO<Pulumi.IO<string>[]> AssociatedRouters { get; set; }

		/// <summary>
		/// A description for the firewall. Changing this
		/// updates the `description` of an existing firewall.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// A name for the firewall. Changing this
		/// updates the `name` of an existing firewall.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// Should this firewall not be associated with any routers
		/// (must be "true" or "false" if provide - defaults to "false").
		/// Conflicts with `associated_routers`.
		/// </summary>
		public Pulumi.IO<bool> NoRouters { get; set; }

		/// <summary>
		/// The policy resource id for the firewall. Changing
		/// this updates the `policy_id` of an existing firewall.
		/// </summary>
		public Pulumi.IO<string> PolicyId { get; set; }

		/// <summary>
		/// The region in which to obtain the v1 networking client.
		/// A networking client is needed to create a firewall. If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// firewall.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// The owner of the floating IP. Required if admin wants
		/// to create a firewall for another tenant. Changing this creates a new
		/// firewall.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// Map of additional options.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> ValueSpecs { get; set; }

	} // FirewallArgs

	/// <summary>
	/// Manages a v1 firewall resource within OpenStack.
	/// </summary>
	public class Firewall : Pulumi.CustomResource {
		/// <summary>
		/// Administrative up/down status for the firewall
		/// (must be "true" or "false" if provided - defaults to "true").
		/// Changing this updates the `admin_state_up` of an existing firewall.
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// Router(s) to associate this firewall instance
		/// with. Must be a list of strings. Changing this updates the associated routers
		/// of an existing firewall. Conflicts with `no_routers`.
		/// </summary>
		public Pulumi.IO<string[]> AssociatedRouters { get; set; }

		/// <summary>
		/// A description for the firewall. Changing this
		/// updates the `description` of an existing firewall.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// A name for the firewall. Changing this
		/// updates the `name` of an existing firewall.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// Should this firewall not be associated with any routers
		/// (must be "true" or "false" if provide - defaults to "false").
		/// Conflicts with `associated_routers`.
		/// </summary>
		public Pulumi.IO<bool> NoRouters { get; set; }

		/// <summary>
		/// The policy resource id for the firewall. Changing
		/// this updates the `policy_id` of an existing firewall.
		/// </summary>
		public Pulumi.IO<string> PolicyId { get; set; }

		/// <summary>
		/// The region in which to obtain the v1 networking client.
		/// A networking client is needed to create a firewall. If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// firewall.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// The owner of the floating IP. Required if admin wants
		/// to create a firewall for another tenant. Changing this creates a new
		/// firewall.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// Map of additional options.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> ValueSpecs { get; set; }

		public Firewall(string name, FirewallArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:firewall/firewall:Firewall", name, SerialiseArgs(args), opts) {
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			AssociatedRouters = Outputs["associatedRouters"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			NoRouters = Outputs["noRouters"].Select(item => Protobuf.ToBool(item));
			PolicyId = Outputs["policyId"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			ValueSpecs = Outputs["valueSpecs"].Select(item => Protobuf.ToMap(item));
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			AssociatedRouters = Outputs["associatedRouters"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			NoRouters = Outputs["noRouters"].Select(item => Protobuf.ToBool(item));
			PolicyId = Outputs["policyId"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			ValueSpecs = Outputs["valueSpecs"].Select(item => Protobuf.ToMap(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(FirewallArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["adminStateUp"] = Protobuf.ToProtobuf(args.AdminStateUp);
			props["associatedRouters"] = Protobuf.ToProtobuf(args.AssociatedRouters, item => Protobuf.ToProtobuf(item));
			props["description"] = Protobuf.ToProtobuf(args.Description);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["noRouters"] = Protobuf.ToProtobuf(args.NoRouters);
			props["policyId"] = Protobuf.ToProtobuf(args.PolicyId);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["tenantId"] = Protobuf.ToProtobuf(args.TenantId);
			props["valueSpecs"] = Protobuf.ToProtobuf(args.ValueSpecs);
			return props;
		} // SerialiseArgs

	} // Firewall
} // Pulumi.Openstack.Firewall