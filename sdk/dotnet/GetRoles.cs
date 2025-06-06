// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront
{
    public static class GetRoles
    {
        /// <summary>
        /// Use this data source to get all Roles in Wavefront.
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
        ///     // Get all Roles
        ///     var roles = Wavefront.GetRoles.Invoke(new()
        ///     {
        ///         Limit = 10,
        ///         Offset = 0,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRolesResult> InvokeAsync(GetRolesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRolesResult>("wavefront:index/getRoles:getRoles", args ?? new GetRolesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get all Roles in Wavefront.
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
        ///     // Get all Roles
        ///     var roles = Wavefront.GetRoles.Invoke(new()
        ///     {
        ///         Limit = 10,
        ///         Offset = 0,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRolesResult> Invoke(GetRolesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRolesResult>("wavefront:index/getRoles:getRoles", args ?? new GetRolesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get all Roles in Wavefront.
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
        ///     // Get all Roles
        ///     var roles = Wavefront.GetRoles.Invoke(new()
        ///     {
        ///         Limit = 10,
        ///         Offset = 0,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRolesResult> Invoke(GetRolesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRolesResult>("wavefront:index/getRoles:getRoles", args ?? new GetRolesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRolesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Limit is the maximum number of results to be returned. Defaults to 100.
        /// </summary>
        [Input("limit")]
        public int? Limit { get; set; }

        /// <summary>
        /// Offset is the offset from the first result to be returned. Defaults to 0.
        /// </summary>
        [Input("offset")]
        public int? Offset { get; set; }

        public GetRolesArgs()
        {
        }
        public static new GetRolesArgs Empty => new GetRolesArgs();
    }

    public sealed class GetRolesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Limit is the maximum number of results to be returned. Defaults to 100.
        /// </summary>
        [Input("limit")]
        public Input<int>? Limit { get; set; }

        /// <summary>
        /// Offset is the offset from the first result to be returned. Defaults to 0.
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        public GetRolesInvokeArgs()
        {
        }
        public static new GetRolesInvokeArgs Empty => new GetRolesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRolesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? Limit;
        public readonly int? Offset;
        /// <summary>
        /// List of Wavefront Roles.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRolesRoleResult> Roles;

        [OutputConstructor]
        private GetRolesResult(
            string id,

            int? limit,

            int? offset,

            ImmutableArray<Outputs.GetRolesRoleResult> roles)
        {
            Id = id;
            Limit = limit;
            Offset = offset;
            Roles = roles;
        }
    }
}
