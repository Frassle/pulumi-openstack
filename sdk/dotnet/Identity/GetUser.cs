// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Identity {

	/// <summary>
	/// A collection of arguments for invoking getUser.
	/// </summary>
	public struct GetUserArgs {
		/// <summary>
		/// The domain this user belongs to.
		/// </summary>
		public string DomainId { get; set; }

		/// <summary>
		/// Whether the user is enabled or disabled. Valid
		/// values are `true` and `false`.
		/// </summary>
		public bool Enabled { get; set; }

		/// <summary>
		/// The identity provider ID of the user.
		/// </summary>
		public string IdpId { get; set; }

		/// <summary>
		/// The name of the user.
		/// </summary>
		public string Name { get; set; }

		/// <summary>
		/// Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
		/// </summary>
		public string PasswordExpiresAt { get; set; }

		/// <summary>
		/// The protocol ID of the user.
		/// </summary>
		public string ProtocolId { get; set; }

		public string Region { get; set; }

		/// <summary>
		/// The unique ID of the user.
		/// </summary>
		public string UniqueId { get; set; }

	} // GetUserArgs

	/// <summary>
	/// A collection of values returned by getUser.
	/// </summary>
	public struct GetUserResult {
		/// <summary>
		/// See Argument Reference above.
		/// </summary>
		public string DefaultProjectId { get; set; }

		/// <summary>
		/// See Argument Reference above.
		/// </summary>
		public string DomainId { get; set; }

		/// <summary>
		/// The region the user is located in.
		/// </summary>
		public string Region { get; set; }

		/// <summary>
		/// id is the provider-assigned unique ID for this managed resource.
		/// </summary>
		public string Id { get; set; }

	} // GetUserResult

	public static partial class IdentityModule {
		/// <summary>
		/// Use this data source to get the ID of an OpenStack user.
		/// </summary>
		public static System.Threading.Tasks.Task<GetUserResult> GetUser(GetUserArgs args, Pulumi.InvokeOptions opts = default(Pulumi.InvokeOptions)) {
			var invokeArgs = new Google.Protobuf.WellKnownTypes.Struct();
			invokeArgs.Fields["domainId"] = Protobuf.ToProtobuf(args.DomainId);
			invokeArgs.Fields["enabled"] = Protobuf.ToProtobuf(args.Enabled);
			invokeArgs.Fields["idpId"] = Protobuf.ToProtobuf(args.IdpId);
			invokeArgs.Fields["name"] = Protobuf.ToProtobuf(args.Name);
			invokeArgs.Fields["passwordExpiresAt"] = Protobuf.ToProtobuf(args.PasswordExpiresAt);
			invokeArgs.Fields["protocolId"] = Protobuf.ToProtobuf(args.ProtocolId);
			invokeArgs.Fields["region"] = Protobuf.ToProtobuf(args.Region);
			invokeArgs.Fields["uniqueId"] = Protobuf.ToProtobuf(args.UniqueId);

			var task = Pulumi.Runtime.InvokeAsync("openstack:identity/getUser:getUser", invokeArgs, opts);

			return task.ContinueWith(response => {
				var protobuf = response.Result;
				var result = new GetUserResult();
				if (protobuf.Fields.ContainsKey("defaultProjectId")) {
					result.DefaultProjectId = Protobuf.ToString(protobuf.Fields["defaultProjectId"]);
				}
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
		} // GetUser

	} // IdentityModule
} // Pulumi.Openstack.Identity