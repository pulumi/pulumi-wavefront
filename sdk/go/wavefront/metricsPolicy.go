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

// Provides a Wavefront Metrics Policy Resource. This allows management of Metrics Policy to control access to time series, histograms, and delta counters
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
//			everyone, err := wavefront.GetDefaultUserGroup(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = wavefront.NewMetricsPolicy(ctx, "main", &wavefront.MetricsPolicyArgs{
//				PolicyRules: wavefront.MetricsPolicyPolicyRuleArray{
//					&wavefront.MetricsPolicyPolicyRuleArgs{
//						Name:        pulumi.String("Allow All Metrics"),
//						Description: pulumi.String("Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."),
//						Prefixes: pulumi.StringArray{
//							pulumi.String("*"),
//						},
//						TagsAnded:  pulumi.Bool(false),
//						AccessType: pulumi.String("ALLOW"),
//						UserGroupIds: pulumi.StringArray{
//							pulumi.String(everyone.GroupId),
//						},
//					},
//				},
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
// ## Data Source
//
// Provides a Wavefront Metrics Policy Data Source. This allows looking up the current policy and associated rules.
//
// ### Example
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
//			policy, err := wavefront.LookupMetricsPolicy(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("policy", policy)
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Users can be imported by using the `updated_epoch_millis`, e.g.:
//
// ```sh
// $ pulumi import wavefront:index/metricsPolicy:MetricsPolicy some_metrics_policy 1651846476678
// ```
type MetricsPolicy struct {
	pulumi.CustomResourceState

	// The customer the user is associated with.
	Customer pulumi.StringOutput `pulumi:"customer"`
	// List of Metrics Policies, must have at least one entry.
	PolicyRules MetricsPolicyPolicyRuleArrayOutput `pulumi:"policyRules"`
	// When the policy was applied in epoch_millis.
	UpdatedEpochMillis pulumi.IntOutput `pulumi:"updatedEpochMillis"`
	// The accountId who applied the current policy.
	UpdaterId pulumi.StringOutput `pulumi:"updaterId"`
}

// NewMetricsPolicy registers a new resource with the given unique name, arguments, and options.
func NewMetricsPolicy(ctx *pulumi.Context,
	name string, args *MetricsPolicyArgs, opts ...pulumi.ResourceOption) (*MetricsPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyRules == nil {
		return nil, errors.New("invalid value for required argument 'PolicyRules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MetricsPolicy
	err := ctx.RegisterResource("wavefront:index/metricsPolicy:MetricsPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricsPolicy gets an existing MetricsPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricsPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricsPolicyState, opts ...pulumi.ResourceOption) (*MetricsPolicy, error) {
	var resource MetricsPolicy
	err := ctx.ReadResource("wavefront:index/metricsPolicy:MetricsPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricsPolicy resources.
type metricsPolicyState struct {
	// The customer the user is associated with.
	Customer *string `pulumi:"customer"`
	// List of Metrics Policies, must have at least one entry.
	PolicyRules []MetricsPolicyPolicyRule `pulumi:"policyRules"`
	// When the policy was applied in epoch_millis.
	UpdatedEpochMillis *int `pulumi:"updatedEpochMillis"`
	// The accountId who applied the current policy.
	UpdaterId *string `pulumi:"updaterId"`
}

type MetricsPolicyState struct {
	// The customer the user is associated with.
	Customer pulumi.StringPtrInput
	// List of Metrics Policies, must have at least one entry.
	PolicyRules MetricsPolicyPolicyRuleArrayInput
	// When the policy was applied in epoch_millis.
	UpdatedEpochMillis pulumi.IntPtrInput
	// The accountId who applied the current policy.
	UpdaterId pulumi.StringPtrInput
}

func (MetricsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsPolicyState)(nil)).Elem()
}

type metricsPolicyArgs struct {
	// List of Metrics Policies, must have at least one entry.
	PolicyRules []MetricsPolicyPolicyRule `pulumi:"policyRules"`
}

// The set of arguments for constructing a MetricsPolicy resource.
type MetricsPolicyArgs struct {
	// List of Metrics Policies, must have at least one entry.
	PolicyRules MetricsPolicyPolicyRuleArrayInput
}

func (MetricsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsPolicyArgs)(nil)).Elem()
}

type MetricsPolicyInput interface {
	pulumi.Input

	ToMetricsPolicyOutput() MetricsPolicyOutput
	ToMetricsPolicyOutputWithContext(ctx context.Context) MetricsPolicyOutput
}

func (*MetricsPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsPolicy)(nil)).Elem()
}

func (i *MetricsPolicy) ToMetricsPolicyOutput() MetricsPolicyOutput {
	return i.ToMetricsPolicyOutputWithContext(context.Background())
}

func (i *MetricsPolicy) ToMetricsPolicyOutputWithContext(ctx context.Context) MetricsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsPolicyOutput)
}

