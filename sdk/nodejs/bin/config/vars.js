"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
let __config = new pulumi.Config("openstack");
/**
 * The Identity authentication URL.
 */
exports.authUrl = __config.get("authUrl");
/**
 * A Custom CA certificate.
 */
exports.cacertFile = __config.get("cacertFile");
/**
 * A client certificate to authenticate with.
 */
exports.cert = __config.get("cert");
/**
 * An entry in a `clouds.yaml` file to use.
 */
exports.cloud = __config.get("cloud");
/**
 * The ID of the Domain to scope to (Identity v3).
 */
exports.domainId = __config.get("domainId");
/**
 * The name of the Domain to scope to (Identity v3).
 */
exports.domainName = __config.get("domainName");
exports.endpointType = __config.get("endpointType");
/**
 * Trust self-signed certificates.
 */
exports.insecure = __config.getObject("insecure");
/**
 * A client private key to authenticate with.
 */
exports.key = __config.get("key");
/**
 * Password to login with.
 */
exports.password = __config.get("password");
/**
 * The ID of the domain where the proejct resides (Identity v3).
 */
exports.projectDomainId = __config.get("projectDomainId");
/**
 * The name of the domain where the project resides (Identity v3).
 */
exports.projectDomainName = __config.get("projectDomainName");
/**
 * The OpenStack region to connect to.
 */
exports.region = __config.get("region");
/**
 * Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
 */
exports.swauth = __config.getObject("swauth");
/**
 * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
exports.tenantId = __config.get("tenantId");
/**
 * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
exports.tenantName = __config.get("tenantName");
/**
 * Authentication token to use as an alternative to username/password.
 */
exports.token = __config.get("token");
/**
 * If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
 */
exports.useOctavia = __config.getObject("useOctavia");
/**
 * The ID of the domain where the user resides (Identity v3).
 */
exports.userDomainId = __config.get("userDomainId");
/**
 * The name of the domain where the user resides (Identity v3).
 */
exports.userDomainName = __config.get("userDomainName");
/**
 * Username to login with.
 */
exports.userId = __config.get("userId");
/**
 * Username to login with.
 */
exports.userName = __config.get("userName");
//# sourceMappingURL=vars.js.map