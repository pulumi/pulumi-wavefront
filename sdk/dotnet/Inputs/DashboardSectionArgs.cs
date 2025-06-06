// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront.Inputs
{

    public sealed class DashboardSectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of this section.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("rows", required: true)]
        private InputList<Inputs.DashboardSectionRowArgs>? _rows;

        /// <summary>
        /// See dashboard section rows.
        /// </summary>
        public InputList<Inputs.DashboardSectionRowArgs> Rows
        {
            get => _rows ?? (_rows = new InputList<Inputs.DashboardSectionRowArgs>());
            set => _rows = value;
        }

        public DashboardSectionArgs()
        {
        }
        public static new DashboardSectionArgs Empty => new DashboardSectionArgs();
    }
}
