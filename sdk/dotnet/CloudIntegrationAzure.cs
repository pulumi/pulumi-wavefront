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
    /// Provides a Wavefront Cloud Integration for Microsoft Azure. This allows Azure cloud integrations to be created,
    /// updated, and deleted.
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
    ///     var azureActivityLog = new Wavefront.CloudIntegrationAzureActivityLog("azure_activity_log", new()
    ///     {
    ///         Name = "Test Integration",
    ///         ClientId = "client-id2",
    ///         ClientSecret = "client-secret2",
    ///         Tenant = "my-tenant2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure Cloud Integrations can be imported by using the `id`, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure azure a411c16b-3cf7-4f03-bf11-8ca05aab898d
    /// ```
    /// </summary>
    [WavefrontResourceType("wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure")]
    public partial class CloudIntegrationAzure : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of point tag key-values to add to every point ingested using this integration.
        /// </summary>
        [Output("additionalTags")]
        public Output<ImmutableDictionary<string, string>?> AdditionalTags { get; private set; } = null!;

        /// <summary>
        /// A list of Azure Activity Log categories.
        /// </summary>
        [Output("categoryFilters")]
        public Output<ImmutableArray<string>> CategoryFilters { get; private set; } = null!;

        /// <summary>
        /// Client ID for an Azure service account within your project.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client secret for an Azure service account within your project.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Forces this resource to save, even if errors are present.
        /// </summary>
        [Output("forceSave")]
        public Output<bool?> ForceSave { get; private set; } = null!;

        /// <summary>
        /// A regular expression that a metric name must match (case-insensitively) in order to be ingested.
        /// </summary>
        [Output("metricFilterRegex")]
        public Output<string?> MetricFilterRegex { get; private set; } = null!;

        /// <summary>
        /// The human-readable name of this integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of Azure resource groups from which to pull metrics.
        /// </summary>
        [Output("resourceGroupFilters")]
        public Output<ImmutableArray<string>> ResourceGroupFilters { get; private set; } = null!;

        /// <summary>
        /// A value denoting which cloud service this service integrates with.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// How often, in minutes, to refresh the service.
        /// </summary>
        [Output("serviceRefreshRateInMinutes")]
        public Output<int?> ServiceRefreshRateInMinutes { get; private set; } = null!;

        /// <summary>
        /// Tenant ID for an Azure service account within your project.
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;


        /// <summary>
        /// Create a CloudIntegrationAzure resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudIntegrationAzure(string name, CloudIntegrationAzureArgs args, CustomResourceOptions? options = null)
            : base("wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure", name, args ?? new CloudIntegrationAzureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudIntegrationAzure(string name, Input<string> id, CloudIntegrationAzureState? state = null, CustomResourceOptions? options = null)
            : base("wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudIntegrationAzure resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudIntegrationAzure Get(string name, Input<string> id, CloudIntegrationAzureState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudIntegrationAzure(name, id, state, options);
        }
    }

    public sealed class CloudIntegrationAzureArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalTags")]
        private InputMap<string>? _additionalTags;

        /// <summary>
        /// A list of point tag key-values to add to every point ingested using this integration.
        /// </summary>
        public InputMap<string> AdditionalTags
        {
            get => _additionalTags ?? (_additionalTags = new InputMap<string>());
            set => _additionalTags = value;
        }

        [Input("categoryFilters")]
        private InputList<string>? _categoryFilters;

        /// <summary>
        /// A list of Azure Activity Log categories.
        /// </summary>
        public InputList<string> CategoryFilters
        {
            get => _categoryFilters ?? (_categoryFilters = new InputList<string>());
            set => _categoryFilters = value;
        }

        /// <summary>
        /// Client ID for an Azure service account within your project.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        private Input<string>? _clientSecret;

        /// <summary>
        /// Client secret for an Azure service account within your project.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Forces this resource to save, even if errors are present.
        /// </summary>
        [Input("forceSave")]
        public Input<bool>? ForceSave { get; set; }

        /// <summary>
        /// A regular expression that a metric name must match (case-insensitively) in order to be ingested.
        /// </summary>
        [Input("metricFilterRegex")]
        public Input<string>? MetricFilterRegex { get; set; }

        /// <summary>
        /// The human-readable name of this integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resourceGroupFilters")]
        private InputList<string>? _resourceGroupFilters;

        /// <summary>
        /// A list of Azure resource groups from which to pull metrics.
        /// </summary>
        public InputList<string> ResourceGroupFilters
        {
            get => _resourceGroupFilters ?? (_resourceGroupFilters = new InputList<string>());
            set => _resourceGroupFilters = value;
        }

        /// <summary>
        /// A value denoting which cloud service this service integrates with.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// How often, in minutes, to refresh the service.
        /// </summary>
        [Input("serviceRefreshRateInMinutes")]
        public Input<int>? ServiceRefreshRateInMinutes { get; set; }

        /// <summary>
        /// Tenant ID for an Azure service account within your project.
        /// </summary>
        [Input("tenant", required: true)]
        public Input<string> Tenant { get; set; } = null!;

        public CloudIntegrationAzureArgs()
        {
        }
        public static new CloudIntegrationAzureArgs Empty => new CloudIntegrationAzureArgs();
    }

    public sealed class CloudIntegrationAzureState : global::Pulumi.ResourceArgs
    {
        [Input("additionalTags")]
        private InputMap<string>? _additionalTags;

        /// <summary>
        /// A list of point tag key-values to add to every point ingested using this integration.
        /// </summary>
        public InputMap<string> AdditionalTags
        {
            get => _additionalTags ?? (_additionalTags = new InputMap<string>());
            set => _additionalTags = value;
        }

        [Input("categoryFilters")]
        private InputList<string>? _categoryFilters;

        /// <summary>
        /// A list of Azure Activity Log categories.
        /// </summary>
        public InputList<string> CategoryFilters
        {
            get => _categoryFilters ?? (_categoryFilters = new InputList<string>());
            set => _categoryFilters = value;
        }

        /// <summary>
        /// Client ID for an Azure service account within your project.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// Client secret for an Azure service account within your project.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Forces this resource to save, even if errors are present.
        /// </summary>
        [Input("forceSave")]
        public Input<bool>? ForceSave { get; set; }

        /// <summary>
        /// A regular expression that a metric name must match (case-insensitively) in order to be ingested.
        /// </summary>
        [Input("metricFilterRegex")]
        public Input<string>? MetricFilterRegex { get; set; }

        /// <summary>
        /// The human-readable name of this integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resourceGroupFilters")]
        private InputList<string>? _resourceGroupFilters;

        /// <summary>
        /// A list of Azure resource groups from which to pull metrics.
        /// </summary>
        public InputList<string> ResourceGroupFilters
        {
            get => _resourceGroupFilters ?? (_resourceGroupFilters = new InputList<string>());
            set => _resourceGroupFilters = value;
        }

        /// <summary>
        /// A value denoting which cloud service this service integrates with.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// How often, in minutes, to refresh the service.
        /// </summary>
        [Input("serviceRefreshRateInMinutes")]
        public Input<int>? ServiceRefreshRateInMinutes { get; set; }

        /// <summary>
        /// Tenant ID for an Azure service account within your project.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        public CloudIntegrationAzureState()
        {
        }
        public static new CloudIntegrationAzureState Empty => new CloudIntegrationAzureState();
    }
}
