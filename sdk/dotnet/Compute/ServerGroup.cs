// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Compute {

	/// <summary>
	/// The set of arguments for constructing a ServerGroup resource.
	/// </summary>
	public struct ServerGroupArgs {
		/// <summary>
		/// A unique name for the server group. Changing this creates
		/// a new server group.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The set of policies for the server group. Only two
		/// two policies are available right now, and both are mutually exclusive. See
		/// the Policies section for more information. Changing this creates a new
		/// server group.
		/// </summary>
		public Pulumi.IO<Pulumi.IO<string>[]> Policies { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Compute client.
		/// If omitted, the `region` argument of the provider is used. Changing
		/// this creates a new server group.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// Map of additional options.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> ValueSpecs { get; set; }

	} // ServerGroupArgs

	/// <summary>
	/// Manages a V2 Server Group resource within OpenStack.
	/// </summary>
	public class ServerGroup : Pulumi.CustomResource {
		/// <summary>
		/// The instances that are part of this server group.
		/// </summary>
		public Pulumi.IO<string[]> Members { get; set; }

		/// <summary>
		/// A unique name for the server group. Changing this creates
		/// a new server group.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The set of policies for the server group. Only two
		/// two policies are available right now, and both are mutually exclusive. See
		/// the Policies section for more information. Changing this creates a new
		/// server group.
		/// </summary>
		public Pulumi.IO<string[]> Policies { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Compute client.
		/// If omitted, the `region` argument of the provider is used. Changing
		/// this creates a new server group.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// Map of additional options.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> ValueSpecs { get; set; }

		public ServerGroup(string name, ServerGroupArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:compute/serverGroup:ServerGroup", name, SerialiseArgs(args), opts) {
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Policies = Outputs["policies"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			ValueSpecs = Outputs["valueSpecs"].Select(item => Protobuf.ToMap(item));
			Members = Outputs["members"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Policies = Outputs["policies"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			ValueSpecs = Outputs["valueSpecs"].Select(item => Protobuf.ToMap(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(ServerGroupArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["policies"] = Protobuf.ToProtobuf(args.Policies, item => Protobuf.ToProtobuf(item));
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["valueSpecs"] = Protobuf.ToProtobuf(args.ValueSpecs);
			props["members"] = null; //out
			return props;
		} // SerialiseArgs

	} // ServerGroup
} // Pulumi.Openstack.Compute