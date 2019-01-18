// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Identity {

	/// <summary>
	/// The set of arguments for constructing a RoleAssignment resource.
	/// </summary>
	public struct RoleAssignmentArgs {
		/// <summary>
		/// The domain to assign the role in.
		/// </summary>
		public Pulumi.IO<string> DomainId { get; set; }

		/// <summary>
		/// The group to assign the role to.
		/// </summary>
		public Pulumi.IO<string> GroupId { get; set; }

		/// <summary>
		/// The project to assign the role in.
		/// </summary>
		public Pulumi.IO<string> ProjectId { get; set; }

		/// <summary>
		/// The role to assign.
		/// </summary>
		public Pulumi.IO<string> RoleId { get; set; }

		/// <summary>
		/// The user to assign the role to.
		/// </summary>
		public Pulumi.IO<string> UserId { get; set; }

	} // RoleAssignmentArgs

	/// <summary>
	/// Manages a V3 Role assignment within OpenStack Keystone.
	/// 
	/// Note: You _must_ have admin privileges in your OpenStack cloud to use
	/// this resource.
	/// </summary>
	public class RoleAssignment : Pulumi.CustomResource {
		/// <summary>
		/// The domain to assign the role in.
		/// </summary>
		public Pulumi.IO<string> DomainId { get; set; }

		/// <summary>
		/// The group to assign the role to.
		/// </summary>
		public Pulumi.IO<string> GroupId { get; set; }

		/// <summary>
		/// The project to assign the role in.
		/// </summary>
		public Pulumi.IO<string> ProjectId { get; set; }

		/// <summary>
		/// The role to assign.
		/// </summary>
		public Pulumi.IO<string> RoleId { get; set; }

		/// <summary>
		/// The user to assign the role to.
		/// </summary>
		public Pulumi.IO<string> UserId { get; set; }

		public RoleAssignment(string name, RoleAssignmentArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:identity/roleAssignment:RoleAssignment", name, SerialiseArgs(args), opts) {
			DomainId = Outputs["domainId"].Select(item => Protobuf.ToString(item));
			GroupId = Outputs["groupId"].Select(item => Protobuf.ToString(item));
			ProjectId = Outputs["projectId"].Select(item => Protobuf.ToString(item));
			RoleId = Outputs["roleId"].Select(item => Protobuf.ToString(item));
			UserId = Outputs["userId"].Select(item => Protobuf.ToString(item));
			DomainId = Outputs["domainId"].Select(item => Protobuf.ToString(item));
			GroupId = Outputs["groupId"].Select(item => Protobuf.ToString(item));
			ProjectId = Outputs["projectId"].Select(item => Protobuf.ToString(item));
			RoleId = Outputs["roleId"].Select(item => Protobuf.ToString(item));
			UserId = Outputs["userId"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(RoleAssignmentArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["domainId"] = Protobuf.ToProtobuf(args.DomainId);
			props["groupId"] = Protobuf.ToProtobuf(args.GroupId);
			props["projectId"] = Protobuf.ToProtobuf(args.ProjectId);
			props["roleId"] = Protobuf.ToProtobuf(args.RoleId);
			props["userId"] = Protobuf.ToProtobuf(args.UserId);
			return props;
		} // SerialiseArgs

	} // RoleAssignment
} // Pulumi.Openstack.Identity