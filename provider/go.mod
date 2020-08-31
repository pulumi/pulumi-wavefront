module github.com/pulumi/pulumi-wavefront/provider

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	//Temporary replacement due to the upstream bug waiting to be fixed
	github.com/terraform-providers/terraform-provider-wavefront => github.com/stack72/terraform-provider-wavefront-1 v0.0.0-20200807193123-b466e196e207
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.11.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.7.3
	github.com/pulumi/pulumi/sdk/v2 v2.9.1-0.20200825190708-910aa96016cd
	github.com/terraform-providers/terraform-provider-wavefront v0.0.0-20200804234730-c9e202329e36
)
