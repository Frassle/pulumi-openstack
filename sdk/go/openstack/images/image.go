// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 Image resource within OpenStack Glance.
type Image struct {
	s *pulumi.ResourceState
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOpt) (*Image, error) {
	if args == nil || args.ContainerFormat == nil {
		return nil, errors.New("missing required argument 'ContainerFormat'")
	}
	if args == nil || args.DiskFormat == nil {
		return nil, errors.New("missing required argument 'DiskFormat'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["containerFormat"] = nil
		inputs["diskFormat"] = nil
		inputs["imageCachePath"] = nil
		inputs["imageSourceUrl"] = nil
		inputs["localFilePath"] = nil
		inputs["minDiskGb"] = nil
		inputs["minRamMb"] = nil
		inputs["name"] = nil
		inputs["properties"] = nil
		inputs["protected"] = nil
		inputs["region"] = nil
		inputs["tags"] = nil
		inputs["verifyChecksum"] = nil
		inputs["visibility"] = nil
	} else {
		inputs["containerFormat"] = args.ContainerFormat
		inputs["diskFormat"] = args.DiskFormat
		inputs["imageCachePath"] = args.ImageCachePath
		inputs["imageSourceUrl"] = args.ImageSourceUrl
		inputs["localFilePath"] = args.LocalFilePath
		inputs["minDiskGb"] = args.MinDiskGb
		inputs["minRamMb"] = args.MinRamMb
		inputs["name"] = args.Name
		inputs["properties"] = args.Properties
		inputs["protected"] = args.Protected
		inputs["region"] = args.Region
		inputs["tags"] = args.Tags
		inputs["verifyChecksum"] = args.VerifyChecksum
		inputs["visibility"] = args.Visibility
	}
	inputs["checksum"] = nil
	inputs["createdAt"] = nil
	inputs["file"] = nil
	inputs["metadata"] = nil
	inputs["owner"] = nil
	inputs["schema"] = nil
	inputs["sizeBytes"] = nil
	inputs["status"] = nil
	inputs["updateAt"] = nil
	s, err := ctx.RegisterResource("openstack:images/image:Image", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Image{s: s}, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ImageState, opts ...pulumi.ResourceOpt) (*Image, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["checksum"] = state.Checksum
		inputs["containerFormat"] = state.ContainerFormat
		inputs["createdAt"] = state.CreatedAt
		inputs["diskFormat"] = state.DiskFormat
		inputs["file"] = state.File
		inputs["imageCachePath"] = state.ImageCachePath
		inputs["imageSourceUrl"] = state.ImageSourceUrl
		inputs["localFilePath"] = state.LocalFilePath
		inputs["metadata"] = state.Metadata
		inputs["minDiskGb"] = state.MinDiskGb
		inputs["minRamMb"] = state.MinRamMb
		inputs["name"] = state.Name
		inputs["owner"] = state.Owner
		inputs["properties"] = state.Properties
		inputs["protected"] = state.Protected
		inputs["region"] = state.Region
		inputs["schema"] = state.Schema
		inputs["sizeBytes"] = state.SizeBytes
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
		inputs["updateAt"] = state.UpdateAt
		inputs["verifyChecksum"] = state.VerifyChecksum
		inputs["visibility"] = state.Visibility
	}
	s, err := ctx.ReadResource("openstack:images/image:Image", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Image{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Image) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Image) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The checksum of the data associated with the image.
func (r *Image) Checksum() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["checksum"])
}

// The container format. Must be one of
// "ami", "ari", "aki", "bare", "ovf".
func (r *Image) ContainerFormat() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["containerFormat"])
}

// The date the image was created.
func (r *Image) CreatedAt() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createdAt"])
}

// The disk format. Must be one of
// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
func (r *Image) DiskFormat() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["diskFormat"])
}

// the trailing path after the glance
// endpoint that represent the location of the image
// or the path to retrieve it.
func (r *Image) File() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["file"])
}

// This is the directory where the images will
// be downloaded. Images will be stored with a filename corresponding to
// the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"
func (r *Image) ImageCachePath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["imageCachePath"])
}

// This is the url of the raw image that will
// be downloaded in the `image_cache_path` before being uploaded to Glance.
// Glance is able to download image from internet but the `gophercloud` library
// does not yet provide a way to do so.
// Conflicts with `local_file_path`.
func (r *Image) ImageSourceUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["imageSourceUrl"])
}

// This is the filepath of the raw image file
// that will be uploaded to Glance. Conflicts with `image_source_url`.
func (r *Image) LocalFilePath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["localFilePath"])
}

