// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wavefront

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Wavefront Cloud Integration for GCP. This allows GCP cloud integrations to be created,
// updated, and deleted.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-wavefront/sdk/go/wavefront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wavefront.NewCloudIntegrationGcp(ctx, "gcp", &wavefront.CloudIntegrationGcpArgs{
// 			JsonKey:   pulumi.String(fmt.Sprintf("%v%v", "{...your gcp key ...}\n", "\n")),
// 			ProjectId: pulumi.String("example-gcp-project"),
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
// GCP Cloud Integrations can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import wavefront:index/cloudIntegrationGcp:CloudIntegrationGcp gcp a411c16b-3cf7-4f03-bf11-8ca05aab898d
// ```
type CloudIntegrationGcp struct {
	pulumi.CustomResourceState

	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags pulumi.StringMapOutput `pulumi:"additionalTags"`
	// A list of Google Cloud Platform (GCP) services.  Valid values are `APPENGINE`,
	// `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
	// `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
	// `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
	// `TPU`, `VPN`
	Categories pulumi.StringArrayOutput `pulumi:"categories"`
	// Forces this resource to save, even if errors are present
	ForceSave pulumi.BoolPtrOutput `pulumi:"forceSave"`
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must be at least granted Monitoring Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey pulumi.StringOutput `pulumi:"jsonKey"`
	// A regular expression that a metric name must match (case-insensitively) in order to be ingested
	MetricFilterRegex pulumi.StringPtrOutput `pulumi:"metricFilterRegex"`
	// The human-readable name of this integration
	Name pulumi.StringOutput `pulumi:"name"`
	// The Google Cloud Platform (GCP) Project Id
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// A value denoting which cloud service this service integrates with
	Service pulumi.StringOutput `pulumi:"service"`
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes pulumi.IntPtrOutput `pulumi:"serviceRefreshRateInMinutes"`
}

// NewCloudIntegrationGcp registers a new resource with the given unique name, arguments, and options.
func NewCloudIntegrationGcp(ctx *pulumi.Context,
	name string, args *CloudIntegrationGcpArgs, opts ...pulumi.ResourceOption) (*CloudIntegrationGcp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JsonKey == nil {
		return nil, errors.New("invalid value for required argument 'JsonKey'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource CloudIntegrationGcp
	err := ctx.RegisterResource("wavefront:index/cloudIntegrationGcp:CloudIntegrationGcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudIntegrationGcp gets an existing CloudIntegrationGcp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudIntegrationGcp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudIntegrationGcpState, opts ...pulumi.ResourceOption) (*CloudIntegrationGcp, error) {
	var resource CloudIntegrationGcp
	err := ctx.ReadResource("wavefront:index/cloudIntegrationGcp:CloudIntegrationGcp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudIntegrationGcp resources.
type cloudIntegrationGcpState struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// A list of Google Cloud Platform (GCP) services.  Valid values are `APPENGINE`,
	// `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
	// `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
	// `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
	// `TPU`, `VPN`
	Categories []string `pulumi:"categories"`
	// Forces this resource to save, even if errors are present
	ForceSave *bool `pulumi:"forceSave"`
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must be at least granted Monitoring Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey *string `pulumi:"jsonKey"`
	// A regular expression that a metric name must match (case-insensitively) in order to be ingested
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The human-readable name of this integration
	Name *string `pulumi:"name"`
	// The Google Cloud Platform (GCP) Project Id
	ProjectId *string `pulumi:"projectId"`
	// A value denoting which cloud service this service integrates with
	Service *string `pulumi:"service"`
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

type CloudIntegrationGcpState struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags pulumi.StringMapInput
	// A list of Google Cloud Platform (GCP) services.  Valid values are `APPENGINE`,
	// `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
	// `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
	// `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
	// `TPU`, `VPN`
	Categories pulumi.StringArrayInput
	// Forces this resource to save, even if errors are present
	ForceSave pulumi.BoolPtrInput
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must be at least granted Monitoring Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey pulumi.StringPtrInput
	// A regular expression that a metric name must match (case-insensitively) in order to be ingested
	MetricFilterRegex pulumi.StringPtrInput
	// The human-readable name of this integration
	Name pulumi.StringPtrInput
	// The Google Cloud Platform (GCP) Project Id
	ProjectId pulumi.StringPtrInput
	// A value denoting which cloud service this service integrates with
	Service pulumi.StringPtrInput
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationGcpState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationGcpState)(nil)).Elem()
}

type cloudIntegrationGcpArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags map[string]string `pulumi:"additionalTags"`
	// A list of Google Cloud Platform (GCP) services.  Valid values are `APPENGINE`,
	// `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
	// `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
	// `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
	// `TPU`, `VPN`
	Categories []string `pulumi:"categories"`
	// Forces this resource to save, even if errors are present
	ForceSave *bool `pulumi:"forceSave"`
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must be at least granted Monitoring Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey string `pulumi:"jsonKey"`
	// A regular expression that a metric name must match (case-insensitively) in order to be ingested
	MetricFilterRegex *string `pulumi:"metricFilterRegex"`
	// The human-readable name of this integration
	Name *string `pulumi:"name"`
	// The Google Cloud Platform (GCP) Project Id
	ProjectId string `pulumi:"projectId"`
	// A value denoting which cloud service this service integrates with
	Service string `pulumi:"service"`
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes *int `pulumi:"serviceRefreshRateInMinutes"`
}

// The set of arguments for constructing a CloudIntegrationGcp resource.
type CloudIntegrationGcpArgs struct {
	// A list of point tag key-values to add to every point ingested using this integration
	AdditionalTags pulumi.StringMapInput
	// A list of Google Cloud Platform (GCP) services.  Valid values are `APPENGINE`,
	// `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
	// `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
	// `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
	// `TPU`, `VPN`
	Categories pulumi.StringArrayInput
	// Forces this resource to save, even if errors are present
	ForceSave pulumi.BoolPtrInput
	// Private key for a Google Cloud Platform (GCP) service account within your project.
	// The account must be at least granted Monitoring Viewer permissions. This key must be in the JSON format generated by GCP.
	JsonKey pulumi.StringInput
	// A regular expression that a metric name must match (case-insensitively) in order to be ingested
	MetricFilterRegex pulumi.StringPtrInput
	// The human-readable name of this integration
	Name pulumi.StringPtrInput
	// The Google Cloud Platform (GCP) Project Id
	ProjectId pulumi.StringInput
	// A value denoting which cloud service this service integrates with
	Service pulumi.StringInput
	// How often, in minutes, to refresh the service
	ServiceRefreshRateInMinutes pulumi.IntPtrInput
}

func (CloudIntegrationGcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudIntegrationGcpArgs)(nil)).Elem()
}

