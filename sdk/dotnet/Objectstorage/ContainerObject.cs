// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Linq;

namespace Pulumi.Openstack.Objectstorage {

	/// <summary>
	/// The set of arguments for constructing a ContainerObject resource.
	/// </summary>
	public struct ContainerObjectArgs {
		/// <summary>
		/// A unique (within an account) name for the container. 
		/// The container name must be from 1 to 256 characters long and can start
		/// with any character and contain any pattern. Character set must be UTF-8.
		/// The container name cannot contain a slash (/) character because this
		/// character delimits the container and object name. For example, the path
		/// /v1/account/www/pages specifies the www container, not the www/pages container.
		/// </summary>
		public Pulumi.IO<string> ContainerName { get; set; }

		/// <summary>
		/// A string representing the content of the object. Conflicts with
		/// `source` and `copy_from`.
		/// </summary>
		public Pulumi.IO<string> Content { get; set; }

		/// <summary>
		/// A string which specifies the override behavior for 
		/// the browser. For example, this header might specify that the browser use a download
		/// program to save this file rather than show the file, which is the default.
		/// </summary>
		public Pulumi.IO<string> ContentDisposition { get; set; }

		/// <summary>
		/// A string representing the value of the Content-Encoding
		/// metadata.
		/// </summary>
		public Pulumi.IO<string> ContentEncoding { get; set; }

		/// <summary>
		/// A string which sets the MIME type for the object.
		/// </summary>
		public Pulumi.IO<string> ContentType { get; set; }

		/// <summary>
		/// A string representing the name of an object 
		/// used to create the new object by copying the `copy_from` object. The value is in form
		/// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
		/// container and object before you include them in the header. Conflicts with `source` and
		/// `content`.
		/// </summary>
		public Pulumi.IO<string> CopyFrom { get; set; }

		/// <summary>
		/// An integer representing the number of seconds after which the
		/// system removes the object. Internally, the Object Storage system stores this value in
		/// the X-Delete-At metadata item.
		/// </summary>
		public Pulumi.IO<int> DeleteAfter { get; set; }

		/// <summary>
		/// An string representing the date when the system removes the object. 
		/// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
		/// </summary>
		public Pulumi.IO<string> DeleteAt { get; set; }

		/// <summary>
		/// If set to true, Object Storage guesses the content 
		/// type based on the file extension and ignores the value sent in the Content-Type
		/// header, if present.
		/// </summary>
		public Pulumi.IO<bool> DetectContentType { get; set; }

		/// <summary>
		/// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
		/// </summary>
		public Pulumi.IO<string> Etag { get; set; }

		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> Metadata { get; set; }

		/// <summary>
		/// A unique name for the object.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// A string set to specify that this is a dynamic large 
		/// object manifest object. The value is the container and object name prefix of the
		/// segment objects in the form container/prefix. You must UTF-8-encode and then
		/// URL-encode the names of the container and prefix before you include them in this
		/// header.
		/// </summary>
		public Pulumi.IO<string> ObjectManifest { get; set; }

		/// <summary>
		/// The region in which to create the container. If
		/// omitted, the `region` argument of the provider is used. Changing this
		/// creates a new container.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// A string representing the local path of a file which will be used
		/// as the object's content. Conflicts with `source` and `copy_from`.
		/// </summary>
		public Pulumi.IO<string> Source { get; set; }

	} // ContainerObjectArgs

	/// <summary>
	/// Manages a V1 container object resource within OpenStack.
	/// </summary>
	public class ContainerObject : Pulumi.CustomResource {
		/// <summary>
		/// A unique (within an account) name for the container. 
		/// The container name must be from 1 to 256 characters long and can start
		/// with any character and contain any pattern. Character set must be UTF-8.
		/// The container name cannot contain a slash (/) character because this
		/// character delimits the container and object name. For example, the path
		/// /v1/account/www/pages specifies the www container, not the www/pages container.
		/// </summary>
		public Pulumi.IO<string> ContainerName { get; set; }

		/// <summary>
		/// A string representing the content of the object. Conflicts with
		/// `source` and `copy_from`.
		/// </summary>
		public Pulumi.IO<string> Content { get; set; }

