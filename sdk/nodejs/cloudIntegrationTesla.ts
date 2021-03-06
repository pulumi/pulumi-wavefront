// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Wavefront Cloud Integration for Tesla. This allows NewRelic cloud integrations to be created,
 * updated, and deleted.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * const tesla = new wavefront.CloudIntegrationTesla("tesla", {
 *     email: "email@example.com",
 *     password: "password",
 * });
 * ```
 *
 * ## Import
 *
 * Tesla Integrations can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import wavefront:index/cloudIntegrationTesla:CloudIntegrationTesla tesla a411c16b-3cf7-4f03-bf11-8ca05aab898d
 * ```
 */
export class CloudIntegrationTesla extends pulumi.CustomResource {
    /**
     * Get an existing CloudIntegrationTesla resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudIntegrationTeslaState, opts?: pulumi.CustomResourceOptions): CloudIntegrationTesla {
        return new CloudIntegrationTesla(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'wavefront:index/cloudIntegrationTesla:CloudIntegrationTesla';

    /**
     * Returns true if the given object is an instance of CloudIntegrationTesla.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudIntegrationTesla {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudIntegrationTesla.__pulumiType;
    }

    /**
     * A list of point tag key-values to add to every point ingested using this integration
     */
    public readonly additionalTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Email address for the Tesla account login
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Forces this resource to save, even if errors are present
     */
    public readonly forceSave!: pulumi.Output<boolean | undefined>;
    /**
     * The human-readable name of this integration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for the Tesla account login
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * A value denoting which cloud service this service integrates with
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * How often, in minutes, to refresh the service
     */
    public readonly serviceRefreshRateInMinutes!: pulumi.Output<number | undefined>;

    /**
     * Create a CloudIntegrationTesla resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudIntegrationTeslaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudIntegrationTeslaArgs | CloudIntegrationTeslaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudIntegrationTeslaState | undefined;
            inputs["additionalTags"] = state ? state.additionalTags : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["forceSave"] = state ? state.forceSave : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["serviceRefreshRateInMinutes"] = state ? state.serviceRefreshRateInMinutes : undefined;
        } else {
            const args = argsOrState as CloudIntegrationTeslaArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            inputs["additionalTags"] = args ? args.additionalTags : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["forceSave"] = args ? args.forceSave : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["serviceRefreshRateInMinutes"] = args ? args.serviceRefreshRateInMinutes : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CloudIntegrationTesla.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudIntegrationTesla resources.
 */
export interface CloudIntegrationTeslaState {
    /**
     * A list of point tag key-values to add to every point ingested using this integration
     */
    readonly additionalTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Email address for the Tesla account login
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Forces this resource to save, even if errors are present
     */
    readonly forceSave?: pulumi.Input<boolean>;
    /**
     * The human-readable name of this integration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Password for the Tesla account login
     */
    readonly password?: pulumi.Input<string>;
    /**
     * A value denoting which cloud service this service integrates with
     */
    readonly service?: pulumi.Input<string>;
    /**
     * How often, in minutes, to refresh the service
     */
    readonly serviceRefreshRateInMinutes?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CloudIntegrationTesla resource.
 */
export interface CloudIntegrationTeslaArgs {
    /**
     * A list of point tag key-values to add to every point ingested using this integration
     */
    readonly additionalTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Email address for the Tesla account login
     */
    readonly email: pulumi.Input<string>;
    /**
     * Forces this resource to save, even if errors are present
     */
    readonly forceSave?: pulumi.Input<boolean>;
    /**
     * The human-readable name of this integration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Password for the Tesla account login
     */
    readonly password: pulumi.Input<string>;
    /**
     * A value denoting which cloud service this service integrates with
     */
    readonly service: pulumi.Input<string>;
    /**
     * How often, in minutes, to refresh the service
     */
    readonly serviceRefreshRateInMinutes?: pulumi.Input<number>;
}
