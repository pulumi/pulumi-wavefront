// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Wavefront External Link Resource. This allows external links to be created, updated, and deleted.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * const basic = new wavefront.ExternalLink("basic", {
 *     name: "External Link",
 *     description: "An external link description",
 *     template: "https://example.com/source={{{source}}}&startTime={{startEpochMillis}}",
 * });
 * ```
 *
 * ## Import
 *
 * Maintenance windows can be imported by using the `id`, e.g.:
 *
 * ```sh
 * $ pulumi import wavefront:index/externalLink:ExternalLink basic fVj6fz6zYC4aBkID
 * ```
 */
export class ExternalLink extends pulumi.CustomResource {
    /**
     * Get an existing ExternalLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExternalLinkState, opts?: pulumi.CustomResourceOptions): ExternalLink {
        return new ExternalLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'wavefront:index/externalLink:ExternalLink';

    /**
     * Returns true if the given object is an instance of ExternalLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalLink.__pulumiType;
    }

    /**
     * Human-readable description for this link.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether this is a "Log Integration" subType of external link.
     */
    public readonly isLogIntegration!: pulumi.Output<boolean | undefined>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
     */
    public readonly metricFilterRegex!: pulumi.Output<string | undefined>;
    /**
     * The name of the external link.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted
     * series. This is a map from string to regular expression. The highlighted series must contain point tags whose
     * keys are present in the keys of this map and whose values match the regular expressions associated with those
     * keys in order for the link to be displayed.
     */
    public readonly pointTagFilterRegexes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
     */
    public readonly sourceFilterRegex!: pulumi.Output<string | undefined>;
    /**
     * The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
     */
    public readonly template!: pulumi.Output<string>;

    /**
     * Create a ExternalLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExternalLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExternalLinkArgs | ExternalLinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExternalLinkState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isLogIntegration"] = state ? state.isLogIntegration : undefined;
            resourceInputs["metricFilterRegex"] = state ? state.metricFilterRegex : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pointTagFilterRegexes"] = state ? state.pointTagFilterRegexes : undefined;
            resourceInputs["sourceFilterRegex"] = state ? state.sourceFilterRegex : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
        } else {
            const args = argsOrState as ExternalLinkArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.template === undefined) && !opts.urn) {
                throw new Error("Missing required property 'template'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isLogIntegration"] = args ? args.isLogIntegration : undefined;
            resourceInputs["metricFilterRegex"] = args ? args.metricFilterRegex : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pointTagFilterRegexes"] = args ? args.pointTagFilterRegexes : undefined;
            resourceInputs["sourceFilterRegex"] = args ? args.sourceFilterRegex : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExternalLink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExternalLink resources.
 */
export interface ExternalLinkState {
    /**
     * Human-readable description for this link.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether this is a "Log Integration" subType of external link.
     */
    isLogIntegration?: pulumi.Input<boolean>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
     */
    metricFilterRegex?: pulumi.Input<string>;
    /**
     * The name of the external link.
     */
    name?: pulumi.Input<string>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted
     * series. This is a map from string to regular expression. The highlighted series must contain point tags whose
     * keys are present in the keys of this map and whose values match the regular expressions associated with those
     * keys in order for the link to be displayed.
     */
    pointTagFilterRegexes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
     */
    sourceFilterRegex?: pulumi.Input<string>;
    /**
     * The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
     */
    template?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExternalLink resource.
 */
export interface ExternalLinkArgs {
    /**
     * Human-readable description for this link.
     */
    description: pulumi.Input<string>;
    /**
     * Whether this is a "Log Integration" subType of external link.
     */
    isLogIntegration?: pulumi.Input<boolean>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
     */
    metricFilterRegex?: pulumi.Input<string>;
    /**
     * The name of the external link.
     */
    name?: pulumi.Input<string>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted
     * series. This is a map from string to regular expression. The highlighted series must contain point tags whose
     * keys are present in the keys of this map and whose values match the regular expressions associated with those
     * keys in order for the link to be displayed.
     */
    pointTagFilterRegexes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
     */
    sourceFilterRegex?: pulumi.Input<string>;
    /**
     * The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
     */
    template: pulumi.Input<string>;
}
