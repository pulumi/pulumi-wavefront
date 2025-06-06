// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a Wavefront Alert resource. This allows alerts to be created, updated, and deleted.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as wavefront from "@pulumi/wavefront";
 *
 * const foobar = new wavefront.Alert("foobar", {
 *     name: "Test Alert",
 *     target: "test@example.com,target:alert-target-id",
 *     condition: "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80",
 *     displayExpression: "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )",
 *     minutes: 5,
 *     resolveAfterMinutes: 5,
 *     severity: "WARN",
 *     tags: [
 *         "terraform",
 *         "test",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Alerts can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import wavefront:index/alert:Alert alert_target 1479868728473
 * ```
 */
export class Alert extends pulumi.CustomResource {
    /**
     * Get an existing Alert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertState, opts?: pulumi.CustomResourceOptions): Alert {
        return new Alert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'wavefront:index/alert:Alert';

    /**
     * Returns true if the given object is an instance of Alert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alert.__pulumiType;
    }

    /**
     * User-supplied additional explanatory information for this alert.
     * Useful for linking runbooks, migrations, etc.
     */
    public readonly additionalInformation!: pulumi.Output<string | undefined>;
    /**
     * A set of user-supplied dashboard and parameters to create dashboard links for triaging alerts.
     */
    public readonly alertTriageDashboards!: pulumi.Output<outputs.AlertAlertTriageDashboard[] | undefined>;
    /**
     * The type of alert in Wavefront. Either `CLASSIC` (default)
     * or `THRESHOLD`.
     */
    public readonly alertType!: pulumi.Output<string | undefined>;
    /**
     * A list of valid users or groups that can modify this resource on a tenant.
     */
    public readonly canModifies!: pulumi.Output<string[]>;
    /**
     * A list of valid users or groups that can view this resource on a tenant. Default is Empty list.
     */
    public readonly canViews!: pulumi.Output<string[] | undefined>;
    /**
     * A Wavefront query that is evaluated at regular intervals (default is 1 minute).
     * The alert fires and notifications are triggered when a data series matching this query evaluates
     * to a non-zero value for a set number of consecutive minutes.
     */
    public readonly condition!: pulumi.Output<string | undefined>;
    /**
     * a string->string map of `severity` to `condition`
     * for which this alert will trigger.
     */
    public readonly conditions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A second query whose results are displayed in the alert user
     * interface instead of the condition query. This field is often used to display a version
     * of the condition query with Boolean operators removed so that numerical values are plotted.
     */
    public readonly displayExpression!: pulumi.Output<string | undefined>;
    /**
     * The number of consecutive minutes that a series matching the condition query must
     * evaluate to "true" (non-zero value) before the alert fires.
     */
    public readonly minutes!: pulumi.Output<number>;
    /**
     * The name of the alert as it is displayed in Wavefront.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * How often to re-trigger a continually failing alert.
     * If absent or <= 0, no re-triggering occurs.
     */
    public readonly notificationResendFrequencyMinutes!: pulumi.Output<number | undefined>;
    /**
     * The specified query is executed every `processRateMinutes` minutes. Default value is 5 minutes.
     */
    public readonly processRateMinutes!: pulumi.Output<number | undefined>;
    /**
     * The number of consecutive minutes that a firing series matching the condition
     * query must evaluate to "false" (zero value) before the alert resolves. When unset, this defaults to
     * the same value as `minutes`.
     */
    public readonly resolveAfterMinutes!: pulumi.Output<number | undefined>;
    /**
     * A list of user-supplied runbook links for this alert.
     */
    public readonly runbookLinks!: pulumi.Output<string[] | undefined>;
    /**
     * Severity of the alert, valid values are `INFO`, `SMOKE`, `WARN`, `SEVERE`.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * A set of tags to assign to this resource.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * A comma-separated list of the email address or integration endpoint
     * (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.
     * Alert target format: ({email}|pd:{pd_key}|target:{alert-target-id}).
     */
    public readonly target!: pulumi.Output<string | undefined>;
    /**
     * A string to string map of Targets for severity.
     */
    public readonly thresholdTargets!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Alert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertArgs | AlertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertState | undefined;
            resourceInputs["additionalInformation"] = state ? state.additionalInformation : undefined;
            resourceInputs["alertTriageDashboards"] = state ? state.alertTriageDashboards : undefined;
            resourceInputs["alertType"] = state ? state.alertType : undefined;
            resourceInputs["canModifies"] = state ? state.canModifies : undefined;
            resourceInputs["canViews"] = state ? state.canViews : undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["displayExpression"] = state ? state.displayExpression : undefined;
            resourceInputs["minutes"] = state ? state.minutes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationResendFrequencyMinutes"] = state ? state.notificationResendFrequencyMinutes : undefined;
            resourceInputs["processRateMinutes"] = state ? state.processRateMinutes : undefined;
            resourceInputs["resolveAfterMinutes"] = state ? state.resolveAfterMinutes : undefined;
            resourceInputs["runbookLinks"] = state ? state.runbookLinks : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["thresholdTargets"] = state ? state.thresholdTargets : undefined;
        } else {
            const args = argsOrState as AlertArgs | undefined;
            if ((!args || args.minutes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'minutes'");
            }
            if ((!args || args.tags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tags'");
            }
            resourceInputs["additionalInformation"] = args ? args.additionalInformation : undefined;
            resourceInputs["alertTriageDashboards"] = args ? args.alertTriageDashboards : undefined;
            resourceInputs["alertType"] = args ? args.alertType : undefined;
            resourceInputs["canModifies"] = args ? args.canModifies : undefined;
            resourceInputs["canViews"] = args ? args.canViews : undefined;
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["displayExpression"] = args ? args.displayExpression : undefined;
            resourceInputs["minutes"] = args ? args.minutes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationResendFrequencyMinutes"] = args ? args.notificationResendFrequencyMinutes : undefined;
            resourceInputs["processRateMinutes"] = args ? args.processRateMinutes : undefined;
            resourceInputs["resolveAfterMinutes"] = args ? args.resolveAfterMinutes : undefined;
            resourceInputs["runbookLinks"] = args ? args.runbookLinks : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["thresholdTargets"] = args ? args.thresholdTargets : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alert resources.
 */
export interface AlertState {
    /**
     * User-supplied additional explanatory information for this alert.
     * Useful for linking runbooks, migrations, etc.
     */
    additionalInformation?: pulumi.Input<string>;
    /**
     * A set of user-supplied dashboard and parameters to create dashboard links for triaging alerts.
     */
    alertTriageDashboards?: pulumi.Input<pulumi.Input<inputs.AlertAlertTriageDashboard>[]>;
    /**
     * The type of alert in Wavefront. Either `CLASSIC` (default)
     * or `THRESHOLD`.
     */
    alertType?: pulumi.Input<string>;
    /**
     * A list of valid users or groups that can modify this resource on a tenant.
     */
    canModifies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of valid users or groups that can view this resource on a tenant. Default is Empty list.
     */
    canViews?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Wavefront query that is evaluated at regular intervals (default is 1 minute).
     * The alert fires and notifications are triggered when a data series matching this query evaluates
     * to a non-zero value for a set number of consecutive minutes.
     */
    condition?: pulumi.Input<string>;
    /**
     * a string->string map of `severity` to `condition`
     * for which this alert will trigger.
     */
    conditions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A second query whose results are displayed in the alert user
     * interface instead of the condition query. This field is often used to display a version
     * of the condition query with Boolean operators removed so that numerical values are plotted.
     */
    displayExpression?: pulumi.Input<string>;
    /**
     * The number of consecutive minutes that a series matching the condition query must
     * evaluate to "true" (non-zero value) before the alert fires.
     */
    minutes?: pulumi.Input<number>;
    /**
     * The name of the alert as it is displayed in Wavefront.
     */
    name?: pulumi.Input<string>;
    /**
     * How often to re-trigger a continually failing alert.
     * If absent or <= 0, no re-triggering occurs.
     */
    notificationResendFrequencyMinutes?: pulumi.Input<number>;
    /**
     * The specified query is executed every `processRateMinutes` minutes. Default value is 5 minutes.
     */
    processRateMinutes?: pulumi.Input<number>;
    /**
     * The number of consecutive minutes that a firing series matching the condition
     * query must evaluate to "false" (zero value) before the alert resolves. When unset, this defaults to
     * the same value as `minutes`.
     */
    resolveAfterMinutes?: pulumi.Input<number>;
    /**
     * A list of user-supplied runbook links for this alert.
     */
    runbookLinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Severity of the alert, valid values are `INFO`, `SMOKE`, `WARN`, `SEVERE`.
     */
    severity?: pulumi.Input<string>;
    /**
     * A set of tags to assign to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A comma-separated list of the email address or integration endpoint
     * (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.
     * Alert target format: ({email}|pd:{pd_key}|target:{alert-target-id}).
     */
    target?: pulumi.Input<string>;
    /**
     * A string to string map of Targets for severity.
     */
    thresholdTargets?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Alert resource.
 */
export interface AlertArgs {
    /**
     * User-supplied additional explanatory information for this alert.
     * Useful for linking runbooks, migrations, etc.
     */
    additionalInformation?: pulumi.Input<string>;
    /**
     * A set of user-supplied dashboard and parameters to create dashboard links for triaging alerts.
     */
    alertTriageDashboards?: pulumi.Input<pulumi.Input<inputs.AlertAlertTriageDashboard>[]>;
    /**
     * The type of alert in Wavefront. Either `CLASSIC` (default)
     * or `THRESHOLD`.
     */
    alertType?: pulumi.Input<string>;
    /**
     * A list of valid users or groups that can modify this resource on a tenant.
     */
    canModifies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of valid users or groups that can view this resource on a tenant. Default is Empty list.
     */
    canViews?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Wavefront query that is evaluated at regular intervals (default is 1 minute).
     * The alert fires and notifications are triggered when a data series matching this query evaluates
     * to a non-zero value for a set number of consecutive minutes.
     */
    condition?: pulumi.Input<string>;
    /**
     * a string->string map of `severity` to `condition`
     * for which this alert will trigger.
     */
    conditions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A second query whose results are displayed in the alert user
     * interface instead of the condition query. This field is often used to display a version
     * of the condition query with Boolean operators removed so that numerical values are plotted.
     */
    displayExpression?: pulumi.Input<string>;
    /**
     * The number of consecutive minutes that a series matching the condition query must
     * evaluate to "true" (non-zero value) before the alert fires.
     */
    minutes: pulumi.Input<number>;
    /**
     * The name of the alert as it is displayed in Wavefront.
     */
    name?: pulumi.Input<string>;
    /**
     * How often to re-trigger a continually failing alert.
     * If absent or <= 0, no re-triggering occurs.
     */
    notificationResendFrequencyMinutes?: pulumi.Input<number>;
    /**
     * The specified query is executed every `processRateMinutes` minutes. Default value is 5 minutes.
     */
    processRateMinutes?: pulumi.Input<number>;
    /**
     * The number of consecutive minutes that a firing series matching the condition
     * query must evaluate to "false" (zero value) before the alert resolves. When unset, this defaults to
     * the same value as `minutes`.
     */
    resolveAfterMinutes?: pulumi.Input<number>;
    /**
     * A list of user-supplied runbook links for this alert.
     */
    runbookLinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Severity of the alert, valid values are `INFO`, `SMOKE`, `WARN`, `SEVERE`.
     */
    severity?: pulumi.Input<string>;
    /**
     * A set of tags to assign to this resource.
     */
    tags: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A comma-separated list of the email address or integration endpoint
     * (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.
     * Alert target format: ({email}|pd:{pd_key}|target:{alert-target-id}).
     */
    target?: pulumi.Input<string>;
    /**
     * A string to string map of Targets for severity.
     */
    thresholdTargets?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
