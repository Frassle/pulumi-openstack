# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class ContainerObject(pulumi.CustomResource):
    """
    Manages a V1 container object resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, container_name=None, content=None, content_disposition=None, content_encoding=None, content_type=None, copy_from=None, delete_after=None, delete_at=None, detect_content_type=None, etag=None, metadata=None, name=None, object_manifest=None, region=None, source=None):
        """Create a ContainerObject resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not container_name:
            raise TypeError('Missing required property container_name')
        elif not isinstance(container_name, basestring):
            raise TypeError('Expected property container_name to be a basestring')
        __self__.container_name = container_name
        """
        A unique (within an account) name for the container. 
        The container name must be from 1 to 256 characters long and can start
        with any character and contain any pattern. Character set must be UTF-8.
        The container name cannot contain a slash (/) character because this
        character delimits the container and object name. For example, the path
        /v1/account/www/pages specifies the www container, not the www/pages container.
        """
        __props__['containerName'] = container_name

        if content and not isinstance(content, basestring):
            raise TypeError('Expected property content to be a basestring')
        __self__.content = content
        """
        A string representing the content of the object. Conflicts with
        `source` and `copy_from`.
        """
        __props__['content'] = content

        if content_disposition and not isinstance(content_disposition, basestring):
            raise TypeError('Expected property content_disposition to be a basestring')
        __self__.content_disposition = content_disposition
        """
        A string which specifies the override behavior for 
        the browser. For example, this header might specify that the browser use a download
        program to save this file rather than show the file, which is the default.
        """
        __props__['contentDisposition'] = content_disposition

        if content_encoding and not isinstance(content_encoding, basestring):
            raise TypeError('Expected property content_encoding to be a basestring')
        __self__.content_encoding = content_encoding
        """
        A string representing the value of the Content-Encoding
        metadata.
        """
        __props__['contentEncoding'] = content_encoding

        if content_type and not isinstance(content_type, basestring):
            raise TypeError('Expected property content_type to be a basestring')
        __self__.content_type = content_type
        """
        A string which sets the MIME type for the object.
        """
        __props__['contentType'] = content_type

        if copy_from and not isinstance(copy_from, basestring):
            raise TypeError('Expected property copy_from to be a basestring')
        __self__.copy_from = copy_from
        """
        A string representing the name of an object 
        used to create the new object by copying the `copy_from` object. The value is in form
        {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
        container and object before you include them in the header. Conflicts with `source` and
        `content`.
        """
        __props__['copyFrom'] = copy_from

        if delete_after and not isinstance(delete_after, int):
            raise TypeError('Expected property delete_after to be a int')
        __self__.delete_after = delete_after
        """
        An integer representing the number of seconds after which the
        system removes the object. Internally, the Object Storage system stores this value in
        the X-Delete-At metadata item.
        """
        __props__['deleteAfter'] = delete_after

        if delete_at and not isinstance(delete_at, basestring):
            raise TypeError('Expected property delete_at to be a basestring')
        __self__.delete_at = delete_at
        """
        An string representing the date when the system removes the object. 
        For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
        """
        __props__['deleteAt'] = delete_at

        if detect_content_type and not isinstance(detect_content_type, bool):
            raise TypeError('Expected property detect_content_type to be a bool')
        __self__.detect_content_type = detect_content_type
        """
        If set to true, Object Storage guesses the content 
        type based on the file extension and ignores the value sent in the Content-Type
        header, if present.
        """
        __props__['detectContentType'] = detect_content_type

        if etag and not isinstance(etag, basestring):
            raise TypeError('Expected property etag to be a basestring')
        __self__.etag = etag
        """
        Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
        """
        __props__['etag'] = etag

        if metadata and not isinstance(metadata, dict):
            raise TypeError('Expected property metadata to be a dict')
        __self__.metadata = metadata
        __props__['metadata'] = metadata

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        A unique name for the object.
        """
        __props__['name'] = name

        if object_manifest and not isinstance(object_manifest, basestring):
            raise TypeError('Expected property object_manifest to be a basestring')
        __self__.object_manifest = object_manifest
        """
        A string set to specify that this is a dynamic large 
        object manifest object. The value is the container and object name prefix of the
        segment objects in the form container/prefix. You must UTF-8-encode and then
        URL-encode the names of the container and prefix before you include them in this
        header.
        """
        __props__['objectManifest'] = object_manifest

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to create the container. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new container.
        """
        __props__['region'] = region

        if source and not isinstance(source, basestring):
            raise TypeError('Expected property source to be a basestring')
        __self__.source = source
        """
        A string representing the local path of a file which will be used
        as the object's content. Conflicts with `source` and `copy_from`.
        """
        __props__['source'] = source

        __self__.content_length = pulumi.runtime.UNKNOWN
        """
        If the operation succeeds, this value is zero (0) or the 
        length of informational or error text in the response body.
        """
        __self__.date = pulumi.runtime.UNKNOWN
        """
        The date and time the system responded to the request, using the preferred 
        format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
        time is always in UTC.
        """
        __self__.last_modified = pulumi.runtime.UNKNOWN
        """
        The date and time when the object was last modified. The date and time 
        stamp format is ISO 8601:
        CCYY-MM-DDThh:mm:ss±hh:mm
        For example, 2015-08-27T09:49:58-05:00.
        The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
        example, the offset value is -05:00.
        """
        __self__.trans_id = pulumi.runtime.UNKNOWN
        """
        A unique transaction ID for this request. Your service provider might 
        need this value if you report a problem.
        """

        super(ContainerObject, __self__).__init__(
            'openstack:objectstorage/containerObject:ContainerObject',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'containerName' in outs:
            self.container_name = outs['containerName']
        if 'content' in outs:
            self.content = outs['content']
        if 'contentDisposition' in outs:
            self.content_disposition = outs['contentDisposition']
        if 'contentEncoding' in outs:
            self.content_encoding = outs['contentEncoding']
        if 'contentLength' in outs:
            self.content_length = outs['contentLength']
        if 'contentType' in outs:
            self.content_type = outs['contentType']
        if 'copyFrom' in outs:
            self.copy_from = outs['copyFrom']
        if 'date' in outs:
            self.date = outs['date']
        if 'deleteAfter' in outs:
            self.delete_after = outs['deleteAfter']
        if 'deleteAt' in outs:
            self.delete_at = outs['deleteAt']
        if 'detectContentType' in outs:
            self.detect_content_type = outs['detectContentType']
        if 'etag' in outs:
            self.etag = outs['etag']
        if 'lastModified' in outs:
            self.last_modified = outs['lastModified']
        if 'metadata' in outs:
            self.metadata = outs['metadata']
        if 'name' in outs:
            self.name = outs['name']
        if 'objectManifest' in outs:
            self.object_manifest = outs['objectManifest']
        if 'region' in outs:
            self.region = outs['region']
        if 'source' in outs:
            self.source = outs['source']
        if 'transId' in outs:
            self.trans_id = outs['transId']
