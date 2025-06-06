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

// Provides a Wavefront Cloud Integration for CloudWatch. This allows CloudWatch cloud integrations to be created,
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
//			extId, err := wavefront.NewCloudIntegrationAwsExternalId(ctx, "ext_id", nil)
//			if err != nil {
//				return err
//			}
//			_, err = wavefront.NewCloudIntegrationCloudWatch(ctx, "cloudwatch", &wavefront.CloudIntegrationCloudWatchArgs{
//				Name:       pulumi.String("Test Integration"),
//				ForceSave:  pulumi.Bool(true),
//				RoleArn:    pulumi.String("arn:aws::1234567:role/example-arn"),
//				ExternalId: extId.ID(),
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
// CloudWatch Cloud Integrations can be imported by using the `id`, e.g.:
//
// ```sh
// $ pulumi import wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch cloudwatch a411c16b-3cf7-4f03-bf11-8ca05aab898d
// ```
type CloudIntegrationCloudWatch struct {
	pulumi.CustomResourceState

	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapOutput `pulumi:"additionalTags"`
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrOutput `pulumi:"forceSave"`
	// A string->string map allow list of instance tag-value pairs (in AWS).
	// If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
	// Multiple entries are OR'ed.
	InstanceSelectionTags pulumi.StringMapOutput `pulumi:"instanceSelectionTags"`
	// A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
	MetricFilterRegex pulumi.StringPtrOutput `pulumi:"metricFilterRegex"`
	// The human-readable name of this integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of namespaces that limit what we query from CloudWatch.
	Namespaces pulumi.StringArrayOutput `pulumi:"namespaces"`
	// A regular expression that AWS tag key name must match (case-insensitively)
	// in order to be ingested.
	PointTagFilterRegex pulumi.StringPtrOutput `pulumi:"pointTagFilterRegex"`
	// The external ID corresponding to the Role ARN.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringOutput `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrOutput `pulumi:"serviceRefreshRateInMinutes"`
	// A string->string map of allow list of volume tag-value pairs (in AWS).
	// If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
	// Multiple entries are OR'ed.
	VolumeSelectionTags pulumi.StringMapOutput `pulumi:"volumeSelectionTags"`
}