// The metadata associated with the image.
// Image metadata allow for meaningfully define the image properties
// and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
func (r *Image) Metadata() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["metadata"])
}

// Amount of disk space (in GB) required to boot image.
// Defaults to 0.
func (r *Image) MinDiskGb() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["minDiskGb"])
}

// Amount of ram (in MB) required to boot image.
// Defauts to 0.
func (r *Image) MinRamMb() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["minRamMb"])
}

// The name of the image.
func (r *Image) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The id of the openstack user who owns the image.
func (r *Image) Owner() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["owner"])
}

// A map of key/value pairs to set freeform
// information about an image.
func (r *Image) Properties() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["properties"])
}

// If true, image will not be deletable.
// Defaults to false.
func (r *Image) Protected() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["protected"])
}

// The region in which to obtain the V2 Glance client.
// A Glance client is needed to create an Image that can be used with
// a compute instance. If omitted, the `region` argument of the provider
// is used. Changing this creates a new Image.
func (r *Image) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The path to the JSON-schema that represent
// the image or image
func (r *Image) Schema() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["schema"])
}

// The size in bytes of the data associated with the image.
func (r *Image) SizeBytes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["sizeBytes"])
}

// The status of the image. It can be "queued", "active"
// or "saving".
func (r *Image) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// The tags of the image. It must be a list of strings.
// At this time, it is not possible to delete all tags of an image.
func (r *Image) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// The date the image was last updated.
func (r *Image) UpdateAt() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["updateAt"])
}

// If false, the checksum will not be verified
// once the image is finished uploading. Defaults to true.
func (r *Image) VerifyChecksum() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["verifyChecksum"])
}

// The visibility of the image. Must be one of
// "public", "private", "community", or "shared". The ability to set the
// visibility depends upon the configuration of the OpenStack cloud.
func (r *Image) Visibility() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["visibility"])
}

// Input properties used for looking up and filtering Image resources.
type ImageState struct {
	// The checksum of the data associated with the image.
	Checksum interface{}
	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	ContainerFormat interface{}
	// The date the image was created.
	CreatedAt interface{}
	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	DiskFormat interface{}
	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File interface{}
	// This is the directory where the images will
	// be downloaded. Images will be stored with a filename corresponding to
	// the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"
	ImageCachePath interface{}
	// This is the url of the raw image that will
	// be downloaded in the `image_cache_path` before being uploaded to Glance.
	// Glance is able to download image from internet but the `gophercloud` library
	// does not yet provide a way to do so.
	// Conflicts with `local_file_path`.
	ImageSourceUrl interface{}
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `image_source_url`.
	LocalFilePath interface{}
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
	Metadata interface{}
	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	MinDiskGb interface{}
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb interface{}
	// The name of the image.
	Name interface{}
	// The id of the openstack user who owns the image.
	Owner interface{}
	// A map of key/value pairs to set freeform
	// information about an image.
	Properties interface{}
	// If true, image will not be deletable.
	// Defaults to false.
	Protected interface{}
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new Image.
	Region interface{}
	// The path to the JSON-schema that represent
	// the image or image
	Schema interface{}
	// The size in bytes of the data associated with the image.
	SizeBytes interface{}
	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status interface{}
	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	Tags interface{}
	// The date the image was last updated.
	UpdateAt interface{}
	// If false, the checksum will not be verified
	// once the image is finished uploading. Defaults to true.
	VerifyChecksum interface{}
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility interface{}
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	ContainerFormat interface{}
	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	DiskFormat interface{}
	// This is the directory where the images will
	// be downloaded. Images will be stored with a filename corresponding to
	// the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"
	ImageCachePath interface{}
	// This is the url of the raw image that will
	// be downloaded in the `image_cache_path` before being uploaded to Glance.
	// Glance is able to download image from internet but the `gophercloud` library
	// does not yet provide a way to do so.
	// Conflicts with `local_file_path`.
	ImageSourceUrl interface{}
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `image_source_url`.
	LocalFilePath interface{}
	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	MinDiskGb interface{}
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb interface{}
	// The name of the image.
	Name interface{}
	// A map of key/value pairs to set freeform
	// information about an image.
	Properties interface{}
	// If true, image will not be deletable.
	// Defaults to false.
	Protected interface{}
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new Image.
	Region interface{}
	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	Tags interface{}
	// If false, the checksum will not be verified
	// once the image is finished uploading. Defaults to true.
	VerifyChecksum interface{}
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility interface{}
}
