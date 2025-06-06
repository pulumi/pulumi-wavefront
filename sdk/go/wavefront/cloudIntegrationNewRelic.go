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

// Provides a Wavefront Cloud Integration for New Relic. This allows New Relic cloud integrations to be created,
// updated, and deleted.
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
//			_, err := wavefront.NewCloudIntegrationNewRelic(ctx, "newrelic", &wavefront.CloudIntegrationNewRelicArgs{
//				Name:   pulumi.String("Test Integration"),
//				ApiKey: pulumi.String("example-api-key"),
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
// NewRelic Integrations can be imported by using the `id`, e.g.:
//
// ```sh
// $ pulumi import wavefront:index/cloudIntegrationNewRelic:CloudIntegrationNewRelic newrelic a411c16b-3cf7-4f03-bf11-8ca05aab898d
// ```
type CloudIntegrationNewRelic struct {
	pulumi.CustomResourceState

	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapOutput `pulumi:"additionalTags"`
	// New Relic REST API key.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// A regular expression that an application name must match (case-insensitively) in order to collect metrics.
	AppFilterRegex pulumi.StringPtrOutput `pulumi:"appFilterRegex"`
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrOutput `pulumi:"forceSave"`
	// A regular expression that a host name must match (case-insensitively) in order to collect metrics.
	HostFilterRegex pulumi.StringPtrOutput `pulumi:"hostFilterRegex"`
	// See Metric Filter.
	MetricFilters CloudIntegrationNewRelicMetricFilterArrayOutput `pulumi:"metricFilters"`
	// The human-readable name of this integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringOutput `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrOutput `pulumi:"serviceRefreshRateInMinutes"`
}

// NewCloudIntegrationNewRelic registers a new resource with the given unique name, arguments, and options.
func NewCloudIntegrationNewRelic(ctx *pulumi.Context,
	name string, args *CloudIntegrationNewRelicArgs, opts ...pulumi.ResourceOption) (*CloudIntegrationNewRelic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudIntegrationNewRelic
	err := ctx.RegisterResource("wavefront:index/cloudIntegrationNewRelic:CloudIntegrationNewRelic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudIntegrationNewRelic gets an existing CloudIntegrationNewRelic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudIntegrationNewRelic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudIntegrationNewRelicState, opts ...pulumi.ResourceOption) (*CloudIntegrationNewRelic, error) {
	var resource CloudIntegrationNewRelic
	err := ctx.ReadResource("wavefront:index/cloudIntegrationNewRelic:CloudIntegrationNewRelic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudIntegrationNewRelic resources.
type cloudIntegrationNewRelicState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// New Relic REST API key.
	ApiKey *string `pulumi:"apiKey"`
	// A regular expression that an application name must match (case-insensitively) in order to collect metrics.
	AppFilterRegex *string `pulumi:"appFilterRegex"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// A regular expression that a host name must match (case-insensitively) in order to collect metrics.
	HostFilterRegex *string `pulumi:"hostFilterRegex"`
	// See Metric Filter.
	MetricFilters []CloudIntegrationNewRelicMetricFilter `pulumi:"metricFilters"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

type CloudIntegrationNewRelicState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// New Relic REST API key.
	ApiKey pulumi.StringPtrInput
	// A regular expression that an application name must match (case-insensitively) in order to collect metrics.
	AppFilterRegex pulumi.StringPtrInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// A regular expression that a host name must match (case-insensitively) in order to collect metrics.
	HostFilterRegex pulumi.StringPtrInput
	// See Metric Filter.
	MetricFilters CloudIntegrationNewRelicMetricFilterArrayInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationNewRelicState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationNewRelicState)(nil)).Elem()
}

type cloudIntegrationNewRelicArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// New Relic REST API key.
	ApiKey string `pulumi:"apiKey"`
	// A regular expression that an application name must match (case-insensitively) in order to collect metrics.
	AppFilterRegex *string `pulumi:"appFilterRegex"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// A regular expression that a host name must match (case-insensitively) in order to collect metrics.
	HostFilterRegex *string `pulumi:"hostFilterRegex"`
	// See Metric Filter.
	MetricFilters []CloudIntegrationNewRelicMetricFilter `pulumi:"metricFilters"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

