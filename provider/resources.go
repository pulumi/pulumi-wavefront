// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wavefront

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi-wavefront/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/vmware/terraform-provider-wavefront/wavefront"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "wavefront"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(wavefront.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "wavefront",
		Description: "A Pulumi package for creating and managing wavefront cloud resources.",
		Keywords:    []string{"pulumi", "wavefront"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-wavefront",
		GitHubOrg:   "vmware",
		Config:      map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"wavefront_alert":                                {Tok: makeResource(mainMod, "Alert")},
			"wavefront_alert_target":                         {Tok: makeResource(mainMod, "AlertTarget")},
			"wavefront_cloud_integration_app_dynamics":       {Tok: makeResource(mainMod, "CloudIntegrationAppDynamics")},
			"wavefront_cloud_integration_aws_external_id":    {Tok: makeResource(mainMod, "CloudIntegrationAwsExternalId")},
			"wavefront_cloud_integration_azure":              {Tok: makeResource(mainMod, "CloudIntegrationAzure")},
			"wavefront_cloud_integration_azure_activity_log": {Tok: makeResource(mainMod, "CloudIntegrationAzureActivityLog")},
			"wavefront_cloud_integration_cloudtrail":         {Tok: makeResource(mainMod, "CloudIntegrationCloudTrail")},
			"wavefront_cloud_integration_cloudwatch":         {Tok: makeResource(mainMod, "CloudIntegrationCloudWatch")},
			"wavefront_cloud_integration_ec2":                {Tok: makeResource(mainMod, "CloudIntegrationEc2")},
			"wavefront_cloud_integration_gcp":                {Tok: makeResource(mainMod, "CloudIntegrationGcp")},
			"wavefront_cloud_integration_gcp_billing":        {Tok: makeResource(mainMod, "CloudIntegrationGcpBilling")},
			"wavefront_cloud_integration_newrelic":           {Tok: makeResource(mainMod, "CloudIntegrationNewRelic")},
			"wavefront_cloud_integration_tesla":              {Tok: makeResource(mainMod, "CloudIntegrationTesla")},
			"wavefront_dashboard":                            {Tok: makeResource(mainMod, "Dashboard")},
			"wavefront_dashboard_json": {
				Tok: makeResource(mainMod, "DashboardJson"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"dashboard_json": {
						CSharpName: "JSON",
					},
				},
			},
			"wavefront_derived_metric":     {Tok: makeResource(mainMod, "DerivedMetric")},
			"wavefront_role":               {Tok: makeResource(mainMod, "Role")},
			"wavefront_user":               {Tok: makeResource(mainMod, "User")},
			"wavefront_user_group":         {Tok: makeResource(mainMod, "UserGroup")},
			"wavefront_service_account":    {Tok: makeResource(mainMod, "ServiceAccount")},
			"wavefront_maintenance_window": {Tok: makeResource(mainMod, "MaintenanceWindow")},
			"wavefront_external_link":      {Tok: makeResource(mainMod, "ExternalLink")},
			"wavefront_ingestion_policy":   {Tok: makeResource(mainMod, "IngestionPolicy")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"wavefront_default_user_group": {Tok: makeDataSource(mainMod, "getDefaultUserGroup")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0-alpha.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0a1,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "3.*-*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				"wavefront": "Wavefront",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
