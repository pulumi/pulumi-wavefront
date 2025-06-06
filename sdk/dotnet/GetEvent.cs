// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront
{
    public static class GetEvent
    {
        /// <summary>
        /// Use this data source to get information about a certain Wavefront event.
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
        ///     // Get the information about a Wavefront event by its ID.
        ///     var example = Wavefront.GetEvent.Invoke(new()
        ///     {
        ///         Id = "sample-event-id",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetEventResult> InvokeAsync(GetEventArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEventResult>("wavefront:index/getEvent:getEvent", args ?? new GetEventArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a certain Wavefront event.
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
        ///     // Get the information about a Wavefront event by its ID.
        ///     var example = Wavefront.GetEvent.Invoke(new()
        ///     {
        ///         Id = "sample-event-id",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEventResult> Invoke(GetEventInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventResult>("wavefront:index/getEvent:getEvent", args ?? new GetEventInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a certain Wavefront event.
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
        ///     // Get the information about a Wavefront event by its ID.
        ///     var example = Wavefront.GetEvent.Invoke(new()
        ///     {
        ///         Id = "sample-event-id",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEventResult> Invoke(GetEventInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventResult>("wavefront:index/getEvent:getEvent", args ?? new GetEventInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID associated with the event data to be fetched.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetEventArgs()
        {
        }
        public static new GetEventArgs Empty => new GetEventArgs();
    }

    public sealed class GetEventInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID associated with the event data to be fetched.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetEventInvokeArgs()
        {
        }
        public static new GetEventInvokeArgs Empty => new GetEventInvokeArgs();
    }


    [OutputType]
    public sealed class GetEventResult
    {
        /// <summary>
        /// Annotations associated with the event.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// The description of the event.
        /// </summary>
        public readonly string Details;
        public readonly int EndtimeKey;
        /// <summary>
        /// The ID of the event in Wavefront.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A Boolean flag. If set to `true`, creates a point-in-time event (i.e. with no duration).
        /// </summary>
        public readonly bool IsEphemeral;
        /// <summary>
        /// The name of the event in Wavefront.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The severity category of the event.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// The start time of the event in epoch milliseconds.
        /// </summary>
        public readonly int StartTime;
        /// <summary>
        /// A set of tags assigned to the event.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of the event.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetEventResult(
            ImmutableDictionary<string, string> annotations,

            string details,

            int endtimeKey,

            string id,

            bool isEphemeral,

            string name,

            string severity,

            int startTime,

            ImmutableArray<string> tags,

            string type)
        {
            Annotations = annotations;
            Details = details;
            EndtimeKey = endtimeKey;
            Id = id;
            IsEphemeral = isEphemeral;
            Name = name;
            Severity = severity;
            StartTime = startTime;
            Tags = tags;
            Type = type;
        }
    }
}
