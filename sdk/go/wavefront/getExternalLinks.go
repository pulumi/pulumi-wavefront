// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wavefront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about all Wavefront external links.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-wavefront/sdk/go/wavefront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = wavefront.GetExternalLinks(ctx, &GetExternalLinksArgs{
//				Limit:  pulumi.IntRef(10),
//				Offset: pulumi.IntRef(0),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetExternalLinks(ctx *pulumi.Context, args *GetExternalLinksArgs, opts ...pulumi.InvokeOption) (*GetExternalLinksResult, error) {
	var rv GetExternalLinksResult
	err := ctx.Invoke("wavefront:index/getExternalLinks:getExternalLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExternalLinks.
type GetExternalLinksArgs struct {
	// Limit is the maximum number of results to be returned. Defaults to 100.
	Limit *int `pulumi:"limit"`
	// Offset is the offset from the first result to be returned. Defaults to 0.
	Offset *int `pulumi:"offset"`
}

// A collection of values returned by getExternalLinks.
type GetExternalLinksResult struct {
	// List of all external links in Wavefront. For each external link you will see a list of attributes.
	ExternalLinks []GetExternalLinksExternalLink `pulumi:"externalLinks"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Limit  *int   `pulumi:"limit"`
	Offset *int   `pulumi:"offset"`
}

func GetExternalLinksOutput(ctx *pulumi.Context, args GetExternalLinksOutputArgs, opts ...pulumi.InvokeOption) GetExternalLinksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetExternalLinksResult, error) {
			args := v.(GetExternalLinksArgs)
			r, err := GetExternalLinks(ctx, &args, opts...)
			var s GetExternalLinksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetExternalLinksResultOutput)
}

// A collection of arguments for invoking getExternalLinks.
type GetExternalLinksOutputArgs struct {
	// Limit is the maximum number of results to be returned. Defaults to 100.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// Offset is the offset from the first result to be returned. Defaults to 0.
	Offset pulumi.IntPtrInput `pulumi:"offset"`
}

func (GetExternalLinksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalLinksArgs)(nil)).Elem()
}

// A collection of values returned by getExternalLinks.
type GetExternalLinksResultOutput struct{ *pulumi.OutputState }

func (GetExternalLinksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalLinksResult)(nil)).Elem()
}

func (o GetExternalLinksResultOutput) ToGetExternalLinksResultOutput() GetExternalLinksResultOutput {
	return o
}

func (o GetExternalLinksResultOutput) ToGetExternalLinksResultOutputWithContext(ctx context.Context) GetExternalLinksResultOutput {
	return o
}

// List of all external links in Wavefront. For each external link you will see a list of attributes.
func (o GetExternalLinksResultOutput) ExternalLinks() GetExternalLinksExternalLinkArrayOutput {
	return o.ApplyT(func(v GetExternalLinksResult) []GetExternalLinksExternalLink { return v.ExternalLinks }).(GetExternalLinksExternalLinkArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetExternalLinksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExternalLinksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetExternalLinksResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetExternalLinksResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetExternalLinksResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetExternalLinksResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExternalLinksResultOutput{})
}