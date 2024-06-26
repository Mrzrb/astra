package helpers

import (
	"os"
	"testing"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/Mrzrb/astra"
	"github.com/Mrzrb/astra/inputs"
	"github.com/Mrzrb/astra/outputs"
	"github.com/stretchr/testify/require"
)

func SetupTestAstraWithDefaultConfig(t *testing.T, r *gin.Engine, options ...astra.Option) (*gabs.Container, error) {
	t.Helper()

	config := &astra.Config{
		Host: "localhost",
		Port: 8000,
	}

	return SetupTestAstra(t, r, config, options...)
}

func SetupTestAstra(t *testing.T, r *gin.Engine, config *astra.Config, options ...astra.Option) (*gabs.Container, error) {
	t.Helper()

	options = append(options, inputs.WithGinInput(r), outputs.WithOpenAPIOutput("./output.json"))

	gen := astra.New(options...)

	gen.SetConfig(config)

	err := gen.Parse()
	require.NoError(t, err)

	fileContents, err := os.ReadFile("./output.json")
	require.NoError(t, err)

	return gabs.ParseJSON(fileContents)
}
