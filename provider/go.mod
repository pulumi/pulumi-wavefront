module github.com/pulumi/pulumi-wavefront/provider

go 1.16

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.22.1
	github.com/pulumi/pulumi/sdk/v2 v2.22.1-0.20210310211618-1f16423ede4c
	github.com/vmware/terraform-provider-wavefront v0.0.0-20210301224956-247110356692
)
