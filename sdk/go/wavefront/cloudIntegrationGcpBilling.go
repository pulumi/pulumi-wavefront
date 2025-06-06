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

// Provides a Wavefront Cloud Integration for Google Cloud Billing. This allows GCP Billing cloud integrations to be created,
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
//			_, err := wavefront.NewCloudIntegrationGcpBilling(ctx, "gcp_billing", &wavefront.CloudIntegrationGcpBillingArgs{
//				Name:      pulumi.String("Test Integration"),
//				ProjectId: pulumi.String("example-gcp-project"),
//				ApiKey:    pulumi.String("example-api-key"),
//				JsonKey:   pulumi.String("{...your gcp key ...}\n"),
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
// GCP Billing Cloud Integrations can be imported by using the `id`, e.g.:
//
// ```sh
// $ pulumi import wavefront:index/cloudIntegrationGcpBilling:CloudIntegrationGcpBilling gcp_billing a411c16b-3cf7-4f03-bf11-8ca05aab898d
// ```
type CloudIntegrationGcpBilling struct {
	pulumi.CustomResourceState

	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapOutput `pulumi:"additionalTags"`
	// API key for Google Cloud Platform (GCP).
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrOutput `pulumi:"forceSave"`
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey pulumi.StringOutput `pulumi:"jsonKey"`
	// The human-readable name of this integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Google Cloud Platform (GCP) Project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringOutput `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrOutput `pulumi:"serviceRefreshRateInMinutes"`
}

// NewCloudIntegrationGcpBilling registers a new resource with the given unique name, arguments, and options.
func NewCloudIntegrationGcpBilling(ctx *pulumi.Context,
	name string, args *CloudIntegrationGcpBillingArgs, opts ...pulumi.ResourceOption) (*CloudIntegrationGcpBilling, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.JsonKey == nil {
		return nil, errors.New("invalid value for required argument 'JsonKey'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringInput)
	}
	if args.JsonKey != nil {
		args.JsonKey = pulumi.ToSecret(args.JsonKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
		"jsonKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudIntegrationGcpBilling
	err := ctx.RegisterResource("wavefront:index/cloudIntegrationGcpBilling:CloudIntegrationGcpBilling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudIntegrationGcpBilling gets an existing CloudIntegrationGcpBilling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudIntegrationGcpBilling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudIntegrationGcpBillingState, opts ...pulumi.ResourceOption) (*CloudIntegrationGcpBilling, error) {
	var resource CloudIntegrationGcpBilling
	err := ctx.ReadResource("wavefront:index/cloudIntegrationGcpBilling:CloudIntegrationGcpBilling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudIntegrationGcpBilling resources.
type cloudIntegrationGcpBillingState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// API key for Google Cloud Platform (GCP).
	ApiKey *string `pulumi:"apiKey"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey *string `pulumi:"jsonKey"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// The Google Cloud Platform (GCP) Project ID.
	ProjectId *string `pulumi:"projectId"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

type CloudIntegrationGcpBillingState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// API key for Google Cloud Platform (GCP).
	ApiKey pulumi.StringPtrInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey pulumi.StringPtrInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// The Google Cloud Platform (GCP) Project ID.
	ProjectId pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationGcpBillingState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationGcpBillingState)(nil)).Elem()
}

type cloudIntegrationGcpBillingArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// API key for Google Cloud Platform (GCP).
	ApiKey string `pulumi:"apiKey"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey string `pulumi:"jsonKey"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// The Google Cloud Platform (GCP) Project ID.
	ProjectId string `pulumi:"projectId"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

// The set of arguments for constructing a CloudIntegrationGcpBilling resource.
type CloudIntegrationGcpBillingArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// API key for Google Cloud Platform (GCP).
	ApiKey pulumi.StringInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey pulumi.StringInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// The Google Cloud Platform (GCP) Project ID.
	ProjectId pulumi.StringInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationGcpBillingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationGcpBillingArgs)(nil)).Elem()
}

type CloudIntegrationGcpBillingInput interface {
	pulumi.Input

	ToCloudIntegrationGcpBillingOutput() CloudIntegrationGcpBillingOutput
	ToCloudIntegrationGcpBillingOutputWithContext(ctx context.Context) CloudIntegrationGcpBillingOutput
}

func (*CloudIntegrationGcpBilling) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationGcpBilling)(nil)).Elem()
}

