// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a Wavefront Dashboard resource. This allows dashboards to be created, updated, and deleted.
 *
 * ## Import
 *
 * Dashboards can be imported by using the `id`, e.g.:
 *
 * ```sh
 * $ pulumi import wavefront:index/dashboard:Dashboard dashboard tftestimport
 * ```
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'wavefront:index/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    /**
     * A list of users/groups/roles that can modify the dashboard.
     */
    public readonly canModifies!: pulumi.Output<string[]>;
    /**
     * A list of users/groups/roles that can view the dashboard.
     */
    public readonly canViews!: pulumi.Output<string[] | undefined>;
    /**
     * Human-readable description of the dashboard.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether the dashboard parameters section is opened by default when the dashboard
     * is shown.
     */
    public readonly displayQueryParameters!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the "pills" quick-linked the sections of the dashboard are
     * displayed by default when the dashboard is shown.
     */
    public readonly displaySectionTableOfContents!: pulumi.Output<boolean | undefined>;
    /**
     * How charts belonging to this dashboard should display events. `BYCHART` is default if
     * unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.
     */
    public readonly eventFilterType!: pulumi.Output<string | undefined>;
    /**
     * Name of the dashboard.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current JSON representation of dashboard parameters. See parameter details.
     */
    public readonly parameterDetails!: pulumi.Output<outputs.DashboardParameterDetail[] | undefined>;
    /**
     * Dashboard chart sections. See dashboard sections.
     */
    public readonly sections!: pulumi.Output<outputs.DashboardSection[]>;
    /**
     * A set of tags to assign to this resource.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Unique identifier, also a URL slug of the dashboard.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardState | undefined;
            resourceInputs["canModifies"] = state ? state.canModifies : undefined;
            resourceInputs["canViews"] = state ? state.canViews : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayQueryParameters"] = state ? state.displayQueryParameters : undefined;
            resourceInputs["displaySectionTableOfContents"] = state ? state.displaySectionTableOfContents : undefined;
            resourceInputs["eventFilterType"] = state ? state.eventFilterType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameterDetails"] = state ? state.parameterDetails : undefined;
            resourceInputs["sections"] = state ? state.sections : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.sections === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sections'");
            }
            if ((!args || args.tags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tags'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["canModifies"] = args ? args.canModifies : undefined;
            resourceInputs["canViews"] = args ? args.canViews : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayQueryParameters"] = args ? args.displayQueryParameters : undefined;
            resourceInputs["displaySectionTableOfContents"] = args ? args.displaySectionTableOfContents : undefined;
            resourceInputs["eventFilterType"] = args ? args.eventFilterType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameterDetails"] = args ? args.parameterDetails : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Dashboard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    /**
     * A list of users/groups/roles that can modify the dashboard.
     */
    canModifies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of users/groups/roles that can view the dashboard.
     */
    canViews?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable description of the dashboard.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the dashboard parameters section is opened by default when the dashboard
     * is shown.
     */
    displayQueryParameters?: pulumi.Input<boolean>;
    /**
     * Whether the "pills" quick-linked the sections of the dashboard are
     * displayed by default when the dashboard is shown.
     */
    displaySectionTableOfContents?: pulumi.Input<boolean>;
    /**
     * How charts belonging to this dashboard should display events. `BYCHART` is default if
     * unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.
     */
    eventFilterType?: pulumi.Input<string>;
    /**
     * Name of the dashboard.
     */
    name?: pulumi.Input<string>;
    /**
     * The current JSON representation of dashboard parameters. See parameter details.
     */
    parameterDetails?: pulumi.Input<pulumi.Input<inputs.DashboardParameterDetail>[]>;
    /**
     * Dashboard chart sections. See dashboard sections.
     */
    sections?: pulumi.Input<pulumi.Input<inputs.DashboardSection>[]>;
    /**
     * A set of tags to assign to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unique identifier, also a URL slug of the dashboard.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * A list of users/groups/roles that can modify the dashboard.
     */
    canModifies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of users/groups/roles that can view the dashboard.
     */
    canViews?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable description of the dashboard.
     */
    description: pulumi.Input<string>;
    /**
     * Whether the dashboard parameters section is opened by default when the dashboard
     * is shown.
     */
    displayQueryParameters?: pulumi.Input<boolean>;
    /**
     * Whether the "pills" quick-linked the sections of the dashboard are
     * displayed by default when the dashboard is shown.
     */
    displaySectionTableOfContents?: pulumi.Input<boolean>;
    /**
     * How charts belonging to this dashboard should display events. `BYCHART` is default if
     * unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.
     */
    eventFilterType?: pulumi.Input<string>;
    /**
     * Name of the dashboard.
     */
    name?: pulumi.Input<string>;
    /**
     * The current JSON representation of dashboard parameters. See parameter details.
     */
    parameterDetails?: pulumi.Input<pulumi.Input<inputs.DashboardParameterDetail>[]>;
    /**
     * Dashboard chart sections. See dashboard sections.
     */
    sections: pulumi.Input<pulumi.Input<inputs.DashboardSection>[]>;
    /**
     * A set of tags to assign to this resource.
     */
    tags: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unique identifier, also a URL slug of the dashboard.
     */
    url: pulumi.Input<string>;
}
