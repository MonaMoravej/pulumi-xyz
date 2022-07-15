// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xyz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Random struct {
	pulumi.CustomResourceState

	Length pulumi.IntOutput    `pulumi:"length"`
	Result pulumi.StringOutput `pulumi:"result"`
}

// NewRandom registers a new resource with the given unique name, arguments, and options.
func NewRandom(ctx *pulumi.Context,
	name string, args *RandomArgs, opts ...pulumi.ResourceOption) (*Random, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	var resource Random
	err := ctx.RegisterResource("xyz:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandom gets an existing Random resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomState, opts ...pulumi.ResourceOption) (*Random, error) {
	var resource Random
	err := ctx.ReadResource("xyz:index:Random", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Random resources.
type randomState struct {
}

type RandomState struct {
}

func (RandomState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomState)(nil)).Elem()
}

type randomArgs struct {
	Length int `pulumi:"length"`
}

// The set of arguments for constructing a Random resource.
type RandomArgs struct {
	Length pulumi.IntInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}

type RandomInput interface {
	pulumi.Input

	ToRandomOutput() RandomOutput
	ToRandomOutputWithContext(ctx context.Context) RandomOutput
}

func (*Random) ElementType() reflect.Type {
	return reflect.TypeOf((**Random)(nil)).Elem()
}

func (i *Random) ToRandomOutput() RandomOutput {
	return i.ToRandomOutputWithContext(context.Background())
}

func (i *Random) ToRandomOutputWithContext(ctx context.Context) RandomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomOutput)
}

// RandomArrayInput is an input type that accepts RandomArray and RandomArrayOutput values.
// You can construct a concrete instance of `RandomArrayInput` via:
//
//          RandomArray{ RandomArgs{...} }
type RandomArrayInput interface {
	pulumi.Input

	ToRandomArrayOutput() RandomArrayOutput
	ToRandomArrayOutputWithContext(context.Context) RandomArrayOutput
}

type RandomArray []RandomInput

func (RandomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Random)(nil)).Elem()
}

func (i RandomArray) ToRandomArrayOutput() RandomArrayOutput {
	return i.ToRandomArrayOutputWithContext(context.Background())
}

func (i RandomArray) ToRandomArrayOutputWithContext(ctx context.Context) RandomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomArrayOutput)
}

// RandomMapInput is an input type that accepts RandomMap and RandomMapOutput values.
// You can construct a concrete instance of `RandomMapInput` via:
//
//          RandomMap{ "key": RandomArgs{...} }
type RandomMapInput interface {
	pulumi.Input

	ToRandomMapOutput() RandomMapOutput
	ToRandomMapOutputWithContext(context.Context) RandomMapOutput
}

type RandomMap map[string]RandomInput

func (RandomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Random)(nil)).Elem()
}

func (i RandomMap) ToRandomMapOutput() RandomMapOutput {
	return i.ToRandomMapOutputWithContext(context.Background())
}

func (i RandomMap) ToRandomMapOutputWithContext(ctx context.Context) RandomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomMapOutput)
}

type RandomOutput struct{ *pulumi.OutputState }

func (RandomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Random)(nil)).Elem()
}

func (o RandomOutput) ToRandomOutput() RandomOutput {
	return o
}

func (o RandomOutput) ToRandomOutputWithContext(ctx context.Context) RandomOutput {
	return o
}

func (o RandomOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *Random) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

func (o RandomOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *Random) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

type RandomArrayOutput struct{ *pulumi.OutputState }

func (RandomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Random)(nil)).Elem()
}

func (o RandomArrayOutput) ToRandomArrayOutput() RandomArrayOutput {
	return o
}

func (o RandomArrayOutput) ToRandomArrayOutputWithContext(ctx context.Context) RandomArrayOutput {
	return o
}

func (o RandomArrayOutput) Index(i pulumi.IntInput) RandomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Random {
		return vs[0].([]*Random)[vs[1].(int)]
	}).(RandomOutput)
}

type RandomMapOutput struct{ *pulumi.OutputState }

func (RandomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Random)(nil)).Elem()
}

func (o RandomMapOutput) ToRandomMapOutput() RandomMapOutput {
	return o
}

func (o RandomMapOutput) ToRandomMapOutputWithContext(ctx context.Context) RandomMapOutput {
	return o
}

func (o RandomMapOutput) MapIndex(k pulumi.StringInput) RandomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Random {
		return vs[0].(map[string]*Random)[vs[1].(string)]
	}).(RandomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomInput)(nil)).Elem(), &Random{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomArrayInput)(nil)).Elem(), RandomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomMapInput)(nil)).Elem(), RandomMap{})
	pulumi.RegisterOutputType(RandomOutput{})
	pulumi.RegisterOutputType(RandomArrayOutput{})
	pulumi.RegisterOutputType(RandomMapOutput{})
}