module github.com/pulumi/pulumi-wavefront/provider

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.30.0
	github.com/pulumi/pulumi/sdk/v3 v3.39.3
	github.com/vmware/terraform-provider-wavefront v0.0.0-20220902122515-594f66d9b135
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20220824175045-450992f2f5b9