		/// <summary>
		/// A string which specifies the override behavior for 
		/// the browser. For example, this header might specify that the browser use a download
		/// program to save this file rather than show the file, which is the default.
		/// </summary>
		public Pulumi.IO<string> ContentDisposition { get; set; }

		/// <summary>
		/// A string representing the value of the Content-Encoding
		/// metadata.
		/// </summary>
		public Pulumi.IO<string> ContentEncoding { get; set; }

		/// <summary>
		/// If the operation succeeds, this value is zero (0) or the 
		/// length of informational or error text in the response body.
		/// </summary>
		public Pulumi.IO<int> ContentLength { get; set; }

		/// <summary>
		/// A string which sets the MIME type for the object.
		/// </summary>
		public Pulumi.IO<string> ContentType { get; set; }

		/// <summary>
		/// A string representing the name of an object 
		/// used to create the new object by copying the `copy_from` object. The value is in form
		/// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
		/// container and object before you include them in the header. Conflicts with `source` and
		/// `content`.
		/// </summary>
		public Pulumi.IO<string> CopyFrom { get; set; }

		/// <summary>
		/// The date and time the system responded to the request, using the preferred 
		/// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
		/// time is always in UTC.
		/// </summary>
		public Pulumi.IO<string> Date { get; set; }

		/// <summary>
		/// An integer representing the number of seconds after which the
		/// system removes the object. Internally, the Object Storage system stores this value in
		/// the X-Delete-At metadata item.
		/// </summary>
		public Pulumi.IO<int> DeleteAfter { get; set; }

		/// <summary>
		/// An string representing the date when the system removes the object. 
		/// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
		/// </summary>
		public Pulumi.IO<string> DeleteAt { get; set; }

		/// <summary>
		/// If set to true, Object Storage guesses the content 
		/// type based on the file extension and ignores the value sent in the Content-Type
		/// header, if present.
		/// </summary>
		public Pulumi.IO<bool> DetectContentType { get; set; }

		/// <summary>
		/// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
		/// </summary>
		public Pulumi.IO<string> Etag { get; set; }

		/// <summary>
		/// The date and time when the object was last modified. The date and time 
		/// stamp format is ISO 8601:
		/// CCYY-MM-DDThh:mm:ss±hh:mm
		/// For example, 2015-08-27T09:49:58-05:00.
		/// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
		/// example, the offset value is -05:00.
		/// </summary>
		public Pulumi.IO<string> LastModified { get; set; }

		public Pulumi.IO<System.Collections.Generic.Dictionary<string, string>> Metadata { get; set; }

		/// <summary>
		/// A unique name for the object.
		/// </summary>
		public Pulumi.IO<string> Name { get; set; }

		/// <summary>
		/// A string set to specify that this is a dynamic large 
		/// object manifest object. The value is the container and object name prefix of the
		/// segment objects in the form container/prefix. You must UTF-8-encode and then
		/// URL-encode the names of the container and prefix before you include them in this
		/// header.
		/// </summary>
		public Pulumi.IO<string> ObjectManifest { get; set; }

		/// <summary>
		/// The region in which to create the container. If
		/// omitted, the `region` argument of the provider is used. Changing this
		/// creates a new container.
		/// </summary>
		public Pulumi.IO<string> Region { get; set; }

		/// <summary>
		/// A string representing the local path of a file which will be used
		/// as the object's content. Conflicts with `source` and `copy_from`.
		/// </summary>
		public Pulumi.IO<string> Source { get; set; }

		/// <summary>
		/// A unique transaction ID for this request. Your service provider might 
		/// need this value if you report a problem.
		/// </summary>
		public Pulumi.IO<string> TransId { get; set; }

