package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	checkWavefrontToken(t)
	return integration.ProgramTestOptions{
		ExpectRefreshChanges:     true,
		AllowEmptyPreviewChanges: true,
	}
}

func checkWavefrontToken(t *testing.T) {
	token := os.Getenv("WAVEFRONT_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing WAVEFRONT_TOKEN environment variable")
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}
