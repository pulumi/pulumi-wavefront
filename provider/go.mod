module github.com/pulumi/pulumi-wavefront/provider

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.18.0
	github.com/pulumi/pulumi/sdk/v3 v3.23.2
	github.com/vmware/terraform-provider-wavefront v0.0.0-20220311175142-5c491df5bb2e
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
