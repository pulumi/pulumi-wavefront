// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information about all Wavefront events.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * // Get the information about all events
 * const example = pulumi.output(wavefront.getEvents({
 *     earliestStartTimeEpochMillis: 1665427195,
 *     latestStartTimeEpochMillis: 1665427195,
 *     limit: 10,
 *     offset: 0,
 * }));
 * ```
 */
export function getEvents(args: GetEventsArgs, opts?: pulumi.InvokeOptions): Promise<GetEventsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("wavefront:index/getEvents:getEvents", {
        "earliestStartTimeEpochMillis": args.earliestStartTimeEpochMillis,
        "latestStartTimeEpochMillis": args.latestStartTimeEpochMillis,
        "limit": args.limit,
        "offset": args.offset,
    }, opts);
}

/**
 * A collection of arguments for invoking getEvents.
 */
export interface GetEventsArgs {
    /**
     * The earliest start time in epoch milliseconds.
     */
    earliestStartTimeEpochMillis: number;
    /**
     * The latest start time in epoch milliseconds.
     */
    latestStartTimeEpochMillis: number;
    /**
     * Limit is the maximum number of results to be returned. Defaults to 100.
     */
    limit?: number;
    /**
     * Offset is the offset from the first result to be returned. Defaults to 0.
     */
    offset?: number;
}

/**
 * A collection of values returned by getEvents.
 */
export interface GetEventsResult {
    /**
     * Earliest start time in epoch milliseconds.
     */
    readonly earliestStartTimeEpochMillis: number;
    /**
     * List of all events in Wavefront. For each event you will see a list of attributes.
     */
    readonly events: outputs.GetEventsEvent[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Latest start time in epoch milliseconds.
     */
    readonly latestStartTimeEpochMillis: number;
    readonly limit?: number;
    readonly offset?: number;
}

export function getEventsOutput(args: GetEventsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEventsResult> {
    return pulumi.output(args).apply(a => getEvents(a, opts))
}

/**
 * A collection of arguments for invoking getEvents.
 */
export interface GetEventsOutputArgs {
    /**
     * The earliest start time in epoch milliseconds.
     */
    earliestStartTimeEpochMillis: pulumi.Input<number>;
    /**
     * The latest start time in epoch milliseconds.
     */
    latestStartTimeEpochMillis: pulumi.Input<number>;
    /**
     * Limit is the maximum number of results to be returned. Defaults to 100.
     */
    limit?: pulumi.Input<number>;
    /**
     * Offset is the offset from the first result to be returned. Defaults to 0.
     */
    offset?: pulumi.Input<number>;
}