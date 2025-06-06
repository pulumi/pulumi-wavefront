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

// Provides a Wavefront Cloud Integration for Azure Activity Logs. This allows Azure activity log cloud integrations to be created,
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
//			_, err := wavefront.NewCloudIntegrationAzureActivityLog(ctx, "azure_activity_log", &wavefront.CloudIntegrationAzureActivityLogArgs{
//				Name: pulumi.String("Test Integration"),
//				CategoryFilters: pulumi.StringArray{
//					pulumi.String("ADMINISTRATIVE"),
//				},
//				ClientId:     pulumi.String("client-id2"),
//				ClientSecret: pulumi.String("client-secret2"),
//				Tenant:       pulumi.String("my-tenant2"),
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
// Azure Activity Log Cloud Integrations can be imported by using the `id`, e.g.:
//
// ```sh
// $ pulumi import wavefront:index/cloudIntegrationAzureActivityLog:CloudIntegrationAzureActivityLog azure_al a411c16b-3cf7-4f03-bf11-8ca05aab898d
// ```
type CloudIntegrationAzureActivityLog struct {
	pulumi.CustomResourceState

	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapOutput `pulumi:"additionalTags"`
	// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics.
	CategoryFilters pulumi.StringArrayOutput `pulumi:"categoryFilters"`
	// Client ID for an Azure service account within your project.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client secret for an Azure service account within your project.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrOutput `pulumi:"forceSave"`
	// The human-readable name of this integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringOutput `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrOutput `pulumi:"serviceRefreshRateInMinutes"`
	// Tenant ID for an Azure service account within your project.
	Tenant pulumi.StringOutput `pulumi:"tenant"`
}

