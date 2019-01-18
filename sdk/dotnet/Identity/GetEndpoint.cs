// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Identity {

	/// <summary>
	/// A collection of arguments for invoking getEndpoint.
	/// </summary>
	public struct GetEndpointArgs {
		/// <summary>
		/// The endpoint interface. Valid values are `public`,
		/// `internal`, and `admin`. Default value is `public`
		/// </summary>
		public string Interface { get; set; }

		/// <summary>
		/// The region the endpoint is located in.
		/// </summary>
		public string Region { get; set; }

		/// <summary>
		/// The service id this endpoint belongs to.
		/// </summary>
		public string ServiceId { get; set; }

		/// <summary>
		/// The service name of the endpoint.
		/// </summary>
		public string ServiceName { get; set; }

	} // GetEndpointArgs

	/// <summary>
	/// A collection of values returned by getEndpoint.
	/// </summary>
	public struct GetEndpointResult {
		/// <summary>
		/// The region the endpoint is located in.
		/// </summary>
		public string Region { get; set; }

		/// <summary>
		/// The endpoint URL
		/// </summary>
		public string Url { get; set; }

		/// <summary>
		/// id is the provider-assigned unique ID for this managed resource.
		/// </summary>
		public string Id { get; set; }

	} // GetEndpointResult

	public static partial class IdentityModule {
		/// <summary>
		/// Use this data source to get the ID of an OpenStack endpoint.
		/// 
		/// Note: This usually requires admin privileges.
		/// </summary>
		public static System.Threading.Tasks.Task<GetEndpointResult> GetEndpoint(GetEndpointArgs args, Pulumi.InvokeOptions opts = default(Pulumi.InvokeOptions)) {
			var invokeArgs = new Google.Protobuf.WellKnownTypes.Struct();
			invokeArgs.Fields["interface"] = Protobuf.ToProtobuf(args.Interface);
			invokeArgs.Fields["region"] = Protobuf.ToProtobuf(args.Region);
			invokeArgs.Fields["serviceId"] = Protobuf.ToProtobuf(args.ServiceId);
			invokeArgs.Fields["serviceName"] = Protobuf.ToProtobuf(args.ServiceName);

			var task = Pulumi.Runtime.InvokeAsync("openstack:identity/getEndpoint:getEndpoint", invokeArgs, opts);

			return task.ContinueWith(response => {
				var protobuf = response.Result;
				var result = new GetEndpointResult();
				if (protobuf.Fields.ContainsKey("region")) {
					result.Region = Protobuf.ToString(protobuf.Fields["region"]);
				}
				if (protobuf.Fields.ContainsKey("url")) {
					result.Url = Protobuf.ToString(protobuf.Fields["url"]);
				}
				if (protobuf.Fields.ContainsKey("id")) {
					result.Id = Protobuf.ToString(protobuf.Fields["id"]);
				}
				return result;
			});
		} // GetEndpoint

	} // IdentityModule
} // Pulumi.Openstack.Identity