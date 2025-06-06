// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wavefront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-wavefront/sdk/v3/go/wavefront/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a Wavefront external link by its ID.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-wavefront/sdk/v3/go/wavefront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Get the information about a specific external links.
//			_, err := wavefront.LookupExternalLink(ctx, &wavefront.LookupExternalLinkArgs{
//				Id: "sample-external-link-id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupExternalLink(ctx *pulumi.Context, args *LookupExternalLinkArgs, opts ...pulumi.InvokeOption) (*LookupExternalLinkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupExternalLinkResult
	err := ctx.Invoke("wavefront:index/getExternalLink:getExternalLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExternalLink.
type LookupExternalLinkArgs struct {
	// The ID of the external link.
	Id string `pulumi:"id"`
}

// A collection of values returned by getExternalLink.
type LookupExternalLinkResult struct {
	// The timestamp in epoch milliseconds indicating when the external link is created.
	CreatedEpochMillis int `pulumi:"createdEpochMillis"`
	// The ID of the user who created the external link.
	CreatorId string `pulumi:"creatorId"`
	// Human-readable description of this link.
	Description string `pulumi:"description"`
	// The ID of the external link.
	Id string `pulumi:"id"`
	// Whether this is a "Log Integration" subType of external link.
	IsLogIntegration bool `pulumi:"isLogIntegration"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex string `pulumi:"metricFilterRegex"`
	// The name of the external link.
	Name string `pulumi:"name"`
	// (Optional) Controls whether a link is displayed in the context menu of a highlighted
	// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
	// keys are present in the keys of this map and whose values match the regular expressions associated with those
	// keys in order for the link to be displayed.
	PointTagFilterRegexes map[string]string `pulumi:"pointTagFilterRegexes"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex string `pulumi:"sourceFilterRegex"`
	// The mustache template for the link. The template must expand to a full URL, including scheme, origin, etc.
	Template string `pulumi:"template"`
	// The timestamp in epoch milliseconds indicating when the external link is updated.
	UpdatedEpochMillis int `pulumi:"updatedEpochMillis"`
	// The ID of the user who updated the external link.
	UpdaterId string `pulumi:"updaterId"`
}

func LookupExternalLinkOutput(ctx *pulumi.Context, args LookupExternalLinkOutputArgs, opts ...pulumi.InvokeOption) LookupExternalLinkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExternalLinkResultOutput, error) {
			args := v.(LookupExternalLinkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("wavefront:index/getExternalLink:getExternalLink", args, LookupExternalLinkResultOutput{}, options).(LookupExternalLinkResultOutput), nil
		}).(LookupExternalLinkResultOutput)
}

// A collection of arguments for invoking getExternalLink.
type LookupExternalLinkOutputArgs struct {
	// The ID of the external link.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupExternalLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalLinkArgs)(nil)).Elem()
}

// A collection of values returned by getExternalLink.
type LookupExternalLinkResultOutput struct{ *pulumi.OutputState }

func (LookupExternalLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalLinkResult)(nil)).Elem()
}

func (o LookupExternalLinkResultOutput) ToLookupExternalLinkResultOutput() LookupExternalLinkResultOutput {
	return o
}

func (o LookupExternalLinkResultOutput) ToLookupExternalLinkResultOutputWithContext(ctx context.Context) LookupExternalLinkResultOutput {
	return o
}

// The timestamp in epoch milliseconds indicating when the external link is created.
func (o LookupExternalLinkResultOutput) CreatedEpochMillis() pulumi.IntOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) int { return v.CreatedEpochMillis }).(pulumi.IntOutput)
}

// The ID of the user who created the external link.
func (o LookupExternalLinkResultOutput) CreatorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.CreatorId }).(pulumi.StringOutput)
}

// Human-readable description of this link.
func (o LookupExternalLinkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.Description }).(pulumi.StringOutput)
}

// The ID of the external link.
func (o LookupExternalLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether this is a "Log Integration" subType of external link.
func (o LookupExternalLinkResultOutput) IsLogIntegration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) bool { return v.IsLogIntegration }).(pulumi.BoolOutput)
}

// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
func (o LookupExternalLinkResultOutput) MetricFilterRegex() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.MetricFilterRegex }).(pulumi.StringOutput)
}

// The name of the external link.
func (o LookupExternalLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// (Optional) Controls whether a link is displayed in the context menu of a highlighted
// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
// keys are present in the keys of this map and whose values match the regular expressions associated with those
// keys in order for the link to be displayed.
func (o LookupExternalLinkResultOutput) PointTagFilterRegexes() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) map[string]string { return v.PointTagFilterRegexes }).(pulumi.StringMapOutput)
}

// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
func (o LookupExternalLinkResultOutput) SourceFilterRegex() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.SourceFilterRegex }).(pulumi.StringOutput)
}

// The mustache template for the link. The template must expand to a full URL, including scheme, origin, etc.
func (o LookupExternalLinkResultOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.Template }).(pulumi.StringOutput)
}

// The timestamp in epoch milliseconds indicating when the external link is updated.
func (o LookupExternalLinkResultOutput) UpdatedEpochMillis() pulumi.IntOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) int { return v.UpdatedEpochMillis }).(pulumi.IntOutput)
}

// The ID of the user who updated the external link.
func (o LookupExternalLinkResultOutput) UpdaterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalLinkResult) string { return v.UpdaterId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExternalLinkResultOutput{})
}
