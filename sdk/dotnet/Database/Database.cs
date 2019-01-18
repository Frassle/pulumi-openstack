// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Database {

	/// <summary>
	/// The set of arguments for constructing a Database resource.
	/// </summary>
	public struct DatabaseArgs {
		/// <summary>
		/// The ID for the database instance.
		/// </summary>
		public Pulumi.IO<string> InstanceId { get; set; }

		/// <summary>
		/// A unique name for the resource.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// Openstack region resource is created in.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

	} // DatabaseArgs

	/// <summary>
	/// Manages a V1 DB database resource within OpenStack.
	/// </summary>
	public class Database : Pulumi.CustomResource {
		/// <summary>
		/// The ID for the database instance.
		/// </summary>
		public Pulumi.IO<string> InstanceId { get; set; }

		/// <summary>
		/// A unique name for the resource.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// Openstack region resource is created in.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		public Database(string name, DatabaseArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:database/database:Database", name, SerialiseArgs(args), opts) {
			InstanceId = Outputs["instanceId"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			InstanceId = Outputs["instanceId"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(DatabaseArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["instanceId"] = Protobuf.ToProtobuf(args.InstanceId);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			return props;
		} // SerialiseArgs

	} // Database
} // Pulumi.Openstack.Database