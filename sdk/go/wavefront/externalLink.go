// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wavefront

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Wavefront External Link Resource. This allows external links to be created, updated, and deleted.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-wavefront/sdk/go/wavefront"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wavefront.NewExternalLink(ctx, "basic", &wavefront.ExternalLinkArgs{
// 			Description: pulumi.String("An external link description"),
// 			Template:    pulumi.String("https://example.com/source={{{source}}}&startTime={{startEpochMillis}}"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ExternalLink struct {
	pulumi.CustomResourceState

	// Human-readable description for this link
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether this is a "Log Integration" subType of external link
	IsLogIntegration pulumi.BoolPtrOutput `pulumi:"isLogIntegration"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex pulumi.StringPtrOutput `pulumi:"metricFilterRegex"`
	// The name of the external link
	Name pulumi.StringOutput `pulumi:"name"`
	// Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed
	PointTagFilterRegexes pulumi.StringMapOutput `pulumi:"pointTagFilterRegexes"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex pulumi.StringPtrOutput `pulumi:"sourceFilterRegex"`
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template pulumi.StringOutput `pulumi:"template"`
}

// NewExternalLink registers a new resource with the given unique name, arguments, and options.
func NewExternalLink(ctx *pulumi.Context,
	name string, args *ExternalLinkArgs, opts ...pulumi.ResourceOption) (*ExternalLink, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Template == nil {
		return nil, errors.New("missing required argument 'Template'")
	}
	if args == nil {
		args = &ExternalLinkArgs{}
	}
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
	// Human-readable description for this link
	Description *string `pulumi:"description"`
	// Whether this is a "Log Integration" subType of external link
	IsLogIntegration *bool `pulumi:"isLogIntegration"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The name of the external link
	Name *string `pulumi:"name"`
	// Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed
	PointTagFilterRegexes map[string]string `pulumi:"pointTagFilterRegexes"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex *string `pulumi:"sourceFilterRegex"`
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template *string `pulumi:"template"`
}

type ExternalLinkState struct {
	// Human-readable description for this link
	Description pulumi.StringPtrInput
	// Whether this is a "Log Integration" subType of external link
	IsLogIntegration pulumi.BoolPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex pulumi.StringPtrInput
	// The name of the external link
	Name pulumi.StringPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed
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
	// Human-readable description for this link
	Description string `pulumi:"description"`
	// Whether this is a "Log Integration" subType of external link
	IsLogIntegration *bool `pulumi:"isLogIntegration"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The name of the external link
	Name *string `pulumi:"name"`
	// Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed
	PointTagFilterRegexes map[string]string `pulumi:"pointTagFilterRegexes"`
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex *string `pulumi:"sourceFilterRegex"`
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template string `pulumi:"template"`
}

// The set of arguments for constructing a ExternalLink resource.
type ExternalLinkArgs struct {
	// Human-readable description for this link
	Description pulumi.StringInput
	// Whether this is a "Log Integration" subType of external link
	IsLogIntegration pulumi.BoolPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
	MetricFilterRegex pulumi.StringPtrInput
	// The name of the external link
	Name pulumi.StringPtrInput
	// Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed
	PointTagFilterRegexes pulumi.StringMapInput
	// Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
	SourceFilterRegex pulumi.StringPtrInput
	// The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
	Template pulumi.StringInput
}

func (ExternalLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalLinkArgs)(nil)).Elem()
}