// NewCloudIntegrationCloudWatch registers a new resource with the given unique name, arguments, and options.
func NewCloudIntegrationCloudWatch(ctx *pulumi.Context,
	name string, args *CloudIntegrationCloudWatchArgs, opts ...pulumi.ResourceOption) (*CloudIntegrationCloudWatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalId == nil {
		return nil, errors.New("invalid value for required argument 'ExternalId'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudIntegrationCloudWatch
	err := ctx.RegisterResource("wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudIntegrationCloudWatch gets an existing CloudIntegrationCloudWatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudIntegrationCloudWatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudIntegrationCloudWatchState, opts ...pulumi.ResourceOption) (*CloudIntegrationCloudWatch, error) {
	var resource CloudIntegrationCloudWatch
	err := ctx.ReadResource("wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudIntegrationCloudWatch resources.
type cloudIntegrationCloudWatchState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
	ExternalId *string `pulumi:"externalId"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// A string->string map allow list of instance tag-value pairs (in AWS).
	// If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
	// Multiple entries are OR'ed.
	InstanceSelectionTags map[string]string `pulumi:"instanceSelectionTags"`
	// A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// A list of namespaces that limit what we query from CloudWatch.
	Namespaces []string `pulumi:"namespaces"`
	// A regular expression that AWS tag key name must match (case-insensitively)
	// in order to be ingested.
	PointTagFilterRegex *string `pulumi:"pointTagFilterRegex"`
	// The external ID corresponding to the Role ARN.
	RoleArn *string `pulumi:"roleArn"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
	// A string->string map of allow list of volume tag-value pairs (in AWS).
	// If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
	// Multiple entries are OR'ed.
	VolumeSelectionTags map[string]string `pulumi:"volumeSelectionTags"`
}

type CloudIntegrationCloudWatchState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
	ExternalId pulumi.StringPtrInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// A string->string map allow list of instance tag-value pairs (in AWS).
	// If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
	// Multiple entries are OR'ed.
	InstanceSelectionTags pulumi.StringMapInput
	// A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
	MetricFilterRegex pulumi.StringPtrInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// A list of namespaces that limit what we query from CloudWatch.
	Namespaces pulumi.StringArrayInput
	// A regular expression that AWS tag key name must match (case-insensitively)
	// in order to be ingested.
	PointTagFilterRegex pulumi.StringPtrInput
	// The external ID corresponding to the Role ARN.
	RoleArn pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
	// A string->string map of allow list of volume tag-value pairs (in AWS).
	// If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
	// Multiple entries are OR'ed.
	VolumeSelectionTags pulumi.StringMapInput
}

func (CloudIntegrationCloudWatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationCloudWatchState)(nil)).Elem()
}

type cloudIntegrationCloudWatchArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
	ExternalId string `pulumi:"externalId"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// A string->string map allow list of instance tag-value pairs (in AWS).
	// If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
	// Multiple entries are OR'ed.
	InstanceSelectionTags map[string]string `pulumi:"instanceSelectionTags"`
	// A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// A list of namespaces that limit what we query from CloudWatch.
	Namespaces []string `pulumi:"namespaces"`
	// A regular expression that AWS tag key name must match (case-insensitively)
	// in order to be ingested.
	PointTagFilterRegex *string `pulumi:"pointTagFilterRegex"`
	// The external ID corresponding to the Role ARN.
	RoleArn string `pulumi:"roleArn"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
	// A string->string map of allow list of volume tag-value pairs (in AWS).
	// If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
	// Multiple entries are OR'ed.
	VolumeSelectionTags map[string]string `pulumi:"volumeSelectionTags"`
}

// The set of arguments for constructing a CloudIntegrationCloudWatch resource.
type CloudIntegrationCloudWatchArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
	ExternalId pulumi.StringInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// A string->string map allow list of instance tag-value pairs (in AWS).
	// If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
	// Multiple entries are OR'ed.
	InstanceSelectionTags pulumi.StringMapInput
	// A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
	MetricFilterRegex pulumi.StringPtrInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// A list of namespaces that limit what we query from CloudWatch.
	Namespaces pulumi.StringArrayInput
	// A regular expression that AWS tag key name must match (case-insensitively)
	// in order to be ingested.
	PointTagFilterRegex pulumi.StringPtrInput
	// The external ID corresponding to the Role ARN.
	RoleArn pulumi.StringInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
	// A string->string map of allow list of volume tag-value pairs (in AWS).
	// If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
	// Multiple entries are OR'ed.
	VolumeSelectionTags pulumi.StringMapInput
}

func (CloudIntegrationCloudWatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationCloudWatchArgs)(nil)).Elem()
}

type CloudIntegrationCloudWatchInput interface {
	pulumi.Input

	ToCloudIntegrationCloudWatchOutput() CloudIntegrationCloudWatchOutput
	ToCloudIntegrationCloudWatchOutputWithContext(ctx context.Context) CloudIntegrationCloudWatchOutput
}

func (*CloudIntegrationCloudWatch) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationCloudWatch)(nil)).Elem()
}

func (i *CloudIntegrationCloudWatch) ToCloudIntegrationCloudWatchOutput() CloudIntegrationCloudWatchOutput {
	return i.ToCloudIntegrationCloudWatchOutputWithContext(context.Background())
}

func (i *CloudIntegrationCloudWatch) ToCloudIntegrationCloudWatchOutputWithContext(ctx context.Context) CloudIntegrationCloudWatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudWatchOutput)
}