// NewCloudIntegrationAzureActivityLog registers a new resource with the given unique name, arguments, and options.
func NewCloudIntegrationAzureActivityLog(ctx *pulumi.Context,
	name string, args *CloudIntegrationAzureActivityLogArgs, opts ...pulumi.ResourceOption) (*CloudIntegrationAzureActivityLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.Tenant == nil {
		return nil, errors.New("invalid value for required argument 'Tenant'")
	}
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudIntegrationAzureActivityLog
	err := ctx.RegisterResource("wavefront:index/cloudIntegrationAzureActivityLog:CloudIntegrationAzureActivityLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudIntegrationAzureActivityLog gets an existing CloudIntegrationAzureActivityLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudIntegrationAzureActivityLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudIntegrationAzureActivityLogState, opts ...pulumi.ResourceOption) (*CloudIntegrationAzureActivityLog, error) {
	var resource CloudIntegrationAzureActivityLog
	err := ctx.ReadResource("wavefront:index/cloudIntegrationAzureActivityLog:CloudIntegrationAzureActivityLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudIntegrationAzureActivityLog resources.
type cloudIntegrationAzureActivityLogState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics.
	CategoryFilters []string `pulumi:"categoryFilters"`
	// Client ID for an Azure service account within your project.
	ClientId *string `pulumi:"clientId"`
	// Client secret for an Azure service account within your project.
	ClientSecret *string `pulumi:"clientSecret"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
	// Tenant ID for an Azure service account within your project.
	Tenant *string `pulumi:"tenant"`
}

type CloudIntegrationAzureActivityLogState struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics.
	CategoryFilters pulumi.StringArrayInput
	// Client ID for an Azure service account within your project.
	ClientId pulumi.StringPtrInput
	// Client secret for an Azure service account within your project.
	ClientSecret pulumi.StringPtrInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
	// Tenant ID for an Azure service account within your project.
	Tenant pulumi.StringPtrInput
}

func (CloudIntegrationAzureActivityLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationAzureActivityLogState)(nil)).Elem()
}

type cloudIntegrationAzureActivityLogArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics.
	CategoryFilters []string `pulumi:"categoryFilters"`
	// Client ID for an Azure service account within your project.
	ClientId string `pulumi:"clientId"`
	// Client secret for an Azure service account within your project.
	ClientSecret string `pulumi:"clientSecret"`
	// Forces this resource to save, even if errors are present.
	ForceSave *bool `pulumi:"forceSave"`
	// The human-readable name of this integration.
	Name *string `pulumi:"name"`
	// A value denoting which cloud service this service integrates with.
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
	// Tenant ID for an Azure service account within your project.
	Tenant string `pulumi:"tenant"`
}

// The set of arguments for constructing a CloudIntegrationAzureActivityLog resource.
type CloudIntegrationAzureActivityLogArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration.
	AdditionalTags pulumi.StringMapInput
	// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics.
	CategoryFilters pulumi.StringArrayInput
	// Client ID for an Azure service account within your project.
	ClientId pulumi.StringInput
	// Client secret for an Azure service account within your project.
	ClientSecret pulumi.StringInput
	// Forces this resource to save, even if errors are present.
	ForceSave pulumi.BoolPtrInput
	// The human-readable name of this integration.
	Name pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with.
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service.
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
	// Tenant ID for an Azure service account within your project.
	Tenant pulumi.StringInput
}

func (CloudIntegrationAzureActivityLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationAzureActivityLogArgs)(nil)).Elem()
}

type CloudIntegrationAzureActivityLogInput interface {
	pulumi.Input

	ToCloudIntegrationAzureActivityLogOutput() CloudIntegrationAzureActivityLogOutput
	ToCloudIntegrationAzureActivityLogOutputWithContext(ctx context.Context) CloudIntegrationAzureActivityLogOutput
}

func (*CloudIntegrationAzureActivityLog) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationAzureActivityLog)(nil)).Elem()
}

func (i *CloudIntegrationAzureActivityLog) ToCloudIntegrationAzureActivityLogOutput() CloudIntegrationAzureActivityLogOutput {
	return i.ToCloudIntegrationAzureActivityLogOutputWithContext(context.Background())
}

func (i *CloudIntegrationAzureActivityLog) ToCloudIntegrationAzureActivityLogOutputWithContext(ctx context.Context) CloudIntegrationAzureActivityLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationAzureActivityLogOutput)
}

// CloudIntegrationAzureActivityLogArrayInput is an input type that accepts CloudIntegrationAzureActivityLogArray and CloudIntegrationAzureActivityLogArrayOutput values.
// You can construct a concrete instance of `CloudIntegrationAzureActivityLogArrayInput` via:
//
//	CloudIntegrationAzureActivityLogArray{ CloudIntegrationAzureActivityLogArgs{...} }
type CloudIntegrationAzureActivityLogArrayInput interface {
	pulumi.Input

	ToCloudIntegrationAzureActivityLogArrayOutput() CloudIntegrationAzureActivityLogArrayOutput
	ToCloudIntegrationAzureActivityLogArrayOutputWithContext(context.Context) CloudIntegrationAzureActivityLogArrayOutput
}

type CloudIntegrationAzureActivityLogArray []CloudIntegrationAzureActivityLogInput

func (CloudIntegrationAzureActivityLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationAzureActivityLog)(nil)).Elem()
}

func (i CloudIntegrationAzureActivityLogArray) ToCloudIntegrationAzureActivityLogArrayOutput() CloudIntegrationAzureActivityLogArrayOutput {
	return i.ToCloudIntegrationAzureActivityLogArrayOutputWithContext(context.Background())
}

func (i CloudIntegrationAzureActivityLogArray) ToCloudIntegrationAzureActivityLogArrayOutputWithContext(ctx context.Context) CloudIntegrationAzureActivityLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationAzureActivityLogArrayOutput)
}

// CloudIntegrationAzureActivityLogMapInput is an input type that accepts CloudIntegrationAzureActivityLogMap and CloudIntegrationAzureActivityLogMapOutput values.
// You can construct a concrete instance of `CloudIntegrationAzureActivityLogMapInput` via:
//
//	CloudIntegrationAzureActivityLogMap{ "key": CloudIntegrationAzureActivityLogArgs{...} }
type CloudIntegrationAzureActivityLogMapInput interface {
	pulumi.Input

	ToCloudIntegrationAzureActivityLogMapOutput() CloudIntegrationAzureActivityLogMapOutput
	ToCloudIntegrationAzureActivityLogMapOutputWithContext(context.Context) CloudIntegrationAzureActivityLogMapOutput
}

type CloudIntegrationAzureActivityLogMap map[string]CloudIntegrationAzureActivityLogInput

func (CloudIntegrationAzureActivityLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationAzureActivityLog)(nil)).Elem()
}

func (i CloudIntegrationAzureActivityLogMap) ToCloudIntegrationAzureActivityLogMapOutput() CloudIntegrationAzureActivityLogMapOutput {
	return i.ToCloudIntegrationAzureActivityLogMapOutputWithContext(context.Background())
}

func (i CloudIntegrationAzureActivityLogMap) ToCloudIntegrationAzureActivityLogMapOutputWithContext(ctx context.Context) CloudIntegrationAzureActivityLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationAzureActivityLogMapOutput)
}

type CloudIntegrationAzureActivityLogOutput struct{ *pulumi.OutputState }

func (CloudIntegrationAzureActivityLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationAzureActivityLog)(nil)).Elem()
}

func (o CloudIntegrationAzureActivityLogOutput) ToCloudIntegrationAzureActivityLogOutput() CloudIntegrationAzureActivityLogOutput {
	return o
}

func (o CloudIntegrationAzureActivityLogOutput) ToCloudIntegrationAzureActivityLogOutputWithContext(ctx context.Context) CloudIntegrationAzureActivityLogOutput {
	return o
}

// A list of point tag key-values to add to every point ingested using this integration.
func (o CloudIntegrationAzureActivityLogOutput) AdditionalTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.StringMapOutput { return v.AdditionalTags }).(pulumi.StringMapOutput)
}

// A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics.
func (o CloudIntegrationAzureActivityLogOutput) CategoryFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.StringArrayOutput { return v.CategoryFilters }).(pulumi.StringArrayOutput)
}

// Client ID for an Azure service account within your project.
func (o CloudIntegrationAzureActivityLogOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// Client secret for an Azure service account within your project.
func (o CloudIntegrationAzureActivityLogOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// Forces this resource to save, even if errors are present.
func (o CloudIntegrationAzureActivityLogOutput) ForceSave() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.BoolPtrOutput { return v.ForceSave }).(pulumi.BoolPtrOutput)
}

// The human-readable name of this integration.
func (o CloudIntegrationAzureActivityLogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A value denoting which cloud service this service integrates with.
func (o CloudIntegrationAzureActivityLogOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// How often, in minutes, to refresh the service.
func (o CloudIntegrationAzureActivityLogOutput) ServiceRefreshRateInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.IntPtrOutput { return v.ServiceRefreshRateInMinutes }).(pulumi.IntPtrOutput)
}

// Tenant ID for an Azure service account within your project.
func (o CloudIntegrationAzureActivityLogOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudIntegrationAzureActivityLog) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

type CloudIntegrationAzureActivityLogArrayOutput struct{ *pulumi.OutputState }

func (CloudIntegrationAzureActivityLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudIntegrationAzureActivityLog)(nil)).Elem()
}

func (o CloudIntegrationAzureActivityLogArrayOutput) ToCloudIntegrationAzureActivityLogArrayOutput() CloudIntegrationAzureActivityLogArrayOutput {
	return o
}

func (o CloudIntegrationAzureActivityLogArrayOutput) ToCloudIntegrationAzureActivityLogArrayOutputWithContext(ctx context.Context) CloudIntegrationAzureActivityLogArrayOutput {
	return o
}

func (o CloudIntegrationAzureActivityLogArrayOutput) Index(i pulumi.IntInput) CloudIntegrationAzureActivityLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudIntegrationAzureActivityLog {
		return vs[0].([]*CloudIntegrationAzureActivityLog)[vs[1].(int)]
	}).(CloudIntegrationAzureActivityLogOutput)
}

type CloudIntegrationAzureActivityLogMapOutput struct{ *pulumi.OutputState }

func (CloudIntegrationAzureActivityLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudIntegrationAzureActivityLog)(nil)).Elem()
}

func (o CloudIntegrationAzureActivityLogMapOutput) ToCloudIntegrationAzureActivityLogMapOutput() CloudIntegrationAzureActivityLogMapOutput {
	return o
}

func (o CloudIntegrationAzureActivityLogMapOutput) ToCloudIntegrationAzureActivityLogMapOutputWithContext(ctx context.Context) CloudIntegrationAzureActivityLogMapOutput {
	return o
}

func (o CloudIntegrationAzureActivityLogMapOutput) MapIndex(k pulumi.StringInput) CloudIntegrationAzureActivityLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudIntegrationAzureActivityLog {
		return vs[0].(map[string]*CloudIntegrationAzureActivityLog)[vs[1].(string)]
	}).(CloudIntegrationAzureActivityLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationAzureActivityLogInput)(nil)).Elem(), &CloudIntegrationAzureActivityLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationAzureActivityLogArrayInput)(nil)).Elem(), CloudIntegrationAzureActivityLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudIntegrationAzureActivityLogMapInput)(nil)).Elem(), CloudIntegrationAzureActivityLogMap{})
	pulumi.RegisterOutputType(CloudIntegrationAzureActivityLogOutput{})
	pulumi.RegisterOutputType(CloudIntegrationAzureActivityLogArrayOutput{})
	pulumi.RegisterOutputType(CloudIntegrationAzureActivityLogMapOutput{})
}
