package inputs

import (
	"github.com/Mrzrb/astra"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWithGinInput(t *testing.T) {
	service := &astra.Service{}

	require.Len(t, service.Inputs, 0)

	WithGinInput(nil)(service)

	require.Len(t, service.Inputs, 1)
}