type CloudIntegrationGcpInput interface {
	pulumi.Input

	ToCloudIntegrationGcpOutput() CloudIntegrationGcpOutput
	ToCloudIntegrationGcpOutputWithContext(ctx context.Context) CloudIntegrationGcpOutput
}

func (*CloudIntegrationGcp) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudIntegrationGcp)(nil))
}

func (i *CloudIntegrationGcp) ToCloudIntegrationGcpOutput() CloudIntegrationGcpOutput {
	return i.ToCloudIntegrationGcpOutputWithContext(context.Background())
}

func (i *CloudIntegrationGcp) ToCloudIntegrationGcpOutputWithContext(ctx context.Context) CloudIntegrationGcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpOutput)
}

func (i *CloudIntegrationGcp) ToCloudIntegrationGcpPtrOutput() CloudIntegrationGcpPtrOutput {
	return i.ToCloudIntegrationGcpPtrOutputWithContext(context.Background())
}

func (i *CloudIntegrationGcp) ToCloudIntegrationGcpPtrOutputWithContext(ctx context.Context) CloudIntegrationGcpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpPtrOutput)
}

type CloudIntegrationGcpPtrInput interface {
	pulumi.Input

	ToCloudIntegrationGcpPtrOutput() CloudIntegrationGcpPtrOutput
	ToCloudIntegrationGcpPtrOutputWithContext(ctx context.Context) CloudIntegrationGcpPtrOutput
}

type cloudIntegrationGcpPtrType CloudIntegrationGcpArgs

func (*cloudIntegrationGcpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationGcp)(nil))
}

func (i *cloudIntegrationGcpPtrType) ToCloudIntegrationGcpPtrOutput() CloudIntegrationGcpPtrOutput {
	return i.ToCloudIntegrationGcpPtrOutputWithContext(context.Background())
}

func (i *cloudIntegrationGcpPtrType) ToCloudIntegrationGcpPtrOutputWithContext(ctx context.Context) CloudIntegrationGcpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpPtrOutput)
}

// CloudIntegrationGcpArrayInput is an input type that accepts CloudIntegrationGcpArray and CloudIntegrationGcpArrayOutput values.
// You can construct a concrete instance of `CloudIntegrationGcpArrayInput` via:
//
//          CloudIntegrationGcpArray{ CloudIntegrationGcpArgs{...} }
type CloudIntegrationGcpArrayInput interface {
	pulumi.Input

	ToCloudIntegrationGcpArrayOutput() CloudIntegrationGcpArrayOutput
	ToCloudIntegrationGcpArrayOutputWithContext(context.Context) CloudIntegrationGcpArrayOutput
}

type CloudIntegrationGcpArray []CloudIntegrationGcpInput

func (CloudIntegrationGcpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CloudIntegrationGcp)(nil))
}

func (i CloudIntegrationGcpArray) ToCloudIntegrationGcpArrayOutput() CloudIntegrationGcpArrayOutput {
	return i.ToCloudIntegrationGcpArrayOutputWithContext(context.Background())
}

