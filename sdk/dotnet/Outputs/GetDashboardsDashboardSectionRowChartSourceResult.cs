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
    public sealed class GetDashboardsDashboardSectionRowChartSourceResult
    {
        /// <summary>
        /// Whether the source is disabled.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// The name of the parameters.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Query expression to plot on the chart.
        /// </summary>
        public readonly string Query;
        public readonly bool QuerybuilderEnabled;
        /// <summary>
        /// For scatter plots, does this query source the X-axis or the Y-axis, `X`, or `Y`.
        /// </summary>
        public readonly string ScatterPlotSource;
        public readonly bool SecondaryAxis;
        public readonly string SourceColor;
        /// <summary>
        /// A description for the purpose of this source.
        /// </summary>
        public readonly string SourceDescription;

        [OutputConstructor]
        private GetDashboardsDashboardSectionRowChartSourceResult(
            bool disabled,

            string name,

            string query,

            bool querybuilderEnabled,

            string scatterPlotSource,

            bool secondaryAxis,

            string sourceColor,

            string sourceDescription)
        {
            Disabled = disabled;
            Name = name;
            Query = query;
            QuerybuilderEnabled = querybuilderEnabled;
            ScatterPlotSource = scatterPlotSource;
            SecondaryAxis = secondaryAxis;
            SourceColor = sourceColor;
            SourceDescription = sourceDescription;
        }
    }
}
