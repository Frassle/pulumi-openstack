// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Loadbalancer {

	/// <summary>
	/// The set of arguments for constructing a Listener resource.
	/// </summary>
	public struct ListenerArgs {
		/// <summary>
		/// The administrative state of the Listener.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// The maximum number of connections allowed
		/// for the Listener.
		/// </summary>
		public Pulumi.IO<int> ConnectionLimit { get; set; }

		/// <summary>
		/// The ID of the default pool with which the
		/// Listener is associated. Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> DefaultPoolId { get; set; }

		/// <summary>
		/// A reference to a Barbican Secrets
		/// container which stores TLS information. This is required if the protocol
		/// is `TERMINATED_HTTPS`. See
		/// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
		/// for more information.
		/// </summary>
		public Pulumi.IO<string> DefaultTlsContainerRef { get; set; }

		/// <summary>
		/// Human-readable description for the Listener.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// The load balancer on which to provision this
		/// Listener. Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> LoadbalancerId { get; set; }

		/// <summary>
		/// Human-readable name for the Listener. Does not have
		/// to be unique.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
		/// Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> Protocol { get; set; }

		/// <summary>
		/// The port on which to listen for client traffic.
		/// Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<int> ProtocolPort { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// Listener.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// A list of references to Barbican Secrets
		/// containers which store SNI information. See
		/// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
		/// for more information.
		/// </summary>
		public Pulumi.IO<Pulumi.IO<string>[]> SniContainerRefs { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the Listener.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

	} // ListenerArgs

	/// <summary>
	/// Manages a V2 listener resource within OpenStack.
	/// </summary>
	public class Listener : Pulumi.CustomResource {
		/// <summary>
		/// The administrative state of the Listener.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// The maximum number of connections allowed
		/// for the Listener.
		/// </summary>
		public Pulumi.IO<int> ConnectionLimit { get; set; }

		/// <summary>
		/// The ID of the default pool with which the
		/// Listener is associated. Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> DefaultPoolId { get; set; }

		/// <summary>
		/// A reference to a Barbican Secrets
		/// container which stores TLS information. This is required if the protocol
		/// is `TERMINATED_HTTPS`. See
		/// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
		/// for more information.
		/// </summary>
		public Pulumi.IO<string> DefaultTlsContainerRef { get; set; }

		/// <summary>
		/// Human-readable description for the Listener.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// The load balancer on which to provision this
		/// Listener. Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> LoadbalancerId { get; set; }

		/// <summary>
		/// Human-readable name for the Listener. Does not have
		/// to be unique.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
		/// Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> Protocol { get; set; }

		/// <summary>
		/// The port on which to listen for client traffic.
		/// Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<int> ProtocolPort { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// Listener.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// A list of references to Barbican Secrets
		/// containers which store SNI information. See
		/// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
		/// for more information.
		/// </summary>
		public Pulumi.IO<string[]> SniContainerRefs { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the Listener.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new Listener.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		public Listener(string name, ListenerArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:loadbalancer/listener:Listener", name, SerialiseArgs(args), opts) {
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			ConnectionLimit = Outputs["connectionLimit"].Select(item => Protobuf.ToInt(item));
			DefaultPoolId = Outputs["defaultPoolId"].Select(item => Protobuf.ToString(item));
			DefaultTlsContainerRef = Outputs["defaultTlsContainerRef"].Select(item => Protobuf.ToString(item));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			LoadbalancerId = Outputs["loadbalancerId"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Protocol = Outputs["protocol"].Select(item => Protobuf.ToString(item));
			ProtocolPort = Outputs["protocolPort"].Select(item => Protobuf.ToInt(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			SniContainerRefs = Outputs["sniContainerRefs"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			ConnectionLimit = Outputs["connectionLimit"].Select(item => Protobuf.ToInt(item));
			DefaultPoolId = Outputs["defaultPoolId"].Select(item => Protobuf.ToString(item));
			DefaultTlsContainerRef = Outputs["defaultTlsContainerRef"].Select(item => Protobuf.ToString(item));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			LoadbalancerId = Outputs["loadbalancerId"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Protocol = Outputs["protocol"].Select(item => Protobuf.ToString(item));
			ProtocolPort = Outputs["protocolPort"].Select(item => Protobuf.ToInt(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			SniContainerRefs = Outputs["sniContainerRefs"].Select(item => Protobuf.ToList(item, item1 => Protobuf.ToString(item1)));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(ListenerArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["adminStateUp"] = Protobuf.ToProtobuf(args.AdminStateUp);
			props["connectionLimit"] = Protobuf.ToProtobuf(args.ConnectionLimit);
			props["defaultPoolId"] = Protobuf.ToProtobuf(args.DefaultPoolId);
			props["defaultTlsContainerRef"] = Protobuf.ToProtobuf(args.DefaultTlsContainerRef);
			props["description"] = Protobuf.ToProtobuf(args.Description);
			props["loadbalancerId"] = Protobuf.ToProtobuf(args.LoadbalancerId);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["protocol"] = Protobuf.ToProtobuf(args.Protocol);
			props["protocolPort"] = Protobuf.ToProtobuf(args.ProtocolPort);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["sniContainerRefs"] = Protobuf.ToProtobuf(args.SniContainerRefs, item => Protobuf.ToProtobuf(item));
			props["tenantId"] = Protobuf.ToProtobuf(args.TenantId);
			return props;
		} // SerialiseArgs

	} // Listener
} // Pulumi.Openstack.Loadbalancer
