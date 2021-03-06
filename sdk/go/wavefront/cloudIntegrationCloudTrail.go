// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wavefront

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Wavefront Cloud Integration for CloudTrail. This allows CloudTrail cloud integrations to be created,
// updated, and deleted.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-wavefront/sdk/go/wavefront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		extId, err := wavefront.NewCloudIntegrationAwsExternalId(ctx, "extId", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = wavefront.NewCloudIntegrationCloudTrail(ctx, "cloudtrail", &wavefront.CloudIntegrationCloudTrailArgs{
// 			RoleArn:    pulumi.String("arn:aws::1234567:role/example-arn"),
// 			ExternalId: extId.ID(),
// 			Region:     pulumi.String("us-west-2"),
// 			BucketName: pulumi.String("example-s3-bucket"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// CloudTrail Cloud Integrations can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import wavefront:index/cloudIntegrationCloudTrail:CloudIntegrationCloudTrail cloudtrail a411c16b-3cf7-4f03-bf11-8ca05aab898d
// ```
type CloudIntegrationCloudTrail struct {
	pulumi.CustomResourceState

	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags pulumi.StringMapOutput `pulumi:"additionalTags"`
	// Name of the S3 bucket where CloudTrail logs are stored
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Rule to filter CloudTrail log event
	FilterRule pulumi.StringPtrOutput `pulumi:"filterRule"`
	// Forces this resource to save, even if errors are present
	ForceSave pulumi.BoolPtrOutput `pulumi:"forceSave"`
	// The human-readable name of this integration
	Name pulumi.StringOutput `pulumi:"name"`
	// The common prefix, if any, appended to all CloudTrail log files.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The AWS region of the S3 bucket where CloudTrail logs are stored
	Region pulumi.StringOutput `pulumi:"region"`
	// The external id corresponding to the Role ARN
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A value denoting which cloud service this service integrates with
	Service pulumi.StringOutput `pulumi:"service"`
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes pulumi.IntPtrOutput `pulumi:"serviceRefreshRateInMinutes"`
}

// NewCloudIntegrationCloudTrail registers a new resource with the given unique name, arguments, and options.
func NewCloudIntegrationCloudTrail(ctx *pulumi.Context,
	name string, args *CloudIntegrationCloudTrailArgs, opts ...pulumi.ResourceOption) (*CloudIntegrationCloudTrail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.ExternalId == nil {
		return nil, errors.New("invalid value for required argument 'ExternalId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource CloudIntegrationCloudTrail
	err := ctx.RegisterResource("wavefront:index/cloudIntegrationCloudTrail:CloudIntegrationCloudTrail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudIntegrationCloudTrail gets an existing CloudIntegrationCloudTrail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudIntegrationCloudTrail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudIntegrationCloudTrailState, opts ...pulumi.ResourceOption) (*CloudIntegrationCloudTrail, error) {
	var resource CloudIntegrationCloudTrail
	err := ctx.ReadResource("wavefront:index/cloudIntegrationCloudTrail:CloudIntegrationCloudTrail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudIntegrationCloudTrail resources.
type cloudIntegrationCloudTrailState struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// Name of the S3 bucket where CloudTrail logs are stored
	BucketName *string `pulumi:"bucketName"`
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront
	ExternalId *string `pulumi:"externalId"`
	// Rule to filter CloudTrail log event
	FilterRule *string `pulumi:"filterRule"`
	// Forces this resource to save, even if errors are present
	ForceSave *bool `pulumi:"forceSave"`
	// The human-readable name of this integration
	Name *string `pulumi:"name"`
	// The common prefix, if any, appended to all CloudTrail log files.
	Prefix *string `pulumi:"prefix"`
	// The AWS region of the S3 bucket where CloudTrail logs are stored
	Region *string `pulumi:"region"`
	// The external id corresponding to the Role ARN
	RoleArn *string `pulumi:"roleArn"`
	// A value denoting which cloud service this service integrates with
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

type CloudIntegrationCloudTrailState struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags pulumi.StringMapInput
	// Name of the S3 bucket where CloudTrail logs are stored
	BucketName pulumi.StringPtrInput
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront
	ExternalId pulumi.StringPtrInput
	// Rule to filter CloudTrail log event
	FilterRule pulumi.StringPtrInput
	// Forces this resource to save, even if errors are present
	ForceSave pulumi.BoolPtrInput
	// The human-readable name of this integration
	Name pulumi.StringPtrInput
	// The common prefix, if any, appended to all CloudTrail log files.
	Prefix pulumi.StringPtrInput
	// The AWS region of the S3 bucket where CloudTrail logs are stored
	Region pulumi.StringPtrInput
	// The external id corresponding to the Role ARN
	RoleArn pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationCloudTrailState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationCloudTrailState)(nil)).Elem()
}

type cloudIntegrationCloudTrailArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// Name of the S3 bucket where CloudTrail logs are stored
	BucketName string `pulumi:"bucketName"`
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront
	ExternalId string `pulumi:"externalId"`
	// Rule to filter CloudTrail log event
	FilterRule *string `pulumi:"filterRule"`
	// Forces this resource to save, even if errors are present
	ForceSave *bool `pulumi:"forceSave"`
	// The human-readable name of this integration
	Name *string `pulumi:"name"`
	// The common prefix, if any, appended to all CloudTrail log files.
	Prefix *string `pulumi:"prefix"`
	// The AWS region of the S3 bucket where CloudTrail logs are stored
	Region string `pulumi:"region"`
	// The external id corresponding to the Role ARN
	RoleArn string `pulumi:"roleArn"`
	// A value denoting which cloud service this service integrates with
	Service string `pulumi:"service"`
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

// The set of arguments for constructing a CloudIntegrationCloudTrail resource.
type CloudIntegrationCloudTrailArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags pulumi.StringMapInput
	// Name of the S3 bucket where CloudTrail logs are stored
	BucketName pulumi.StringInput
	// The Role ARN that the customer has created in AWS IAM to allow access to Wavefront
	ExternalId pulumi.StringInput
	// Rule to filter CloudTrail log event
	FilterRule pulumi.StringPtrInput
	// Forces this resource to save, even if errors are present
	ForceSave pulumi.BoolPtrInput
	// The human-readable name of this integration
	Name pulumi.StringPtrInput
	// The common prefix, if any, appended to all CloudTrail log files.
	Prefix pulumi.StringPtrInput
	// The AWS region of the S3 bucket where CloudTrail logs are stored
	Region pulumi.StringInput
	// The external id corresponding to the Role ARN
	RoleArn pulumi.StringInput
	// A value denoting which cloud service this service integrates with
	Service pulumi.StringInput
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationCloudTrailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationCloudTrailArgs)(nil)).Elem()
}