		public ContainerObject(string name, ContainerObjectArgs args, Pulumi.ResourceOptions opts = default(Pulumi.ResourceOptions))
			: base("openstack:objectstorage/containerObject:ContainerObject", name, SerialiseArgs(args), opts) {
			ContainerName = Outputs["containerName"].Select(item => Protobuf.ToString(item));
			Content = Outputs["content"].Select(item => Protobuf.ToString(item));
			ContentDisposition = Outputs["contentDisposition"].Select(item => Protobuf.ToString(item));
			ContentEncoding = Outputs["contentEncoding"].Select(item => Protobuf.ToString(item));
			ContentType = Outputs["contentType"].Select(item => Protobuf.ToString(item));
			CopyFrom = Outputs["copyFrom"].Select(item => Protobuf.ToString(item));
			DeleteAfter = Outputs["deleteAfter"].Select(item => Protobuf.ToInt(item));
			DeleteAt = Outputs["deleteAt"].Select(item => Protobuf.ToString(item));
			DetectContentType = Outputs["detectContentType"].Select(item => Protobuf.ToBool(item));
			Etag = Outputs["etag"].Select(item => Protobuf.ToString(item));
			Metadata = Outputs["metadata"].Select(item => Protobuf.ToMap(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			ObjectManifest = Outputs["objectManifest"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			Source = Outputs["source"].Select(item => Protobuf.ToString(item));
			ContainerName = Outputs["containerName"].Select(item => Protobuf.ToString(item));
			Content = Outputs["content"].Select(item => Protobuf.ToString(item));
			ContentDisposition = Outputs["contentDisposition"].Select(item => Protobuf.ToString(item));
			ContentEncoding = Outputs["contentEncoding"].Select(item => Protobuf.ToString(item));
			ContentLength = Outputs["contentLength"].Select(item => Protobuf.ToInt(item));
			ContentType = Outputs["contentType"].Select(item => Protobuf.ToString(item));
			CopyFrom = Outputs["copyFrom"].Select(item => Protobuf.ToString(item));
			Date = Outputs["date"].Select(item => Protobuf.ToString(item));
			DeleteAfter = Outputs["deleteAfter"].Select(item => Protobuf.ToInt(item));
			DeleteAt = Outputs["deleteAt"].Select(item => Protobuf.ToString(item));
			DetectContentType = Outputs["detectContentType"].Select(item => Protobuf.ToBool(item));
			Etag = Outputs["etag"].Select(item => Protobuf.ToString(item));
			LastModified = Outputs["lastModified"].Select(item => Protobuf.ToString(item));
			Metadata = Outputs["metadata"].Select(item => Protobuf.ToMap(item));
			Name = Outputs["name"].Select(item => Protobuf.ToString(item));
			ObjectManifest = Outputs["objectManifest"].Select(item => Protobuf.ToString(item));
			Region = Outputs["region"].Select(item => Protobuf.ToString(item));
			Source = Outputs["source"].Select(item => Protobuf.ToString(item));
			TransId = Outputs["transId"].Select(item => Protobuf.ToString(item));
		} // ctor

		private static Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>> SerialiseArgs(ContainerObjectArgs args) {
			var props = new Dictionary<string, Pulumi.IO<Google.Protobuf.WellKnownTypes.Value>>();
			props["containerName"] = Protobuf.ToProtobuf(args.ContainerName);
			props["content"] = Protobuf.ToProtobuf(args.Content);
			props["contentDisposition"] = Protobuf.ToProtobuf(args.ContentDisposition);
			props["contentEncoding"] = Protobuf.ToProtobuf(args.ContentEncoding);
			props["contentType"] = Protobuf.ToProtobuf(args.ContentType);
			props["copyFrom"] = Protobuf.ToProtobuf(args.CopyFrom);
			props["deleteAfter"] = Protobuf.ToProtobuf(args.DeleteAfter);
			props["deleteAt"] = Protobuf.ToProtobuf(args.DeleteAt);
			props["detectContentType"] = Protobuf.ToProtobuf(args.DetectContentType);
			props["etag"] = Protobuf.ToProtobuf(args.Etag);
			props["metadata"] = Protobuf.ToProtobuf(args.Metadata);
			props["name"] = Protobuf.ToProtobuf(args.Name);
			props["objectManifest"] = Protobuf.ToProtobuf(args.ObjectManifest);
			props["region"] = Protobuf.ToProtobuf(args.Region);
			props["source"] = Protobuf.ToProtobuf(args.Source);
			props["contentLength"] = null; //out
			props["date"] = null; //out
			props["lastModified"] = null; //out
			props["transId"] = null; //out
			return props;
		} // SerialiseArgs

	} // ContainerObject
} // Pulumi.Openstack.Objectstorage