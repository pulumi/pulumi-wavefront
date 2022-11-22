module github.com/pulumi/pulumi-wavefront/provider

go 1.16

require (
	github.com/hashicorp/go-hclog v1.3.1 // indirect
	github.com/hashicorp/terraform-plugin-go v0.14.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.24.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.33.0
	github.com/pulumi/pulumi/sdk/v3 v3.44.2
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/vmware/terraform-provider-wavefront v0.0.0-20221109122518-8c834edeaf49
	golang.org/x/net v0.2.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20221122203342-430f685de305
