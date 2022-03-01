module github.com/pulumi/pulumi-wavefront/provider

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.19.1
	github.com/pulumi/pulumi/sdk/v3 v3.25.0
	github.com/vmware/terraform-provider-wavefront v0.0.0-20210610195536-769fabacff12
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
