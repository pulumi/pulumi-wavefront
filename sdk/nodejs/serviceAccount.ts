// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Wavefront Service Account Resource. This allows service accounts to be created, updated, and deleted.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * const basic = new wavefront.ServiceAccount("basic", {
 *     active: true,
 *     identifier: "sa::tftesting",
 * });
 * ```
 */
export class ServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceAccountState, opts?: pulumi.CustomResourceOptions): ServiceAccount {
        return new ServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'wavefront:index/serviceAccount:ServiceAccount';

    /**
     * Returns true if the given object is an instance of ServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAccount.__pulumiType;
    }

    /**
     * Whether or not the service account is active
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the service account
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The (unique) identifier of the service account to create. Must start with sa::
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * List of permission to grant to this service account.  Valid options are
     * `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
     * `hostTagManagement`, `metricsManagement`, `userManagement`
     */
    public readonly permissions!: pulumi.Output<string[]>;
    /**
     * List of user groups for this service account
     */
    public readonly userGroups!: pulumi.Output<string[]>;

    /**
     * Create a ServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceAccountArgs | ServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceAccountState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["identifier"] = state ? state.identifier : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["userGroups"] = state ? state.userGroups : undefined;
        } else {
            const args = argsOrState as ServiceAccountArgs | undefined;
            if (!args || args.identifier === undefined) {
                throw new Error("Missing required property 'identifier'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["userGroups"] = args ? args.userGroups : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceAccount resources.
 */
export interface ServiceAccountState {
    /**
     * Whether or not the service account is active
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * The description of the service account
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The (unique) identifier of the service account to create. Must start with sa::
     */
    readonly identifier?: pulumi.Input<string>;
    /**
     * List of permission to grant to this service account.  Valid options are
     * `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
     * `hostTagManagement`, `metricsManagement`, `userManagement`
     */
    readonly permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of user groups for this service account
     */
    readonly userGroups?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ServiceAccount resource.
 */
export interface ServiceAccountArgs {
    /**
     * Whether or not the service account is active
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * The description of the service account
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The (unique) identifier of the service account to create. Must start with sa::
     */
    readonly identifier: pulumi.Input<string>;
    /**
     * List of permission to grant to this service account.  Valid options are
     * `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
     * `hostTagManagement`, `metricsManagement`, `userManagement`
     */
    readonly permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of user groups for this service account
     */
    readonly userGroups?: pulumi.Input<pulumi.Input<string>[]>;
}