func (i CloudIntegrationGcpArray) ToCloudIntegrationGcpArrayOutputWithContext(ctx context.Context) CloudIntegrationGcpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpArrayOutput)
}

// CloudIntegrationGcpMapInput is an input type that accepts CloudIntegrationGcpMap and CloudIntegrationGcpMapOutput values.
// You can construct a concrete instance of `CloudIntegrationGcpMapInput` via:
//
//          CloudIntegrationGcpMap{ "key": CloudIntegrationGcpArgs{...} }
type CloudIntegrationGcpMapInput interface {
	pulumi.Input

	ToCloudIntegrationGcpMapOutput() CloudIntegrationGcpMapOutput
	ToCloudIntegrationGcpMapOutputWithContext(context.Context) CloudIntegrationGcpMapOutput
}

type CloudIntegrationGcpMap map[string]CloudIntegrationGcpInput

func (CloudIntegrationGcpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CloudIntegrationGcp)(nil))
}

func (i CloudIntegrationGcpMap) ToCloudIntegrationGcpMapOutput() CloudIntegrationGcpMapOutput {
	return i.ToCloudIntegrationGcpMapOutputWithContext(context.Background())
}

func (i CloudIntegrationGcpMap) ToCloudIntegrationGcpMapOutputWithContext(ctx context.Context) CloudIntegrationGcpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudIntegrationGcpMapOutput)
}

type CloudIntegrationGcpOutput struct {
	*pulumi.OutputState
}

func (CloudIntegrationGcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudIntegrationGcp)(nil))
}

func (o CloudIntegrationGcpOutput) ToCloudIntegrationGcpOutput() CloudIntegrationGcpOutput {
	return o
}

func (o CloudIntegrationGcpOutput) ToCloudIntegrationGcpOutputWithContext(ctx context.Context) CloudIntegrationGcpOutput {
	return o
}

func (o CloudIntegrationGcpOutput) ToCloudIntegrationGcpPtrOutput() CloudIntegrationGcpPtrOutput {
	return o.ToCloudIntegrationGcpPtrOutputWithContext(context.Background())
}

func (o CloudIntegrationGcpOutput) ToCloudIntegrationGcpPtrOutputWithContext(ctx context.Context) CloudIntegrationGcpPtrOutput {
	return o.ApplyT(func(v CloudIntegrationGcp) *CloudIntegrationGcp {
		return &v
	}).(CloudIntegrationGcpPtrOutput)
}

type CloudIntegrationGcpPtrOutput struct {
	*pulumi.OutputState
}

func (CloudIntegrationGcpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudIntegrationGcp)(nil))
}

func (o CloudIntegrationGcpPtrOutput) ToCloudIntegrationGcpPtrOutput() CloudIntegrationGcpPtrOutput {
	return o
}

func (o CloudIntegrationGcpPtrOutput) ToCloudIntegrationGcpPtrOutputWithContext(ctx context.Context) CloudIntegrationGcpPtrOutput {
	return o
}

type CloudIntegrationGcpArrayOutput struct{ *pulumi.OutputState }

func (CloudIntegrationGcpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudIntegrationGcp)(nil))
}

func (o CloudIntegrationGcpArrayOutput) ToCloudIntegrationGcpArrayOutput() CloudIntegrationGcpArrayOutput {
	return o
}

func (o CloudIntegrationGcpArrayOutput) ToCloudIntegrationGcpArrayOutputWithContext(ctx context.Context) CloudIntegrationGcpArrayOutput {
	return o
}

func (o CloudIntegrationGcpArrayOutput) Index(i pulumi.IntInput) CloudIntegrationGcpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudIntegrationGcp {
		return vs[0].([]CloudIntegrationGcp)[vs[1].(int)]
	}).(CloudIntegrationGcpOutput)
}

type CloudIntegrationGcpMapOutput struct{ *pulumi.OutputState }

func (CloudIntegrationGcpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CloudIntegrationGcp)(nil))
}

func (o CloudIntegrationGcpMapOutput) ToCloudIntegrationGcpMapOutput() CloudIntegrationGcpMapOutput {
	return o
}

func (o CloudIntegrationGcpMapOutput) ToCloudIntegrationGcpMapOutputWithContext(ctx context.Context) CloudIntegrationGcpMapOutput {
	return o
}

func (o CloudIntegrationGcpMapOutput) MapIndex(k pulumi.StringInput) CloudIntegrationGcpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CloudIntegrationGcp {
		return vs[0].(map[string]CloudIntegrationGcp)[vs[1].(string)]
	}).(CloudIntegrationGcpOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudIntegrationGcpOutput{})
	pulumi.RegisterOutputType(CloudIntegrationGcpPtrOutput{})
	pulumi.RegisterOutputType(CloudIntegrationGcpArrayOutput{})
	pulumi.RegisterOutputType(CloudIntegrationGcpMapOutput{})
}
