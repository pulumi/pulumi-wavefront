// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront
{
    public static class GetMaintenanceWindow
    {
        /// <summary>
        /// Use this data source to get information about a Wavefront maintenance window by its ID.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Wavefront = Pulumi.Wavefront;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Wavefront.GetMaintenanceWindow.Invoke(new()
        ///     {
        ///         Id = "sample-maintenance-window-id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMaintenanceWindowResult> InvokeAsync(GetMaintenanceWindowArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMaintenanceWindowResult>("wavefront:index/getMaintenanceWindow:getMaintenanceWindow", args ?? new GetMaintenanceWindowArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a Wavefront maintenance window by its ID.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Wavefront = Pulumi.Wavefront;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Wavefront.GetMaintenanceWindow.Invoke(new()
        ///     {
        ///         Id = "sample-maintenance-window-id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMaintenanceWindowResult> Invoke(GetMaintenanceWindowInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMaintenanceWindowResult>("wavefront:index/getMaintenanceWindow:getMaintenanceWindow", args ?? new GetMaintenanceWindowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMaintenanceWindowArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the maintenance window.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetMaintenanceWindowArgs()
        {
        }
        public static new GetMaintenanceWindowArgs Empty => new GetMaintenanceWindowArgs();
    }

    public sealed class GetMaintenanceWindowInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the maintenance window.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetMaintenanceWindowInvokeArgs()
        {
        }
        public static new GetMaintenanceWindowInvokeArgs Empty => new GetMaintenanceWindowInvokeArgs();
    }


    [OutputType]
    public sealed class GetMaintenanceWindowResult
    {
        /// <summary>
        /// The timestamp in epoch milliseconds indicating when the maintenance window is created.
        /// </summary>
        public readonly int CreatedEpochMillis;
        /// <summary>
        /// The ID of the user who created the maintenance window.
        /// </summary>
        public readonly string CreatorId;
        /// <summary>
        /// The ID of the customer in Wavefront.
        /// </summary>
        public readonly string CustomerId;
        /// <summary>
        /// The end time in seconds after 1 Jan 1970 GMT.
        /// </summary>
        public readonly int EndTimeInSeconds;
        /// <summary>
        /// The event name of the maintenance window.
        /// </summary>
        public readonly string EventName;
        /// <summary>
        /// If set to `true`, the source or host must be in `relevant_host_names` and must have tags matching the specification formed by `relevant_host_tags` and `relevant_host_tags_anded` in for this maintenance window to apply. 
        /// If set to false, the source or host must either be in `relevant_host_names` or match `relevant_host_tags` and `relevant_host_tags_anded`. Default value is `false`.
        /// </summary>
        public readonly bool HostTagGroupHostNamesGroupAnded;
        /// <summary>
        /// The ID of the maintenance window.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The reason for the maintenance window.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// The list of alert tags whose matching alerts will be put into maintenance because
        /// of this maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or `relevant_host_names`
        /// is required.
        /// </summary>
        public readonly ImmutableArray<string> RelevantCustomerTags;
        /// <summary>
        /// The list of source or host names that will be put into maintenance because of this
        /// maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or `relevant_host_names`
        /// is required.
        /// </summary>
        public readonly ImmutableArray<string> RelevantHostNames;
        /// <summary>
        /// The list of source or host tags whose matching sources or hosts will be put into maintenance
        /// because of this maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or
        /// `relevant_host_names` is required.
        /// </summary>
        public readonly ImmutableArray<string> RelevantHostTags;
        /// <summary>
        /// Whether to AND source or host tags listed in `relevant_host_tags`.
        /// If set to `true`, the source or host must contain all tags for the maintenance window to apply. If set to `false`,
        /// the tags are OR'ed, and the source or host must contain one of the tags. Default value is `false`.
        /// </summary>
        public readonly bool RelevantHostTagsAnded;
        /// <summary>
        /// The running state of the maintenance window.
        /// </summary>
        public readonly string RunningState;
        public readonly int SortAttr;
        /// <summary>
        /// The start time in seconds after 1 Jan 1970 GMT.
        /// </summary>
        public readonly int StartTimeInSeconds;
        /// <summary>
        /// The title of the maintenance window.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// The timestamp in epoch milliseconds indicating when the maintenance window is updated.
        /// </summary>
        public readonly int UpdatedEpochMillis;
        /// <summary>
        /// The ID of the user who updated the maintenance window.
        /// </summary>
        public readonly string UpdaterId;

        [OutputConstructor]
        private GetMaintenanceWindowResult(
            int createdEpochMillis,

            string creatorId,

            string customerId,

            int endTimeInSeconds,

            string eventName,

            bool hostTagGroupHostNamesGroupAnded,

            string id,

            string reason,

            ImmutableArray<string> relevantCustomerTags,

            ImmutableArray<string> relevantHostNames,

            ImmutableArray<string> relevantHostTags,

            bool relevantHostTagsAnded,

            string runningState,

            int sortAttr,

            int startTimeInSeconds,

            string title,

            int updatedEpochMillis,

            string updaterId)
        {
            CreatedEpochMillis = createdEpochMillis;
            CreatorId = creatorId;
            CustomerId = customerId;
            EndTimeInSeconds = endTimeInSeconds;
            EventName = eventName;
            HostTagGroupHostNamesGroupAnded = hostTagGroupHostNamesGroupAnded;
            Id = id;
            Reason = reason;
            RelevantCustomerTags = relevantCustomerTags;
            RelevantHostNames = relevantHostNames;
            RelevantHostTags = relevantHostTags;
            RelevantHostTagsAnded = relevantHostTagsAnded;
            RunningState = runningState;
            SortAttr = sortAttr;
            StartTimeInSeconds = startTimeInSeconds;
            Title = title;
            UpdatedEpochMillis = updatedEpochMillis;
            UpdaterId = updaterId;
        }
    }
}