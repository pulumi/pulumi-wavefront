// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront
{
    public static class GetEvents
    {
        /// <summary>
        /// Use this data source to get information about all Wavefront events.
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
        ///     // Get the information about all events
        ///     var example = Wavefront.GetEvents.Invoke(new()
        ///     {
        ///         Limit = 10,
        ///         Offset = 0,
        ///         LatestStartTimeEpochMillis = 1665427195,
        ///         EarliestStartTimeEpochMillis = 1665427195,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetEventsResult> InvokeAsync(GetEventsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEventsResult>("wavefront:index/getEvents:getEvents", args ?? new GetEventsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about all Wavefront events.
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
        ///     // Get the information about all events
        ///     var example = Wavefront.GetEvents.Invoke(new()
        ///     {
        ///         Limit = 10,
        ///         Offset = 0,
        ///         LatestStartTimeEpochMillis = 1665427195,
        ///         EarliestStartTimeEpochMillis = 1665427195,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEventsResult> Invoke(GetEventsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventsResult>("wavefront:index/getEvents:getEvents", args ?? new GetEventsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about all Wavefront events.
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
        ///     // Get the information about all events
        ///     var example = Wavefront.GetEvents.Invoke(new()
        ///     {
        ///         Limit = 10,
        ///         Offset = 0,
        ///         LatestStartTimeEpochMillis = 1665427195,
        ///         EarliestStartTimeEpochMillis = 1665427195,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEventsResult> Invoke(GetEventsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventsResult>("wavefront:index/getEvents:getEvents", args ?? new GetEventsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The earliest start time in epoch milliseconds.
        /// </summary>
        [Input("earliestStartTimeEpochMillis", required: true)]
        public int EarliestStartTimeEpochMillis { get; set; }

        /// <summary>
        /// The latest start time in epoch milliseconds.
        /// </summary>
        [Input("latestStartTimeEpochMillis", required: true)]
        public int LatestStartTimeEpochMillis { get; set; }

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

        public GetEventsArgs()
        {
        }
        public static new GetEventsArgs Empty => new GetEventsArgs();
    }

    public sealed class GetEventsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The earliest start time in epoch milliseconds.
        /// </summary>
        [Input("earliestStartTimeEpochMillis", required: true)]
        public Input<int> EarliestStartTimeEpochMillis { get; set; } = null!;

        /// <summary>
        /// The latest start time in epoch milliseconds.
        /// </summary>
        [Input("latestStartTimeEpochMillis", required: true)]
        public Input<int> LatestStartTimeEpochMillis { get; set; } = null!;

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

        public GetEventsInvokeArgs()
        {
        }
        public static new GetEventsInvokeArgs Empty => new GetEventsInvokeArgs();
    }


    [OutputType]
    public sealed class GetEventsResult
    {
        /// <summary>
        /// Earliest start time in epoch milliseconds.
        /// </summary>
        public readonly int EarliestStartTimeEpochMillis;
        /// <summary>
        /// List of all events in Wavefront. For each event you will see a list of attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEventsEventResult> Events;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Latest start time in epoch milliseconds.
        /// </summary>
        public readonly int LatestStartTimeEpochMillis;
        public readonly int? Limit;
        public readonly int? Offset;

        [OutputConstructor]
        private GetEventsResult(
            int earliestStartTimeEpochMillis,

            ImmutableArray<Outputs.GetEventsEventResult> events,

            string id,

            int latestStartTimeEpochMillis,

            int? limit,

            int? offset)
        {
            EarliestStartTimeEpochMillis = earliestStartTimeEpochMillis;
            Events = events;
            Id = id;
            LatestStartTimeEpochMillis = latestStartTimeEpochMillis;
            Limit = limit;
            Offset = offset;
        }
    }
}