// MetricsPolicyArrayInput is an input type that accepts MetricsPolicyArray and MetricsPolicyArrayOutput values.
// You can construct a concrete instance of `MetricsPolicyArrayInput` via:
//
//	MetricsPolicyArray{ MetricsPolicyArgs{...} }
type MetricsPolicyArrayInput interface {
	pulumi.Input

	ToMetricsPolicyArrayOutput() MetricsPolicyArrayOutput
	ToMetricsPolicyArrayOutputWithContext(context.Context) MetricsPolicyArrayOutput
}

type MetricsPolicyArray []MetricsPolicyInput

func (MetricsPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricsPolicy)(nil)).Elem()
}

func (i MetricsPolicyArray) ToMetricsPolicyArrayOutput() MetricsPolicyArrayOutput {
	return i.ToMetricsPolicyArrayOutputWithContext(context.Background())
}

func (i MetricsPolicyArray) ToMetricsPolicyArrayOutputWithContext(ctx context.Context) MetricsPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsPolicyArrayOutput)
}

// MetricsPolicyMapInput is an input type that accepts MetricsPolicyMap and MetricsPolicyMapOutput values.
// You can construct a concrete instance of `MetricsPolicyMapInput` via:
//
//	MetricsPolicyMap{ "key": MetricsPolicyArgs{...} }
type MetricsPolicyMapInput interface {
	pulumi.Input

	ToMetricsPolicyMapOutput() MetricsPolicyMapOutput
	ToMetricsPolicyMapOutputWithContext(context.Context) MetricsPolicyMapOutput
}

type MetricsPolicyMap map[string]MetricsPolicyInput

func (MetricsPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricsPolicy)(nil)).Elem()
}

func (i MetricsPolicyMap) ToMetricsPolicyMapOutput() MetricsPolicyMapOutput {
	return i.ToMetricsPolicyMapOutputWithContext(context.Background())
}

func (i MetricsPolicyMap) ToMetricsPolicyMapOutputWithContext(ctx context.Context) MetricsPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsPolicyMapOutput)
}

type MetricsPolicyOutput struct{ *pulumi.OutputState }

func (MetricsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsPolicy)(nil)).Elem()
}

func (o MetricsPolicyOutput) ToMetricsPolicyOutput() MetricsPolicyOutput {
	return o
}

func (o MetricsPolicyOutput) ToMetricsPolicyOutputWithContext(ctx context.Context) MetricsPolicyOutput {
	return o
}

// The customer the user is associated with.
func (o MetricsPolicyOutput) Customer() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsPolicy) pulumi.StringOutput { return v.Customer }).(pulumi.StringOutput)
}

// List of Metrics Policies, must have at least one entry.
func (o MetricsPolicyOutput) PolicyRules() MetricsPolicyPolicyRuleArrayOutput {
	return o.ApplyT(func(v *MetricsPolicy) MetricsPolicyPolicyRuleArrayOutput { return v.PolicyRules }).(MetricsPolicyPolicyRuleArrayOutput)
}

// When the policy was applied in epoch_millis.
func (o MetricsPolicyOutput) UpdatedEpochMillis() pulumi.IntOutput {
	return o.ApplyT(func(v *MetricsPolicy) pulumi.IntOutput { return v.UpdatedEpochMillis }).(pulumi.IntOutput)
}

// The accountId who applied the current policy.
func (o MetricsPolicyOutput) UpdaterId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsPolicy) pulumi.StringOutput { return v.UpdaterId }).(pulumi.StringOutput)
}

type MetricsPolicyArrayOutput struct{ *pulumi.OutputState }

func (MetricsPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricsPolicy)(nil)).Elem()
}

func (o MetricsPolicyArrayOutput) ToMetricsPolicyArrayOutput() MetricsPolicyArrayOutput {
	return o
}

func (o MetricsPolicyArrayOutput) ToMetricsPolicyArrayOutputWithContext(ctx context.Context) MetricsPolicyArrayOutput {
	return o
}

func (o MetricsPolicyArrayOutput) Index(i pulumi.IntInput) MetricsPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetricsPolicy {
		return vs[0].([]*MetricsPolicy)[vs[1].(int)]
	}).(MetricsPolicyOutput)
}

type MetricsPolicyMapOutput struct{ *pulumi.OutputState }

func (MetricsPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricsPolicy)(nil)).Elem()
}

func (o MetricsPolicyMapOutput) ToMetricsPolicyMapOutput() MetricsPolicyMapOutput {
	return o
}

func (o MetricsPolicyMapOutput) ToMetricsPolicyMapOutputWithContext(ctx context.Context) MetricsPolicyMapOutput {
	return o
}

func (o MetricsPolicyMapOutput) MapIndex(k pulumi.StringInput) MetricsPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetricsPolicy {
		return vs[0].(map[string]*MetricsPolicy)[vs[1].(string)]
	}).(MetricsPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricsPolicyInput)(nil)).Elem(), &MetricsPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricsPolicyArrayInput)(nil)).Elem(), MetricsPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricsPolicyMapInput)(nil)).Elem(), MetricsPolicyMap{})
	pulumi.RegisterOutputType(MetricsPolicyOutput{})
	pulumi.RegisterOutputType(MetricsPolicyArrayOutput{})
	pulumi.RegisterOutputType(MetricsPolicyMapOutput{})
}
