package petstore

import (
	"github.com/Mrzrb/astra"
	"github.com/Mrzrb/astra/tests/integration/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubstituteTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	r := setupRouter()

	testAstra, err := helpers.SetupTestAstraWithDefaultConfig(t, r, astra.WithCustomTypeMapping(map[string]astra.TypeFormat{
		"github.com/Mrzrb/astra/tests/petstore.Tag": astra.TypeFormat{
			Type:   "string",
			Format: "tag",
		},
	}))
	require.NoError(t, err)

	schemas := testAstra.Path("components.schemas")

	require.Equal(t, "string", schemas.Search("petstore.Tag", "type").Data().(string))
	require.Equal(t, "tag", schemas.Search("petstore.Tag", "format").Data().(string))
}