func (i *CloudIntegrationGcpBilling) ToCloudIntegrationGcpBillingOutput() CloudIntegrationGcpBillingOutput {
	return i.ToCloudIntegrationGcpBillingOutputWithContext(context.Background())
}

func (i *CloudIntegrationGcpBilling) ToCloudIntegrationGcpBillingOutputWithContext(ctx context.Context) CloudIntegrationGcpBillingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpBillingOutput)
}

// CloudIntegrationGcpBillingArrayInput is an input type that accepts CloudIntegrationGcpBillingArray and CloudIntegrationGcpBillingArrayOutput values.
// You can construct a concrete instance of `CloudIntegrationGcpBillingArrayInput` via:
//
//	CloudIntegrationGcpBillingArray{ CloudIntegrationGcpBillingArgs{...} }
type CloudIntegrationGcpBillingArrayInput interface {
	pulumi.Input

	ToCloudIntegrationGcpBillingArrayOutput() CloudIntegrationGcpBillingArrayOutput
	ToCloudIntegrationGcpBillingArrayOutputWithContext(context.Context) CloudIntegrationGcpBillingArrayOutput
}

type CloudIntegrationGcpBillingArray []CloudIntegrationGcpBillingInput

func (CloudIntegrationGcpBillingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationGcpBilling)(nil)).Elem()
}

func (i CloudIntegrationGcpBillingArray) ToCloudIntegrationGcpBillingArrayOutput() CloudIntegrationGcpBillingArrayOutput {
	return i.ToCloudIntegrationGcpBillingArrayOutputWithContext(context.Background())
}

func (i CloudIntegrationGcpBillingArray) ToCloudIntegrationGcpBillingArrayOutputWithContext(ctx context.Context) CloudIntegrationGcpBillingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpBillingArrayOutput)
}

// CloudIntegrationGcpBillingMapInput is an input type that accepts CloudIntegrationGcpBillingMap and CloudIntegrationGcpBillingMapOutput values.
// You can construct a concrete instance of `CloudIntegrationGcpBillingMapInput` via:
//
//	CloudIntegrationGcpBillingMap{ "key": CloudIntegrationGcpBillingArgs{...} }
type CloudIntegrationGcpBillingMapInput interface {
	pulumi.Input

	ToCloudIntegrationGcpBillingMapOutput() CloudIntegrationGcpBillingMapOutput
	ToCloudIntegrationGcpBillingMapOutputWithContext(context.Context) CloudIntegrationGcpBillingMapOutput
}

type CloudIntegrationGcpBillingMap map[string]CloudIntegrationGcpBillingInput

func (CloudIntegrationGcpBillingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationGcpBilling)(nil)).Elem()
}

func (i CloudIntegrationGcpBillingMap) ToCloudIntegrationGcpBillingMapOutput() CloudIntegrationGcpBillingMapOutput {
	return i.ToCloudIntegrationGcpBillingMapOutputWithContext(context.Background())
}

func (i CloudIntegrationGcpBillingMap) ToCloudIntegrationGcpBillingMapOutputWithContext(ctx context.Context) CloudIntegrationGcpBillingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpBillingMapOutput)
}

type CloudIntegrationGcpBillingOutput struct{ *pulumi.OutputState }

func (CloudIntegrationGcpBillingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationGcpBilling)(nil)).Elem()
}

func (o CloudIntegrationGcpBillingOutput) ToCloudIntegrationGcpBillingOutput() CloudIntegrationGcpBillingOutput {
	return o
}

