// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wavefront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-wavefront/sdk/v3/go/wavefront/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Wavefront External Link Resource. This allows external links to be created, updated, and deleted.
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
//			_, err := wavefront.NewExternalLink(ctx, "basic", &wavefront.ExternalLinkArgs{
//				Name:        pulumi.String("External Link"),
//				Description: pulumi.String("An external link description"),
//				Template:    pulumi.String("https://example.com/source={{{source}}}&startTime={{startEpochMillis}}"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Maintenance windows can be imported by using the `id`, e.g.:
//
// ```sh
// $ pulumi import wavefront:index/externalLink:ExternalLink basic fVj6fz6zYC4aBkID
// ```
type ExternalLink struct {
	pulumi.CustomResourceState

	// Human-readable description for this link.
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether this is a "Log Integration" subType of external link.
	IsLogIntegration pulumi.BoolPtrOutput `pulumi:"isLogIntegration"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex pulumi.StringPtrOutput `pulumi:"metricFilterRegex"`
	// The name of the external link.
	Name pulumi.StringOutput `pulumi:"name"`
	// Controls whether a link is displayed in the context menu of a highlighted
	// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
	// keys are present in the keys of this map and whose values match the regular expressions associated with those
	// keys in order for the link to be displayed.
	PointTagFilterRegexes pulumi.StringMapOutput `pulumi:"pointTagFilterRegexes"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex pulumi.StringPtrOutput `pulumi:"sourceFilterRegex"`
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template pulumi.StringOutput `pulumi:"template"`
}

// NewExternalLink registers a new resource with the given unique name, arguments, and options.
func NewExternalLink(ctx *pulumi.Context,
	name string, args *ExternalLinkArgs, opts ...pulumi.ResourceOption) (*ExternalLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExternalLink
	err := ctx.RegisterResource("wavefront:index/externalLink:ExternalLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalLink gets an existing ExternalLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalLinkState, opts ...pulumi.ResourceOption) (*ExternalLink, error) {
	var resource ExternalLink
	err := ctx.ReadResource("wavefront:index/externalLink:ExternalLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalLink resources.
type externalLinkState struct {
	// Human-readable description for this link.
	Description *string `pulumi:"description"`
	// Whether this is a "Log Integration" subType of external link.
	IsLogIntegration *bool `pulumi:"isLogIntegration"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The name of the external link.
	Name *string `pulumi:"name"`
	// Controls whether a link is displayed in the context menu of a highlighted
	// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
	// keys are present in the keys of this map and whose values match the regular expressions associated with those
	// keys in order for the link to be displayed.
	PointTagFilterRegexes map[string]string `pulumi:"pointTagFilterRegexes"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex *string `pulumi:"sourceFilterRegex"`
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template *string `pulumi:"template"`
}

type ExternalLinkState struct {
	// Human-readable description for this link.
	Description pulumi.StringPtrInput
	// Whether this is a "Log Integration" subType of external link.
	IsLogIntegration pulumi.BoolPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex pulumi.StringPtrInput
	// The name of the external link.
	Name pulumi.StringPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted
	// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
	// keys are present in the keys of this map and whose values match the regular expressions associated with those
	// keys in order for the link to be displayed.
	PointTagFilterRegexes pulumi.StringMapInput
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex pulumi.StringPtrInput
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template pulumi.StringPtrInput
}

func (ExternalLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalLinkState)(nil)).Elem()
}

type externalLinkArgs struct {
	// Human-readable description for this link.
	Description string `pulumi:"description"`
	// Whether this is a "Log Integration" subType of external link.
	IsLogIntegration *bool `pulumi:"isLogIntegration"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The name of the external link.
	Name *string `pulumi:"name"`
	// Controls whether a link is displayed in the context menu of a highlighted
	// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
	// keys are present in the keys of this map and whose values match the regular expressions associated with those
	// keys in order for the link to be displayed.
	PointTagFilterRegexes map[string]string `pulumi:"pointTagFilterRegexes"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex *string `pulumi:"sourceFilterRegex"`
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template string `pulumi:"template"`
}

// The set of arguments for constructing a ExternalLink resource.
type ExternalLinkArgs struct {
	// Human-readable description for this link.
	Description pulumi.StringInput
	// Whether this is a "Log Integration" subType of external link.
	IsLogIntegration pulumi.BoolPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex pulumi.StringPtrInput
	// The name of the external link.
	Name pulumi.StringPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted
	// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
	// keys are present in the keys of this map and whose values match the regular expressions associated with those
	// keys in order for the link to be displayed.
	PointTagFilterRegexes pulumi.StringMapInput
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex pulumi.StringPtrInput
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template pulumi.StringInput
}

func (ExternalLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalLinkArgs)(nil)).Elem()
}

type ExternalLinkInput interface {
	pulumi.Input

	ToExternalLinkOutput() ExternalLinkOutput
	ToExternalLinkOutputWithContext(ctx context.Context) ExternalLinkOutput
}

func (*ExternalLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalLink)(nil)).Elem()
}

func (i *ExternalLink) ToExternalLinkOutput() ExternalLinkOutput {
	return i.ToExternalLinkOutputWithContext(context.Background())
}

