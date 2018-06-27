import * as pulumi from "@pulumi/pulumi";
/**
 * Use this data source to get the ID of an available OpenStack image.
 */
export declare function getImage(args?: GetImageArgs): Promise<GetImageResult>;
/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * The status of the image. Must be one of
     * "accepted", "pending", "rejected", or "all".
     */
    readonly memberStatus?: pulumi.Input<string>;
    /**
     * If more than one result is returned, use the most
     * recent image.
     */
    readonly mostRecent?: pulumi.Input<boolean>;
    /**
     * The name of the image.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The owner (UUID) of the image.
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * a map of key/value pairs to match an image with.
     * All specified properties must be matched.
     */
    readonly properties?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The maximum size (in bytes) of the image to return.
     */
    readonly sizeMax?: pulumi.Input<number>;
    /**
     * The minimum size (in bytes) of the image to return.
     */
    readonly sizeMin?: pulumi.Input<number>;
    /**
     * Order the results in either `asc` or `desc`.
     */
    readonly sortDirection?: pulumi.Input<string>;
    /**
     * Sort images based on a certain key. Defaults to `name`.
     */
    readonly sortKey?: pulumi.Input<string>;
    /**
     * Search for images with a specific tag.
     */
    readonly tag?: pulumi.Input<string>;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". Defaults to "private".
     */
    readonly visibility?: pulumi.Input<string>;
}
/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    /**
     * The checksum of the data associated with the image.
     */
    readonly checksum: string;
    readonly containerFormat: string;
    readonly diskFormat: string;
    /**
     * the trailing path after the glance endpoint that represent the
     * location of the image or the path to retrieve it.
     */
    readonly file: string;
    /**
     * The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
     */
    readonly metadata: {
        [key: string]: any;
    };
    /**
     * The minimum amount of disk space required to use the image.
     */
    readonly minDiskGb: number;
    /**
     * The minimum amount of ram required to use the image.
     */
    readonly minRamMb: number;
    /**
     * Whether or not the image is protected.
     */
    readonly protected: boolean;
    readonly region: string;
    /**
     * The path to the JSON-schema that represent
     * the image or image
     */
    readonly schema: string;
    /**
     * The size of the image (in bytes).
     */
    readonly sizeBytes: number;
    readonly updatedAt: string;
}