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

// Provides a Wavefront User Resource. This allows user accounts to be created, updated, and deleted.
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
//			_, err := wavefront.NewUser(ctx, "basic", &wavefront.UserArgs{
//				Email: pulumi.String("test+tftesting@example.com"),
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
// Users can be imported by using the `id`, e.g.:
//
// ```sh
// $ pulumi import wavefront:index/user:User some_user test@example.com
// ```
type User struct {
	pulumi.CustomResourceState

	// The customer the user is associated with.
	Customer pulumi.StringOutput `pulumi:"customer"`
	// The unique identifier of the user account to create. Must be a valid email address.
	Email pulumi.StringOutput `pulumi:"email"`
	// List of permission to grant to this user. Valid options are
	// `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
	// `hostTagManagement`, `metricsManagement`, and `userManagement`.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// List of user groups to this user.
	UserGroups pulumi.StringArrayOutput `pulumi:"userGroups"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("wavefront:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("wavefront:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The customer the user is associated with.
	Customer *string `pulumi:"customer"`
	// The unique identifier of the user account to create. Must be a valid email address.
	Email *string `pulumi:"email"`
	// List of permission to grant to this user. Valid options are
	// `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
	// `hostTagManagement`, `metricsManagement`, and `userManagement`.
	Permissions []string `pulumi:"permissions"`
	// List of user groups to this user.
	UserGroups []string `pulumi:"userGroups"`
}

type UserState struct {
	// The customer the user is associated with.
	Customer pulumi.StringPtrInput
	// The unique identifier of the user account to create. Must be a valid email address.
	Email pulumi.StringPtrInput
	// List of permission to grant to this user. Valid options are
	// `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
	// `hostTagManagement`, `metricsManagement`, and `userManagement`.
	Permissions pulumi.StringArrayInput
	// List of user groups to this user.
	UserGroups pulumi.StringArrayInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The customer the user is associated with.
	Customer *string `pulumi:"customer"`
	// The unique identifier of the user account to create. Must be a valid email address.
	Email string `pulumi:"email"`
	// List of permission to grant to this user. Valid options are
	// `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
	// `hostTagManagement`, `metricsManagement`, and `userManagement`.
	Permissions []string `pulumi:"permissions"`
	// List of user groups to this user.
	UserGroups []string `pulumi:"userGroups"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The customer the user is associated with.
	Customer pulumi.StringPtrInput
	// The unique identifier of the user account to create. Must be a valid email address.
	Email pulumi.StringInput
	// List of permission to grant to this user. Valid options are
	// `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
	// `hostTagManagement`, `metricsManagement`, and `userManagement`.
	Permissions pulumi.StringArrayInput
	// List of user groups to this user.
	UserGroups pulumi.StringArrayInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// The customer the user is associated with.
func (o UserOutput) Customer() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Customer }).(pulumi.StringOutput)
}

// The unique identifier of the user account to create. Must be a valid email address.
func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// List of permission to grant to this user. Valid options are
// `agentManagement`, `alertsManagement`, `dashboardManagement`, `embeddedCharts`, `eventsManagement`, `externalLinksManagement`,
// `hostTagManagement`, `metricsManagement`, and `userManagement`.
func (o UserOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.Permissions }).(pulumi.StringArrayOutput)
}

// List of user groups to this user.
func (o UserOutput) UserGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.UserGroups }).(pulumi.StringArrayOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
