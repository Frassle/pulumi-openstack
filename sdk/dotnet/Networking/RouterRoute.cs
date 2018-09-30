// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Networking {

	/// <summary>
	/// The set of arguments for constructing a RouterRoute resource.
	/// </summary>
	public struct RouterRouteArgs {
		/// <summary>
		/// CIDR block to match on the packet’s destination IP. Changing
		/// this creates a new routing entry.
		/// </summary>
		public Pulumi.IO<string> DestinationCidr { get; set; }

		/// <summary>
		/// IP address of the next hop gateway.  Changing
		/// this creates a new routing entry.
		/// </summary>
		public Pulumi.IO<string> NextHop { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 networking client.
		/// A networking client is needed to configure a routing entry on a router. If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// routing entry.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// ID of the router this routing entry belongs to. Changing
		/// this creates a new routing entry.
		/// </summary>
		public Pulumi.IO<string> RouterId { get; set; }

	} // RouterRouteArgs

	/// <summary>
	/// Creates a routing entry on a OpenStack V2 router.
	/// </summary>
	public class RouterRoute : Pulumi.CustomResource {
		/// <summary>
		/// CIDR block to match on the packet’s destination IP. Changing
		/// this creates a new routing entry.
		/// </summary>
		public Pulumi.IO<string> DestinationCidr { get; set; }

		/// <summary>
		/// IP address of the next hop gateway.  Changing
		/// this creates a new routing entry.
		/// </summary>
		public Pulumi.IO<string> NextHop { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 networking client.
		/// A networking client is needed to configure a routing entry on a router. If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// routing entry.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// ID of the router this routing entry belongs to. Changing
		/// this creates a new routing entry.
		/// </summary>
		public Pulumi.IO<string> RouterId { get; set; }

		public RouterRoute(string name, RouterRouteArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:networking/routerRoute:RouterRoute", name, SerialiseArgs(args), opts) {
			DestinationCidr = Outputs["destinationCidr"].Select(item => Protobuf.ToString(item));
			NextHop = Outputs["nextHop"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			RouterId = Outputs["routerId"].Select(item => Protobuf.ToString(item));
			DestinationCidr = Outputs["destinationCidr"].Select(item => Protobuf.ToString(item));
			NextHop = Outputs["nextHop"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			RouterId = Outputs["routerId"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(RouterRouteArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["destinationCidr"] = Protobuf.ToProtobuf(args.DestinationCidr);
			props["nextHop"] = Protobuf.ToProtobuf(args.NextHop);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["routerId"] = Protobuf.ToProtobuf(args.RouterId);
			return props;
		} // SerialiseArgs

	} // RouterRoute
} // Pulumi.Openstack.Networking