func (i *ExternalLink) ToExternalLinkOutputWithContext(ctx context.Context) ExternalLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalLinkOutput)
}

// ExternalLinkArrayInput is an input type that accepts ExternalLinkArray and ExternalLinkArrayOutput values.
// You can construct a concrete instance of `ExternalLinkArrayInput` via:
//
//	ExternalLinkArray{ ExternalLinkArgs{...} }
type ExternalLinkArrayInput interface {
	pulumi.Input

	ToExternalLinkArrayOutput() ExternalLinkArrayOutput
	ToExternalLinkArrayOutputWithContext(context.Context) ExternalLinkArrayOutput
}

type ExternalLinkArray []ExternalLinkInput

func (ExternalLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalLink)(nil)).Elem()
}

func (i ExternalLinkArray) ToExternalLinkArrayOutput() ExternalLinkArrayOutput {
	return i.ToExternalLinkArrayOutputWithContext(context.Background())
}

func (i ExternalLinkArray) ToExternalLinkArrayOutputWithContext(ctx context.Context) ExternalLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalLinkArrayOutput)
}

// ExternalLinkMapInput is an input type that accepts ExternalLinkMap and ExternalLinkMapOutput values.
// You can construct a concrete instance of `ExternalLinkMapInput` via:
//
//	ExternalLinkMap{ "key": ExternalLinkArgs{...} }
type ExternalLinkMapInput interface {
	pulumi.Input

	ToExternalLinkMapOutput() ExternalLinkMapOutput
	ToExternalLinkMapOutputWithContext(context.Context) ExternalLinkMapOutput
}

type ExternalLinkMap map[string]ExternalLinkInput

func (ExternalLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalLink)(nil)).Elem()
}

func (i ExternalLinkMap) ToExternalLinkMapOutput() ExternalLinkMapOutput {
	return i.ToExternalLinkMapOutputWithContext(context.Background())
}

func (i ExternalLinkMap) ToExternalLinkMapOutputWithContext(ctx context.Context) ExternalLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalLinkMapOutput)
}

type ExternalLinkOutput struct{ *pulumi.OutputState }

func (ExternalLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalLink)(nil)).Elem()
}

func (o ExternalLinkOutput) ToExternalLinkOutput() ExternalLinkOutput {
	return o
}

func (o ExternalLinkOutput) ToExternalLinkOutputWithContext(ctx context.Context) ExternalLinkOutput {
	return o
}

// Human-readable description for this link.
func (o ExternalLinkOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Whether this is a "Log Integration" subType of external link.
func (o ExternalLinkOutput) IsLogIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.BoolPtrOutput { return v.IsLogIntegration }).(pulumi.BoolPtrOutput)
}

// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
func (o ExternalLinkOutput) MetricFilterRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringPtrOutput { return v.MetricFilterRegex }).(pulumi.StringPtrOutput)
}

// The name of the external link.
func (o ExternalLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Controls whether a link is displayed in the context menu of a highlighted
// series. This is a map from string to regular expression. The highlighted series must contain point tags whose
// keys are present in the keys of this map and whose values match the regular expressions associated with those
// keys in order for the link to be displayed.
func (o ExternalLinkOutput) PointTagFilterRegexes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringMapOutput { return v.PointTagFilterRegexes }).(pulumi.StringMapOutput)
}

// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
func (o ExternalLinkOutput) SourceFilterRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringPtrOutput { return v.SourceFilterRegex }).(pulumi.StringPtrOutput)
}

// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
func (o ExternalLinkOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

type ExternalLinkArrayOutput struct{ *pulumi.OutputState }

func (ExternalLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalLink)(nil)).Elem()
}

func (o ExternalLinkArrayOutput) ToExternalLinkArrayOutput() ExternalLinkArrayOutput {
	return o
}

func (o ExternalLinkArrayOutput) ToExternalLinkArrayOutputWithContext(ctx context.Context) ExternalLinkArrayOutput {
	return o
}

func (o ExternalLinkArrayOutput) Index(i pulumi.IntInput) ExternalLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalLink {
		return vs[0].([]*ExternalLink)[vs[1].(int)]
	}).(ExternalLinkOutput)
}

type ExternalLinkMapOutput struct{ *pulumi.OutputState }

func (ExternalLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalLink)(nil)).Elem()
}

func (o ExternalLinkMapOutput) ToExternalLinkMapOutput() ExternalLinkMapOutput {
	return o
}

func (o ExternalLinkMapOutput) ToExternalLinkMapOutputWithContext(ctx context.Context) ExternalLinkMapOutput {
	return o
}

func (o ExternalLinkMapOutput) MapIndex(k pulumi.StringInput) ExternalLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalLink {
		return vs[0].(map[string]*ExternalLink)[vs[1].(string)]
	}).(ExternalLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalLinkInput)(nil)).Elem(), &ExternalLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalLinkArrayInput)(nil)).Elem(), ExternalLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalLinkMapInput)(nil)).Elem(), ExternalLinkMap{})
	pulumi.RegisterOutputType(ExternalLinkOutput{})
	pulumi.RegisterOutputType(ExternalLinkArrayOutput{})
	pulumi.RegisterOutputType(ExternalLinkMapOutput{})
}
