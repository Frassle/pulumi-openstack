// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Dns {

	/// <summary>
	/// A collection of arguments for invoking getDnsZone.
	/// </summary>
	public struct GetDnsZoneArgs {
		public System.Collections.Generic.Dictionary<string, string> Attributes { get; set; }

		public string CreatedAt { get; set; }

		/// <summary>
		/// A description of the zone.
		/// </summary>
		public string Description { get; set; }

		/// <summary>
		/// The email contact for the zone record.
		/// </summary>
		public string Email { get; set; }

		public string[] Masters { get; set; }

		/// <summary>
		/// The name of the zone.
		/// </summary>
		public string Name { get; set; }

		public string PoolId { get; set; }

		public string ProjectId { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 DNS client.
		/// A DNS client is needed to retrieve zone ids. If omitted, the
		/// `region` argument of the provider is used.
		/// </summary>
		public string Region { get; set; }

		public int Serial { get; set; }

		/// <summary>
		/// The zone's status.
		/// </summary>
		public string Status { get; set; }

		public string TransferredAt { get; set; }

		/// <summary>
		/// The time to live (TTL) of the zone.
		/// </summary>
		public int Ttl { get; set; }

		/// <summary>
		/// The type of the zone. Can either be `PRIMARY` or `SECONDARY`.
		/// </summary>
		public string Type { get; set; }

		public string UpdatedAt { get; set; }

		public int Version { get; set; }

	} // GetDnsZoneArgs

	/// <summary>
	/// A collection of values returned by getDnsZone.
	/// </summary>
	public struct GetDnsZoneResult {
		/// <summary>
		/// Attributes of the DNS Service scheduler.
		/// </summary>
		public System.Collections.Generic.Dictionary<string, string> Attributes { get; set; }

		/// <summary>
		/// The time the zone was created.
		/// </summary>
		public string CreatedAt { get; set; }

		/// <summary>
		/// An array of master DNS servers. When `type` is  `SECONDARY`.
		/// </summary>
		public string[] Masters { get; set; }

		/// <summary>
		/// The ID of the pool hosting the zone.
		/// </summary>
		public string PoolId { get; set; }

		/// <summary>
		/// The project ID that owns the zone.
		/// </summary>
		public string ProjectId { get; set; }

		/// <summary>
		/// See Argument Reference above.
		/// </summary>
		public string Region { get; set; }

		/// <summary>
		/// The serial number of the zone.
		/// </summary>
		public int Serial { get; set; }

		/// <summary>
		/// The time the zone was transferred.
		/// </summary>
		public string TransferredAt { get; set; }

		/// <summary>
		/// The time the zone was last updated.
		/// </summary>
		public string UpdatedAt { get; set; }

		/// <summary>
		/// The version of the zone.
		/// </summary>
		public int Version { get; set; }

		/// <summary>
		/// id is the provider-assigned unique ID for this managed resource.
		/// </summary>
		public string Id { get; set; }

	} // GetDnsZoneResult

	public static partial class DnsModule {
		/// <summary>
		/// Use this data source to get the ID of an available OpenStack DNS zone.
		/// </summary>
		public static System.Threading.Tasks.Task<GetDnsZoneResult> GetDnsZone(GetDnsZoneArgs args, Pulumi.InvokeOptions opts = default(Pulumi.InvokeOptions)) {
			var invokeArgs = new Google.Protobuf.WellKnownTypes.Struct();
			invokeArgs.Fields["attributes"] = Protobuf.ToProtobuf(args.Attributes);
			invokeArgs.Fields["createdAt"] = Protobuf.ToProtobuf(args.CreatedAt);
			invokeArgs.Fields["description"] = Protobuf.ToProtobuf(args.Description);
			invokeArgs.Fields["email"] = Protobuf.ToProtobuf(args.Email);
			invokeArgs.Fields["masters"] = Protobuf.ToProtobuf(args.Masters, item => Protobuf.ToProtobuf(item));
			invokeArgs.Fields["name"] = Protobuf.ToProtobuf(args.Name);
			invokeArgs.Fields["poolId"] = Protobuf.ToProtobuf(args.PoolId);
			invokeArgs.Fields["projectId"] = Protobuf.ToProtobuf(args.ProjectId);
			invokeArgs.Fields["region"] = Protobuf.ToProtobuf(args.Region);
			invokeArgs.Fields["serial"] = Protobuf.ToProtobuf(args.Serial);
			invokeArgs.Fields["status"] = Protobuf.ToProtobuf(args.Status);
			invokeArgs.Fields["transferredAt"] = Protobuf.ToProtobuf(args.TransferredAt);
			invokeArgs.Fields["ttl"] = Protobuf.ToProtobuf(args.Ttl);
			invokeArgs.Fields["type"] = Protobuf.ToProtobuf(args.Type);
			invokeArgs.Fields["updatedAt"] = Protobuf.ToProtobuf(args.UpdatedAt);
			invokeArgs.Fields["version"] = Protobuf.ToProtobuf(args.Version);

			var task = Pulumi.Runtime.InvokeAsync("openstack:dns/getDnsZone:getDnsZone", invokeArgs, opts);

			return task.ContinueWith(response => {
				var protobuf = response.Result;
				var result = new GetDnsZoneResult();
				if (protobuf.Fields.ContainsKey("attributes")) {
					result.Attributes = Protobuf.ToMap(protobuf.Fields["attributes"]);
				}
				if (protobuf.Fields.ContainsKey("createdAt")) {
					result.CreatedAt = Protobuf.ToString(protobuf.Fields["createdAt"]);
				}
				if (protobuf.Fields.ContainsKey("masters")) {
					result.Masters = Protobuf.ToList(protobuf.Fields["masters"], item => Protobuf.ToString(item));
				}
				if (protobuf.Fields.ContainsKey("poolId")) {
					result.PoolId = Protobuf.ToString(protobuf.Fields["poolId"]);
				}
				if (protobuf.Fields.ContainsKey("projectId")) {
					result.ProjectId = Protobuf.ToString(protobuf.Fields["projectId"]);
				}
				if (protobuf.Fields.ContainsKey("region")) {
					result.Region = Protobuf.ToString(protobuf.Fields["region"]);
				}
				if (protobuf.Fields.ContainsKey("serial")) {
					result.Serial = Protobuf.ToInt(protobuf.Fields["serial"]);
				}
				if (protobuf.Fields.ContainsKey("transferredAt")) {
					result.TransferredAt = Protobuf.ToString(protobuf.Fields["transferredAt"]);
				}
				if (protobuf.Fields.ContainsKey("updatedAt")) {
					result.UpdatedAt = Protobuf.ToString(protobuf.Fields["updatedAt"]);
				}
				if (protobuf.Fields.ContainsKey("version")) {
					result.Version = Protobuf.ToInt(protobuf.Fields["version"]);
				}
				if (protobuf.Fields.ContainsKey("id")) {
					result.Id = Protobuf.ToString(protobuf.Fields["id"]);
				}
				return result;
			});
		} // GetDnsZone

	} // DnsModule
} // Pulumi.Openstack.Dns