type CloudIntegrationCloudTrailInput interface {
	pulumi.Input

	ToCloudIntegrationCloudTrailOutput() CloudIntegrationCloudTrailOutput
	ToCloudIntegrationCloudTrailOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailOutput
}

func (*CloudIntegrationCloudTrail) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudIntegrationCloudTrail)(nil))
}

func (i *CloudIntegrationCloudTrail) ToCloudIntegrationCloudTrailOutput() CloudIntegrationCloudTrailOutput {
	return i.ToCloudIntegrationCloudTrailOutputWithContext(context.Background())
}

func (i *CloudIntegrationCloudTrail) ToCloudIntegrationCloudTrailOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudTrailOutput)
}

func (i *CloudIntegrationCloudTrail) ToCloudIntegrationCloudTrailPtrOutput() CloudIntegrationCloudTrailPtrOutput {
	return i.ToCloudIntegrationCloudTrailPtrOutputWithContext(context.Background())
}

func (i *CloudIntegrationCloudTrail) ToCloudIntegrationCloudTrailPtrOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudTrailPtrOutput)
}

type CloudIntegrationCloudTrailPtrInput interface {
	pulumi.Input

	ToCloudIntegrationCloudTrailPtrOutput() CloudIntegrationCloudTrailPtrOutput
	ToCloudIntegrationCloudTrailPtrOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailPtrOutput
}

type cloudIntegrationCloudTrailPtrType CloudIntegrationCloudTrailArgs

func (*cloudIntegrationCloudTrailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationCloudTrail)(nil))
}

func (i *cloudIntegrationCloudTrailPtrType) ToCloudIntegrationCloudTrailPtrOutput() CloudIntegrationCloudTrailPtrOutput {
	return i.ToCloudIntegrationCloudTrailPtrOutputWithContext(context.Background())
}

func (i *cloudIntegrationCloudTrailPtrType) ToCloudIntegrationCloudTrailPtrOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudTrailPtrOutput)
}

// CloudIntegrationCloudTrailArrayInput is an input type that accepts CloudIntegrationCloudTrailArray and CloudIntegrationCloudTrailArrayOutput values.
// You can construct a concrete instance of `CloudIntegrationCloudTrailArrayInput` via:
//
//          CloudIntegrationCloudTrailArray{ CloudIntegrationCloudTrailArgs{...} }
type CloudIntegrationCloudTrailArrayInput interface {
	pulumi.Input

	ToCloudIntegrationCloudTrailArrayOutput() CloudIntegrationCloudTrailArrayOutput
	ToCloudIntegrationCloudTrailArrayOutputWithContext(context.Context) CloudIntegrationCloudTrailArrayOutput
}

type CloudIntegrationCloudTrailArray []CloudIntegrationCloudTrailInput

func (CloudIntegrationCloudTrailArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CloudIntegrationCloudTrail)(nil))
}

func (i CloudIntegrationCloudTrailArray) ToCloudIntegrationCloudTrailArrayOutput() CloudIntegrationCloudTrailArrayOutput {
	return i.ToCloudIntegrationCloudTrailArrayOutputWithContext(context.Background())
}

func (i CloudIntegrationCloudTrailArray) ToCloudIntegrationCloudTrailArrayOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudTrailArrayOutput)
}

