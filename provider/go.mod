module github.com/pulumi/pulumi-wavefront/provider

go 1.14

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.19.0
	github.com/pulumi/pulumi/sdk/v2 v2.20.1-0.20210212181059-f4b0fa86fedc
	github.com/vmware/terraform-provider-wavefront v0.0.0-20210201162339-95440ce2c2f6
)
