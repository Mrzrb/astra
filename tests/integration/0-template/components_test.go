package petstore

import (
	"github.com/Mrzrb/astra/tests/integration/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchemas(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	r := setupRouter()

	testAstra, err := helpers.SetupTestAstraWithDefaultConfig(t, r)
	require.NoError(t, err)

	require.NotNil(t, testAstra)

	// Placeholder for integration test
}
