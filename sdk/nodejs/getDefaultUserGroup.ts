// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get the Group ID of the `Everyone` group in Wavefront.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * // Get the default user group "Everyone"
 * const everyoneGroup = wavefront.getDefaultUserGroup({});
 * ```
 */
export function getDefaultUserGroup(opts?: pulumi.InvokeOptions): Promise<GetDefaultUserGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("wavefront:index/getDefaultUserGroup:getDefaultUserGroup", {
    }, opts);
}

/**
 * A collection of values returned by getDefaultUserGroup.
 */
export interface GetDefaultUserGroupResult {
    /**
     * Set to the Group ID of the `Everyone` group, suitable for referencing
     * in other resources that support group memberships.
     */
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * Use this data source to get the Group ID of the `Everyone` group in Wavefront.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * // Get the default user group "Everyone"
 * const everyoneGroup = wavefront.getDefaultUserGroup({});
 * ```
 */
export function getDefaultUserGroupOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDefaultUserGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("wavefront:index/getDefaultUserGroup:getDefaultUserGroup", {
    }, opts);
}
