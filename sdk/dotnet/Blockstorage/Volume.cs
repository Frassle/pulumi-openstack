// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Blockstorage {

	public sealed class VolumeAttachment : Pulumi.IIOProtobuf {
		public Pulumi.IO<string> Device { get; set; }
		public Pulumi.IO<string> Id { get; set; }
		public Pulumi.IO<string> InstanceId { get; set; }

		public Pulumi.IO<Google.Protobuf.WellKnownTypes.Value> ToProtobuf() {
			return Pulumi.Protobuf.ToProtobuf(
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("device", Protobuf.ToProtobuf(Device)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("id", Protobuf.ToProtobuf(Id)),
				new KeyValuePair<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>("instance_id", Protobuf.ToProtobuf(InstanceId)));
		} // ToProtobuf

		public static VolumeAttachment FromProtobuf(Google.Protobuf.WellKnownTypes.Value value) {
			var obj = value.StructValue;
			return new VolumeAttachment() {
				Device = Protobuf.ToString(obj.Fields["device"]),
				Id = Protobuf.ToString(obj.Fields["id"]),
				InstanceId = Protobuf.ToString(obj.Fields["instance_id"]),
			};
		} // FromProtobuf

	} // VolumeAttachment

	/// <summary>
	/// The set of arguments for constructing a Volume resource.
	/// </summary>
	public struct VolumeArgs {
		/// <summary>
		/// The availability zone for the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> AvailabilityZone { get; set; }

		/// <summary>
		/// The consistency group to place the volume
		/// in.
		/// </summary>
		public Pulumi.IO<string> ConsistencyGroupId { get; set; }

		/// <summary>
		/// A description of the volume. Changing this updates
		/// the volume's description.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// When this option is set it allows extending
		/// attached volumes. Note: updating size of an attached volume requires Cinder
		/// support for version 3.42 and a compatible storage driver.
		/// </summary>
		public Pulumi.IO<bool> EnableOnlineResize { get; set; }

		/// <summary>
		/// The image ID from which to create the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> ImageId { get; set; }

		/// <summary>
		/// Metadata key/value pairs to associate with the volume.
		/// Changing this updates the existing volume metadata.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> Metadata { get; set; }

		/// <summary>
		/// A unique name for the volume. Changing this updates the
		/// volume's name.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The region in which to create the volume. If
		/// omitted, the `region` argument of the provider is used. Changing this
		/// creates a new volume.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// The size of the volume to create (in gigabytes).
		/// </summary>
		public Pulumi.IO<int> Size { get; set; }

		/// <summary>
		/// The snapshot ID from which to create the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> SnapshotId { get; set; }

		/// <summary>
		/// The volume ID to replicate with.
		/// </summary>
		public Pulumi.IO<string> SourceReplica { get; set; }

		/// <summary>
		/// The volume ID from which to create the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> SourceVolId { get; set; }

		/// <summary>
		/// The type of volume to create.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> VolumeType { get; set; }

	} // VolumeArgs

	/// <summary>
	/// Manages a V3 volume resource within OpenStack.
	/// </summary>
	public class Volume : Pulumi.CustomResource {
		/// <summary>
		/// If a volume is attached to an instance, this attribute will
		/// display the Attachment ID, Instance ID, and the Device as the Instance
		/// sees it.
		/// </summary>
		public Pulumi.IO<VolumeAttachment[]> Attachments { get; set; }

		/// <summary>
		/// The availability zone for the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> AvailabilityZone { get; set; }

		/// <summary>
		/// The consistency group to place the volume
		/// in.
		/// </summary>
		public Pulumi.IO<string> ConsistencyGroupId { get; set; }

		/// <summary>
		/// A description of the volume. Changing this updates
		/// the volume's description.
		/// </summary>
		public Pulumi.IO<string> Description { get; set; }

		/// <summary>
		/// When this option is set it allows extending
		/// attached volumes. Note: updating size of an attached volume requires Cinder
		/// support for version 3.42 and a compatible storage driver.
		/// </summary>
		public Pulumi.IO<bool> EnableOnlineResize { get; set; }

		/// <summary>
		/// The image ID from which to create the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> ImageId { get; set; }

		/// <summary>
		/// Metadata key/value pairs to associate with the volume.
		/// Changing this updates the existing volume metadata.
		/// </summary>
		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> Metadata { get; set; }

		/// <summary>
		/// A unique name for the volume. Changing this updates the
		/// volume's name.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// The region in which to create the volume. If
		/// omitted, the `region` argument of the provider is used. Changing this
		/// creates a new volume.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// The size of the volume to create (in gigabytes).
		/// </summary>
		public Pulumi.IO<int> Size { get; set; }

		/// <summary>
		/// The snapshot ID from which to create the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> SnapshotId { get; set; }

		/// <summary>
		/// The volume ID to replicate with.
		/// </summary>
		public Pulumi.IO<string> SourceReplica { get; set; }

		/// <summary>
		/// The volume ID from which to create the volume.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> SourceVolId { get; set; }

		/// <summary>
		/// The type of volume to create.
		/// Changing this creates a new volume.
		/// </summary>
		public Pulumi.IO<string> VolumeType { get; set; }

		public Volume(string name, VolumeArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:blockstorage/volume:Volume", name, SerialiseArgs(args), opts) {
			AvailabilityZone = Outputs["availabilityZone"].Select(item => Protobuf.ToString(item));
			ConsistencyGroupId = Outputs["consistencyGroupId"].Select(item => Protobuf.ToString(item));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			EnableOnlineResize = Outputs["enableOnlineResize"].Select(item => Protobuf.ToBool(item));
			ImageId = Outputs["imageId"].Select(item => Protobuf.ToString(item));
			Metadata = Outputs["metadata"].Select(item => Protobuf.ToMap(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			Size = Outputs["size"].Select(item => Protobuf.ToInt(item));
			SnapshotId = Outputs["snapshotId"].Select(item => Protobuf.ToString(item));
			SourceReplica = Outputs["sourceReplica"].Select(item => Protobuf.ToString(item));
			SourceVolId = Outputs["sourceVolId"].Select(item => Protobuf.ToString(item));
			VolumeType = Outputs["volumeType"].Select(item => Protobuf.ToString(item));
			Attachments = Outputs["attachments"].Select(item => Protobuf.ToList(item, item1 => VolumeAttachment.FromProtobuf(item1)));
			AvailabilityZone = Outputs["availabilityZone"].Select(item => Protobuf.ToString(item));
			ConsistencyGroupId = Outputs["consistencyGroupId"].Select(item => Protobuf.ToString(item));
			Description = Outputs["description"].Select(item => Protobuf.ToString(item));
			EnableOnlineResize = Outputs["enableOnlineResize"].Select(item => Protobuf.ToBool(item));
			ImageId = Outputs["imageId"].Select(item => Protobuf.ToString(item));
			Metadata = Outputs["metadata"].Select(item => Protobuf.ToMap(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			Size = Outputs["size"].Select(item => Protobuf.ToInt(item));
			SnapshotId = Outputs["snapshotId"].Select(item => Protobuf.ToString(item));
			SourceReplica = Outputs["sourceReplica"].Select(item => Protobuf.ToString(item));
			SourceVolId = Outputs["sourceVolId"].Select(item => Protobuf.ToString(item));
			VolumeType = Outputs["volumeType"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(VolumeArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["availabilityZone"] = Protobuf.ToProtobuf(args.AvailabilityZone);
			props["consistencyGroupId"] = Protobuf.ToProtobuf(args.ConsistencyGroupId);
			props["description"] = Protobuf.ToProtobuf(args.Description);
			props["enableOnlineResize"] = Protobuf.ToProtobuf(args.EnableOnlineResize);
			props["imageId"] = Protobuf.ToProtobuf(args.ImageId);
			props["metadata"] = Protobuf.ToProtobuf(args.Metadata);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["size"] = Protobuf.ToProtobuf(args.Size);
			props["snapshotId"] = Protobuf.ToProtobuf(args.SnapshotId);
			props["sourceReplica"] = Protobuf.ToProtobuf(args.SourceReplica);
			props["sourceVolId"] = Protobuf.ToProtobuf(args.SourceVolId);
			props["volumeType"] = Protobuf.ToProtobuf(args.VolumeType);
			props["attachments"] = null; //out
			return props;
		} // SerialiseArgs

	} // Volume
} // Pulumi.Openstack.Blockstorage