// The set of arguments for constructing a CloudIntegrationNewRelic resource.
type CloudIntegrationNewRelicArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// New Relic REST API key.
	ApiKey pulumi.StringInput
	// A regular expression that an application name must match (case-insensitively) in order to collect metrics.
	AppFilterRegex pulumi.StringPtrInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// A regular expression that a host name must match (case-insensitively) in order to collect metrics.
	HostFilterRegex pulumi.StringPtrInput
	// See Metric Filter.
	MetricFilters CloudIntegrationNewRelicMetricFilterArrayInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationNewRelicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationNewRelicArgs)(nil)).Elem()
}

type CloudIntegrationNewRelicInput interface {
	pulumi.Input

	ToCloudIntegrationNewRelicOutput() CloudIntegrationNewRelicOutput
	ToCloudIntegrationNewRelicOutputWithContext(ctx context.Context) CloudIntegrationNewRelicOutput
}

func (*CloudIntegrationNewRelic) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationNewRelic)(nil)).Elem()
}

func (i *CloudIntegrationNewRelic) ToCloudIntegrationNewRelicOutput() CloudIntegrationNewRelicOutput {
	return i.ToCloudIntegrationNewRelicOutputWithContext(context.Background())
}

func (i *CloudIntegrationNewRelic) ToCloudIntegrationNewRelicOutputWithContext(ctx context.Context) CloudIntegrationNewRelicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationNewRelicOutput)
}

// CloudIntegrationNewRelicArrayInput is an input type that accepts CloudIntegrationNewRelicArray and CloudIntegrationNewRelicArrayOutput values.
// You can construct a concrete instance of `CloudIntegrationNewRelicArrayInput` via:
//
//	CloudIntegrationNewRelicArray{ CloudIntegrationNewRelicArgs{...} }
type CloudIntegrationNewRelicArrayInput interface {
	pulumi.Input

	ToCloudIntegrationNewRelicArrayOutput() CloudIntegrationNewRelicArrayOutput
	ToCloudIntegrationNewRelicArrayOutputWithContext(context.Context) CloudIntegrationNewRelicArrayOutput
}

type CloudIntegrationNewRelicArray []CloudIntegrationNewRelicInput

func (CloudIntegrationNewRelicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationNewRelic)(nil)).Elem()
}

func (i CloudIntegrationNewRelicArray) ToCloudIntegrationNewRelicArrayOutput() CloudIntegrationNewRelicArrayOutput {
	return i.ToCloudIntegrationNewRelicArrayOutputWithContext(context.Background())
}

func (i CloudIntegrationNewRelicArray) ToCloudIntegrationNewRelicArrayOutputWithContext(ctx context.Context) CloudIntegrationNewRelicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationNewRelicArrayOutput)
}

// CloudIntegrationNewRelicMapInput is an input type that accepts CloudIntegrationNewRelicMap and CloudIntegrationNewRelicMapOutput values.
// You can construct a concrete instance of `CloudIntegrationNewRelicMapInput` via:
//
//	CloudIntegrationNewRelicMap{ "key": CloudIntegrationNewRelicArgs{...} }
type CloudIntegrationNewRelicMapInput interface {
	pulumi.Input

	ToCloudIntegrationNewRelicMapOutput() CloudIntegrationNewRelicMapOutput
	ToCloudIntegrationNewRelicMapOutputWithContext(context.Context) CloudIntegrationNewRelicMapOutput
}

type CloudIntegrationNewRelicMap map[string]CloudIntegrationNewRelicInput

func (CloudIntegrationNewRelicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationNewRelic)(nil)).Elem()
}

func (i CloudIntegrationNewRelicMap) ToCloudIntegrationNewRelicMapOutput() CloudIntegrationNewRelicMapOutput {
	return i.ToCloudIntegrationNewRelicMapOutputWithContext(context.Background())
}

func (i CloudIntegrationNewRelicMap) ToCloudIntegrationNewRelicMapOutputWithContext(ctx context.Context) CloudIntegrationNewRelicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationNewRelicMapOutput)
}

type CloudIntegrationNewRelicOutput struct{ *pulumi.OutputState }

func (CloudIntegrationNewRelicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationNewRelic)(nil)).Elem()
}

func (o CloudIntegrationNewRelicOutput) ToCloudIntegrationNewRelicOutput() CloudIntegrationNewRelicOutput {
	return o
}

