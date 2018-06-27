# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Image(pulumi.CustomResource):
    """
    Manages a V2 Image resource within OpenStack Glance.
    """
    def __init__(__self__, __name__, __opts__=None, container_format=None, disk_format=None, image_cache_path=None, image_source_url=None, local_file_path=None, min_disk_gb=None, min_ram_mb=None, name=None, properties=None, protected=None, region=None, tags=None, verify_checksum=None, visibility=None):
        """Create a Image resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not container_format:
            raise TypeError('Missing required property container_format')
        elif not isinstance(container_format, basestring):
            raise TypeError('Expected property container_format to be a basestring')
        __self__.container_format = container_format
        """
        The container format. Must be one of
        "ami", "ari", "aki", "bare", "ovf".
        """
        __props__['containerFormat'] = container_format

        if not disk_format:
            raise TypeError('Missing required property disk_format')
        elif not isinstance(disk_format, basestring):
            raise TypeError('Expected property disk_format to be a basestring')
        __self__.disk_format = disk_format
        """
        The disk format. Must be one of
        "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
        """
        __props__['diskFormat'] = disk_format

        if image_cache_path and not isinstance(image_cache_path, basestring):
            raise TypeError('Expected property image_cache_path to be a basestring')
        __self__.image_cache_path = image_cache_path
        """
        This is the directory where the images will
        be downloaded. Images will be stored with a filename corresponding to
        the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"
        """
        __props__['imageCachePath'] = image_cache_path

        if image_source_url and not isinstance(image_source_url, basestring):
            raise TypeError('Expected property image_source_url to be a basestring')
        __self__.image_source_url = image_source_url
        """
        This is the url of the raw image that will
        be downloaded in the `image_cache_path` before being uploaded to Glance.
        Glance is able to download image from internet but the `gophercloud` library
        does not yet provide a way to do so.
        Conflicts with `local_file_path`.
        """
        __props__['imageSourceUrl'] = image_source_url

        if local_file_path and not isinstance(local_file_path, basestring):
            raise TypeError('Expected property local_file_path to be a basestring')
        __self__.local_file_path = local_file_path
        """
        This is the filepath of the raw image file
        that will be uploaded to Glance. Conflicts with `image_source_url`.
        """
        __props__['localFilePath'] = local_file_path

        if min_disk_gb and not isinstance(min_disk_gb, int):
            raise TypeError('Expected property min_disk_gb to be a int')
        __self__.min_disk_gb = min_disk_gb
        """
        Amount of disk space (in GB) required to boot image.
        Defaults to 0.
        """
        __props__['minDiskGb'] = min_disk_gb

        if min_ram_mb and not isinstance(min_ram_mb, int):
            raise TypeError('Expected property min_ram_mb to be a int')
        __self__.min_ram_mb = min_ram_mb
        """
        Amount of ram (in MB) required to boot image.
        Defauts to 0.
        """
        __props__['minRamMb'] = min_ram_mb

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the image.
        """
        __props__['name'] = name

        if properties and not isinstance(properties, dict):
            raise TypeError('Expected property properties to be a dict')
        __self__.properties = properties
        """
        A map of key/value pairs to set freeform
        information about an image.
        """
        __props__['properties'] = properties

        if protected and not isinstance(protected, bool):
            raise TypeError('Expected property protected to be a bool')
        __self__.protected = protected
        """
        If true, image will not be deletable.
        Defaults to false.
        """
        __props__['protected'] = protected

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V2 Glance client.
        A Glance client is needed to create an Image that can be used with
        a compute instance. If omitted, the `region` argument of the provider
        is used. Changing this creates a new Image.
        """
        __props__['region'] = region

        if tags and not isinstance(tags, list):
            raise TypeError('Expected property tags to be a list')
        __self__.tags = tags
        """
        The tags of the image. It must be a list of strings.
        At this time, it is not possible to delete all tags of an image.
        """
        __props__['tags'] = tags

        if verify_checksum and not isinstance(verify_checksum, bool):
            raise TypeError('Expected property verify_checksum to be a bool')
        __self__.verify_checksum = verify_checksum
        """
        If false, the checksum will not be verified
        once the image is finished uploading. Defaults to true.
        """
        __props__['verifyChecksum'] = verify_checksum

        if visibility and not isinstance(visibility, basestring):
            raise TypeError('Expected property visibility to be a basestring')
        __self__.visibility = visibility
        """
        The visibility of the image. Must be one of
        "public", "private", "community", or "shared". The ability to set the
        visibility depends upon the configuration of the OpenStack cloud.
        """
        __props__['visibility'] = visibility

        __self__.checksum = pulumi.runtime.UNKNOWN
        """
        The checksum of the data associated with the image.
        """
        __self__.created_at = pulumi.runtime.UNKNOWN
        """
        The date the image was created.
        """
        __self__.file = pulumi.runtime.UNKNOWN
        """
        the trailing path after the glance
        endpoint that represent the location of the image
        or the path to retrieve it.
        """
        __self__.metadata = pulumi.runtime.UNKNOWN
        """
        The metadata associated with the image.
        Image metadata allow for meaningfully define the image properties
        and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
        """
        __self__.owner = pulumi.runtime.UNKNOWN
        """
        The id of the openstack user who owns the image.
        """
        __self__.schema = pulumi.runtime.UNKNOWN
        """
        The path to the JSON-schema that represent
        the image or image
        """
        __self__.size_bytes = pulumi.runtime.UNKNOWN
        """
        The size in bytes of the data associated with the image.
        """
        __self__.status = pulumi.runtime.UNKNOWN
        """
        The status of the image. It can be "queued", "active"
        or "saving".
        """
        __self__.update_at = pulumi.runtime.UNKNOWN
        """
        The date the image was last updated.
        """

        super(Image, __self__).__init__(
            'openstack:images/image:Image',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'checksum' in outs:
            self.checksum = outs['checksum']
        if 'containerFormat' in outs:
            self.container_format = outs['containerFormat']
        if 'createdAt' in outs:
            self.created_at = outs['createdAt']
        if 'diskFormat' in outs:
            self.disk_format = outs['diskFormat']
        if 'file' in outs:
            self.file = outs['file']
        if 'imageCachePath' in outs:
            self.image_cache_path = outs['imageCachePath']
        if 'imageSourceUrl' in outs:
            self.image_source_url = outs['imageSourceUrl']
        if 'localFilePath' in outs:
            self.local_file_path = outs['localFilePath']
        if 'metadata' in outs:
            self.metadata = outs['metadata']
        if 'minDiskGb' in outs:
            self.min_disk_gb = outs['minDiskGb']
        if 'minRamMb' in outs:
            self.min_ram_mb = outs['minRamMb']
        if 'name' in outs:
            self.name = outs['name']
        if 'owner' in outs:
            self.owner = outs['owner']
        if 'properties' in outs:
            self.properties = outs['properties']
        if 'protected' in outs:
            self.protected = outs['protected']
        if 'region' in outs:
            self.region = outs['region']
        if 'schema' in outs:
            self.schema = outs['schema']
        if 'sizeBytes' in outs:
            self.size_bytes = outs['sizeBytes']
        if 'status' in outs:
            self.status = outs['status']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'updateAt' in outs:
            self.update_at = outs['updateAt']
        if 'verifyChecksum' in outs:
            self.verify_checksum = outs['verifyChecksum']
        if 'visibility' in outs:
            self.visibility = outs['visibility']