func (o CloudIntegrationGcpBillingOutput) ToCloudIntegrationGcpBillingOutputWithContext(ctx context.Context) CloudIntegrationGcpBillingOutput {
	return o
}

// A list of point tag key-values to add to every point ingested using this integration.
func (o CloudIntegrationGcpBillingOutput) AdditionalTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.StringMapOutput { return v.AdditionalTags }).(pulumi.StringMapOutput)
}

// API key for Google Cloud Platform (GCP).
func (o CloudIntegrationGcpBillingOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

// Forces this resource to save, even if errors are present.
func (o CloudIntegrationGcpBillingOutput) ForceSave() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.BoolPtrOutput { return v.ForceSave }).(pulumi.BoolPtrOutput)
}

// Private key for a Google Cloud Platform (GCP) service account within your project.
// The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
func (o CloudIntegrationGcpBillingOutput) JsonKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.StringOutput { return v.JsonKey }).(pulumi.StringOutput)
}

// The human-readable name of this integration.
func (o CloudIntegrationGcpBillingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Google Cloud Platform (GCP) Project ID.
func (o CloudIntegrationGcpBillingOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// A value denoting which cloud service this service integrates with.
func (o CloudIntegrationGcpBillingOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// How often, in minutes, to refresh the service.
func (o CloudIntegrationGcpBillingOutput) ServiceRefreshRateInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationGcpBilling) pulumi.IntPtrOutput { return v.ServiceRefreshRateInMinutes }).(pulumi.IntPtrOutput)
}

type CloudIntegrationGcpBillingArrayOutput struct{ *pulumi.OutputState }

func (CloudIntegrationGcpBillingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationGcpBilling)(nil)).Elem()
}

func (o CloudIntegrationGcpBillingArrayOutput) ToCloudIntegrationGcpBillingArrayOutput() CloudIntegrationGcpBillingArrayOutput {
	return o
}

func (o CloudIntegrationGcpBillingArrayOutput) ToCloudIntegrationGcpBillingArrayOutputWithContext(ctx context.Context) CloudIntegrationGcpBillingArrayOutput {
	return o
}

func (o CloudIntegrationGcpBillingArrayOutput) Index(i pulumi.IntInput) CloudIntegrationGcpBillingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudIntegrationGcpBilling {
		return vs[0].([]*CloudIntegrationGcpBilling)[vs[1].(int)]
	}).(CloudIntegrationGcpBillingOutput)
}

type CloudIntegrationGcpBillingMapOutput struct{ *pulumi.OutputState }

func (CloudIntegrationGcpBillingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationGcpBilling)(nil)).Elem()
}

func (o CloudIntegrationGcpBillingMapOutput) ToCloudIntegrationGcpBillingMapOutput() CloudIntegrationGcpBillingMapOutput {
	return o
}

func (o CloudIntegrationGcpBillingMapOutput) ToCloudIntegrationGcpBillingMapOutputWithContext(ctx context.Context) CloudIntegrationGcpBillingMapOutput {
	return o
}

func (o CloudIntegrationGcpBillingMapOutput) MapIndex(k pulumi.StringInput) CloudIntegrationGcpBillingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudIntegrationGcpBilling {
		return vs[0].(map[string]*CloudIntegrationGcpBilling)[vs[1].(string)]
	}).(CloudIntegrationGcpBillingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationGcpBillingInput)(nil)).Elem(), &CloudIntegrationGcpBilling{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationGcpBillingArrayInput)(nil)).Elem(), CloudIntegrationGcpBillingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationGcpBillingMapInput)(nil)).Elem(), CloudIntegrationGcpBillingMap{})
	pulumi.RegisterOutputType(CloudIntegrationGcpBillingOutput{})
	pulumi.RegisterOutputType(CloudIntegrationGcpBillingArrayOutput{})
	pulumi.RegisterOutputType(CloudIntegrationGcpBillingMapOutput{})
}
