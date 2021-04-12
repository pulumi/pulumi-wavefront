module github.com/pulumi/pulumi-wavefront/provider

go 1.16

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.23.0
	github.com/pulumi/pulumi/sdk/v2 v2.24.1
	github.com/vmware/terraform-provider-wavefront v0.0.0-20210301224956-247110356692
)