// CloudIntegrationCloudTrailMapInput is an input type that accepts CloudIntegrationCloudTrailMap and CloudIntegrationCloudTrailMapOutput values.
// You can construct a concrete instance of `CloudIntegrationCloudTrailMapInput` via:
//
//          CloudIntegrationCloudTrailMap{ "key": CloudIntegrationCloudTrailArgs{...} }
type CloudIntegrationCloudTrailMapInput interface {
	pulumi.Input

	ToCloudIntegrationCloudTrailMapOutput() CloudIntegrationCloudTrailMapOutput
	ToCloudIntegrationCloudTrailMapOutputWithContext(context.Context) CloudIntegrationCloudTrailMapOutput
}

type CloudIntegrationCloudTrailMap map[string]CloudIntegrationCloudTrailInput

func (CloudIntegrationCloudTrailMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CloudIntegrationCloudTrail)(nil))
}

func (i CloudIntegrationCloudTrailMap) ToCloudIntegrationCloudTrailMapOutput() CloudIntegrationCloudTrailMapOutput {
	return i.ToCloudIntegrationCloudTrailMapOutputWithContext(context.Background())
}

func (i CloudIntegrationCloudTrailMap) ToCloudIntegrationCloudTrailMapOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationCloudTrailMapOutput)
}

type CloudIntegrationCloudTrailOutput struct {
	*pulumi.OutputState
}

func (CloudIntegrationCloudTrailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudIntegrationCloudTrail)(nil))
}

func (o CloudIntegrationCloudTrailOutput) ToCloudIntegrationCloudTrailOutput() CloudIntegrationCloudTrailOutput {
	return o
}

func (o CloudIntegrationCloudTrailOutput) ToCloudIntegrationCloudTrailOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailOutput {
	return o
}

func (o CloudIntegrationCloudTrailOutput) ToCloudIntegrationCloudTrailPtrOutput() CloudIntegrationCloudTrailPtrOutput {
	return o.ToCloudIntegrationCloudTrailPtrOutputWithContext(context.Background())
}

func (o CloudIntegrationCloudTrailOutput) ToCloudIntegrationCloudTrailPtrOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailPtrOutput {
	return o.ApplyT(func(v CloudIntegrationCloudTrail) *CloudIntegrationCloudTrail {
		return &v
	}).(CloudIntegrationCloudTrailPtrOutput)
}

type CloudIntegrationCloudTrailPtrOutput struct {
	*pulumi.OutputState
}

func (CloudIntegrationCloudTrailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationCloudTrail)(nil))
}

func (o CloudIntegrationCloudTrailPtrOutput) ToCloudIntegrationCloudTrailPtrOutput() CloudIntegrationCloudTrailPtrOutput {
	return o
}

func (o CloudIntegrationCloudTrailPtrOutput) ToCloudIntegrationCloudTrailPtrOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailPtrOutput {
	return o
}

type CloudIntegrationCloudTrailArrayOutput struct{ *pulumi.OutputState }

func (CloudIntegrationCloudTrailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudIntegrationCloudTrail)(nil))
}

func (o CloudIntegrationCloudTrailArrayOutput) ToCloudIntegrationCloudTrailArrayOutput() CloudIntegrationCloudTrailArrayOutput {
	return o
}

func (o CloudIntegrationCloudTrailArrayOutput) ToCloudIntegrationCloudTrailArrayOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailArrayOutput {
	return o
}

func (o CloudIntegrationCloudTrailArrayOutput) Index(i pulumi.IntInput) CloudIntegrationCloudTrailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudIntegrationCloudTrail {
		return vs[0].([]CloudIntegrationCloudTrail)[vs[1].(int)]
	}).(CloudIntegrationCloudTrailOutput)
}

type CloudIntegrationCloudTrailMapOutput struct{ *pulumi.OutputState }

func (CloudIntegrationCloudTrailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CloudIntegrationCloudTrail)(nil))
}

func (o CloudIntegrationCloudTrailMapOutput) ToCloudIntegrationCloudTrailMapOutput() CloudIntegrationCloudTrailMapOutput {
	return o
}

func (o CloudIntegrationCloudTrailMapOutput) ToCloudIntegrationCloudTrailMapOutputWithContext(ctx context.Context) CloudIntegrationCloudTrailMapOutput {
	return o
}

func (o CloudIntegrationCloudTrailMapOutput) MapIndex(k pulumi.StringInput) CloudIntegrationCloudTrailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CloudIntegrationCloudTrail {
		return vs[0].(map[string]CloudIntegrationCloudTrail)[vs[1].(string)]
	}).(CloudIntegrationCloudTrailOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudIntegrationCloudTrailOutput{})
	pulumi.RegisterOutputType(CloudIntegrationCloudTrailPtrOutput{})
	pulumi.RegisterOutputType(CloudIntegrationCloudTrailArrayOutput{})
	pulumi.RegisterOutputType(CloudIntegrationCloudTrailMapOutput{})
}