func (o CloudIntegrationNewRelicOutput) ToCloudIntegrationNewRelicOutputWithContext(ctx context.Context) CloudIntegrationNewRelicOutput {
	return o
}

// A list of point tag key-values to add to every point ingested using this integration.
func (o CloudIntegrationNewRelicOutput) AdditionalTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.StringMapOutput { return v.AdditionalTags }).(pulumi.StringMapOutput)
}

// New Relic REST API key.
func (o CloudIntegrationNewRelicOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

// A regular expression that an application name must match (case-insensitively) in order to collect metrics.
func (o CloudIntegrationNewRelicOutput) AppFilterRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.StringPtrOutput { return v.AppFilterRegex }).(pulumi.StringPtrOutput)
}

// Forces this resource to save, even if errors are present.
func (o CloudIntegrationNewRelicOutput) ForceSave() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.BoolPtrOutput { return v.ForceSave }).(pulumi.BoolPtrOutput)
}

// A regular expression that a host name must match (case-insensitively) in order to collect metrics.
func (o CloudIntegrationNewRelicOutput) HostFilterRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.StringPtrOutput { return v.HostFilterRegex }).(pulumi.StringPtrOutput)
}

// See Metric Filter.
func (o CloudIntegrationNewRelicOutput) MetricFilters() CloudIntegrationNewRelicMetricFilterArrayOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) CloudIntegrationNewRelicMetricFilterArrayOutput {
		return v.MetricFilters
	}).(CloudIntegrationNewRelicMetricFilterArrayOutput)
}

// The human-readable name of this integration.
func (o CloudIntegrationNewRelicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A value denoting which cloud service this service integrates with.
func (o CloudIntegrationNewRelicOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// How often, in minutes, to refresh the service.
func (o CloudIntegrationNewRelicOutput) ServiceRefreshRateInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationNewRelic) pulumi.IntPtrOutput { return v.ServiceRefreshRateInMinutes }).(pulumi.IntPtrOutput)
}

type CloudIntegrationNewRelicArrayOutput struct{ *pulumi.OutputState }

func (CloudIntegrationNewRelicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationNewRelic)(nil)).Elem()
}

func (o CloudIntegrationNewRelicArrayOutput) ToCloudIntegrationNewRelicArrayOutput() CloudIntegrationNewRelicArrayOutput {
	return o
}

func (o CloudIntegrationNewRelicArrayOutput) ToCloudIntegrationNewRelicArrayOutputWithContext(ctx context.Context) CloudIntegrationNewRelicArrayOutput {
	return o
}

func (o CloudIntegrationNewRelicArrayOutput) Index(i pulumi.IntInput) CloudIntegrationNewRelicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudIntegrationNewRelic {
		return vs[0].([]*CloudIntegrationNewRelic)[vs[1].(int)]
	}).(CloudIntegrationNewRelicOutput)
}

type CloudIntegrationNewRelicMapOutput struct{ *pulumi.OutputState }

func (CloudIntegrationNewRelicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationNewRelic)(nil)).Elem()
}

func (o CloudIntegrationNewRelicMapOutput) ToCloudIntegrationNewRelicMapOutput() CloudIntegrationNewRelicMapOutput {
	return o
}

func (o CloudIntegrationNewRelicMapOutput) ToCloudIntegrationNewRelicMapOutputWithContext(ctx context.Context) CloudIntegrationNewRelicMapOutput {
	return o
}

func (o CloudIntegrationNewRelicMapOutput) MapIndex(k pulumi.StringInput) CloudIntegrationNewRelicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudIntegrationNewRelic {
		return vs[0].(map[string]*CloudIntegrationNewRelic)[vs[1].(string)]
	}).(CloudIntegrationNewRelicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationNewRelicInput)(nil)).Elem(), &CloudIntegrationNewRelic{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationNewRelicArrayInput)(nil)).Elem(), CloudIntegrationNewRelicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationNewRelicMapInput)(nil)).Elem(), CloudIntegrationNewRelicMap{})
	pulumi.RegisterOutputType(CloudIntegrationNewRelicOutput{})
	pulumi.RegisterOutputType(CloudIntegrationNewRelicArrayOutput{})
	pulumi.RegisterOutputType(CloudIntegrationNewRelicMapOutput{})
}
