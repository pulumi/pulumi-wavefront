// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AlertAlertTriageDashboard {
    /**
     * Dashboard ID
     */
    dashboardId: pulumi.Input<string>;
    /**
     * Dashboard Description
     */
    description: pulumi.Input<string>;
    parameters?: pulumi.Input<inputs.AlertAlertTriageDashboardParameters>;
}

export interface AlertAlertTriageDashboardParameters {
    constants?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface AlertTargetRoute {
    /**
     * (Required) String that filters the route. Space delimited. Currently only allows a single key value pair.
     * (e.g. `env prod`)
     */
    filter?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The notification method used for notification target. One of `WEBHOOK`, `EMAIL`, `PAGERDUTY`.
     */
    method: pulumi.Input<string>;
    /**
     * (Required) The endpoint for the alert route. `EMAIL`: email address. `PAGERDUTY`: PagerDuty routing
     * key. `WEBHOOK`: URL endpoint.
     */
    target: pulumi.Input<string>;
}

export interface CloudIntegrationNewRelicMetricFilter {
    /**
     * The name of a NewRelic App.
     */
    appName: pulumi.Input<string>;
    /**
     * A regular expression that a metric name must match (case-insensitively) in order to be ingested.
     */
    metricFilterRegex: pulumi.Input<string>;
}

export interface DashboardParameterDetail {
    /**
     * The default value of the parameter.
     */
    defaultValue: pulumi.Input<string>;
    /**
     * For `DYNAMIC` parameter types, the type of the field. Valid options are `SOURCE`,
     * `SOURCE_TAG`, `METRIC_NAME`, `TAG_KEY`, and `MATCHING_SOURCE_TAG`.
     */
    dynamicFieldType?: pulumi.Input<string>;
    /**
     * If `true` the parameter will only be shown on the edit view of the dashboard.
     */
    hideFromView: pulumi.Input<boolean>;
    /**
     * The label for the parameter.
     */
    label: pulumi.Input<string>;
    /**
     * The name of the parameters.
     */
    name: pulumi.Input<string>;
    /**
     * The type of the parameter. `SIMPLE`, `LIST`, or `DYNAMIC`.
     */
    parameterType: pulumi.Input<string>;
    /**
     * For `DYNAMIC` parameter types, the query to execute to return values.
     */
    queryValue?: pulumi.Input<string>;
    /**
     * for `TAG_KEY` dynamic field types, the tag key to return.
     */
    tagKey?: pulumi.Input<string>;
    /**
     * A string->string map. At least one of the keys must match the value of
     * `defaultValue`.
     */
    valuesToReadableStrings: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface DashboardSection {
    /**
     * Name of this section.
     */
    name: pulumi.Input<string>;
    /**
     * See dashboard section rows.
     */
    rows: pulumi.Input<pulumi.Input<inputs.DashboardSectionRow>[]>;
}

export interface DashboardSectionRow {
    /**
     * Charts in this section. See dashboard chart.
     */
    charts: pulumi.Input<pulumi.Input<inputs.DashboardSectionRowChart>[]>;
}

export interface DashboardSectionRowChart {
    /**
     * The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
     */
    base?: pulumi.Input<number>;
    chartAttribute?: pulumi.Input<string>;
    /**
     * Chart settings. See chart settings.
     */
    chartSetting: pulumi.Input<inputs.DashboardSectionRowChartChartSetting>;
    /**
     * Description of the chart.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the source.
     */
    name: pulumi.Input<string>;
    /**
     * Show events related to the sources included in queries
     */
    noDefaultEvents?: pulumi.Input<boolean>;
    /**
     * Query expression to plot on the chart. See chart source queries.
     */
    sources: pulumi.Input<pulumi.Input<inputs.DashboardSectionRowChartSource>[]>;
    /**
     * Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
     * `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
     */
    summarization: pulumi.Input<string>;
    /**
     * String to label the units of the chart on the Y-Axis.
     */
    units: pulumi.Input<string>;
}

export interface DashboardSectionRowChartChartSetting {
    /**
     * This setting is deprecated.
     */
    autoColumnTags?: pulumi.Input<boolean>;
    /**
     * This setting is deprecated.
     */
    columnTags?: pulumi.Input<string>;
    /**
     * For the tabular view, a list of point tags to display when using the `custom` tag display mode.
     */
    customTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Threshold (in seconds) for time delta between consecutive points in a series
     * above which a dotted line will replace a solid in in line plots. Default is 60.
     */
    expectedDataSpacing?: pulumi.Input<number>;
    /**
     * For a chart with a fixed legend, a list of statistics to display in the legend.
     */
    fixedLegendDisplayStats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable a fixed tabular legend adjacent to the chart.
     */
    fixedLegendEnabled?: pulumi.Input<boolean>;
    /**
     * Statistic to use for determining whether a series is displayed on the fixed legend.
     * Valid options are `CURRENT`, `MEAN`, `MEDIAN`, `SUM`, `MIN`, `MAX`, and `COUNT`.
     */
    fixedLegendFilterField?: pulumi.Input<string>;
    /**
     * Number of series to include in the fixed legend.
     */
    fixedLegendFilterLimit?: pulumi.Input<number>;
    /**
     * Whether to display `TOP` or `BOTTOM` ranked series in a fixed legend. Valid options
     * are `TOP`, and `BOTTOM`.
     */
    fixedLegendFilterSort?: pulumi.Input<string>;
    /**
     * This setting is deprecated.
     */
    fixedLegendHideLabel?: pulumi.Input<boolean>;
    /**
     * Where the fixed legend should be displayed with respect to the chart.
     * Valid options are `RIGHT`, `TOP`, `LEFT`, `BOTTOM`.
     */
    fixedLegendPosition?: pulumi.Input<string>;
    /**
     * If `true`, the legend uses non-summarized stats instead of summarized.
     */
    fixedLegendUseRawStats?: pulumi.Input<boolean>;
    /**
     * For the tabular view, whether to group multi metrics into a single row by a common source.
     * If `false`, each source is displayed in its own row. if `true`, multiple metrics for the same host are displayed as different
     * columns in the same row.
     */
    groupBySource?: pulumi.Input<boolean>;
    /**
     * Whether to disable the display of the floating legend (but
     * reenable it when the ctrl-key is pressed).
     */
    invertDynamicLegendHoverControl?: pulumi.Input<boolean>;
    /**
     * Plot interpolation type.  `linear` is default. Valid options are `linear`, `step-before`,
     * `step-after`, `basis`, `cardinal`, and `monotone`.
     */
    lineType?: pulumi.Input<string>;
    /**
     * Max value of the Y-axis. Set to null or leave blank for auto.
     */
    max?: pulumi.Input<number>;
    /**
     * Min value of the Y-axis. Set to null or leave blank for auto.
     */
    min?: pulumi.Input<number>;
    /**
     * For the tabular view defines how many point tags to display.
     */
    numTags?: pulumi.Input<number>;
    /**
     * The markdown content for a Markdown display, in plain text.
     */
    plainMarkdownContent?: pulumi.Input<string>;
    /**
     * For the tabular view, whether to display sources. Default is `true`.
     */
    showHosts?: pulumi.Input<boolean>;
    /**
     * For the tabular view, whether to display labels. Default is `true`.
     */
    showLabels?: pulumi.Input<boolean>;
    /**
     * For the tabular view, whether to display raw values. Default is `false`.
     */
    showRawValues?: pulumi.Input<boolean>;
    /**
     * For the tabular view, whether to display values in descending order. Default is `false`.
     */
    sortValuesDescending?: pulumi.Input<boolean>;
    /**
     * For the single stat view, the decimal precision of the displayed number.
     */
    sparklineDecimalPrecision?: pulumi.Input<number>;
    /**
     * For the single stat view, the color of the displayed text (when not dynamically determined).
     * Values should be in `rgba(,,,,)` format.
     */
    sparklineDisplayColor?: pulumi.Input<string>;
    /**
     * For the single stat view, the font size of the displayed text, in percent.
     */
    sparklineDisplayFontSize?: pulumi.Input<string>;
    /**
     * For the single stat view, the horizontal position of the displayed text.
     * Valid options are `MIDDLE`, `LEFT`, `RIGHT`.
     */
    sparklineDisplayHorizontalPosition?: pulumi.Input<string>;
    /**
     * For the single stat view, a string to append to the displayed text.
     */
    sparklineDisplayPostfix?: pulumi.Input<string>;
    /**
     * For the single stat view, a string to add before the displayed text.
     */
    sparklineDisplayPrefix?: pulumi.Input<string>;
    /**
     * For the single stat view, where to display the name of the query or the value of the query.
     * Valid options are `VALUE` or `LABEL`.
     */
    sparklineDisplayValueType?: pulumi.Input<string>;
    /**
     * This setting is deprecated.
     */
    sparklineDisplayVerticalPosition?: pulumi.Input<string>;
    /**
     * For the single stat view, the color of the background fill. Values should be
     * in `rgba(,,,,)`.
     */
    sparklineFillColor?: pulumi.Input<string>;
    /**
     * For the single stat view, the color of the line. Values should be in `rgba(,,,,)` format.
     */
    sparklineLineColor?: pulumi.Input<string>;
    /**
     * For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart.
     * Valid options are `BACKGROUND`, `BOTTOM`, `NONE`.
     */
    sparklineSize?: pulumi.Input<string>;
    /**
     * For the single stat view, whether to apply dynamic color settings to
     * the displayed `TEXT` or `BACKGROUND`. Valid options are `TEXT` or `BACKGROUND`.
     */
    sparklineValueColorMapApplyTo?: pulumi.Input<string>;
    /**
     * For the single stat view, A list of colors that differing query values map to.
     * Must contain one more element than `sparklineValueColorMapValuesV2`. Values should be in `rgba(,,,,)`.
     */
    sparklineValueColorMapColors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This setting is deprecated.
     */
    sparklineValueColorMapValues?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * For the single stat view, a list of boundaries for mapping different
     * query values to colors. Must contain one element less than `sparklineValueColorMapColors`.
     */
    sparklineValueColorMapValuesV2s?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * For the single stat view, a list of display text values that different query
     * values map to. Must contain one more element than `sparklineValueTextMapThresholds`.
     */
    sparklineValueTextMapTexts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * For the single stat view, a list of threshold boundaries for
     * mapping different query values to display text. Must contain one element less than `sparklineValueTextMapText`.
     */
    sparklineValueTextMapThresholds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Type of stacked chart (applicable only if chart type is `stacked`). `zero` (default) means
     * stacked from y=0. `expand` means normalized from 0 to 1.  `wiggle` means minimize weighted changes. `silhouette` means to
     * center the stream. Valid options are `zero`, `expand`, `wiggle`, `silhouette`, and `bars`.
     */
    stackType?: pulumi.Input<string>;
    /**
     * For the tabular view, which mode to use to determine which point tags to display.
     * Valid options are `all`, `top`, or `custom`.
     */
    tagMode?: pulumi.Input<string>;
    /**
     * For x-y scatterplots, whether to color more recent points as darker than older points.
     */
    timeBasedColoring?: pulumi.Input<boolean>;
    /**
     * Chart Type. `line` refers to the Line Plot, `scatter` to the Point Plot, `stacked-area` to
     * the Stacked Area plot, `table` to the Tabular View, `scatterplot-xy` to Scatter Plot, `markdown-widget` to the
     * Markdown display, and `sparkline` to the Single Stat view. Valid options are`line`, `scatterplot`,
     * `stacked-area`, `stacked-column`, `table`, `scatterplot-xy`, `markdown-widget`, `sparkline`, `globe`, `nodemap`,
     * `top-k`, `status-list`, and `histogram`.
     */
    type: pulumi.Input<string>;
    /**
     * Width, in minutes, of the time window to use for `last` windowing.
     */
    windowSize?: pulumi.Input<number>;
    /**
     * For the tabular view, whether to use the full time window for the query or the last X minutes.
     * Valid options are `full` or `last`.
     */
    windowing?: pulumi.Input<string>;
    /**
     * For x-y scatterplots, max value for the X-axis. Set to null for auto.
     */
    xmax?: pulumi.Input<number>;
    /**
     * For x-y scatterplots, min value for the X-axis. Set to null for auto.
     */
    xmin?: pulumi.Input<number>;
    /**
     * Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).
     */
    y0ScaleSiBy1024?: pulumi.Input<boolean>;
    /**
     * Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.
     */
    y0UnitAutoscaling?: pulumi.Input<boolean>;
    /**
     * Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).
     */
    y1ScaleSiBy1024?: pulumi.Input<boolean>;
    /**
     * Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.
     */
    y1UnitAutoscaling?: pulumi.Input<boolean>;
    /**
     * For plots with multiple Y-axes, units for right side Y-axis.
     */
    y1Units?: pulumi.Input<string>;
    /**
     * For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.
     */
    y1max?: pulumi.Input<number>;
    /**
     * For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.
     */
    y1min?: pulumi.Input<number>;
    /**
     * For x-y scatterplots, max value for the Y-axis. Set to null for auto.
     */
    ymax?: pulumi.Input<number>;
    /**
     * For x-y scatterplots, min value for the Y-axis. Set to null for auto.
     */
    ymin?: pulumi.Input<number>;
}

export interface DashboardSectionRowChartSource {
    /**
     * Whether the source is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Name of the source.
     */
    name: pulumi.Input<string>;
    /**
     * Query expression to plot on the chart.
     */
    query: pulumi.Input<string>;
    /**
     * Whether or not this source line should have the query builder enabled.
     */
    queryBuilderEnabled?: pulumi.Input<boolean>;
    /**
     * For scatter plots, does this query source the X-axis or the Y-axis, `X`, or `Y`.
     */
    scatterPlotSource?: pulumi.Input<string>;
    /**
     * A description for the purpose of this source.
     */
    sourceDescription?: pulumi.Input<string>;
}

export interface IngestionPolicyTag {
    key: pulumi.Input<string>;
    value: pulumi.Input<string>;
}

export interface MetricsPolicyPolicyRule {
    /**
     * Valid options are `ALLOW` and `BLOCK`.
     */
    accessType: pulumi.Input<string>;
    /**
     * List of account ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
     */
    accountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A detailed description of the Metrics Policy. The description is visible only when you edit the rule.
     */
    description: pulumi.Input<string>;
    /**
     * The unique name identifier for a Metrics Policy. The name is visible on the Metrics Security Policy page.
     */
    name: pulumi.Input<string>;
    /**
     * List of prefixes to match metrics on. You can specify the full metric name or use a wildcard character in metric names, sources, or point tags. The wildcard character alone (*) means all metrics.
     */
    prefixes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of role ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
     */
    roleIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Key/Value tags to select target metrics for policy.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.MetricsPolicyPolicyRuleTag>[]>;
    /**
     * Bool where `true` require all tags are met by selected metrics, else `false` select metrics that match any give tag.
     */
    tagsAnded: pulumi.Input<boolean>;
    /**
     * List of user group ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
     */
    userGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface MetricsPolicyPolicyRuleTag {
    /**
     * The tag's key.
     */
    key: pulumi.Input<string>;
    /**
     * The tag's value.
     */
    value: pulumi.Input<string>;
}
