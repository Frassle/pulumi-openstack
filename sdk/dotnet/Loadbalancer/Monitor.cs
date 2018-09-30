// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Loadbalancer {

	/// <summary>
	/// The set of arguments for constructing a Monitor resource.
	/// </summary>
	public struct MonitorArgs {
		/// <summary>
		/// The administrative state of the monitor.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// The time, in seconds, between sending probes to members.
		/// </summary>
		public Pulumi.IO<int> Delay { get; set; }

		/// <summary>
		/// Required for HTTP(S) types. Expected HTTP codes
		/// for a passing HTTP(S) monitor. You can either specify a single status like
		/// "200", or a range like "200-202".
		/// </summary>
		public Pulumi.IO<string> ExpectedCodes { get; set; }

		/// <summary>
		/// Required for HTTP(S) types. The HTTP method used
		/// for requests by the monitor. If this attribute is not specified, it
		/// defaults to "GET".
		/// </summary>
		public Pulumi.IO<string> HttpMethod { get; set; }

		/// <summary>
		/// Number of permissible ping failures before
		/// changing the member's status to INACTIVE. Must be a number between 1
		/// and 10..
		/// </summary>
		public Pulumi.IO<int> MaxRetries { get; set; }

		/// <summary>
		/// The Name of the Monitor.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The id of the pool that this monitor will be assigned to.
		/// </summary>
		public Pulumi.IO<string> PoolId { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// monitor.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the monitor.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new monitor.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// Maximum number of seconds for a monitor to wait for a
		/// ping reply before it times out. The value must be less than the delay
		/// value.
		/// </summary>
		public Pulumi.IO<int> Timeout { get; set; }

		/// <summary>
		/// The type of probe, which is PING, TCP, HTTP, or HTTPS,
		/// that is sent by the load balancer to verify the member state. Changing this
		/// creates a new monitor.
		/// </summary>
		public Pulumi.IO<string> Type { get; set; }

		/// <summary>
		/// Required for HTTP(S) types. URI path that will be
		/// accessed if monitor type is HTTP or HTTPS.
		/// </summary>
		public Pulumi.IO<string> UrlPath { get; set; }

	} // MonitorArgs

	/// <summary>
	/// Manages a V2 monitor resource within OpenStack.
	/// </summary>
	public class Monitor : Pulumi.CustomResource {
		/// <summary>
		/// The administrative state of the monitor.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// The time, in seconds, between sending probes to members.
		/// </summary>
		public Pulumi.IO<int> Delay { get; set; }

		/// <summary>
		/// Required for HTTP(S) types. Expected HTTP codes
		/// for a passing HTTP(S) monitor. You can either specify a single status like
		/// "200", or a range like "200-202".
		/// </summary>
		public Pulumi.IO<string> ExpectedCodes { get; set; }

		/// <summary>
		/// Required for HTTP(S) types. The HTTP method used
		/// for requests by the monitor. If this attribute is not specified, it
		/// defaults to "GET".
		/// </summary>
		public Pulumi.IO<string> HttpMethod { get; set; }

		/// <summary>
		/// Number of permissible ping failures before
		/// changing the member's status to INACTIVE. Must be a number between 1
		/// and 10..
		/// </summary>
		public Pulumi.IO<int> MaxRetries { get; set; }

		/// <summary>
		/// The Name of the Monitor.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The id of the pool that this monitor will be assigned to.
		/// </summary>
		public Pulumi.IO<string> PoolId { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// monitor.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the monitor.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new monitor.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		/// <summary>
		/// Maximum number of seconds for a monitor to wait for a
		/// ping reply before it times out. The value must be less than the delay
		/// value.
		/// </summary>
		public Pulumi.IO<int> Timeout { get; set; }

		/// <summary>
		/// The type of probe, which is PING, TCP, HTTP, or HTTPS,
		/// that is sent by the load balancer to verify the member state. Changing this
		/// creates a new monitor.
		/// </summary>
		public Pulumi.IO<string> Type { get; set; }

		/// <summary>
		/// Required for HTTP(S) types. URI path that will be
		/// accessed if monitor type is HTTP or HTTPS.
		/// </summary>
		public Pulumi.IO<string> UrlPath { get; set; }

		public Monitor(string name, MonitorArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:loadbalancer/monitor:Monitor", name, SerialiseArgs(args), opts) {
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			Delay = Outputs["delay"].Select(item => Protobuf.ToInt(item));
			ExpectedCodes = Outputs["expectedCodes"].Select(item => Protobuf.ToString(item));
			HttpMethod = Outputs["httpMethod"].Select(item => Protobuf.ToString(item));
			MaxRetries = Outputs["maxRetries"].Select(item => Protobuf.ToInt(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			PoolId = Outputs["poolId"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			Timeout = Outputs["timeout"].Select(item => Protobuf.ToInt(item));
			Type = Outputs["type"].Select(item => Protobuf.ToString(item));
			UrlPath = Outputs["urlPath"].Select(item => Protobuf.ToString(item));
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			Delay = Outputs["delay"].Select(item => Protobuf.ToInt(item));
			ExpectedCodes = Outputs["expectedCodes"].Select(item => Protobuf.ToString(item));
			HttpMethod = Outputs["httpMethod"].Select(item => Protobuf.ToString(item));
			MaxRetries = Outputs["maxRetries"].Select(item => Protobuf.ToInt(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			PoolId = Outputs["poolId"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			Timeout = Outputs["timeout"].Select(item => Protobuf.ToInt(item));
			Type = Outputs["type"].Select(item => Protobuf.ToString(item));
			UrlPath = Outputs["urlPath"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(MonitorArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["adminStateUp"] = Protobuf.ToProtobuf(args.AdminStateUp);
			props["delay"] = Protobuf.ToProtobuf(args.Delay);
			props["expectedCodes"] = Protobuf.ToProtobuf(args.ExpectedCodes);
			props["httpMethod"] = Protobuf.ToProtobuf(args.HttpMethod);
			props["maxRetries"] = Protobuf.ToProtobuf(args.MaxRetries);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["poolId"] = Protobuf.ToProtobuf(args.PoolId);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["tenantId"] = Protobuf.ToProtobuf(args.TenantId);
			props["timeout"] = Protobuf.ToProtobuf(args.Timeout);
			props["type"] = Protobuf.ToProtobuf(args.Type);
			props["urlPath"] = Protobuf.ToProtobuf(args.UrlPath);
			return props;
		} // SerialiseArgs

	} // Monitor
} // Pulumi.Openstack.Loadbalancer
