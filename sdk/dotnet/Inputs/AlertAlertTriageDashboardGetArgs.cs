// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront.Inputs
{

    public sealed class AlertAlertTriageDashboardGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dashboard ID
        /// </summary>
        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        /// <summary>
        /// Dashboard Description
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("parameters")]
        public Input<Inputs.AlertAlertTriageDashboardParametersGetArgs>? Parameters { get; set; }

        public AlertAlertTriageDashboardGetArgs()
        {
        }
        public static new AlertAlertTriageDashboardGetArgs Empty => new AlertAlertTriageDashboardGetArgs();
    }
}
