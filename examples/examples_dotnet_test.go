// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build dotnet || all
// +build dotnet all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Wavefront",
		},
	})

	return baseCsharp
}
