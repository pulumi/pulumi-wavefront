// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GetAddress(ctx *pulumi.Context) string {
	return config.Get(ctx, "wavefront:address")
}
func GetHttpProxy(ctx *pulumi.Context) string {
	return config.Get(ctx, "wavefront:httpProxy")
}
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "wavefront:token")
}
