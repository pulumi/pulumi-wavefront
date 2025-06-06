// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront.Outputs
{

    [OutputType]
    public sealed class DashboardSectionRowChart
    {
        /// <summary>
        /// The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
        /// </summary>
        public readonly int? Base;
        public readonly string? ChartAttribute;
        /// <summary>
        /// Chart settings. See chart settings.
        /// </summary>
        public readonly Outputs.DashboardSectionRowChartChartSetting ChartSetting;
        /// <summary>
        /// Description of the chart.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Name of the source.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Show events related to the sources included in queries
        /// </summary>
        public readonly bool? NoDefaultEvents;
        /// <summary>
        /// Query expression to plot on the chart. See chart source queries.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardSectionRowChartSource> Sources;
        /// <summary>
        /// Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
        /// `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
        /// </summary>
        public readonly string Summarization;
        /// <summary>
        /// String to label the units of the chart on the Y-Axis.
        /// </summary>
        public readonly string Units;

        [OutputConstructor]
        private DashboardSectionRowChart(
            int? @base,

            string? chartAttribute,

            Outputs.DashboardSectionRowChartChartSetting chartSetting,

            string? description,

            string name,

            bool? noDefaultEvents,

            ImmutableArray<Outputs.DashboardSectionRowChartSource> sources,

            string summarization,

            string units)
        {
            Base = @base;
            ChartAttribute = chartAttribute;
            ChartSetting = chartSetting;
            Description = description;
            Name = name;
            NoDefaultEvents = noDefaultEvents;
            Sources = sources;
            Summarization = summarization;
            Units = units;
        }
    }
}
