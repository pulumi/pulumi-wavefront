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
    /// Provides a Wavefront Service Account Resource. This allows service accounts to be created, updated, and deleted.
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
    ///         var basic = new Wavefront.ServiceAccount("basic", new Wavefront.ServiceAccountArgs
    ///         {
    ///             Active = true,
    ///             Identifier = "sa::tftesting",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ServiceAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether or not the service account is active
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// The description of the service account
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The (unique) identifier of the service account to create. Must start with sa::
        /// </summary>
        [Output("identifier")]
        public Output<string> Identifier { get; private set; } = null!;

        /// <summary>
        /// List of permission to grant to this service account.  Valid options are
        /// `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
        /// `host_tag_management`, `metrics_management`, `user_management`
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// List of user groups for this service account
        /// </summary>
        [Output("userGroups")]
        public Output<ImmutableArray<string>> UserGroups { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceAccount(string name, ServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("wavefront:index/serviceAccount:ServiceAccount", name, args ?? new ServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceAccount(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("wavefront:index/serviceAccount:ServiceAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceAccount Get(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceAccount(name, id, state, options);
        }
    }

    public sealed class ServiceAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the service account is active
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The description of the service account
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The (unique) identifier of the service account to create. Must start with sa::
        /// </summary>
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// List of permission to grant to this service account.  Valid options are
        /// `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
        /// `host_tag_management`, `metrics_management`, `user_management`
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        [Input("userGroups")]
        private InputList<string>? _userGroups;

        /// <summary>
        /// List of user groups for this service account
        /// </summary>
        public InputList<string> UserGroups
        {
            get => _userGroups ?? (_userGroups = new InputList<string>());
            set => _userGroups = value;
        }

        public ServiceAccountArgs()
        {
        }
    }

    public sealed class ServiceAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the service account is active
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The description of the service account
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The (unique) identifier of the service account to create. Must start with sa::
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// List of permission to grant to this service account.  Valid options are
        /// `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
        /// `host_tag_management`, `metrics_management`, `user_management`
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        [Input("userGroups")]
        private InputList<string>? _userGroups;

        /// <summary>
        /// List of user groups for this service account
        /// </summary>
        public InputList<string> UserGroups
        {
            get => _userGroups ?? (_userGroups = new InputList<string>());
            set => _userGroups = value;
        }

        public ServiceAccountState()
        {
        }
    }
}