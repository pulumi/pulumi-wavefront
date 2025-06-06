// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront
{
    /// <summary>
    /// Provides a Wavefront External Link Resource. This allows external links to be created, updated, and deleted.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Wavefront = Pulumi.Wavefront;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basic = new Wavefront.ExternalLink("basic", new()
    ///     {
    ///         Name = "External Link",
    ///         Description = "An external link description",
    ///         Template = "https://example.com/source={{{source}}}&amp;startTime={{startEpochMillis}}",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Maintenance windows can be imported by using the `id`, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import wavefront:index/externalLink:ExternalLink basic fVj6fz6zYC4aBkID
    /// ```
    /// </summary>
    [WavefrontResourceType("wavefront:index/externalLink:ExternalLink")]
    public partial class ExternalLink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Human-readable description for this link.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Whether this is a "Log Integration" subType of external link.
        /// </summary>
        [Output("isLogIntegration")]
        public Output<bool?> IsLogIntegration { get; private set; } = null!;

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
        /// </summary>
        [Output("metricFilterRegex")]
        public Output<string?> MetricFilterRegex { get; private set; } = null!;

        /// <summary>
        /// The name of the external link.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted
        /// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
        /// keys are present in the keys of this map and whose values match the regular expressions associated with those
        /// keys in order for the link to be displayed.
        /// </summary>
        [Output("pointTagFilterRegexes")]
        public Output<ImmutableDictionary<string, string>?> PointTagFilterRegexes { get; private set; } = null!;

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
        /// </summary>
        [Output("sourceFilterRegex")]
        public Output<string?> SourceFilterRegex { get; private set; } = null!;

        /// <summary>
        /// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalLink(string name, ExternalLinkArgs args, CustomResourceOptions? options = null)
            : base("wavefront:index/externalLink:ExternalLink", name, args ?? new ExternalLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalLink(string name, Input<string> id, ExternalLinkState? state = null, CustomResourceOptions? options = null)
            : base("wavefront:index/externalLink:ExternalLink", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExternalLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalLink Get(string name, Input<string> id, ExternalLinkState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalLink(name, id, state, options);
        }
    }

    public sealed class ExternalLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-readable description for this link.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Whether this is a "Log Integration" subType of external link.
        /// </summary>
        [Input("isLogIntegration")]
        public Input<bool>? IsLogIntegration { get; set; }

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
        /// </summary>
        [Input("metricFilterRegex")]
        public Input<string>? MetricFilterRegex { get; set; }

        /// <summary>
        /// The name of the external link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pointTagFilterRegexes")]
        private InputMap<string>? _pointTagFilterRegexes;

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted
        /// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
        /// keys are present in the keys of this map and whose values match the regular expressions associated with those
        /// keys in order for the link to be displayed.
        /// </summary>
        public InputMap<string> PointTagFilterRegexes
        {
            get => _pointTagFilterRegexes ?? (_pointTagFilterRegexes = new InputMap<string>());
            set => _pointTagFilterRegexes = value;
        }

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
        /// </summary>
        [Input("sourceFilterRegex")]
        public Input<string>? SourceFilterRegex { get; set; }

        /// <summary>
        /// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
        /// </summary>
        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        public ExternalLinkArgs()
        {
        }
        public static new ExternalLinkArgs Empty => new ExternalLinkArgs();
    }

    public sealed class ExternalLinkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-readable description for this link.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this is a "Log Integration" subType of external link.
        /// </summary>
        [Input("isLogIntegration")]
        public Input<bool>? IsLogIntegration { get; set; }

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
        /// </summary>
        [Input("metricFilterRegex")]
        public Input<string>? MetricFilterRegex { get; set; }

        /// <summary>
        /// The name of the external link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pointTagFilterRegexes")]
        private InputMap<string>? _pointTagFilterRegexes;

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted
        /// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
        /// keys are present in the keys of this map and whose values match the regular expressions associated with those
        /// keys in order for the link to be displayed.
        /// </summary>
        public InputMap<string> PointTagFilterRegexes
        {
            get => _pointTagFilterRegexes ?? (_pointTagFilterRegexes = new InputMap<string>());
            set => _pointTagFilterRegexes = value;
        }

        /// <summary>
        /// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
        /// </summary>
        [Input("sourceFilterRegex")]
        public Input<string>? SourceFilterRegex { get; set; }

        /// <summary>
        /// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public ExternalLinkState()
        {
        }
        public static new ExternalLinkState Empty => new ExternalLinkState();
    }
}
