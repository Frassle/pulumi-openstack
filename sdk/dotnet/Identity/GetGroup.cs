// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Identity {

	/// <summary>
	/// A collection of arguments for invoking getGroup.
	/// </summary>
	public struct GetGroupArgs {
		/// <summary>
		/// The domain the group belongs to.
		/// </summary>
		public string DomainId { get; set; }

		/// <summary>
		/// The name of the group.
		/// </summary>
		public string Name { get; set; }

		/// <summary>
		/// The region in which to obtain the V3 Keystone client.
		/// If omitted, the `region` argument of the provider is used.
		/// </summary>
		public string Region { get; set; }

	} // GetGroupArgs

	/// <summary>
	/// A collection of values returned by getGroup.
	/// </summary>
	public struct GetGroupResult {
		/// <summary>
		/// See Argument Reference above.
		/// </summary>
		public string DomainId { get; set; }

		/// <summary>
		/// See Argument Reference above.
		/// </summary>
		public string Region { get; set; }

		/// <summary>
		/// id is the provider-assigned unique ID for this managed resource.
		/// </summary>
		public string Id { get; set; }

	} // GetGroupResult

	public static partial class IdentityModule {
		/// <summary>
		/// Use this data source to get the ID of an OpenStack group.
		/// 
		/// Note: This usually requires admin privileges.
		/// </summary>
		public static System.Threading.Tasks.Task<GetGroupResult> GetGroup(GetGroupArgs args, Pulumi.InvokeOptions opts = default(Pulumi.InvokeOptions)) {
			var invokeArgs = new Google.Protobuf.WellKnownTypes.Struct();
			invokeArgs.Fields["domainId"] = Protobuf.ToProtobuf(args.DomainId);
			invokeArgs.Fields["name"] = Protobuf.ToProtobuf(args.Name);
			invokeArgs.Fields["region"] = Protobuf.ToProtobuf(args.Region);

			var task = Pulumi.Runtime.InvokeAsync("openstack:identity/getGroup:getGroup", invokeArgs, opts);

			return task.ContinueWith(response => {
				var protobuf = response.Result;
				var result = new GetGroupResult();
				if (protobuf.Fields.ContainsKey("domainId")) {
					result.DomainId = Protobuf.ToString(protobuf.Fields["domainId"]);
				}
				if (protobuf.Fields.ContainsKey("region")) {
					result.Region = Protobuf.ToString(protobuf.Fields["region"]);
				}
				if (protobuf.Fields.ContainsKey("id")) {
					result.Id = Protobuf.ToString(protobuf.Fields["id"]);
				}
				return result;
			});
		} // GetGroup

	} // IdentityModule
} // Pulumi.Openstack.Identity