// CloudIntegrationCloudWatchArrayInput is an input type that accepts CloudIntegrationCloudWatchArray and CloudIntegrationCloudWatchArrayOutput values.
// You can construct a concrete instance of `CloudIntegrationCloudWatchArrayInput` via:
//
//	CloudIntegrationCloudWatchArray{ CloudIntegrationCloudWatchArgs{...} }
type CloudIntegrationCloudWatchArrayInput interface {
	pulumi.Input

	ToCloudIntegrationCloudWatchArrayOutput() CloudIntegrationCloudWatchArrayOutput
	ToCloudIntegrationCloudWatchArrayOutputWithContext(context.Context) CloudIntegrationCloudWatchArrayOutput
}

type CloudIntegrationCloudWatchArray []CloudIntegrationCloudWatchInput

func (CloudIntegrationCloudWatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationCloudWatch)(nil)).Elem()
}

func (i CloudIntegrationCloudWatchArray) ToCloudIntegrationCloudWatchArrayOutput() CloudIntegrationCloudWatchArrayOutput {
	return i.ToCloudIntegrationCloudWatchArrayOutputWithContext(context.Background())
}

func (i CloudIntegrationCloudWatchArray) ToCloudIntegrationCloudWatchArrayOutputWithContext(ctx context.Context) CloudIntegrationCloudWatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudWatchArrayOutput)
}

// CloudIntegrationCloudWatchMapInput is an input type that accepts CloudIntegrationCloudWatchMap and CloudIntegrationCloudWatchMapOutput values.
// You can construct a concrete instance of `CloudIntegrationCloudWatchMapInput` via:
//
//	CloudIntegrationCloudWatchMap{ "key": CloudIntegrationCloudWatchArgs{...} }
type CloudIntegrationCloudWatchMapInput interface {
	pulumi.Input

	ToCloudIntegrationCloudWatchMapOutput() CloudIntegrationCloudWatchMapOutput
	ToCloudIntegrationCloudWatchMapOutputWithContext(context.Context) CloudIntegrationCloudWatchMapOutput
}

type CloudIntegrationCloudWatchMap map[string]CloudIntegrationCloudWatchInput

func (CloudIntegrationCloudWatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationCloudWatch)(nil)).Elem()
}

func (i CloudIntegrationCloudWatchMap) ToCloudIntegrationCloudWatchMapOutput() CloudIntegrationCloudWatchMapOutput {
	return i.ToCloudIntegrationCloudWatchMapOutputWithContext(context.Background())
}

func (i CloudIntegrationCloudWatchMap) ToCloudIntegrationCloudWatchMapOutputWithContext(ctx context.Context) CloudIntegrationCloudWatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudWatchMapOutput)
}

type CloudIntegrationCloudWatchOutput struct{ *pulumi.OutputState }

func (CloudIntegrationCloudWatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationCloudWatch)(nil)).Elem()
}

func (o CloudIntegrationCloudWatchOutput) ToCloudIntegrationCloudWatchOutput() CloudIntegrationCloudWatchOutput {
	return o
}

func (o CloudIntegrationCloudWatchOutput) ToCloudIntegrationCloudWatchOutputWithContext(ctx context.Context) CloudIntegrationCloudWatchOutput {
	return o
}

// A list of point tag key-values to add to every point ingested using this integration.
func (o CloudIntegrationCloudWatchOutput) AdditionalTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringMapOutput { return v.AdditionalTags }).(pulumi.StringMapOutput)
}

// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
func (o CloudIntegrationCloudWatchOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// Forces this resource to save, even if errors are present.
func (o CloudIntegrationCloudWatchOutput) ForceSave() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.BoolPtrOutput { return v.ForceSave }).(pulumi.BoolPtrOutput)
}

// A string->string map allow list of instance tag-value pairs (in AWS).
// If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
// Multiple entries are OR'ed.
func (o CloudIntegrationCloudWatchOutput) InstanceSelectionTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringMapOutput { return v.InstanceSelectionTags }).(pulumi.StringMapOutput)
}

