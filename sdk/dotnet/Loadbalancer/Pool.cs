// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Loadbalancer {

	public sealed class PoolArgsPersistence : Pulumi.IIOProtobuf {
		public Pulumi.IO<string> CookieName { get; set; }
		public Pulumi.IO<string> Type { get; set; }

		public Pulumi.IO<Google.Protobuf.WellKnownTypes.Value> ToProtobuf() {
			return Pulumi.Protobuf.ToProtobuf(
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("cookie_name", Protobuf.ToProtobuf(CookieName)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("type", Protobuf.ToProtobuf(Type)));
		} // ToProtobuf

	} // PoolArgsPersistence

	public sealed class PoolPersistence : Pulumi.IIOProtobuf {
		public Pulumi.IO<string> CookieName { get; set; }
		public Pulumi.IO<string> Type { get; set; }

		public Pulumi.IO<Google.Protobuf.WellKnownTypes.Value> ToProtobuf() {
			return Pulumi.Protobuf.ToProtobuf(
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("cookie_name", Protobuf.ToProtobuf(CookieName)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("type", Protobuf.ToProtobuf(Type)));
		} // ToProtobuf

		public static PoolPersistence FromProtobuf(Google.Protobuf.WellKnownTypes.Value value) {
			var obj = value.StructValue;
			return new PoolPersistence() {
				CookieName = Protobuf.ToString(obj.Fields["cookie_name"]),
				Type = Protobuf.ToString(obj.Fields["type"]),
			};
		} // FromProtobuf

	} // PoolPersistence

	/// <summary>
	/// The set of arguments for constructing a Pool resource.
	/// </summary>
	public struct PoolArgs {
		/// <summary>
		/// The administrative state of the pool.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// Human-readable description for the pool.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// The load balancing algorithm to
		/// distribute traffic to the pool's members. Must be one of
		/// ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
		/// </summary>
		public Pulumi.IO<string> LbMethod { get; set; }

		/// <summary>
		/// The Listener on which the members of the pool
		/// will be associated with. Changing this creates a new pool.
		/// Note:  One of LoadbalancerID or ListenerID must be provided.
		/// </summary>
		public Pulumi.IO<string> ListenerId { get; set; }

		/// <summary>
		/// The load balancer on which to provision this
		/// pool. Changing this creates a new pool.
		/// Note:  One of LoadbalancerID or ListenerID must be provided.
		/// </summary>
		public Pulumi.IO<string> LoadbalancerId { get; set; }

		/// <summary>
		/// Human-readable name for the pool.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// Omit this field to prevent session persistence.  Indicates
		/// whether connections in the same session will be processed by the same Pool
		/// member or not. Changing this creates a new pool.
		/// </summary>
		public Pulumi.IO<Pulumi.IO<PoolArgsPersistence>[]> Persistences { get; set; }

		/// <summary>
		/// See Argument Reference above.
		/// </summary>
		public Pulumi.IO<string> Protocol { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// pool.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the pool.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new pool.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

	} // PoolArgs

	/// <summary>
	/// Manages a V2 pool resource within OpenStack.
	/// </summary>
	public class Pool : Pulumi.CustomResource {
		/// <summary>
		/// The administrative state of the pool.
		/// A valid value is true (UP) or false (DOWN).
		/// </summary>
		public Pulumi.IO<bool> AdminStateUp { get; set; }

		/// <summary>
		/// Human-readable description for the pool.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// The load balancing algorithm to
		/// distribute traffic to the pool's members. Must be one of
		/// ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
		/// </summary>
		public Pulumi.IO<string> LbMethod { get; set; }

		/// <summary>
		/// The Listener on which the members of the pool
		/// will be associated with. Changing this creates a new pool.
		/// Note:  One of LoadbalancerID or ListenerID must be provided.
		/// </summary>
		public Pulumi.IO<string> ListenerId { get; set; }

		/// <summary>
		/// The load balancer on which to provision this
		/// pool. Changing this creates a new pool.
		/// Note:  One of LoadbalancerID or ListenerID must be provided.
		/// </summary>
		public Pulumi.IO<string> LoadbalancerId { get; set; }

		/// <summary>
		/// Human-readable name for the pool.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// Omit this field to prevent session persistence.  Indicates
		/// whether connections in the same session will be processed by the same Pool
		/// member or not. Changing this creates a new pool.
		/// </summary>
		public Pulumi.IO<PoolPersistence[]> Persistences { get; set; }

		/// <summary>
		/// See Argument Reference above.
		/// </summary>
		public Pulumi.IO<string> Protocol { get; set; }

		/// <summary>
		/// The region in which to obtain the V2 Networking client.
		/// A Networking client is needed to create an . If omitted, the
		/// `region` argument of the provider is used. Changing this creates a new
		/// pool.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// Required for admins. The UUID of the tenant who owns
		/// the pool.  Only administrative users can specify a tenant UUID
		/// other than their own. Changing this creates a new pool.
		/// </summary>
		public Pulumi.IO<string> TenantId { get; set; }

		public Pool(string name, PoolArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:loadbalancer/pool:Pool", name, SerialiseArgs(args), opts) {
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			LbMethod = Outputs["lbMethod"].Select(item => Protobuf.ToString(item));
			ListenerId = Outputs["listenerId"].Select(item => Protobuf.ToString(item));
			LoadbalancerId = Outputs["loadbalancerId"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Persistences = Outputs["persistences"].Select(item => Protobuf.ToList(item, item1 => PoolPersistence.FromProtobuf(item1)));
			Protocol = Outputs["protocol"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
			AdminStateUp = Outputs["adminStateUp"].Select(item => Protobuf.ToBool(item));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			LbMethod = Outputs["lbMethod"].Select(item => Protobuf.ToString(item));
			ListenerId = Outputs["listenerId"].Select(item => Protobuf.ToString(item));
			LoadbalancerId = Outputs["loadbalancerId"].Select(item => Protobuf.ToString(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Persistences = Outputs["persistences"].Select(item => Protobuf.ToList(item, item1 => PoolPersistence.FromProtobuf(item1)));
			Protocol = Outputs["protocol"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			TenantId = Outputs["tenantId"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(PoolArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["adminStateUp"] = Protobuf.ToProtobuf(args.AdminStateUp);
			props["description"] = Protobuf.ToProtobuf(args.Description);
			props["lbMethod"] = Protobuf.ToProtobuf(args.LbMethod);
			props["listenerId"] = Protobuf.ToProtobuf(args.ListenerId);
			props["loadbalancerId"] = Protobuf.ToProtobuf(args.LoadbalancerId);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["persistences"] = Protobuf.ToProtobuf(args.Persistences, item => Protobuf.ToProtobuf(item));
			props["protocol"] = Protobuf.ToProtobuf(args.Protocol);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["tenantId"] = Protobuf.ToProtobuf(args.TenantId);
			return props;
		} // SerialiseArgs

	} // Pool
} // Pulumi.Openstack.Loadbalancer
