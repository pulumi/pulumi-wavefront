// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get all users in Wavefront.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * // Get all users
 * const users = pulumi.output(wavefront.getUsers());
 * ```
 */
export function getUsers(opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("wavefront:index/getUsers:getUsers", {
    }, opts);
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of all users in Wavefront.
     */
    readonly users: outputs.GetUsersUser[];
}