// A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
func (o CloudIntegrationCloudWatchOutput) MetricFilterRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringPtrOutput { return v.MetricFilterRegex }).(pulumi.StringPtrOutput)
}

// The human-readable name of this integration.
func (o CloudIntegrationCloudWatchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of namespaces that limit what we query from CloudWatch.
func (o CloudIntegrationCloudWatchOutput) Namespaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringArrayOutput { return v.Namespaces }).(pulumi.StringArrayOutput)
}

// A regular expression that AWS tag key name must match (case-insensitively)
// in order to be ingested.
func (o CloudIntegrationCloudWatchOutput) PointTagFilterRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringPtrOutput { return v.PointTagFilterRegex }).(pulumi.StringPtrOutput)
}

// The external ID corresponding to the Role ARN.
func (o CloudIntegrationCloudWatchOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// A value denoting which cloud service this service integrates with.
func (o CloudIntegrationCloudWatchOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// How often, in minutes, to refresh the service.
func (o CloudIntegrationCloudWatchOutput) ServiceRefreshRateInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.IntPtrOutput { return v.ServiceRefreshRateInMinutes }).(pulumi.IntPtrOutput)
}

// A string->string map of allow list of volume tag-value pairs (in AWS).
// If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
// Multiple entries are OR'ed.
func (o CloudIntegrationCloudWatchOutput) VolumeSelectionTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudIntegrationCloudWatch) pulumi.StringMapOutput { return v.VolumeSelectionTags }).(pulumi.StringMapOutput)
}

type CloudIntegrationCloudWatchArrayOutput struct{ *pulumi.OutputState }

func (CloudIntegrationCloudWatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationCloudWatch)(nil)).Elem()
}

func (o CloudIntegrationCloudWatchArrayOutput) ToCloudIntegrationCloudWatchArrayOutput() CloudIntegrationCloudWatchArrayOutput {
	return o
}

func (o CloudIntegrationCloudWatchArrayOutput) ToCloudIntegrationCloudWatchArrayOutputWithContext(ctx context.Context) CloudIntegrationCloudWatchArrayOutput {
	return o
}

func (o CloudIntegrationCloudWatchArrayOutput) Index(i pulumi.IntInput) CloudIntegrationCloudWatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudIntegrationCloudWatch {
		return vs[0].([]*CloudIntegrationCloudWatch)[vs[1].(int)]
	}).(CloudIntegrationCloudWatchOutput)
}

type CloudIntegrationCloudWatchMapOutput struct{ *pulumi.OutputState }

func (CloudIntegrationCloudWatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationCloudWatch)(nil)).Elem()
}

func (o CloudIntegrationCloudWatchMapOutput) ToCloudIntegrationCloudWatchMapOutput() CloudIntegrationCloudWatchMapOutput {
	return o
}

func (o CloudIntegrationCloudWatchMapOutput) ToCloudIntegrationCloudWatchMapOutputWithContext(ctx context.Context) CloudIntegrationCloudWatchMapOutput {
	return o
}

func (o CloudIntegrationCloudWatchMapOutput) MapIndex(k pulumi.StringInput) CloudIntegrationCloudWatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudIntegrationCloudWatch {
		return vs[0].(map[string]*CloudIntegrationCloudWatch)[vs[1].(string)]
	}).(CloudIntegrationCloudWatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationCloudWatchInput)(nil)).Elem(), &CloudIntegrationCloudWatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationCloudWatchArrayInput)(nil)).Elem(), CloudIntegrationCloudWatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationCloudWatchMapInput)(nil)).Elem(), CloudIntegrationCloudWatchMap{})
	pulumi.RegisterOutputType(CloudIntegrationCloudWatchOutput{})
	pulumi.RegisterOutputType(CloudIntegrationCloudWatchArrayOutput{})
	pulumi.RegisterOutputType(CloudIntegrationCloudWatchMapOutput{})
}
