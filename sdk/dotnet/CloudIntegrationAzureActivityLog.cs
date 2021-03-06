// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront
{
    /// <summary>
    /// Provides a Wavefront Cloud Integration for Azure Activity Logs. This allows azure activity log cloud integrations to be created,
    /// updated, and deleted.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Wavefront = Pulumi.Wavefront;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var azureActivityLog = new Wavefront.CloudIntegrationAzureActivityLog("azureActivityLog", new Wavefront.CloudIntegrationAzureActivityLogArgs
    ///         {
    ///             CategoryFilters = 
    ///             {
    ///                 "ADMINISTRATIVE",
    ///             },
    ///             ClientId = "client-id2",
    ///             ClientSecret = "client-secret2",
    ///             Tenant = "my-tenant2",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure Activity Log Cloud Integrations can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import wavefront:index/cloudIntegrationAzureActivityLog:CloudIntegrationAzureActivityLog azure_al a411c16b-3cf7-4f03-bf11-8ca05aab898d
    /// ```
    /// </summary>
    [WavefrontResourceType("wavefront:index/cloudIntegrationAzureActivityLog:CloudIntegrationAzureActivityLog")]
    public partial class CloudIntegrationAzureActivityLog : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of point tag key-values to add to every point ingested using this integration
        /// </summary>
        [Output("additionalTags")]
        public Output<ImmutableDictionary<string, string>?> AdditionalTags { get; private set; } = null!;

        /// <summary>
        /// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics
        /// </summary>
        [Output("categoryFilters")]
        public Output<ImmutableArray<string>> CategoryFilters { get; private set; } = null!;

        /// <summary>
        /// Client id for an azure service account within your project
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client secret for an Azure service account within your project
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Forces this resource to save, even if errors are present
        /// </summary>
        [Output("forceSave")]
        public Output<bool?> ForceSave { get; private set; } = null!;

        /// <summary>
        /// The human-readable name of this integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A value denoting which cloud service this service integrates with
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// How often, in minutes, to refresh the service
        /// </summary>
        [Output("serviceRefreshRateInMinutes")]
        public Output<int?> ServiceRefreshRateInMinutes { get; private set; } = null!;

        /// <summary>
        /// Tenant Id for an Azure service account within your project
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;


        /// <summary>
        /// Create a CloudIntegrationAzureActivityLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudIntegrationAzureActivityLog(string name, CloudIntegrationAzureActivityLogArgs args, CustomResourceOptions? options = null)
            : base("wavefront:index/cloudIntegrationAzureActivityLog:CloudIntegrationAzureActivityLog", name, args ?? new CloudIntegrationAzureActivityLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudIntegrationAzureActivityLog(string name, Input<string> id, CloudIntegrationAzureActivityLogState? state = null, CustomResourceOptions? options = null)
            : base("wavefront:index/cloudIntegrationAzureActivityLog:CloudIntegrationAzureActivityLog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CloudIntegrationAzureActivityLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudIntegrationAzureActivityLog Get(string name, Input<string> id, CloudIntegrationAzureActivityLogState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudIntegrationAzureActivityLog(name, id, state, options);
        }
    }

    public sealed class CloudIntegrationAzureActivityLogArgs : Pulumi.ResourceArgs
    {
        [Input("additionalTags")]
        private InputMap<string>? _additionalTags;

        /// <summary>
        /// A list of point tag key-values to add to every point ingested using this integration
        /// </summary>
        public InputMap<string> AdditionalTags
        {
            get => _additionalTags ?? (_additionalTags = new InputMap<string>());
            set => _additionalTags = value;
        }

        [Input("categoryFilters")]
        private InputList<string>? _categoryFilters;

        /// <summary>
        /// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics
        /// </summary>
        public InputList<string> CategoryFilters
        {
            get => _categoryFilters ?? (_categoryFilters = new InputList<string>());
            set => _categoryFilters = value;
        }

        /// <summary>
        /// Client id for an azure service account within your project
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Client secret for an Azure service account within your project
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// Forces this resource to save, even if errors are present
        /// </summary>
        [Input("forceSave")]
        public Input<bool>? ForceSave { get; set; }

        /// <summary>
        /// The human-readable name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A value denoting which cloud service this service integrates with
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// How often, in minutes, to refresh the service
        /// </summary>
        [Input("serviceRefreshRateInMinutes")]
        public Input<int>? ServiceRefreshRateInMinutes { get; set; }

        /// <summary>
        /// Tenant Id for an Azure service account within your project
        /// </summary>
        [Input("tenant", required: true)]
        public Input<string> Tenant { get; set; } = null!;

        public CloudIntegrationAzureActivityLogArgs()
        {
        }
    }

    public sealed class CloudIntegrationAzureActivityLogState : Pulumi.ResourceArgs
    {
        [Input("additionalTags")]
        private InputMap<string>? _additionalTags;

        /// <summary>
        /// A list of point tag key-values to add to every point ingested using this integration
        /// </summary>
        public InputMap<string> AdditionalTags
        {
            get => _additionalTags ?? (_additionalTags = new InputMap<string>());
            set => _additionalTags = value;
        }

        [Input("categoryFilters")]
        private InputList<string>? _categoryFilters;

        /// <summary>
        /// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics
        /// </summary>
        public InputList<string> CategoryFilters
        {
            get => _categoryFilters ?? (_categoryFilters = new InputList<string>());
            set => _categoryFilters = value;
        }

        /// <summary>
        /// Client id for an azure service account within your project
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client secret for an Azure service account within your project
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Forces this resource to save, even if errors are present
        /// </summary>
        [Input("forceSave")]
        public Input<bool>? ForceSave { get; set; }

        /// <summary>
        /// The human-readable name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A value denoting which cloud service this service integrates with
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// How often, in minutes, to refresh the service
        /// </summary>
        [Input("serviceRefreshRateInMinutes")]
        public Input<int>? ServiceRefreshRateInMinutes { get; set; }

        /// <summary>
        /// Tenant Id for an Azure service account within your project
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        public CloudIntegrationAzureActivityLogState()
        {
        }
    }
}
